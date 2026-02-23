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

// 精确字符位置格式管理对象
const ExactPositionFormatManager = {
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
                
            } catch (error) {
                console.error("获取格式失败:", error);
                throw error;
            }
        });
        
        return formatInfo;
    },
    
    // 按字符逐个获取格式
    async getCharacterByCharacterFormats(context, selection) {
        const text = selection.text;
        const characterFormats = [];
        
        try {
            // 方法1: 尝试按字符获取
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
        if (!originalFormat || !originalFormat.characterFormats || originalFormat.characterFormats.length === 0) {
            // 没有逐字符格式，尽量应用段落或整体字体
            if (originalFormat && originalFormat.paragraphFormat) {
                ExactPositionFormatManager.applyParagraphFormat(range.paragraphFormat, originalFormat.paragraphFormat);
            }
            if (originalFormat && originalFormat.characterFormats && originalFormat.characterFormats[0]) {
                ExactPositionFormatManager.applyFontFormat(range.font, originalFormat.characterFormats[0].font);
            }
            return;
        }

        const originalFormats = originalFormat.characterFormats;
        const originalLength = originalFormat.originalText.length || originalFormats.length;
        const newLength = newText.length;

        // 尽量应用段落格式
        if (originalFormat.paragraphFormat) {
            try {
                ExactPositionFormatManager.applyParagraphFormat(range.paragraphFormat, originalFormat.paragraphFormat);
            } catch (err) {
                console.warn("applyCharacterWiseFormats: 应用段落格式失败", err);
            }
        }

        // 按字符逐个应用字体样式。为性能考虑，避免每次 sync；一次性设置后 await context.sync()
        try {
            for (let j = 0; j < newLength; j++) {
                const mappedIndex = originalLength === 0 ? 0 : Math.min(originalLength - 1, Math.round(j * (originalLength / newLength)));
                const srcFormat = originalFormats[mappedIndex] ? originalFormats[mappedIndex].font : (originalFormats[0] ? originalFormats[0].font : null);
                if (!srcFormat) continue;

                try {
                    // 获取新文本对应位置的范围（以 range 开始为起点的相对位置）
                    const charRange = range.getRange(j, j);
                    // 直接应用字体格式到单个字符范围
                    ExactPositionFormatManager.applyFontFormat(charRange.font, srcFormat);
                } catch (error) {
                    // 如果单个字符应用失败，则忽略继续
                    console.warn(`applyCharacterWiseFormats: 应用第 ${j} 个字符格式失败:`, error);
                }
            }
            await context.sync();
        } catch (error) {
            console.warn("applyCharacterWiseFormats: 批量应用格式失败，尝试逐个同步:", error);
            // 失败时退化为每字符同步以增强成功率（但性能差）
            for (let j = 0; j < newLength; j++) {
                const mappedIndex = originalLength === 0 ? 0 : Math.min(originalLength - 1, Math.round(j * (originalLength / newLength)));
                const srcFormat = originalFormats[mappedIndex] ? originalFormats[mappedIndex].font : (originalFormats[0] ? originalFormats[0].font : null);
                if (!srcFormat) continue;
                try {
                    const charRange = range.getRange(j, j);
                    ExactPositionFormatManager.applyFontFormat(charRange.font, srcFormat);
                    await context.sync();
                } catch (err) {
                    console.warn(`applyCharacterWiseFormats: 回退应用第 ${j} 个字符失败:`, err);
                }
            }
        }
    },

    // 新增：在给定范围上执行逐字符替换并保留原文样式的函数（用于段落/选区）
    async exactPositionReplaceOnRange(context, range, newText, originalFormat) {
        try {
            // 替换文本
            range.insertText(newText, Word.InsertLocation.replace);
            await context.sync();

            // 应用段落样式与逐字符字体样式
            if (originalFormat && originalFormat.paragraphFormat) {
                try {
                    ExactPositionFormatManager.applyParagraphFormat(range.paragraphFormat, originalFormat.paragraphFormat);
                } catch (e) {
                    console.warn("exactPositionReplaceOnRange: 应用段落样式失败:", e);
                }
            }

            await ExactPositionFormatManager.applyCharacterWiseFormats(context, range, newText, originalFormat);
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

    // 新增：获取文档中每个段落的文本和格式信息（段落与字符级）
    async getDocumentParagraphFormats() {
        const paragraphFormats = [];
        await Word.run(async (context) => {
            const body = context.document.body;
            const paragraphs = body.paragraphs;
            paragraphs.load("items");
            await context.sync();

            for (let i = 0; i < paragraphs.items.length; i++) {
                try {
                    const paragraph = paragraphs.items[i];
                    const range = paragraph.getRange();
                    range.load([
                        "text",
                        "paragraphFormat/alignment", "paragraphFormat/lineSpacing",
                        "paragraphFormat/firstLineIndent"
                    ]);
                    await context.sync();

                    // 获取该段落的逐字符格式（可能降级为单词或整体）
                    const charFormats = await ExactPositionFormatManager.getCharacterByCharacterFormats(context, range);

                    paragraphFormats.push({
                        originalText: range.text,
                        paragraphFormat: {
                            alignment: safeGet(range.paragraphFormat, 'alignment', Word.Alignment.left),
                            lineSpacing: safeGet(range.paragraphFormat, 'lineSpacing'),
                            firstLineIndent: safeGet(range.paragraphFormat, 'firstLineIndent')
                        },
                        characterFormats: charFormats
                    });
                } catch (err) {
                    console.warn("getDocumentParagraphFormats: 获取段落格式失败，继续下一个段落:", err);
                    paragraphFormats.push({
                        originalText: "",
                        paragraphFormat: null,
                        characterFormats: []
                    });
                }
            }
        });
        return paragraphFormats;
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
    
    try {
        showResult("正在保存API密钥...");
        
        const result = await apiSaveApiKey();
        
        if (result.success) {
            updateKeyStatus(true);
            showResult(result.message);
            updateButtonStates(true);
            
            const apiKeyInput = document.getElementById("api-key-input");
            apiKeyInput.type = "password";
            apiKeyInput.value = maskApiKey(result.apiKey);
            document.getElementById("toggle-key-visibility").textContent = "🔓️";
        } else {
            showResult(result.message);
        }
    } catch (error) {
        console.error("保存API密钥失败:", error);
        showResult("保存失败: " + error.message);
    }
}

async function loadApiKey() {
    console.log("开始加载API密钥...");
    
    try {
        const result = await apiLoadApiKey();
        
        if (result.success) {
            document.getElementById("api-key-input").value = result.maskedKey;
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
        const result = await apiClearApiKey();
        
        if (result.success) {
            document.getElementById("api-key-input").value = "";
            updateKeyStatus(false);
            updateButtonStates(false);
            showResult(result.message);
            console.log("API密钥清除成功");
        } else {
            showResult(result.message);
        }
    } catch (error) {
        console.error("清除API密钥失败:", error);
        showResult("清除失败: " + error.message);
    }
}

function toggleKeyVisibility() {
    const apiKeyInput = document.getElementById("api-key-input");
    const toggleBtn = document.getElementById("toggle-key-visibility");
    
    if (!apiKeyInput || !toggleBtn) {
        console.error("切换密钥显示状态失败：找不到输入框或按钮");
        return;
    }
    
    if (apiKeyInput.type === "password") {
        // 显示密钥
        apiKeyInput.type = "text";
        toggleBtn.textContent = "🔒";  // 锁定图标
        
        // 如果是掩码显示，重新加载真实密钥
        const currentValue = apiKeyInput.value;
        if (currentValue.includes('*')) {
            try {
                // 检查Office环境是否就绪
                if (Office.context && Office.context.document && Office.context.document.settings) {
                    const realKey = Office.context.document.settings.get(STORAGE_KEY);
                    if (realKey) {
                        apiKeyInput.value = realKey;
                    }
                } else {
                    console.warn("Office环境未就绪，无法获取真实密钥");
                }
            } catch (error) {
                console.error("获取真实密钥失败:", error);
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


// 核心功能函数 - 使用精确位置格式管理
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
            
            // 检查是否包含图片占位符
            const imagePlaceholders = extractImagePlaceholders(result);
            if (imagePlaceholders.length > 0) {
                showResult(`✅ 内容生成完成！检测到 ${imagePlaceholders.length} 个图片占位符。\n\n图片位置已标记，请手动插入相应图片。\n\n图片描述：\n${imagePlaceholders.map((p, i) => `${i+1}. ${p}`).join('\n')}`);
            } else {
                showResult("✅ 内容生成完成，格式已保留");
            }
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
        // 保存精确位置格式
        const originalFormat = await ExactPositionFormatManager.getExactPositionFormats();
        
        if (!originalFormat.originalText.trim()) {
            showResult("请先选择要重写的文本");
            return;
        }
        
        showResult("正在重写文本，保留字符位置格式...");
        const prompt = `请重写以下文本，保持原意但改进表达：`;
        const result = await callDeepSeekAPI(prompt, originalFormat.originalText);
        
        // 使用精确位置替换方法
        await ExactPositionFormatManager.exactPositionReplace(result, originalFormat);
        
        showResult("✅ 文本重写完成，字符位置格式已保留");
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
        // 保存精确位置格式
        const originalFormat = await ExactPositionFormatManager.getExactPositionFormats();
        
        if (!originalFormat.originalText.trim()) {
            showResult("请先选择要翻译的文本");
            return;
        }
        
        const targetLanguage = document.getElementById("target-language").value;
        const languageName = TRANSLATION_LANGUAGES[targetLanguage] || targetLanguage;
        
        showResult(`正在翻译到${languageName}，保留字符位置格式...`);
        const prompt = `请将以下文本翻译成${languageName}：`;
        const result = await callDeepSeekAPI(prompt, originalFormat.originalText);
        
        // 使用精确位置替换方法
        await ExactPositionFormatManager.exactPositionReplace(result, originalFormat);
        
        showResult(`✅ 翻译完成到${languageName}，字符位置格式已保留`);
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
                const para = paragraphs.items[i]; // <-- 修复：使用 items[i] 访问
                const range = para.getRange();
                const newText = translatedParagraphs[i] || "";
                const originalFormat = paragraphFormats[i] || { originalText: para.getRange().text, characterFormats: [], paragraphFormat: null };

                await ExactPositionFormatManager.exactPositionReplaceOnRange(context, range, newText, originalFormat);
            }

            // 翻译段落更多：插入到文档末尾
            if (translatedParagraphs.length > docParaCount) {
                const lastSourceFormat = paragraphFormats[paragraphFormats.length - 1] || null;
                for (let j = docParaCount; j < translatedParagraphs.length; j++) {
                    const inserted = body.insertParagraph(translatedParagraphs[j], Word.InsertLocation.end);
                    if (lastSourceFormat) {
                        ExactPositionFormatManager.applyParagraphFormat(inserted.paragraphFormat, lastSourceFormat.paragraphFormat);
                        ExactPositionFormatManager.applyFontFormat(inserted.font, lastSourceFormat.characterFormats?.[0]?.font || {});
                    }
                }
            } else if (translatedParagraphs.length < docParaCount) {
                // 翻译段落更少：清空多余原段落
                for (let k = translatedParagraphs.length; k < docParaCount; k++) {
                    const para = paragraphs.items[k]; // <-- 修复：使用 items[k] 访问
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

// 提取图片占位符函数
function extractImagePlaceholders(text) {
    const imagePattern = /\[图片:\s*([^\]]+)\]/g;
    const placeholders = [];
    let match;
    
    while ((match = imagePattern.exec(text)) !== null) {
        placeholders.push(match[1].trim());
    }
    
    return placeholders;
}

// 调试信息
console.log("DeepSeek for Word插件脚本加载完成");
