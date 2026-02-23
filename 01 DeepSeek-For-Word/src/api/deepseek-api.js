 // DeepSeek API 模块
// 将API相关功能分离到独立文件中

// API配置常量
export const STORAGE_KEY = 'deepseek-api-key';
export const DEEPSEEK_API_URL = 'https://api.deepseek.com/v1/chat/completions';

// 安全获取属性函数
export function safeGet(obj, prop, defaultValue = null) {
    try {
        const value = obj[prop];
        return value !== null && value !== undefined ? value : defaultValue;
    } catch (error) {
        return defaultValue;
    }
}

// API密钥管理函数
export function maskApiKey(apiKey) {
    if (!apiKey || apiKey.length < 8) return apiKey;
    const visibleStart = apiKey.substring(0, 6);
    const visibleEnd = apiKey.substring(apiKey.length - 4);
    return `${visibleStart}${'*'.repeat(Math.max(apiKey.length - 10, 4))}${visibleEnd}`;
}

export function isValidApiKeyFormat(apiKey) {
    return apiKey && apiKey.length >= 20 && apiKey.startsWith('sk-');
}

export function validateApiKeyFormat() {
    const apiKeyInput = document.getElementById("api-key-input");
    if (!apiKeyInput) return;
    
    const apiKey = apiKeyInput.value.trim();
    const saveBtn = document.getElementById("save-key-btn");
    
    if (saveBtn) {
        if (apiKey && isValidApiKeyFormat(apiKey)) {
            saveBtn.disabled = false;
            saveBtn.style.backgroundColor = "#0078d4";
        } else {
            saveBtn.disabled = true;
            saveBtn.style.backgroundColor = "#ccc";
        }
    }
}

export function updateKeyStatus(isValid) {
    const statusText = document.getElementById("key-status-text");
    
    if (statusText) {
        if (isValid) {
            statusText.textContent = "✅ 密钥已设置";
            statusText.className = "key-status status-valid";
        } else {
            statusText.textContent = "❌ 密钥未设置";
            statusText.className = "key-status status-not-set";
        }
    }
}

export function updateButtonStates(hasValidKey) {
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
}

export function validatePrompt() {
    const prompt = document.getElementById("prompt-input")?.value.trim();
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

export function getCurrentApiKey() {
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

// API调用函数
export async function callDeepSeekAPI(prompt, context = "") {
    const apiKey = getCurrentApiKey();
    
    if (!apiKey) {
        throw new Error("API密钥未设置，请先设置有效的API密钥");
    }
    
    try {
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
                            content: `你是一个专业的写作助手，帮助用户在Word文档中进行写作和编辑。当用户要求"图文并茂"时，请按照以下格式生成内容：

1. 生成完整的论文结构，包括标题、摘要、目录、正文、结论、参考文献等
2. 在需要插入图片的位置，使用以下格式的占位符：
   [图片: 描述性标题]
   例如：[图片: 视觉引导流程图]
3. 在占位符下方，提供详细的图片描述，说明图片应该展示什么内容
4. 回复要专业、结构清晰，适合学术论文

如果用户没有明确要求图文并茂，则按常规方式回复。`
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
    }
}

// API密钥管理函数（需要Office上下文）
export async function saveApiKey() {
    console.log("开始保存API密钥...");
    const apiKeyInput = document.getElementById("api-key-input");
    if (!apiKeyInput) {
        return { success: false, message: "API密钥输入框未找到" };
    }
    
    const apiKey = apiKeyInput.value.trim();
    
    if (!apiKey) {
        return { success: false, message: "请输入API密钥" };
    }
    
    if (!isValidApiKeyFormat(apiKey)) {
        return { success: false, message: "API密钥格式不正确，应以sk-开头且长度至少20位" };
    }
    
    try {
        Office.context.document.settings.set(STORAGE_KEY, apiKey);
        
        return new Promise((resolve) => {
            Office.context.document.settings.saveAsync((result) => {
                console.log("保存结果:", result);
                if (result.status === Office.AsyncResultStatus.Succeeded) {
                    resolve({ 
                        success: true, 
                        message: "✅ API密钥保存成功",
                        apiKey: apiKey
                    });
                } else {
                    resolve({ 
                        success: false, 
                        message: "❌ 保存失败，请重试",
                        error: result.error
                    });
                }
            });
        });
    } catch (error) {
        console.error("保存API密钥失败:", error);
        return { 
            success: false, 
            message: "保存失败: " + error.message,
            error: error
        };
    }
}

export async function loadApiKey() {
    console.log("开始加载API密钥...");
    
    try {
        if (!Office.context || !Office.context.document || !Office.context.document.settings) {
            console.log("Office环境未就绪，稍后重试");
            return { success: false, message: "Office环境未就绪" };
        }
        
        const apiKey = Office.context.document.settings.get(STORAGE_KEY);
        console.log("从存储中获取的密钥:", apiKey ? "存在" : "不存在");
        
        if (apiKey && isValidApiKeyFormat(apiKey)) {
            return { 
                success: true, 
                apiKey: apiKey,
                maskedKey: maskApiKey(apiKey)
            };
        } else {
            return { 
                success: false, 
                message: "未找到有效API密钥"
            };
        }
    } catch (error) {
        console.error("加载API密钥失败:", error);
        return { 
            success: false, 
            message: "加载失败: " + error.message,
            error: error
        };
    }
}

export async function clearApiKey() {
    console.log("清除API密钥...");
    
    try {
        Office.context.document.settings.remove(STORAGE_KEY);
        
        return new Promise((resolve) => {
            Office.context.document.settings.saveAsync((result) => {
                if (result.status === Office.AsyncResultStatus.Succeeded) {
                    resolve({ 
                        success: true, 
                        message: "🗑️ API密钥已清除"
                    });
                } else {
                    resolve({ 
                        success: false, 
                        message: "清除失败，请重试",
                        error: result.error
                    });
                }
            });
        });
    } catch (error) {
        console.error("清除API密钥失败:", error);
        return { 
            success: false, 
            message: "清除失败: " + error.message,
            error: error
        };
    }
}

// 默认导出
export default {
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
    saveApiKey,
    loadApiKey,
    clearApiKey
};
