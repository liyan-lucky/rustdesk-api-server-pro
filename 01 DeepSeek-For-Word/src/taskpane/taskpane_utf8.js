// 导入DeepSeek API模块
import {
    STORAGE_KEY,
    DEEPSEEK_API_URL,
    safeGet,
    maskApiKey,
    isValidApiKeyFormat,
    validateApiKeyFormat,
    updateKeyStatus,
    updateButtonStates,
    validatePrompt,
    getCurrentApiKey,
    callDeepSeekAPI,
    saveApiKey as apiSaveApiKey,
    loadApiKey as apiLoadApiKey,
    clearApiKey as apiClearApiKey
} from '../api/deepseek-api.js';

const TRANSLATION_LANGUAGES = {
    'zh': '中文',
    'en': '英语',
    'ja': '日语',
    'ko': '韩语',
    'fr': '法语',
    'de': '德语',
    'es': '西班牙语',
    'ru': '俄语'
};

// 统一文本替换管理器 - 专门处理居中检测和字符样式反向还原，并保留嵌入对象
const UnifiedTextReplacer = {
    // 获取选中文本的完整格式信息，包括居中检测和嵌入对象检测
    async getSelectionFormatInfo() {
        let formatInfo = {};
        
        await Word.run(async (context) => {
            const selection = context.document.getSelection();
            
            try {
                // 获取整体文本和段落格式
                selection.load([
                    "text",
                    "paragraphFormat/alignment", "paragraphFormat/lineSpacing", 
                    "paragraphFormat/firstLineIndent"
                ]);
                await context.sync();
                
                formatInfo.originalText = selection.text;
                formatInfo.paragraphFormat = {
                    alignment: safeGet(selection.paragraphFormat, 'alignment', Word.Alignment.left),
                    lineSpacing: safeGet(selection.paragraphFormat, 'lineSpacing'),
                    firstLineIndent: safeGet(selection.paragraphFormat, 'firstLineIndent')
                };
                
                // 检测是否居中
                formatInfo.isCentered = safeGet(selection.paragraphFormat, 'alignment', Word.Alignment.left) === Word.Alignment.centered;
                
                // 获取字符级格式
                formatInfo.characterFormats = await this.getCharacterFormats(context, selection);
                
                // 检测嵌入对象（图片、线条、表格等）
                formatInfo.hasEmbeddedObjects = await this.detectEmbeddedObjects(context, selection);
                
                console.log("格式信息获取成功:", {
                    textLength: formatInfo.originalText.length,
                    isCentered: formatInfo.isCentered,
                    characterFormats: formatInfo.characterFormats.length,
                    hasEmbeddedObjects: formatInfo.hasEmbeddedObjects
                });
                
            } catch (error) {
                console.error("获取格式信息失败:", error);
                throw error;
            }
        });
        
        return formatInfo;
    },
    
    // 检测选区中是否包含嵌入对象
    async detectEmbeddedObjects(context, selection) {
        try {
            // 方法1: 检查内联图片
            const inlinePictures = selection.inlinePictures;
            inlinePictures.load("items");
            await context.sync();
            
            if (inlinePictures.items.length > 0) {
                console.log("检测到内联图片:", inlinePictures.items.length);
                return true;
            }
            
            // 方法2: 检查形状（线条等）
            const shapes = selection.shapeParent ? selection.shapeParent.shapes : null;
            if (shapes) {
                shapes.load("items");
                await context.sync();
                if (shapes.items.length > 0) {
                    console.log("检测到形状:", shapes.items.length);
                    return true;
                }
            }
            
            // 方法3: 检查表格
            const tables = selection.tables;
            tables.load("items");
            await context.sync();
            if (tables.items.length > 0) {
                console.log("检测到表格:", tables.items.length);
                return true;
            }
            
            // 方法4: 检查内容控件
            const contentControls = selection.contentControls;
            contentControls.load("items");
            await context.sync();
            if (contentControls.items.length > 0) {
                console.log("检测到内容控件:", contentControls.items.length);
                return true;
            }
            
            return false;
        } catch (error) {
            console.warn("嵌入对象检测失败:", error);
            return false;
        }
    },
    
    // 获取字符格式（支持降级）
    async getCharacterFormats(context, selection) {
        const text = selection.text;
        const characterFormats = [];
        
        try {
            // 方法1: 按字符获取
            for (let i = 0; i < text.length; i++) {
                try {
                    const charRange = selection.getRange('Character', i);
                    charRange.load([
                        "text",
                        "font/name", "font/size", "font/color", "font/bold", "font/italic", 
                        "font/underline", "font/highlightColor"
                    ]);
                    
                    await context.sync();
                    
                    characterFormats.push({
                        position: i,
                        text: charRange.text,
                        font: {
                            name: safeGet(charRange.font, 'name'),
                            size: safeGet(charRange.font, 'size'),
                            color: safeGet(charRange.font, 'color'),
                            bold: safeGet(charRange.font, 'bold', false),
                            italic: safeGet(charRange.font, 'italic', false),
                            underline: safeGet(charRange.font, 'underline', false),
                            highlightColor: safeGet(charRange.font, 'highlightColor')
                        }
                    });
                } catch (error) {
                    console.warn(`获取第${i}个字符格式失败:`, error);
                    // 使用前一个字符的格式作为后备
                    const prevFormat = characterFormats[i - 1] || { font: {} };
                    characterFormats.push({
                        position: i,
                        text: text[i],
                        font: { ...prevFormat.font }
                    });
                }
            }
        } catch (error) {
            console.error("字符级格式获取失败，使用单词级格式:", error);
            // 方法2: 降级到单词级格式
            return await this.getWordLevelFormats(context, selection, text);
        }
        
        return characterFormats;
    },
    
    // 获取单词级格式（降级方案）
    async getWordLevelFormats(context, selection, text) {
        const characterFormats = [];
        
        try {
            const ranges = selection.getTextRanges([' '], false);
            ranges.load([
                "text",
                "font/name", "font/size", "font/color", "font/bold", "font/italic", 
                "font/underline", "font/highlightColor"
            ]);
            await context.sync();
            
            let currentPosition = 0;
            ranges.items.forEach(range => {
                const wordText = range.text;
                for (let i = 0; i < wordText.length; i++) {
                    characterFormats.push({
                        position: currentPosition + i,
                        text: wordText[i],
                        font: {
                            name: safeGet(range.font, 'name'),
                            size: safeGet(range.font, 'size'),
                            color: safeGet(range.font, 'color'),
                            bold: safeGet(range.font, 'bold', false),
                            italic: safeGet(range.font, 'italic', false),
                            underline: safeGet(range.font, 'underline', false),
                            highlightColor: safeGet(range.font, 'highlightColor')
                        }
                    });
                }
                currentPosition += wordText.length;
            });
            
            console.log("单词级格式获取成功");
        } catch (error) {
            console.error("单词级格式获取失败，使用整体格式:", error);
            // 方法3: 最终降级到整体格式
            selection.load([
                "font/name", "font/size", "font/color", "font/bold", "font/italic", 
                "font/underline", "font/highlightColor"
            ]);
            await context.sync();
            
            for (let i = 0; i < text.length; i++) {
                characterFormats.push({
                    position: i,
                    text: text[i],
                    font: {
                        name: safeGet(selection.font, 'name'),
                        size: safeGet(selection.font, 'size'),
                        color: safeGet(selection.font, 'color'),
                        bold: safeGet(selection.font, 'bold', false),
                        italic: safeGet(selection.font, 'italic', false),
                        underline: safeGet(selection.font, 'underline', false),
                        highlightColor: safeGet(selection.font, 'highlightColor')
                    }
                });
            }
        }
        
        return characterFormats;
    },
    
    // 统一替换方法 - 倒着还原原样式，并处理嵌入对象
    async unifiedReplace(newText, originalFormat) {
        await Word.run(async (context) => {
            const selection = context.document.getSelection();
            
            try {
                console.log("开始统一替换:", {
                    originalLength: originalFormat.originalText.length,
                    newLength: newText.length,
                    isCentered: originalFormat.isCentered,
                    hasEmbeddedObjects: originalFormat.hasEmbeddedObjects
                });
                
                // 如果选区包含嵌入对象，使用安全替换方法
                if (originalFormat.hasEmbeddedObjects) {
                    console.log("检测到嵌入对象，使用安全替换方法");
                    await this.safeReplaceWithEmbeddedObjects(context, selection, newText, originalFormat);
                } else {
                    // 没有嵌入对象，使用标准替换方法
                    selection.insertText(newText, Word.InsertLocation.replace);
                    await context.sync();
                    
                    // 重新选择新文本
                    const newSelection = context.document.getSelection();
                    
                    // 应用段落格式（特别处理居中）
                    this.applyParagraphFormat(newSelection.paragraphFormat, originalFormat.paragraphFormat);
                    
                    // 倒着应用字符格式 - 从后往前应用
                    await this.applyReverseCharacterFormats(context, newSelection, newText, originalFormat);
                }
                
                await context.sync();
                
            } catch (error) {
                console.error("统一替换失败:", error);
                await this.fallbackReplace(selection, newText, originalFormat);
            }
        });
    },
    
    // 安全替换方法 - 保留嵌入对象
    async safeReplaceWithEmbeddedObjects(context, selection, newText, originalFormat) {
        try {
            console.log("执行安全替换，保留嵌入对象");
            
            // 方法1: 尝试使用范围替换，只替换文本内容
            const ranges = selection.getTextRanges([' '], false);
            ranges.load(["text"]);
            await context.sync();
            
            // 计算总文本长度
            let totalTextLength = 0;
            ranges.items.forEach(range => {
                totalTextLength += range.text.length;
            });
            
            // 如果文本长度匹配，逐个替换文本范围
            if (totalTextLength > 0) {
                console.log("使用文本范围替换，保留嵌入对象");
                await this.replaceTextRanges(context, ranges, newText, originalFormat);
            } else {
                // 方法2: 如果无法获取文本范围，使用警告和降级方法
                console.warn("无法获取文本范围，使用降级方法");
                await this.fallbackReplaceWithWarning(selection, newText, originalFormat);
            }
            
        } catch (error) {
            console.error("安全替换失败:", error);
            await this.fallbackReplaceWithWarning(selection, newText, originalFormat);
        }
    },
    
    // 替换文本范围，保留嵌入对象
    async replaceTextRanges(context, ranges, newText, originalFormat) {
        const textParts = newText.split(' ');
        let currentPosition = 0;
        
        for (let i = 0; i < ranges.items.length; i++) {
            const range = ranges.items[i];
            const originalText = range.text;
            
            // 获取对应的新文本部分
            const newTextPart = i < textParts.length ? textParts[i] : '';
            
            if (newTextPart !== originalText) {
                // 只替换文本内容，不删除范围
                range.insertText(newTextPart, Word.InsertLocation.replace);
            }
            
            // 应用格式到当前范围
            await this.applyFormatToRange(context, range, originalFormat, currentPosition, originalText.length);
            
            currentPosition += originalText.length;
            
            // 在单词之间添加空格（除了最后一个单词）
            if (i < ranges.items.length - 1 && i < textParts.length - 1) {
                const spaceRange = range.getRange('After', 0);
                spaceRange.insertText(' ', Word.InsertLocation.replace);
            }
        }
        
        await context.sync();
    },
    
    // 应用格式到范围
    async applyFormatToRange(context, range, originalFormat, startPosition, length) {
        try {
            // 应用段落格式
            if (originalFormat.paragraphFormat) {
                this.applyParagraphFormat(range.paragraphFormat, originalFormat.paragraphFormat);
            }
            
            // 应用字符格式
            for (let j = 0; j < length; j++) {
                const globalPosition = startPosition + j;
                if (globalPosition < originalFormat.characterFormats.length) {
                    const charFormat = originalFormat.characterFormats[globalPosition];
                    if (charFormat) {
                        try {
                            const charRange = range.getRange(j, j);
                            this.applyFontFormat(charRange.font, charFormat.font);
                        } catch (error) {
                            console.warn(`应用字符格式失败:`, error);
                        }
                    }
                }
            }
        } catch (error) {
            console.warn("应用范围格式失败:", error);
        }
    },
    
    // 带警告的降级替换
    async fallbackReplaceWithWarning(selection, newText, originalFormat) {
        console.warn("使用带警告的降级替换方法");
        
        // 显示警告信息
        showResult("⚠️ 检测到嵌入对象，正在尝试保留...");
        
        try {
            // 先保存嵌入对象信息
            const embeddedObjectsInfo = await this.getEmbeddedObjectsInfo(selection);
            
            // 执行标准替换
            selection.insertText(newText, Word.InsertLocation.replace);
            await selection.context.sync();
            
            // 重新选择新文本
            const newSelection = selection.context.document.getSelection();
            
            // 应用格式
            this.applyParagraphFormat(newSelection.paragraphFormat, originalFormat.paragraphFormat);
            await this.applyReverseCharacterFormats(selection.context, newSelection, newText, originalFormat);
            
            // 显示成功信息
            if (embeddedObjectsInfo.count > 0) {
                showResult(`✅ 文本替换完成，但 ${embeddedObjectsInfo.count} 个嵌入对象可能已丢失`);
            } else {
                showResult("✅ 文本替换完成");
            }
            
        } catch (error) {
            console.error("带警告的降级替换失败:", error);
            throw error;
        }
    },
    
    // 获取嵌入对象信息
    async getEmbeddedObjectsInfo(selection) {
        let count = 0;
        let types = [];
        
        try {
            // 检查内联图片
            const inlinePictures = selection.inlinePictures;
            inlinePictures.load("items");
            await selection.context.sync();
            count += inlinePictures.items.length;
            if (inlinePictures.items.length > 0) types.push('图片');
            
            // 检查形状
            const shapes = selection.shapeParent ? selection.shapeParent.shapes : null;
            if (shapes) {
                shapes.load("items");
                await selection.context.sync();
                count += shapes.items.length;
                if (shapes.items.length > 0) types.push('形状');
            }
            
            // 检查表格
            const tables = selection.tables;
            tables.load("items");
            await selection.context.sync();
            count += tables.items.length;
            if (tables.items.length > 0) types.push('表格');
            
        } catch (error) {
            console.warn("获取嵌入对象信息失败:", error);
        }
        
        return { count, types };
    },
    
    // 倒着应用字符格式 - 从后往前还原样式
    async applyReverseCharacterFormats(context, range, newText, originalFormat) {
        if (!originalFormat || !originalFormat.characterFormats || originalFormat.characterFormats.length === 0) {
            // 没有逐字符格式，尽量应用段落或整体字体
            if (originalFormat && originalFormat.paragraphFormat) {
                this.applyParagraphFormat(range.paragraphFormat, originalFormat.paragraphFormat);
            }
            if (originalFormat && originalFormat.characterFormats && originalFormat.characterFormats[0]) {
                this.applyFontFormat(range.font, originalFormat.characterFormats[0].font);
            }
            return;
        }

        const originalFormats = originalFormat.characterFormats;
        const originalLength = originalFormat.originalText.length || originalFormats.length;
        const newLength = newText.length;

        // 应用段落格式
        if (originalFormat.paragraphFormat) {
            try {
                this.applyParagraphFormat(range.paragraphFormat, originalFormat.paragraphFormat);
            } catch (err) {
                console.warn("应用段落格式失败", err);
            }
        }

        // 倒着应用字符格式 - 从后往前
        try {
            for (let j = newLength - 1; j >= 0; j--) {
                // 计算映射位置 - 从后往前映射
                const mappedIndex = originalLength === 0 ? 0 : 
                    Math.min(originalLength - 1, Math.round((newLength - 1 - j) * (originalLength / newLength)));
                
                const srcFormat = originalFormats[mappedIndex] ? 
                    originalFormats[mappedIndex].font : 
                    (originalFormats[0] ? originalFormats[0].font : null);
                
                if (!srcFormat) continue;

                try {
                    const charRange = range.getRange(j, j);
                    this.applyFontFormat(charRange.font, srcFormat);
                } catch (error) {
                    console.warn(`应用第 ${j} 个字符格式失败:`, error);
                }
            }
            await context.sync();
        } catch (error) {
            console.warn("批量应用格式失败，尝试逐个同步:", error);
            // 失败时退化为每字符同步
            for (let j = newLength - 1; j >= 0; j--) {
                const mappedIndex = originalLength === 0 ? 0 : 
                    Math.min(originalLength - 1, Math.round((newLength - 1 - j) * (originalLength / newLength)));
                
                const srcFormat = originalFormats[mappedIndex] ? 
                    originalFormats[mappedIndex].font : 
                    (originalFormats[0] ? originalFormats[0].font : null);
                
                if (!srcFormat) continue;
                
                try {
                    const charRange = range.getRange(j, j);
                    this.applyFontFormat(charRange.font, srcFormat);
                    await context.sync();
                } catch (err) {
                    console.warn(`回退应用第 ${j} 个字符失败:`, err);
                }
            }
        }
    },
    
    // 应用段落格式
    applyParagraphFormat(targetParagraphFormat, sourceParagraphFormat) {
        if (!sourceParagraphFormat) return;
        
        try {
            if (sourceParagraphFormat.alignment !== undefined) {
                targetParagraphFormat.alignment = sourceParagraphFormat.alignment;
            }
            if (sourceParagraphFormat.lineSpacing !== undefined) {
                targetParagraphFormat.lineSpacing = sourceParagraphFormat.lineSpacing;
            }
            if (sourceParagraphFormat.firstLineIndent !== undefined) {
                targetParagraphFormat.firstLineIndent = sourceParagraphFormat.firstLineIndent;
            }
        } catch (error) {
            console.warn("应用段落格式失败:", error);
        }
    },

    // 应用字体格式
    applyFontFormat(targetFont, sourceFont) {
        if (!sourceFont) return;
        
        try {
            if (sourceFont.name !== undefined) {
                targetFont.name = sourceFont.name;
            }
            if (sourceFont.size !== undefined) {
                targetFont.size = sourceFont.size;
            }
            if (sourceFont.color !== undefined) {
                targetFont.color = sourceFont.color;
            }
            if (sourceFont.bold !== undefined) {
                targetFont.bold = sourceFont.bold;
            }
            if (sourceFont.italic !== undefined) {
                targetFont.italic = sourceFont.italic;
            }
            if (sourceFont.underline !== undefined) {
                targetFont.underline = sourceFont.underline;
            }
            if (sourceFont.highlightColor !== undefined) {
                targetFont.highlightColor = sourceFont.highlightColor;
            }
        } catch (error) {
            console.warn("应用字体格式失败:", error);
        }
    },

    // 降级替换方法
    async fallbackReplace(selection, newText, originalFormat) {
        try {
            selection.insertText(newText, Word.InsertLocation.replace);
            
            // 尽量应用格式
            if (originalFormat && originalFormat.paragraphFormat) {
                this.applyParagraphFormat(selection.paragraphFormat, originalFormat.paragraphFormat);
            }
            if (originalFormat && originalFormat.characterFormats && originalFormat.characterFormats.length > 0) {
                const firstFormat = originalFormat.characterFormats.find(cf => cf.text && cf.text.trim()) || originalFormat.characterFormats[0];
                this.applyFontFormat(selection.font, firstFormat.font);
            }
            
            await selection.context.sync();
        } catch (error) {
            console.error("降级替换失败:", error);
            // 最终降级：只替换文本，不保留格式
            selection.insertText(newText, Word.InsertLocation.replace);
            await selection.context.sync();
        }
    }
};

// 精确字符位置格式管理对象（保留向后兼容）
const ExactPositionFormatManager = {
    // 调试函数：检查Word.Alignment枚举值
    debugAlignmentValues() {
        console.log("Word.Alignment 枚举值检查:", {
            centered: Word.Alignment.centered,
            left: Word.Alignment.left,
            right: Word.Alignment.right,
            justified: Word.Alignment.justified,
            distributed: Word.Alignment.distributed,
            undefined: Word.Alignment.undefined
        });
        
        // 检查居中对齐的数值
        console.log("Word.Alignment.centered 数值类型:", typeof Word.Alignment.centered, "值:", Word.Alignment.centered);
        
        // 尝试转换为字符串查看
        try {
            console.log("Word.Alignment.centered 字符串表示:", Word.Alignment.centered.toString());
        } catch (e) {
            console.log("无法转换为字符串:", e);
        }
    },
    
    // 获取选中文本的精确字符位置格式
    async getExactPositionFormats() {
        let formatInfo = {};
        
        await Word.run(async (context) => {
            const selection = context.document.getSelection();
            
            try {
                // 获取整体文本和段落格式
                selection.load([
                    "text",
                    "paragraphFormat/alignment", "paragraphFormat/lineSpacing", 
                    "paragraphFormat/firstLineIndent"
                ]);
                await context.sync();
                
                formatInfo.originalText = selection.text;
                formatInfo.paragraphFormat = {
                    alignment: safeGet(selection.paragraphFormat, 'alignment', Word.Alignment.left),
                    lineSpacing: safeGet(selection.paragraphFormat, 'lineSpacing'),
                    firstLineIndent: safeGet(selection.paragraphFormat, 'firstLineIndent')
                };
                
                // 获取字符级格式
                formatInfo.characterFormats = await this.getCharacterByCharacterFormats(context, selection);
                
                console.log("字符位置格式获取成功:", formatInfo.characterFormats.length, "个字符");
                
                // 调试：检查对齐方式
                console.log("获取格式时对齐方式检查:", {
                    原始对齐方式: selection.paragraphFormat.alignment,
                    原始对齐方式值: selection.paragraphFormat.alignment,
                    是否居中: selection.paragraphFormat.alignment === Word.Alignment.centered ? '是' : '否',
                    存储的对齐方式: formatInfo.paragraphFormat.alignment,
                    存储的对齐方式值: formatInfo.paragraphFormat.alignment,
                    存储的是否居中: formatInfo.paragraphFormat.alignment === Word.Alignment.centered ? '是' : '否'
                });
                
            } catch (error) {
                console.error("获取格式失败:", error);
                throw error;
            }
        });
        
        return formatInfo;
    },
    
    // 按字符逐个获取格式 - 优化版本，批量处理字符
    async getCharacterByCharacterFormats(context, selection) {
        const text = selection.text;
        const characterFormats = [];
        
        // 如果文本太长，使用单词级格式以提高性能
        if (text.length > 500) {
            console.log("文本过长，使用单词级格式以提高性能");
            return await this.getWordLevelFormats(context, selection, text);
        }
        
        try {
            // 批量获取字符格式
            const charRanges = [];
            for (let i = 0; i < text.length; i++) {
                try {
                    const charRange = selection.getRange('Character', i);
                    charRange.load([
                        "text",
                        "font/name", "font/size", "font/color", "font/bold", "font/italic", 
                        "font/underline", "font/highlightColor"
                    ]);
                    charRanges.push(charRange);
                } catch (error) {
                    console.warn(`创建第${i}个字符范围失败:`, error);
                    // 创建占位符
                    charRanges.push(null);
                }
            }
            
            // 批量同步
            await context.sync();
            
            // 处理结果
            for (let i = 0; i < text.length; i++) {
                const charRange = charRanges[i];
                if (charRange) {
                    try {
                        characterFormats.push({
                            position: i,
                            text: charRange.text,
                            font: {
                                name: safeGet(charRange.font, 'name'),
                                size: safeGet(charRange.font, 'size'),
                                color: safeGet(charRange.font, 'color'),
                                bold: safeGet(charRange.font, 'bold', false),
                                italic: safeGet(charRange.font, 'italic', false),
                                underline: safeGet(charRange.font, 'underline', false),
                                highlightColor: safeGet(charRange.font, 'highlightColor')
                            }
                        });
                    } catch (error) {
                        console.warn(`处理第${i}个字符格式失败:`, error);
                        // 使用前一个字符的格式作为后备
                        const prevFormat = characterFormats[i - 1] || { font: {} };
                        characterFormats.push({
                            position: i,
                            text: text[i],
                            font: { ...prevFormat.font }
                        });
                    }
                } else {
                    // 使用前一个字符的格式作为后备
                    const prevFormat = characterFormats[i - 1] || { font: {} };
                    characterFormats.push({
                        position: i,
                        text: text[i],
                        font: { ...prevFormat.font }
                    });
                }
            }
        } catch (error) {
            console.error("字符级格式获取失败，使用单词级格式:", error);
            // 方法2: 降级到单词级格式
            return await this.getWordLevelFormats(context, selection, text);
        }
        
        return characterFormats;
    },
    
    // 获取单词级格式（降级方案）
    async getWordLevelFormats(context, selection, text) {
        const characterFormats = [];
        
        try {
            const ranges = selection.getTextRanges([' '], false);
            ranges.load([
                "text",
                "font/name", "font/size", "font/color", "font/bold", "font/italic", 
                "font/underline", "font/highlightColor"
            ]);
            await context.sync();
            
            let currentPosition = 0;
            ranges.items.forEach(range => {
                const wordText = range.text;
                for (let i = 0; i < wordText.length; i++) {
                    characterFormats.push({
                        position: currentPosition + i,
                        text: wordText[i],
                        font: {
                            name: safeGet(range.font, 'name'),
                            size: safeGet(range.font, 'size'),
                            color: safeGet(range.font, 'color'),
                            bold: safeGet(range.font, 'bold', false),
                            italic: safeGet(range.font, 'italic', false),
                            underline: safeGet(range.font, 'underline', false),
                            highlightColor: safeGet(range.font, 'highlightColor')
                        }
                    });
                }
                currentPosition += wordText.length;
            });
            
            console.log("单词级格式获取成功");
        } catch (error) {
            console.error("单词级格式获取失败，使用整体格式:", error);
            // 方法3: 最终降级到整体格式
            selection.load([
                "font/name", "font/size", "font/color", "font/bold", "font/italic", 
                "font/underline", "font/highlightColor"
            ]);
            await context.sync();
            
            for (let i = 0; i < text.length; i++) {
                characterFormats.push({
                    position: i,
                    text: text[i],
                    font: {
                        name: safeGet(selection.font, 'name'),
                        size: safeGet(selection.font, 'size'),
                        color: safeGet(selection.font, 'color'),
                        bold: safeGet(selection.font, 'bold', false),
                        italic: safeGet(selection.font, 'italic', false),
                        underline: safeGet(selection.font, 'underline', false),
                        highlightColor: safeGet(selection.font, 'highlightColor')
                    }
                });
            }
        }
        
        return characterFormats;
    },
    
    // 精确按位置替换文本 - 保持字符位置格式
    async exactPositionReplace(newText, originalFormat) {
        await Word.run(async (context) => {
            const selection = context.document.getSelection();
            
            try {
                console.log("开始精确位置替换");
                console.log("原始文本长度:", originalFormat.originalText.length);
                console.log("新文本长度:", newText.length);
                console.log("字符格式数量:", originalFormat.characterFormats.length);
                
                // 先插入新文本
                selection.insertText(newText, Word.InsertLocation.replace);
                await context.sync();
                
                // 重新选择新文本
                const newSelection = context.document.getSelection();
                
                // 应用段落格式
                ExactPositionFormatManager.applyParagraphFormat(newSelection.paragraphFormat, originalFormat.paragraphFormat);
                
                // 按位置应用字符格式（逐字符方式）
                await ExactPositionFormatManager.applyCharacterWiseFormats(context, newSelection, newText, originalFormat);
                
                await context.sync();
                
            } catch (error) {
                console.error("精确位置替换失败:", error);
                // 在降级替换时传入 originalFormat 以便尽量保留格式
                await ExactPositionFormatManager.fallbackReplace(selection, newText, originalFormat);
            }
        });
    },
    
    // 按字符逐个应用字体格式，保证原文每个字符样式尽量保留
    async applyCharacterWiseFormats(context, range, newText, originalFormat) {
        console.log("applyCharacterWiseFormats 开始:", {
            新文本长度: newText.length,
            字符格式数量: originalFormat?.characterFormats?.length || 0,
            段落格式存在: !!(originalFormat && originalFormat.paragraphFormat),
            第一个字符粗体: originalFormat?.characterFormats?.[0]?.font?.bold || false,
            段落字体粗体: originalFormat?.paragraphFont?.bold || false,
            段落格式对齐方式: originalFormat?.paragraphFormat?.alignment,
            是否居中: originalFormat?.paragraphFormat?.alignment === Word.Alignment.centered ? '是' : '否',
            段落格式对齐方式值: originalFormat?.paragraphFormat?.alignment,
            Word_Alignment_centered值: Word.Alignment.centered,
            段落格式对象: originalFormat?.paragraphFormat
        });

        if (!originalFormat || !originalFormat.characterFormats || originalFormat.characterFormats.length === 0) {
            console.log("没有逐字符格式，尝试应用段落或整体字体");
            
            // 首先应用段落格式（包括对齐方式）
            if (originalFormat && originalFormat.paragraphFormat) {
                console.log("应用段落格式:", {
                    对齐方式: originalFormat.paragraphFormat.alignment,
                    对齐方式值: originalFormat.paragraphFormat.alignment,
                    是否居中: originalFormat.paragraphFormat.alignment === Word.Alignment.centered ? '是' : '否',
                    行间距: originalFormat.paragraphFormat.lineSpacing,
                    首行缩进: originalFormat.paragraphFormat.firstLineIndent,
                    Word_Alignment_centered值: Word.Alignment.centered
                });
                
                // 确保段落格式被正确应用
                console.log("调用applyParagraphFormat前，目标段落格式对象:", range.paragraphFormat);
                ExactPositionFormatManager.applyParagraphFormat(range.paragraphFormat, originalFormat.paragraphFormat);
                
                // 立即同步以确保段落格式生效
                await context.sync();
                
                // 验证段落格式是否应用成功
                range.paragraphFormat.load("alignment", "lineSpacing", "firstLineIndent");
                await context.sync();
                console.log("段落格式应用后验证:", {
                    实际对齐方式: range.paragraphFormat.alignment,
                    实际对齐方式值: range.paragraphFormat.alignment,
                    是否居中: range.paragraphFormat.alignment === Word.Alignment.centered ? '是' : '否',
                    实际行间距: range.paragraphFormat.lineSpacing,
                    实际首行缩进: range.paragraphFormat.firstLineIndent
                });
            }
            
            // 应用字体格式 - 优先使用段落整体字体，然后使用第一个字符的字体
            let fontToApply = null;
            if (originalFormat && originalFormat.paragraphFont) {
                console.log("使用段落整体字体格式:", {
                    粗体: originalFormat.paragraphFont.bold || false,
                    斜体: originalFormat.paragraphFont.italic || false,
                    字体名称: originalFormat.paragraphFont.name,
                    字体大小: originalFormat.paragraphFont.size,
                    字体颜色: originalFormat.paragraphFont.color
                });
                fontToApply = originalFormat.paragraphFont;
            } else if (originalFormat && originalFormat.characterFormats && originalFormat.characterFormats[0]) {
                console.log("使用第一个字符的字体格式:", {
                    粗体: originalFormat.characterFormats[0].font?.bold || false,
                    斜体: originalFormat.characterFormats[0].font?.italic || false,
                    字体名称: originalFormat.characterFormats[0].font?.name,
                    字体大小: originalFormat.characterFormats[0].font?.size
                });
                fontToApply = originalFormat.characterFormats[0].font;
            }
            
            if (fontToApply) {
                console.log("应用整体字体格式到整个范围:", {
                    粗体: fontToApply.bold || false,
                    斜体: fontToApply.italic || false,
                    字体名称: fontToApply.name,
                    字体大小: fontToApply.size
                });
                ExactPositionFormatManager.applyFontFormat(range.font, fontToApply);
                
                // 立即同步以确保字体格式生效
                await context.sync();
                
                // 验证字体格式是否应用成功
                range.font.load(["bold", "italic", "name", "size"]);
                await context.sync();
                console.log("字体格式应用后验证:", {
                    实际粗体: range.font.bold,
                    实际斜体: range.font.italic,
                    实际字体名称: range.font.name,
                    实际字体大小: range.font.size
                });
            }
            
            return;
        }

        const originalFormats = originalFormat.characterFormats;
        const originalLength = originalFormat.originalText.length || originalFormats.length;
        const newLength = newText.length;

        // 尽量应用段落格式
        if (originalFormat.paragraphFormat) {
            try {
                console.log("应用段落格式:", {
                    对齐方式: originalFormat.paragraphFormat.alignment,
                    对齐方式值: originalFormat.paragraphFormat.alignment,
                    是否居中: originalFormat.paragraphFormat.alignment === Word.Alignment.centered ? '是' : '否',
                    Word_Alignment_centered值: Word.Alignment.centered,
                    段落格式对象: originalFormat.paragraphFormat
                });
                console.log("调用applyParagraphFormat前，目标段落格式对象:", range.paragraphFormat);
                ExactPositionFormatManager.applyParagraphFormat(range.paragraphFormat, originalFormat.paragraphFormat);
                
                // 立即同步以确保段落格式生效
                await context.sync();
                
                // 验证段落格式是否应用成功
                range.paragraphFormat.load("alignment");
                await context.sync();
                console.log("段落格式应用后验证:", {
                    实际对齐方式: range.paragraphFormat.alignment,
                    实际对齐方式值: range.paragraphFormat.alignment,
                    是否居中: range.paragraphFormat.alignment === Word.Alignment.centered ? '是' : '否'
                });
            } catch (err) {
                console.warn("applyCharacterWiseFormats: 应用段落格式失败", err);
            }
        }

        // 按字符逐个应用字体样式。为性能考虑，避免每次 sync；一次性设置后 await context.sync()
        try {
            let appliedCount = 0;
            let boldAppliedCount = 0;
            let italicAppliedCount = 0;
            
            for (let j = 0; j < newLength; j++) {
                // 改进映射逻辑：使用比例映射而不是简单的四舍五入
                const mappedIndex = originalLength === 0 ? 0 : 
                    Math.min(originalLength - 1, Math.floor(j * (originalLength / newLength)));
                
                // 优先使用映射位置的格式，如果没有则使用第一个字符的格式
                let srcFormat = null;
                if (originalFormats[mappedIndex] && originalFormats[mappedIndex].font) {
                    srcFormat = originalFormats[mappedIndex].font;
                } else if (originalFormats[0] && originalFormats[0].font) {
                    srcFormat = originalFormats[0].font;
                } else if (originalFormat.paragraphFont) {
                    srcFormat = originalFormat.paragraphFont;
                }
                
                if (!srcFormat) continue;

                try {
                    // 获取新文本对应位置的范围（以 range 开始为起点的相对位置）
                    const charRange = range.getRange(j, j);
                    // 直接应用字体格式到单个字符范围
                    console.log(`应用第 ${j} 个字符格式:`, {
                        映射索引: mappedIndex,
                        粗体: srcFormat.bold || false,
                        斜体: srcFormat.italic || false,
                        字体名称: srcFormat.name,
                        字体大小: srcFormat.size
                    });
                    ExactPositionFormatManager.applyFontFormat(charRange.font, srcFormat);
                    appliedCount++;
                    
                    // 记录粗体和斜体应用情况
                    if (srcFormat.bold === true) {
                        boldAppliedCount++;
                    }
                    if (srcFormat.italic === true) {
                        italicAppliedCount++;
                    }
                } catch (error) {
                    // 如果单个字符应用失败，则忽略继续
                    console.warn(`applyCharacterWiseFormats: 应用第 ${j} 个字符格式失败:`, error);
                }
            }
            await context.sync();
            console.log(`applyCharacterWiseFormats: 成功应用 ${appliedCount}/${newLength} 个字符格式，其中 ${boldAppliedCount} 个字符应用了粗体，${italicAppliedCount} 个字符应用了斜体`);
            
            // 最终验证字体格式
            range.font.load(["bold", "italic", "name", "size"]);
            await context.sync();
            console.log("最终字体格式验证:", {
                实际粗体: range.font.bold,
                实际斜体: range.font.italic,
                实际字体名称: range.font.name,
                实际字体大小: range.font.size
            });
        }
    },

    // 新增：在给定范围上执行逐字符替换并保留原文样式的函数（用于段落/选区）
    async exactPositionReplaceOnRange(context, range, newText, originalFormat) {
        try {
            // 详细的调试日志
            console.log("exactPositionReplaceOnRange 开始:", {
                段落索引: "需要从调用方传递",
                原文长度: originalFormat?.originalText?.length || 0,
                新文本长度: newText.length,
                段落格式存在: !!(originalFormat && originalFormat.paragraphFormat),
                对齐方式: originalFormat?.paragraphFormat?.alignment,
                是否居中: originalFormat?.paragraphFormat?.alignment === Word.Alignment.centered ? '是' : '否',
                字体粗体: originalFormat?.characterFormats?.[0]?.font?.bold || false,
                段落字体粗体: originalFormat?.paragraphFont?.bold || false,
                字符格式数量: originalFormat?.characterFormats?.length || 0
            });

            // 替换文本
            range.insertText(newText, Word.InsertLocation.replace);
            await context.sync();

            // 应用段落样式与逐字符字体样式
            // 注意：这里不直接设置段落格式，而是让applyCharacterWiseFormats统一处理
            // 这样可以避免重复设置和同步问题
            
            await ExactPositionFormatManager.applyCharacterWiseFormats(context, range, newText, originalFormat);
            
            // 额外验证：确保段落格式和字体格式正确应用
            try {
                // 验证段落格式
                range.paragraphFormat.load("alignment", "lineSpacing", "firstLineIndent");
                await context.sync();
                console.log("exactPositionReplaceOnRange 最终段落格式验证:", {
                    实际对齐方式: range.paragraphFormat.alignment,
                    实际对齐方式值: range.paragraphFormat.alignment,
                    是否居中: range.paragraphFormat.alignment === Word.Alignment.centered ? '是' : '否',
                    实际行间距: range.paragraphFormat.lineSpacing,
                    实际首行缩进: range.paragraphFormat.firstLineIndent
                });
                
                // 验证字体格式
                range.font.load(["bold", "italic", "name", "size"]);
                await context.sync();
                console.log("exactPositionReplaceOnRange 最终字体格式验证:", {
                    实际粗体: range.font.bold,
                    实际斜体: range.font.italic,
                    实际字体名称: range.font.name,
                    实际字体大小: range.font.size
                });
            } catch (verifyError) {
                console.warn("格式验证失败:", verifyError);
            }
        } catch (error) {
            console.error("exactPositionReplaceOnRange: 替换失败，回退为整体替换并尽量保留样式:", error);
            try {
                range.insertText(newText, Word.InsertLocation.replace);
                if (originalFormat && originalFormat.characterFormats && originalFormat.characterFormats.length > 0) {
                    const firstFormat = originalFormat.characterFormats.find(cf => cf.text && cf.text.trim()) || originalFormat.characterFormats[0];
                    ExactPositionFormatManager.applyFontFormat(range.font, firstFormat.font);
                }
                if (originalFormat && originalFormat.paragraphFormat) {
                    ExactPositionFormatManager.applyParagraphFormat(range.paragraphFormat, originalFormat.paragraphFormat);
                }
                await context.sync();
            } catch (err) {
                console.error("exactPositionReplaceOnRange fallback 失败:", err);
                throw err;
            }
        }
    },

    // 插入新内容
    async insertContent(newText) {
        await Word.run(async (context) => {
            const selection = context.document.getSelection();
            
            try {
                // 先插入新文本
                selection.insertText(newText, Word.InsertLocation.replace);
                await context.sync();
                
                // 重新选择新文本
                const newSelection = context.document.getSelection();
                
                // 应用段落格式
                ExactPositionFormatManager.applyParagraphFormat(newSelection.paragraphFormat, {
                    alignment: Word.Alignment.left,
                    lineSpacing: 1.15,
                    firstLineIndent: 0
                });
                
                // 按位置应用字符格式（逐字符方式）
                await ExactPositionFormatManager.applyCharacterWiseFormats(context, newSelection, newText, {
                    originalText: newText,
                    characterFormats: [],
                    paragraphFormat: null
                });
                
                await context.sync();
                
            } catch (error) {
                console.error("插入新内容失败:", error);
            }
        });
    },

    // 新增：获取文档中每个段落的文本和格式信息（段落与字符级）- 增强版本
    async getDocumentParagraphFormats() {
        const paragraphFormats = [];
        await Word.run(async (context) => {
            const body = context.document.body;
            const paragraphs = body.paragraphs;
            paragraphs.load("items");
            await context.sync();

            // 调试：检查Word.Alignment枚举值
            this.debugAlignmentValues();

            for (let i = 0; i < paragraphs.items.length; i++) {
                try {
                    const paragraph = paragraphs.items[i];
                    const range = paragraph.getRange();
                    
                    // 先获取段落基本信息和格式
                    range.load([
                        "text",
                        "paragraphFormat/alignment", "paragraphFormat/lineSpacing",
                        "paragraphFormat/firstLineIndent",
                        "font/name", "font/size", "font/color", "font/bold", "font/italic", 
                        "font/underline", "font/highlightColor"
                    ]);
                    await context.sync();

                    // 获取该段落的逐字符格式（可能降级为单词或整体）
                    let charFormats = [];
                    try {
                        charFormats = await ExactPositionFormatManager.getCharacterByCharacterFormats(context, range);
                    } catch (charError) {
                        console.warn(`getDocumentParagraphFormats: 获取第${i}段字符格式失败，使用段落整体格式:`, charError);
                        // 使用段落整体字体格式作为后备
                        for (let j = 0; j < range.text.length; j++) {
                            charFormats.push({
                                position: j,
                                text: range.text[j] || "",
                                font: {
                                    name: safeGet(range.font, 'name'),
                                    size: safeGet(range.font, 'size'),
                                    color: safeGet(range.font, 'color'),
                                    bold: safeGet(range.font, 'bold', false),
                                    italic: safeGet(range.font, 'italic', false),
                                    underline: safeGet(range.font, 'underline', false),
                                    highlightColor: safeGet(range.font, 'highlightColor')
                                }
                            });
                        }
                    }

                    // 确保对齐方式正确获取
                    const alignment = safeGet(range.paragraphFormat, 'alignment', Word.Alignment.left);
                    const isBold = safeGet(range.font, 'bold', false);
                    
                    // 详细调试信息 - 增强版
                    console.log(`获取第${i}段格式成功:`, {
                        文本长度: range.text.length,
                        对齐方式: alignment,
                        对齐方式值: alignment,
                        对齐方式类型: typeof alignment,
                        是否居中: alignment === Word.Alignment.centered ? '是' : '否',
                        Word_Alignment_centered值: Word.Alignment.centered,
                        Word_Alignment_centered类型: typeof Word.Alignment.centered,
                        相等性检查: `alignment (${alignment}) === Word.Alignment.centered (${Word.Alignment.centered}) = ${alignment === Word.Alignment.centered}`,
                        粗体: isBold,
                        粗体类型: typeof isBold,
                        斜体: safeGet(range.font, 'italic', false),
                        字体名称: safeGet(range.font, 'name'),
                        字体大小: safeGet(range.font, 'size'),
                        字体颜色: safeGet(range.font, 'color')
                    });
                    
                    // 检查Word.Alignment.centered的值
                    console.log(`Word.Alignment.centered = ${Word.Alignment.centered}, 类型 = ${typeof Word.Alignment.centered}, 当前段落对齐方式 = ${alignment}, 类型 = ${typeof alignment}`);
                    
                    // 检查是否真的居中对齐
                    if (alignment === Word.Alignment.centered) {
                        console.log(`✅ 第${i}段确认居中对齐！`);
                    } else if (alignment === 1) { // 有时Word.Alignment.centered可能是数值1
                        console.log(`⚠️ 第${i}段对齐方式为数值1，可能是居中对齐的数值表示`);
                    }
                    
                    paragraphFormats.push({
                        originalText: range.text,
                        paragraphFormat: {
                            alignment: alignment,
                            lineSpacing: safeGet(range.paragraphFormat, 'lineSpacing'),
                            firstLineIndent: safeGet(range.paragraphFormat, 'firstLineIndent')
                        },
                        characterFormats: charFormats,
                        // 添加段落整体字体信息作为后备
                        paragraphFont: {
                            name: safeGet(range.font, 'name'),
                            size: safeGet(range.font, 'size'),
                            color: safeGet(range.font, 'color'),
                            bold: isBold,
                            italic: safeGet(range.font, 'italic', false),
                            underline: safeGet(range.font, 'underline', false),
                            highlightColor: safeGet(range.font, 'highlightColor')
                        }
                    });
                    
                } catch (err) {
                    console.error(`getDocumentParagraphFormats: 获取第${i}段格式完全失败:`, err);
                    // 尝试获取段落的基本信息，即使格式获取失败
                    try {
                        const paragraph = paragraphs.items[i];
                        const range = paragraph.getRange();
                        range.load(["text"]);
                        await context.sync();
                        
                        // 提供一个更合理的默认格式对象，尽量保留可能的信息
                        paragraphFormats.push({
                            originalText: range.text || "",
                            paragraphFormat: {
                                alignment: Word.Alignment.left, // 默认左对齐，但用户应该知道这可能不正确
                                lineSpacing: 1.15,
                                firstLineIndent: 0
                            },
                            characterFormats: [],
                            paragraphFont: {
                                name: "Calibri",
                                size: 11,
                                color: "black",
                                bold: false,
                                italic: false,
                                underline: false,
                                highlightColor: ""
                            }
                        });
                    } catch (innerErr) {
                        console.error(`getDocumentParagraphFormats: 获取第${i}段基本信息也失败:`, innerErr);
                        // 提供最小化的默认格式对象
                        paragraphFormats.push({
                            originalText: "",
                            paragraphFormat: {
                                alignment: Word.Alignment.left,
                                lineSpacing: 1.15,
                                firstLineIndent: 0
                            },
                            characterFormats: [],
                            paragraphFont: {
                                name: "Calibri",
                                size: 11,
                                color: "black",
                                bold: false,
                                italic: false,
                                underline: false,
                                highlightColor: ""
                            }
                        });
                    }
                }
            }
            
            // 返回段落格式数组
            return paragraphFormats;
        });
    },

    // 应用段落格式
    applyParagraphFormat(targetParagraphFormat, sourceParagraphFormat) {
        if (!sourceParagraphFormat) return;
        
        try {
            if (sourceParagraphFormat.alignment !== undefined) {
                targetParagraphFormat.alignment = sourceParagraphFormat.alignment;
            }
            if (sourceParagraphFormat.lineSpacing !== undefined) {
                targetParagraphFormat.lineSpacing = sourceParagraphFormat.lineSpacing;
            }
            if (sourceParagraphFormat.firstLineIndent !== undefined) {
                targetParagraphFormat.firstLineIndent = sourceParagraphFormat.firstLineIndent;
            }
        } catch (error) {
            console.warn("应用段落格式失败:", error);
        }
    },

    // 应用字体格式
    applyFontFormat(targetFont, sourceFont) {
        if (!sourceFont) return;
        
        try {
            if (sourceFont.name !== undefined) {
                targetFont.name = sourceFont.name;
            }
            if (sourceFont.size !== undefined) {
                targetFont.size = sourceFont.size;
            }
            if (sourceFont.color !== undefined) {
                targetFont.color = sourceFont.color;
            }
            if (sourceFont.bold !== undefined) {
                targetFont.bold = sourceFont.bold;
            }
            if (sourceFont.italic !== undefined) {
                targetFont.italic = sourceFont.italic;
            }
            if (sourceFont.underline !== undefined) {
                targetFont.underline = sourceFont.underline;
            }
            if (sourceFont.highlightColor !== undefined) {
                targetFont.highlightColor = sourceFont.highlightColor;
            }
        } catch (error) {
            console.warn("应用字体格式失败:", error);
        }
    },

    // 降级替换方法
    async fallbackReplace(selection, newText, originalFormat) {
        try {
            selection.insertText(newText, Word.InsertLocation.replace);
            
            // 尽量应用格式
            if (originalFormat && originalFormat.paragraphFormat) {
                this.applyParagraphFormat(selection.paragraphFormat, originalFormat.paragraphFormat);
            }
            if (originalFormat && originalFormat.characterFormats && originalFormat.characterFormats.length > 0) {
                const firstFormat = originalFormat.characterFormats.find(cf => cf.text && cf.text.trim()) || originalFormat.characterFormats[0];
                this.applyFontFormat(selection.font, firstFormat.font);
            }
            
            await selection.context.sync();
        } catch (error) {
            console.error("降级替换失败:", error);
            // 最终降级：只替换文本，不保留格式
            selection.insertText(newText, Word.InsertLocation.replace);
            await selection.context.sync();
        }
    }
};

// 按钮状态管理器 - 防止闪烁
const ButtonStateManager = {
    // 设置按钮为处理状态（无闪烁）
    setProcessingState(button, originalText = null) {
        if (!button) return;
        
        // 保存原始状态
        if (originalText !== null) {
            button._originalText = originalText;
        }
        button._originalDisabled = button.disabled;
        button._originalBackgroundColor = button.style.backgroundColor;
        
        // 设置处理状态
        button.disabled = true;
        button.style.backgroundColor = "#0078d4"; // 保持蓝色，不变成灰色
        button.style.opacity = "0.8";
        button.style.cursor = "wait";
        
        if (button._originalText) {
            button.textContent = "处理中...";
        }
    },
    
    // 恢复按钮原始状态
    restoreOriginalState(button) {
        if (!button) return;
        
        button.disabled = button._originalDisabled !== undefined ? button._originalDisabled : false;
        button.style.backgroundColor = button._originalBackgroundColor || "#0078d4";
        button.style.opacity = "1";
        button.style.cursor = button.disabled ? "not-allowed" : "pointer";
        
        if (button._originalText) {
            button.textContent = button._originalText;
        }
    },
    
    // 批量设置处理状态
    setBatchProcessing(buttonIds, excludeButtonIds = []) {
        const excludeSet = new Set(excludeButtonIds);
        
        buttonIds.forEach(buttonId => {
            if (!excludeSet.has(buttonId)) {
                const button = document.getElementById(buttonId);
                if (button) {
                    const originalText = button.textContent;
                    this.setProcessingState(button, originalText);
                }
            }
        });
    },
    
    // 批量恢复状态
    restoreBatchState(buttonIds) {
        buttonIds.forEach(buttonId => {
            const button = document.getElementById(buttonId);
            if (button) {
                this.restoreOriginalState(button);
            }
        });
    }
};

// 应用初始化
Office.onReady((info) => {
    if (info.host === Office.HostType.Word) {
        console.log("Word加载完成，初始化插件...");
        initializeApp();
    }
});

// 初始化应用
function initializeApp() {
    console.log("开始初始化应用...");
    
    try {
        // 绑定事件
        document.getElementById("save-key-btn").onclick = saveApiKey;
        document.getElementById("clear-key-btn").onclick = clearApiKey;
        document.getElementById("toggle-key-visibility").onclick = toggleKeyVisibility;
        document.getElementById("generate-btn").onclick = generateContent;
        document.getElementById("rewrite-btn").onclick = rewriteSelection;
        document.getElementById("summarize-btn").onclick = summarizeDocument;
        document.getElementById("translate-btn").onclick = translateSelection;
        document.getElementById("translate-full-btn").onclick = translateFullDocument;
        document.getElementById("preview-format-btn").onclick = previewFormatting;
        
        console.log("事件绑定完成");
        
        // 输入框变化监听
        document.getElementById("api-key-input").addEventListener('input', validateApiKeyFormat);
        document.getElementById("prompt-input").addEventListener('input', validatePrompt);
        
        // 加载保存的API密钥
        setTimeout(() => {
            loadApiKey();
        }, 500);
        
    } catch (error) {
        console.error("初始化失败:", error);
        showResult("初始化失败: " + error.message);
    }
}

// API密钥管理函数
async function saveApiKey() {
    console.log("开始保存API密钥...");
    const apiKeyInput = document.getElementById("api-key-input");
    const apiKey = apiKeyInput.value.trim();
    
    if (!apiKey) {
        showResult("请输入API密钥");
        return;
    }
    
    if (!isValidApiKeyFormat(apiKey)) {
        showResult("API密钥格式不正确，应以sk-开头且长度至少20位");
        return;
    }
    
    try {
        showResult("正在保存API密钥...");
        
        Office.context.document.settings.set(STORAGE_KEY, apiKey);
        
        Office.context.document.settings.saveAsync((result) => {
            console.log("保存结果:", result);
            if (result.status === Office.AsyncResultStatus.Succeeded) {
                updateKeyStatus(true);
                showResult("✅ API密钥保存成功");
                updateButtonStates(true);
                
                apiKeyInput.type = "password";
                apiKeyInput.value = maskApiKey(apiKey);
                document.getElementById("toggle-key-visibility").textContent = "🔓️";
                
            } else {
                showResult("❌ 保存失败，请重试");
                console.error("保存失败:", result.error);
            }
        });
    } catch (error) {
        console.error("保存API密钥失败:", error);
        showResult("保存失败: " + error.message);
    }
}

async function loadApiKey() {
    console.log("开始加载API密钥...");
    
    try {
        if (!Office.context || !Office.context.document || !Office.context.document.settings) {
            console.log("Office环境未就绪，稍后重试");
            setTimeout(loadApiKey, 1000);
            return;
        }
        
        const apiKey = Office.context.document.settings.get(STORAGE_KEY);
        console.log("从存储中获取的密钥:", apiKey ? "存在" : "不存在");
        
        if (apiKey && isValidApiKeyFormat(apiKey)) {
            document.getElementById("api-key-input").value = maskApiKey(apiKey);
            updateKeyStatus(true);
            updateButtonStates(true);
            console.log("API密钥加载成功，启用功能按钮");
        } else {
            updateKeyStatus(false);
            updateButtonStates(false);
            console.log("未找到有效API密钥，禁用功能按钮");
        }
    } catch (error) {
        console.error("加载API密钥失败:", error);
        updateKeyStatus(false);
        updateButtonStates(false);
    }
}

async function clearApiKey() {
    console.log("清除API密钥...");
    
    try {
        Office.context.document.settings.remove(STORAGE_KEY);
        
        Office.context.document.settings.saveAsync((result) => {
            if (result.status === Office.AsyncResultStatus.Succeeded) {
                document.getElementById("api-key-input").value = "";
                updateKeyStatus(false);
                updateButtonStates(false);
                showResult("🗑️ API密钥已清除");
                console.log("API密钥清除成功");
            } else {
                showResult("清除失败，请重试");
            }
        });
    } catch (error) {
        console.error("清除API密钥失败:", error);
        showResult("清除失败: " + error.message);
    }
}

function toggleKeyVisibility() {
    const apiKeyInput = document.getElementById("api-key-input");
    const toggleBtn = document.getElementById("toggle-key-visibility");
    
    if (apiKeyInput.type === "password") {
        // 显示密钥
        apiKeyInput.type = "text";
        toggleBtn.textContent = "🔒";  // 锁定图标
        
        // 如果是掩码显示，重新加载真实密钥
        const currentValue = apiKeyInput.value;
        if (currentValue.includes('*')) {
            const realKey = Office.context.document.settings.get(STORAGE_KEY);
            if (realKey) {
                apiKeyInput.value = realKey;
            }
        }
    } else {
        // 隐藏密钥
        apiKeyInput.type = "password";
        toggleBtn.textContent = "🔓";  // 解锁图标
        
        // 如果是真实密钥，转换为掩码显示
        const currentValue = apiKeyInput.value;
        if (currentValue && !currentValue.includes('*')) {
            apiKeyInput.value = maskApiKey(currentValue);
        }
    }
}

// 工具函数
function maskApiKey(apiKey) {
    if (!apiKey || apiKey.length < 8) return apiKey;
    const visibleStart = apiKey.substring(0, 6);
    const visibleEnd = apiKey.substring(apiKey.length - 4);
    return `${visibleStart}${'*'.repeat(Math.max(apiKey.length - 10, 4))}${visibleEnd}`;
}

function isValidApiKeyFormat(apiKey) {
    return apiKey && apiKey.length >= 20 && apiKey.startsWith('sk-');
}

function validateApiKeyFormat() {
    const apiKey = document.getElementById("api-key-input").value.trim();
    const saveBtn = document.getElementById("save-key-btn");
    
    if (apiKey && isValidApiKeyFormat(apiKey)) {
        saveBtn.disabled = false;
        saveBtn.style.backgroundColor = "#0078d4";
    } else {
        saveBtn.disabled = true;
        saveBtn.style.backgroundColor = "#ccc";
    }
}

function updateKeyStatus(isValid) {
    const statusText = document.getElementById("key-status-text");
    
    if (isValid) {
        statusText.textContent = "✅ 密钥已设置";
        statusText.className = "key-status status-valid";
    } else {
        statusText.textContent = "❌ 密钥未设置";
        statusText.className = "key-status status-not-set";
    }
}

function updateButtonStates(hasValidKey) {
    console.log("更新按钮状态，hasValidKey:", hasValidKey);
    
    const buttons = [
        "generate-btn",
        "rewrite-btn", 
        "summarize-btn",
        "translate-btn",
        "translate-full-btn",
        "preview-format-btn"
    ];
    
    buttons.forEach(btnId => {
        const btn = document.getElementById(btnId);
        if (btn) {
            // 如果按钮正在处理中，不改变其状态
            if (btn.style.cursor === 'wait') {
                return;
            }
            
            btn.disabled = !hasValidKey;
            if (hasValidKey) {
                if (btnId === "preview-format-btn") {
                    btn.style.backgroundColor = "#28a745";
                } else {
                    btn.style.backgroundColor = "#0078d4";
                }
                btn.style.cursor = "pointer";
            } else {
                btn.style.backgroundColor = "#ccc";
                btn.style.cursor = "not-allowed";
            }
        }
    });
    
    validatePrompt();
}

function validatePrompt() {
    const prompt = document.getElementById("prompt-input").value.trim();
    const generateBtn = document.getElementById("generate-btn");
    
    let hasValidKey = false;
    try {
        const apiKey = Office.context.document.settings.get(STORAGE_KEY);
        hasValidKey = apiKey && isValidApiKeyFormat(apiKey);
    } catch (error) {
        hasValidKey = false;
    }
    
    if (generateBtn) {
        // 如果按钮正在处理中，不改变其状态
        if (generateBtn.style.cursor === 'wait') {
            return;
        }
        
        generateBtn.disabled = !prompt || !hasValidKey;
        if (prompt && hasValidKey) {
            generateBtn.style.backgroundColor = "#0078d4";
            generateBtn.style.cursor = "pointer";
        } else {
            generateBtn.style.backgroundColor = "#ccc";
            generateBtn.style.cursor = "not-allowed";
        }
    }
}

function getCurrentApiKey() {
    try {
        const apiKey = Office.context.document.settings.get(STORAGE_KEY);
        if (apiKey && isValidApiKeyFormat(apiKey)) {
            return apiKey;
        }
        return null;
    } catch (error) {
        console.error("获取API密钥失败:", error);
        return null;
    }
}

// 安全获取属性函数
function safeGet(obj, prop, defaultValue = null) {
    try {
        const value = obj[prop];
        return value !== null && value !== undefined ? value : defaultValue;
    } catch (error) {
        return defaultValue;
    }
}

// 核心功能函数 - 使用统一文本替换管理器
async function generateContent() {
    console.log("生成内容功能被调用");
    
    if (!getCurrentApiKey()) {
        showResult("❌ 请先设置有效的API密钥");
        return;
    }
    
    const prompt = document.getElementById("prompt-input").value;
    if (!prompt) {
        showResult("请输入您的需求");
        return;
    }
    
    showResult("正在生成内容，请稍候...");
    
    try {
        const result = await callDeepSeekAPI(prompt);
        
        // 使用精确插入方法
        await ExactPositionFormatManager.insertContent(result);
        
        showResult("✅ 内容生成完成，格式已保留");
    } catch (error) {
        showResult("❌ 生成失败: " + error.message);
    }
}

async function rewriteSelection() {
    console.log("重写选中文本功能被调用");
    
    if (!getCurrentApiKey()) {
        showResult("❌ 请先设置有效的API密钥");
        return;
    }
    
    try {
        // 使用统一替换管理器获取格式信息（包含嵌入对象检测）
        const originalFormat = await UnifiedTextReplacer.getSelectionFormatInfo();
        
        if (!originalFormat.originalText.trim()) {
            showResult("请先选择要重写的文本");
            return;
        }
        
        showResult("正在重写文本，保留字符位置格式和嵌入对象...");
        const prompt = `请重写以下文本，保持原意但改进表达：`;
        const result = await callDeepSeekAPI(prompt, originalFormat.originalText);
        
        // 使用统一替换方法 - 自动处理嵌入对象保留
        await UnifiedTextReplacer.unifiedReplace(result, originalFormat);
        
        if (originalFormat.hasEmbeddedObjects) {
            showResult("✅ 文本重写完成，字符位置格式和嵌入对象已保留");
        } else {
            showResult("✅ 文本重写完成，字符位置格式已保留");
        }
    } catch (error) {
        showResult("❌ 重写失败: " + error.message);
    }
}

async function summarizeDocument() {
    console.log("总结文档功能被调用");
    
    if (!getCurrentApiKey()) {
        showResult("❌ 请先设置有效的API密钥");
        return;
    }
    
    try {
        let documentText = "";
        
        await Word.run(async (context) => {
            const body = context.document.body;
            body.load("text");
            await context.sync();
            
            documentText = body.text;
        });
        
        if (!documentText.trim()) {
            showResult("文档内容为空");
            return;
        }
        
        showResult("正在生成摘要...");
        const prompt = "请为以下文档生成简洁的摘要：";
        const result = await callDeepSeekAPI(prompt, documentText.substring(0, 4000));
        
        // 在文档末尾插入摘要
        await Word.run(async (context) => {
            const body = context.document.body;
            const summaryParagraph = body.insertParagraph(result, Word.InsertLocation.end);
            
            // 设置摘要格式
            summaryParagraph.font.color = "blue";
            summaryParagraph.font.bold = true;
            summaryParagraph.alignment = Word.Alignment.left;
            
            await context.sync();
        });
        
        showResult(`📝 文档摘要已添加到文档末尾`);
    } catch (error) {
        showResult("❌ 摘要生成失败: " + error.message);
    }
}

async function translateSelection() {
    console.log("翻译选中文本功能被调用");
    
    if (!getCurrentApiKey()) {
        showResult("❌ 请先设置有效的API密钥");
        return;
    }
    
    try {
        // 使用统一替换管理器获取格式信息（包含嵌入对象检测）
        const originalFormat = await UnifiedTextReplacer.getSelectionFormatInfo();
        
        if (!originalFormat.originalText.trim()) {
            showResult("请先选择要翻译的文本");
            return;
        }
        
        const targetLanguage = document.getElementById("target-language").value;
        const languageName = TRANSLATION_LANGUAGES[targetLanguage] || targetLanguage;
        
        showResult(`正在翻译到${languageName}，保留字符位置格式和嵌入对象...`);
        const prompt = `请将以下文本翻译成${languageName}：`;
        const result = await callDeepSeekAPI(prompt, originalFormat.originalText);
        
        // 使用统一替换方法 - 自动处理嵌入对象保留
        await UnifiedTextReplacer.unifiedReplace(result, originalFormat);
        
        if (originalFormat.hasEmbeddedObjects) {
            showResult(`✅ 翻译完成到${languageName}，字符位置格式和嵌入对象已保留`);
        } else {
            showResult(`✅ 翻译完成到${languageName}，字符位置格式已保留`);
        }
    } catch (error) {
        showResult("❌ 翻译失败: " + error.message);
    }
}

// 翻译全文功能 - 优化按钮闪烁
async function translateFullDocument() {
    console.log("翻译全文功能被调用");

    if (!getCurrentApiKey()) {
        showResult("❌ 请先设置有效的API密钥");
        return;
    }

    const translateBtn = document.getElementById("translate-full-btn");

    try {
        const targetLanguage = document.getElementById("target-language").value;
        const languageName = TRANSLATION_LANGUAGES[targetLanguage] || targetLanguage;

        showResult(`开始翻译全文到${languageName}...`);

        // 设置按钮为处理状态（无闪烁）
        ButtonStateManager.setProcessingState(translateBtn, "翻译全文");

        // 获取文档每段的格式信息
        const paragraphFormats = await ExactPositionFormatManager.getDocumentParagraphFormats();

        // 合并所有段落
        const documentText = paragraphFormats.map(p => p.originalText || "").join("\n");
        if (!documentText.trim()) {
            showResult("文档内容为空");
            ButtonStateManager.restoreOriginalState(translateBtn);
            return;
        }

        showResult(`文档内容获取成功，开始翻译...`);

        // 创建进度显示
        const progressInfo = createProgressInfo();
        progressInfo.update(1, 3);

        // 调用 API 获取翻译。要求返回与原文段落数对齐（尽量）或可用换行表示段落分割
        const prompt = `请将以下文档内容完整地翻译成${languageName}，保持段落数量和段落顺序不变，请不要引入额外换行：`;
        const translatedContent = await callDeepSeekAPI(prompt, documentText);

        progressInfo.update(2, 3);

        // 保留空段的分割
        const translatedParagraphs = translatedContent.split(/\r?\n/);
        console.log("translateFullDocument: translatedParagraphs length:", translatedParagraphs.length);

        // 逐段替换：使用 Word.run 批次化替换
        await Word.run(async (context) => {
            const body = context.document.body;
            const paragraphs = body.paragraphs;
            paragraphs.load("items");
            await context.sync();

            const docParaCount = paragraphs.items.length;
            const minCount = Math.min(docParaCount, translatedParagraphs.length);

            // 逐段替换对应段落
            for (let i = 0; i < minCount; i++) {
                const para = paragraphs.items[i];
                const range = para.getRange();
                const newText = translatedParagraphs[i] || "";
                
                // 使用预先获取的段落格式信息，确保格式信息完整
                const originalFormat = paragraphFormats[i] || { 
                    originalText: "", 
                    characterFormats: [], 
                    paragraphFormat: {
                        alignment: Word.Alignment.left,
                        lineSpacing: 1.15,
                        firstLineIndent: 0
                    },
                    paragraphFont: {
                        name: "Calibri",
                        size: 11,
                        color: "black",
                        bold: false,
                        italic: false,
                        underline: false,
                        highlightColor: ""
                    }
                };

                // 确保段落格式不为null
                if (!originalFormat.paragraphFormat) {
                    originalFormat.paragraphFormat = {
                        alignment: Word.Alignment.left,
                        lineSpacing: 1.15,
                        firstLineIndent: 0
                    };
                }

                // 确保字符格式数组存在
                if (!originalFormat.characterFormats || originalFormat.characterFormats.length === 0) {
                    // 如果没有逐字符格式，使用段落整体字体格式
                    originalFormat.characterFormats = [];
                    for (let j = 0; j < (originalFormat.originalText?.length || newText.length); j++) {
                        originalFormat.characterFormats.push({
                            position: j,
                            text: (originalFormat.originalText?.[j] || newText[j]) || "",
                            font: originalFormat.paragraphFont || {
                                name: "Calibri",
                                size: 11,
                                color: "black",
                                bold: false,
                                italic: false,
                                underline: false,
                                highlightColor: ""
                            }
                        });
                    }
                }

                // 详细的调试日志 - 增强版，特别关注居中和粗体
                console.log(`处理第${i}段:`, {
                    原文长度: originalFormat.originalText?.length || 0,
                    新文本长度: newText.length,
                    对齐方式: originalFormat.paragraphFormat?.alignment || 'left',
                    对齐方式值: originalFormat.paragraphFormat?.alignment,
                    是否居中: originalFormat.paragraphFormat?.alignment === Word.Alignment.centered ? '是' : '否',
                    Word_Alignment_centered值: Word.Alignment.centered,
                    段落字体粗体: originalFormat.paragraphFont?.bold || false,
                    段落字体斜体: originalFormat.paragraphFont?.italic || false,
                    字符格式数量: originalFormat.characterFormats?.length || 0,
                    第一个字符粗体: originalFormat.characterFormats?.[0]?.font?.bold || false,
                    第一个字符斜体: originalFormat.characterFormats?.[0]?.font?.italic || false,
                    段落格式对象: originalFormat.paragraphFormat,
                    段落字体对象: originalFormat.paragraphFont
                });

                // 确保格式信息完整
                if (!originalFormat.paragraphFormat) {
                    console.warn(`第${i}段: 段落格式为空，使用默认值`);
                    originalFormat.paragraphFormat = {
                        alignment: Word.Alignment.left,
                        lineSpacing: 1.15,
                        firstLineIndent: 0
                    };
                }
                
                // 特别处理居中对齐
                if (originalFormat.paragraphFormat.alignment === Word.Alignment.centered) {
                    console.log(`第${i}段: 检测到居中对齐，对齐方式值 = ${originalFormat.paragraphFormat.alignment}`);
                }
                
                // 特别处理粗体
                if (originalFormat.paragraphFont?.bold === true || 
                    (originalFormat.characterFormats?.[0]?.font?.bold === true)) {
                    console.log(`第${i}段: 检测到粗体格式`);
                }

                await ExactPositionFormatManager.exactPositionReplaceOnRange(context, range, newText, originalFormat);
            }

            // 翻译段落更多：插入到文档末尾
            if (translatedParagraphs.length > docParaCount) {
                const lastSourceFormat = paragraphFormats[paragraphFormats.length - 1] || null;
                for (let j = docParaCount; j < translatedParagraphs.length; j++) {
                    const inserted = body.insertParagraph(translatedParagraphs[j], Word.InsertLocation.end);
                    if (lastSourceFormat) {
                        // 确保格式信息完整
                        const formatToApply = {
                            paragraphFormat: lastSourceFormat.paragraphFormat || {
                                alignment: Word.Alignment.left,
                                lineSpacing: 1.15,
                                firstLineIndent: 0
                            },
                            characterFormats: lastSourceFormat.characterFormats || []
                        };
                        
                        ExactPositionFormatManager.applyParagraphFormat(inserted.paragraphFormat, formatToApply.paragraphFormat);
                        
                        // 应用字体格式
                        if (lastSourceFormat.characterFormats && lastSourceFormat.characterFormats.length > 0) {
                            const firstCharFormat = lastSourceFormat.characterFormats[0];
                            if (firstCharFormat && firstCharFormat.font) {
                                ExactPositionFormatManager.applyFontFormat(inserted.font, firstCharFormat.font);
                            }
                        } else if (lastSourceFormat.paragraphFont) {
                            ExactPositionFormatManager.applyFontFormat(inserted.font, lastSourceFormat.paragraphFont);
                        }
                    }
                }
            } else if (translatedParagraphs.length < docParaCount) {
                // 翻译段落更少：清空多余原段落
                for (let k = translatedParagraphs.length; k < docParaCount; k++) {
                    const para = paragraphs.items[k];
                    const range = para.getRange();
                    range.insertText("", Word.InsertLocation.replace);
                }
            }

            await context.sync();
        });

        progressInfo.update(3, 3);

        showResult(`✅ 全文翻译完成！已成功将文档翻译为${languageName} 并尽量保留原格式`);
    } catch (error) {
        console.error("翻译全文失败:", error);
        showResult("❌ 翻译全文失败: " + error.message);
    } finally {
        // 恢复按钮状态
        ButtonStateManager.restoreOriginalState(translateBtn);
    }
}

// 创建进度信息
function createProgressInfo() {
    const progressContainer = document.createElement('div');
    progressContainer.className = 'progress-container';
    
    const progressBar = document.createElement('div');
    progressBar.className = 'progress-bar';
    
    const progressText = document.createElement('div');
    progressText.className = 'progress-text';
    
    progressContainer.appendChild(progressBar);
    
    const resultElement = document.getElementById("result");
    resultElement.innerHTML = '';
    resultElement.appendChild(progressContainer);
    resultElement.appendChild(progressText);
    
    return {
        update: (current, total) => {
            const percentage = Math.round((current / total) * 100);
            progressBar.style.width = `${percentage}%`;
            progressText.textContent = `翻译进度: ${current}/${total} (${percentage}%)`;
        },
        complete: () => {
            progressBar.style.width = '100%';
            progressText.textContent = '翻译完成！';
        }
    };
}

// API调用函数 - 优化按钮状态管理
async function callDeepSeekAPI(prompt, context = "") {
    const apiKey = getCurrentApiKey();
    
    if (!apiKey) {
        throw new Error("API密钥未设置，请先设置有效的API密钥");
    }
    
    try {
        // 设置所有功能按钮为处理状态
        const functionButtons = [
            "generate-btn",
            "rewrite-btn", 
            "summarize-btn",
            "translate-btn", 
            "translate-full-btn",
            "preview-format-btn"
        ];
        ButtonStateManager.setBatchProcessing(functionButtons);
        
        const response = await fetch(DEEPSEEK_API_URL, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${apiKey}`
            },
            body: JSON.stringify({
                model: "deepseek-chat",
                messages: [
                    {
                        role: "system",
                        content: "你是一个专业的写作助手，帮助用户在Word文档中进行写作和编辑。回复要简洁专业。"
                    },
                    {
                        role: "user", 
                        content: `${prompt} ${context ? `\n上下文：${context}` : ''}`
                    }
                ],
                max_tokens: 4000,
                temperature: 0.7
            })
        });
        
        if (!response.ok) {
            const errorData = await response.json().catch(() => ({}));
            throw new Error(errorData.error?.message || `API请求失败: ${response.status}`);
        }
        
        const data = await response.json();
        return data.choices[0].message.content;
    } catch (error) {
        console.error("API调用失败:", error);
        throw error;
    } finally {
        // 恢复所有功能按钮状态
        const functionButtons = [
            "generate-btn",
            "rewrite-btn", 
            "summarize-btn",
            "translate-btn",
            "translate-full-btn",
            "preview-format-btn"
        ];
        ButtonStateManager.restoreBatchState(functionButtons);
        
        // 重新验证按钮状态
        const hasValidKey = !!getCurrentApiKey();
        updateButtonStates(hasValidKey);
        validatePrompt();
    }
}

// 格式预览功能
async function previewFormatting() {
    try {
        const formattingInfo = await checkCurrentFormatting();
        const previewElement = document.getElementById("format-preview");
        previewElement.innerHTML = `
            <h4>📋 当前选中文本的格式信息：</h4>
            <pre>${formattingInfo}</pre>
        `;
        previewElement.style.display = 'block';
    } catch (error) {
        showResult("无法获取格式信息: " + error.message);
    }
}

async function checkCurrentFormatting() {
    let formattingInfo = "";
    
    await Word.run(async (context) => {
        const selection = context.document.getSelection();
        
        selection.load([
            "text",
            "font/name", "font/size", "font/color", "font/bold", "font/italic", 
            "font/underline",
            "paragraphFormat/alignment", "paragraphFormat/lineSpacing", 
            "paragraphFormat/firstLineIndent"
        ]);
        await context.sync();
        
        formattingInfo = `
文本内容: "${selection.text.substring(0, 100)}${selection.text.length > 100 ? '...' : ''}"
文本长度: ${selection.text.length} 字符

字体信息:
  字体名称: ${safeGet(selection.font, 'name', "默认")}
  字体大小: ${safeGet(selection.font, 'size', "默认")} pt
  字体颜色: ${safeGet(selection.font, 'color', "默认")}

字体样式:
  粗体: ${safeGet(selection.font, 'bold', false) ? "是" : "否"}
  斜体: ${safeGet(selection.font, 'italic', false) ? "是" : "否"}
  下划线: ${safeGet(selection.font, 'underline', false) ? "是" : "否"}

段落格式:
  对齐方式: ${getAlignmentText(safeGet(selection.paragraphFormat, 'alignment', Word.Alignment.left))}
  行间距: ${safeGet(selection.paragraphFormat, 'lineSpacing', "默认")}
  首行缩进: ${safeGet(selection.paragraphFormat, 'firstLineIndent', "无")}
        `.trim();
    });
    
    return formattingInfo;
}

function getAlignmentText(alignment) {
    const alignments = {
        [Word.Alignment.centered]: "居中",
        [Word.Alignment.left]: "左对齐", 
        [Word.Alignment.right]: "右对齐",
        [Word.Alignment.justified]: "两端对齐",
        [Word.Alignment.distributed]: "分散对齐",
        [Word.Alignment.undefined]: "未定义"
    };
    return alignments[alignment] || "未知";
}

// 工具函数
function showResult(message) {
    const resultElement = document.getElementById("result");
    if (resultElement) {
        resultElement.textContent = message;
        
        if (message.includes("✅") || message.includes("成功") || message.includes("完成")) {
            resultElement.style.color = "#28a745";
            resultElement.className = "result-area success-message";
        } else if (message.includes("❌") || message.includes("失败") || message.includes("错误")) {
            resultElement.style.color = "#dc3545";
            resultElement.className = "result-area error-message";
        } else {
            resultElement.style.color = "#333";
            resultElement.className = "result-area";
        }
    }
}

function showLoadingState(isLoading) {
    // 不再使用全局loading状态，由ButtonStateManager单独控制
    if (isLoading) {
        document.body.classList.add('loading');
    } else {
        document.body.classList.remove('loading');
    }
}

// 调试信息
console.log("DeepSeek for Word插件脚本加载完成");
