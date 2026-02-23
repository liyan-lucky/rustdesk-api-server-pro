Option Explicit
' DeepSeek-API.bas
' DeepSeek for Word - 完全重写版
' 功能: 将 DeepSeek AI 能力深度集成到 Microsoft Word
' 兼容: Microsoft Word VBA (Office 2010 及以上)
' 版本: 3.0

' ==================== 全局配置 ====================

' API 配置常量
' 请在此处填写您的 DeepSeek API 密钥（从 https://platform.deepseek.com/ 获取）
Private Const DEEPSEEK_API_KEY As String = "sk-de73a46bc94c433fa9fb0372b92e9c82"
Private Const DEEPSEEK_API_URL As String = "https://api.deepseek.com/v1/chat/completions"
Private Const DEFAULT_MODEL As String = "deepseek-chat"
Private Const DEFAULT_MAX_TOKENS As Long = 4096
Private Const DEFAULT_TEMPERATURE As Double = 0.7
Private Const DEFAULT_TOP_P As Double = 0.95

' 错误信息存储
Private LastDeepSeekError As String
Private ProcessingStatus As Boolean

' ==================== 核心功能命令 ====================

' 智能改写 - 增强版
Public Sub DS_EnhancedRewrite()
    Dim options As String
    options = "请选择改写风格：" & vbCrLf & _
              "1. 正式商务风格" & vbCrLf & _
              "2. 学术论文风格" & vbCrLf & _
              "3. 创意文学风格" & vbCrLf & _
              "4. 技术文档风格" & vbCrLf & _
              "5. 营销文案风格"
    
    Dim choice As String
    choice = InputBox(options, "DeepSeek 智能改写", "1")
    
    Dim style As String
    Select Case choice
        Case "1": style = "正式商务"
        Case "2": style = "学术论文"
        Case "3": style = "创意文学"
        Case "4": style = "技术文档"
        Case "5": style = "营销文案"
        Case Else: style = "通用"
    End Select
    
    DeepSeek_AdvancedProcess "rewrite:" & style
End Sub

' 多语言翻译 - 增强版
Public Sub DS_MultiTranslate()
    Dim languages As String
    languages = "请选择目标语言：" & vbCrLf & _
                "zh - 中文" & vbCrLf & _
                "en - 英语" & vbCrLf & _
                "ja - 日语" & vbCrLf & _
                "ko - 韩语" & vbCrLf & _
                "fr - 法语" & vbCrLf & _
                "de - 德语" & vbCrLf & _
                "es - 西班牙语" & vbCrLf & _
                "ru - 俄语" & vbCrLf & _
                "ar - 阿拉伯语"
    
    Dim targetLang As String
    targetLang = InputBox(languages, "DeepSeek 多语言翻译", "en")
    
    If Trim(targetLang) <> "" Then
        DeepSeek_AdvancedProcess "translate:" & targetLang
    End If
End Sub

' 专业润色 - 增强版
Public Sub DS_ProfessionalPolish()
    Dim options As String
    options = "请选择润色强度：" & vbCrLf & _
              "1. 轻度润色（仅修正语法）" & vbCrLf & _
              "2. 标准润色（语法+表达优化）" & vbCrLf & _
              "3. 深度润色（全面优化+风格提升）"
    
    Dim choice As String
    choice = InputBox(options, "DeepSeek 专业润色", "2")
    
    Dim intensity As String
    Select Case choice
        Case "1": intensity = "轻度"
        Case "2": intensity = "标准"
        Case "3": intensity = "深度"
        Case Else: intensity = "标准"
    End Select
    
    DeepSeek_AdvancedProcess "polish:" & intensity
End Sub

' 智能总结 - 增强版
Public Sub DS_SmartSummarize()
    Dim options As String
    options = "请选择总结类型：" & vbCrLf & _
              "1. 要点总结（提取关键点）" & vbCrLf & _
              "2. 段落总结（保持段落结构）" & vbCrLf & _
              "3. 详细总结（保留更多细节）" & vbCrLf & _
              "4. 执行摘要（适合商务文档）"
    
    Dim choice As String
    choice = InputBox(options, "DeepSeek 智能总结", "1")
    
    Dim summaryType As String
    Select Case choice
        Case "1": summaryType = "要点"
        Case "2": summaryType = "段落"
        Case "3": summaryType = "详细"
        Case "4": summaryType = "执行摘要"
        Case Else: summaryType = "要点"
    End Select
    
    DeepSeek_AdvancedProcess "summarize:" & summaryType
End Sub

' 内容扩写 - 增强版
Public Sub DS_ContentExpand()
    Dim options As String
    options = "请选择扩写方向：" & vbCrLf & _
              "1. 补充细节" & vbCrLf & _
              "2. 添加示例" & vbCrLf & _
              "3. 深入分析" & vbCrLf & _
              "4. 背景说明" & vbCrLf & _
              "5. 多角度阐述"
    
    Dim choice As String
    choice = InputBox(options, "DeepSeek 内容扩写", "1")
    
    Dim expandType As String
    Select Case choice
        Case "1": expandType = "细节"
        Case "2": expandType = "示例"
        Case "3": expandType = "分析"
        Case "4": expandType = "背景"
        Case "5": expandType = "多角度"
        Case Else: expandType = "细节"
    End Select
    
    DeepSeek_AdvancedProcess "expand:" & expandType
End Sub

' 语法检查 - 新增功能
Public Sub DS_GrammarCheck()
    DeepSeek_AdvancedProcess "grammar"
End Sub

' 风格转换 - 新增功能
Public Sub DS_StyleConvert()
    Dim options As String
    options = "请选择目标风格：" & vbCrLf & _
              "1. 正式 → 口语" & vbCrLf & _
              "2. 口语 → 正式" & vbCrLf & _
              "3. 技术 → 通俗" & vbCrLf & _
              "4. 通俗 → 专业" & vbCrLf & _
              "5. 中文 → 英文风格"
    
    Dim choice As String
    choice = InputBox(options, "DeepSeek 风格转换", "1")
    
    Dim conversion As String
    Select Case choice
        Case "1": conversion = "正式转口语"
        Case "2": conversion = "口语转正式"
        Case "3": conversion = "技术转通俗"
        Case "4": conversion = "通俗转专业"
        Case "5": conversion = "中文转英文风格"
        Case Else: conversion = "正式转口语"
    End Select
    
    DeepSeek_AdvancedProcess "style:" & conversion
End Sub

' 代码解释 - 新增功能（针对技术文档）
Public Sub DS_CodeExplain()
    DeepSeek_AdvancedProcess "code_explain"
End Sub

' 自定义处理 - 增强版
Public Sub DS_AdvancedCustom()
    Dim presetPrompts As String
    presetPrompts = "预设指令：" & vbCrLf & _
                    "1. 转换为诗歌" & vbCrLf & _
                    "2. 制作演讲稿" & vbCrLf & _
                    "3. 写邮件草稿" & vbCrLf & _
                    "4. 生成产品描述" & vbCrLf & _
                    "5. 创建社交媒体文案" & vbCrLf & _
                    "6. 自定义指令"
    
    Dim choice As String
    choice = InputBox(presetPrompts, "DeepSeek 高级自定义", "6")
    
    Dim customPrompt As String
    Select Case choice
        Case "1": customPrompt = "将文本转换为优美的诗歌形式"
        Case "2": customPrompt = "将文本转换为适合演讲的格式"
        Case "3": customPrompt = "将文本转换为专业的邮件格式"
        Case "4": customPrompt = "将文本转换为吸引人的产品描述"
        Case "5": customPrompt = "将文本转换为适合社交媒体的文案"
        Case "6": 
            customPrompt = InputBox("请输入自定义处理指令：", "DeepSeek 自定义处理")
        Case Else: customPrompt = "根据文本内容进行智能处理"
    End Select
    
    If Trim(customPrompt) <> "" Then
        DeepSeek_AdvancedProcess "custom:" & customPrompt
    End If
End Sub

' ==================== 批量处理命令 ====================

' 批量处理文档 - 新增功能
Public Sub DS_BatchProcess()
    If ActiveDocument Is Nothing Then
        MsgBox "请先打开一个文档。", vbExclamation, "DeepSeek"
        Exit Sub
    End If
    
    Dim options As String
    options = "请选择批量处理模式：" & vbCrLf & _
              "1. 处理所有段落" & vbCrLf & _
              "2. 处理所有标题" & vbCrLf & _
              "3. 处理选中的多个段落" & vbCrLf & _
              "4. 处理整个文档"
    
    Dim choice As String
    choice = InputBox(options, "DeepSeek 批量处理", "1")
    
    Dim processType As String
    Select Case choice
        Case "1": processType = "paragraphs"
        Case "2": processType = "headings"
        Case "3": processType = "selected"
        Case "4": processType = "document"
        Case Else: Exit Sub
    End Select
    
    BatchProcessDocument processType
End Sub

' ==================== 配置管理命令 ====================

' 智能配置向导 - 增强版
Public Sub DS_SmartConfigure()
    SmartConfigurationWizard
End Sub

' API 密钥管理 - 增强版
Public Sub DS_KeyManager()
    APIKeyManager
End Sub

' 模型设置 - 新增功能
Public Sub DS_ModelSettings()
    ModelConfiguration
End Sub

' 使用统计 - 新增功能
Public Sub DS_UsageStats()
    ShowUsageStatistics
End Sub

' 帮助系统 - 增强版
Public Sub DS_ComprehensiveHelp()
    ShowComprehensiveHelp
End Sub

' ==================== 核心处理引擎 ====================

' 高级处理函数 - 完全重写
Private Sub DeepSeek_AdvancedProcess(mode As String)
    On Error GoTo ErrHandler
    
    ' 检查处理状态
    If ProcessingStatus Then
        MsgBox "已有处理任务正在进行，请稍后再试。", vbExclamation, "DeepSeek"
        Exit Sub
    End If
    
    ProcessingStatus = True
    
    ' 验证选区
    If Selection Is Nothing Then
        MsgBox "请先选中要处理的文本。", vbExclamation, "DeepSeek"
        ProcessingStatus = False
        Exit Sub
    End If
    
    Dim selRange As Range
    Set selRange = Selection.Range
    
    If Trim(selRange.Text) = "" Then
        MsgBox "选区为空，请选择要处理的文本内容。", vbExclamation, "DeepSeek"
        ProcessingStatus = False
        Exit Sub
    End If
    
    ' 显示进度指示器
    Application.StatusBar = "DeepSeek AI 正在处理中，请稍候..."
    Application.ScreenUpdating = False
    
    ' 保存完整的格式信息
    Dim formatInfo As FormatPreservation
    Set formatInfo = New FormatPreservation
    formatInfo.SaveFormatting selRange
    
    ' 构建智能提示词
    Dim systemPrompt As String
    Dim userPrompt As String
    
    systemPrompt = BuildSmartPrompt(mode, selRange.Text)
    userPrompt = selRange.Text
    
    ' 调用增强版 API
    Dim resultText As String
    resultText = CallEnhancedDeepSeekAPI(systemPrompt, userPrompt)
    
    ' 清除状态
    Application.StatusBar = ""
    Application.ScreenUpdating = True
    ProcessingStatus = False
    
    ' 处理结果
    If resultText = "" Then
        ShowEnhancedError LastDeepSeekError
        Exit Sub
    End If
    
    ' 智能格式恢复
    formatInfo.RestoreFormatting selRange, resultText
    
    ' 显示成功消息
    ShowSuccessMessage mode, Len(selRange.Text), Len(resultText)
    
    Exit Sub

ErrHandler:
    Application.StatusBar = ""
    Application.ScreenUpdating = True
    ProcessingStatus = False
    MsgBox "处理过程中发生错误：" & Err.Number & " - " & Err.Description, vbCritical, "DeepSeek"
End Sub

' ==================== 智能提示词构建 ====================

' 构建智能提示词
Private Function BuildSmartPrompt(mode As String, text As String) As String
    Dim basePrompt As String
    Dim additionalInstructions As String
    
    ' 分析文本特征
    Dim textAnalysis As String
    textAnalysis = AnalyzeTextCharacteristics(text)
    
    ' 构建基础提示词
    basePrompt = "你是一个专业的文本处理助手。请根据以下指令处理文本：" & vbCrLf
    
    ' 根据模式添加具体指令
    Select Case True
        Case mode Like "rewrite:*"
            Dim rewriteStyle As String
            rewriteStyle = Split(mode, ":")(1)
            basePrompt = basePrompt & "对文本进行智能改写，采用" & rewriteStyle & "风格。" & vbCrLf & _
                        "保持原意不变，优化表达方式，使文本更加流畅自然。"
            
        Case mode Like "translate:*"
            Dim targetLang As String
            targetLang = Split(mode, ":")(1)
            basePrompt = basePrompt & "将文本准确翻译为" & GetLanguageName(targetLang) & "。" & vbCrLf & _
                        "保持专业术语准确性，确保翻译自然流畅。"
            
        Case mode Like "polish:*"
            Dim polishIntensity As String
            polishIntensity = Split(mode, ":")(1)
            basePrompt = basePrompt & "对文本进行" & polishIntensity & "润色。" & vbCrLf & _
                        "修正语法错误，优化表达，提升文本质量。"
            
        Case mode Like "summarize:*"
            Dim summaryStyle As String
            summaryStyle = Split(mode, ":")(1)
            basePrompt = basePrompt & "对文本进行" & summaryStyle & "总结。" & vbCrLf & _
                        "提取核心要点，保持逻辑结构清晰。"
            
        Case mode Like "expand:*"
            Dim expandDirection As String
            expandDirection = Split(mode, ":")(1)
            basePrompt = basePrompt & "按照" & expandDirection & "方向扩写文本。" & vbCrLf & _
                        "基于原文内容进行合理扩展和丰富。"
            
        Case mode Like "style:*"
            Dim styleConversion As String
            styleConversion = Split(mode, ":")(1)
            basePrompt = basePrompt & "将文本风格转换为" & styleConversion & "。" & vbCrLf & _
                        "保持内容一致性，仅调整表达风格。"
            
        Case mode Like "custom:*"
            Dim customInstruction As String
            customInstruction = Split(mode, ":", 2)(1)
            basePrompt = basePrompt & customInstruction
            
        Case mode = "grammar"
            basePrompt = basePrompt & "检查并修正文本中的语法错误。" & vbCrLf & _
                        "保持原意不变，仅进行语法校正。"
            
        Case mode = "code_explain"
            basePrompt = basePrompt & "解释代码的功能和逻辑。" & vbCrLf & _
                        "用通俗易懂的语言说明代码作用。"
            
        Case Else
            basePrompt = basePrompt & "根据文本内容进行智能处理。"
    End Select
    
    ' 添加通用指令
    additionalInstructions = vbCrLf & vbCrLf & "重要要求：" & vbCrLf & _
                           "1. 只返回处理后的结果文本" & vbCrLf & _
                           "2. 不要添加任何额外说明或注释" & vbCrLf & _
                           "3. 保持原文的段落结构和逻辑顺序" & vbCrLf & _
                           "4. 确保专业术语的准确性"
    
    ' 添加文本分析结果
    If textAnalysis <> "" Then
        additionalInstructions = additionalInstructions & vbCrLf & "5. 文本特征：" & textAnalysis
    End If
    
    BuildSmartPrompt = basePrompt & additionalInstructions
End Function

' ==================== 增强版 API 交互 ====================

' 增强版 API 调用
Private Function CallEnhancedDeepSeekAPI(systemPrompt As String, userPrompt As String) As String
    On Error GoTo ErrHandler
    
    Dim http As Object
    Set http = CreateObject("MSXML2.XMLHTTP.6.0")
    
    Dim apiKey As String
    apiKey = GetEnhancedAPIKey()
    If Trim(apiKey) = "" Then
        LastDeepSeekError = "未配置有效的 API 密钥。"
        CallEnhancedDeepSeekAPI = ""
        Exit Function
    End If
    
    ' 构建增强版请求负载
    Dim payload As String
    payload = BuildEnhancedPayload(systemPrompt, userPrompt)
    
    ' 清除之前错误
    LastDeepSeekError = ""
    
    ' 发送请求
    http.Open "POST", DEEPSEEK_API_URL, False
    http.setRequestHeader "Content-Type", "application/json"
    http.setRequestHeader "Authorization", "Bearer " & apiKey
    http.setRequestHeader "User-Agent", "DeepSeek-Word-Integration/3.0"
    
    On Error Resume Next
    http.send payload
    If Err.Number <> 0 Then
        LastDeepSeekError = "网络请求失败：" & Err.Number & " - " & Err.Description
        CallEnhancedDeepSeekAPI = ""
        Exit Function
    End If
    On Error GoTo ErrHandler
    
    ' 检查状态码
    If http.Status < 200 Or http.Status >= 300 Then
        LastDeepSeekError = "API 请求失败，状态码: " & http.Status & _
                           vbCrLf & "响应: " & Left$(http.responseText, 500)
        CallEnhancedDeepSeekAPI = ""
        Exit Function
    End If
    
    Dim resp As String
    resp = http.responseText
    
    ' 解析响应
    Dim result As String
    result = ParseEnhancedAPIResponse(resp)
    
    If result = "" Then
        LastDeepSeekError = "无法解析 API 响应: " & Left$(resp, 500)
    End If
    
    CallEnhancedDeepSeekAPI = result
    Exit Function

ErrHandler:
    LastDeepSeekError = "API 调用异常: " & Err.Number & " - " & Err.Description
    CallEnhancedDeepSeekAPI = ""
End Function

' 构建增强版请求负载
Private Function BuildEnhancedPayload(systemPrompt As String, userPrompt As String) As String
    Dim messages As String
    messages = "["
    messages = messages & "{""role"": ""system"", ""content"": " & JsonEscape(systemPrompt) & "},"
    messages = messages & "{""role"": ""user"", ""content"": " & JsonEscape(userPrompt) & "}"
    messages = messages & "]"
    
    Dim payload As String
    payload = "{"
    payload = payload & """model"": " & JsonEscape(GetEnhancedModel()) & ","
    payload = payload & """messages"": " & messages & ","
    payload = payload & """max_tokens"": " & GetEnhancedMaxTokens() & ","
    payload = payload & """temperature"": " & Replace(Format(GetEnhancedTemperature(), "0.0#"), ",", ".") & ","
    payload = payload & """top_p"": " & Replace(Format(GetEnhancedTopP(), "0.0#"), ",", ".") & ","
    payload = payload & """stream"": false"
    payload = payload & "}"
    
    BuildEnhancedPayload = payload
End Function

' 解析增强版 API 响应
Private Function ParseEnhancedAPIResponse(jsonResponse As String) As String
    On Error GoTo ErrHandler
    
    ' 使用更健壮的 JSON 解析方法
    Dim json As Object
    Set json = CreateObject("Scripting.Dictionary")
    
    ' 简单的 JSON 解析逻辑
    Dim choicesPos As Long
    choicesPos = InStr(1, jsonResponse, """choices""", vbTextCompare)
    If choicesPos = 0 Then
        ParseEnhancedAPIResponse = ""
        Exit Function
    End If
    
    ' 查找第一个 message 对象
    Dim messagePos As Long
    messagePos = InStr(choicesPos, jsonResponse, """message""", vbTextCompare)
    If messagePos = 0 Then
        ParseEnhancedAPIResponse = ""
        Exit Function
    End If
    
    ' 查找 content 字段
    Dim contentPos As Long
    contentPos = InStr(messagePos, jsonResponse, """content""", vbTextCompare)
    If contentPos = 0 Then
        ParseEnhancedAPIResponse = ""
        Exit Function
    End If
    
    ' 查找 content 值的开始位置
    Dim valueStart As Long
    valueStart = InStr(contentPos, jsonResponse, ":", vbTextCompare)
    If valueStart = 0 Then
        ParseEnhancedAPIResponse = ""
        Exit Function
    End If
    
    ' 跳过空白字符
    valueStart = valueStart + 1
    Do While valueStart <= Len(jsonResponse) And Mid$(jsonResponse, valueStart, 1) Like "[ :" & vbCr & vbLf & vbTab & "]" 
        valueStart = valueStart + 1
    Loop
    
    ' 提取 content 值
    If Mid$(jsonResponse, valueStart, 1) = """" Then
        ' 字符串值
        valueStart = valueStart + 1
        Dim valueEnd As Long
        valueEnd = InStr(valueStart, jsonResponse, """")
        If valueEnd = 0 Then
            ParseEnhancedAPIResponse = ""
            Exit Function
        End If
        ParseEnhancedAPIResponse = Mid$(jsonResponse, valueStart, valueEnd - valueStart)
        ' 处理转义字符
        ParseEnhancedAPIResponse = Replace(ParseEnhancedAPIResponse, "\n", vbCrLf)
        ParseEnhancedAPIResponse = Replace(ParseEnhancedAPIResponse, "\t", vbTab)
        ParseEnhancedAPIResponse = Replace(ParseEnhancedAPIResponse, "\""", """")
        ParseEnhancedAPIResponse = Replace(ParseEnhancedAPIResponse, "\\", "\")
        ParseEnhancedAPIResponse = Replace(ParseEnhancedAPIResponse, "\r", vbCr)
    Else
        ' 其他类型的值（数字、布尔值等），直接返回空
        ParseEnhancedAPIResponse = ""
    End If
    
    Exit Function
ErrHandler:
    ParseEnhancedAPIResponse = ""
End Function

' ==================== 增强版配置管理 ====================

' 获取增强版 API 密钥
Private Function GetEnhancedAPIKey() As String
    On Error Resume Next
    Dim k As String
    
    ' 多层级密钥获取策略
    ' 1. 检查文档变量
    If Not ActiveDocument Is Nothing Then
        If DocVariableExists("DeepSeekApiKey") Then
            k = Trim(ActiveDocument.Variables("DeepSeekApiKey").Value)
            If k <> "" Then
                GetEnhancedAPIKey = k
                Exit Function
            End If
        End If
    End If
    
    ' 2. 检查环境变量
    k = Trim(Environ("DEEPSEEK_API_KEY"))
    If k <> "" Then
        GetEnhancedAPIKey = k
        Exit Function
    End If
    
    ' 3. 检查注册表（Windows）
    On Error Resume Next
    Dim wsh As Object
    Set wsh = CreateObject("WScript.Shell")
    k = wsh.RegRead("HKEY_CURRENT_USER\Software\DeepSeek\APIKey")
    If Err.Number = 0 And Trim(k) <> "" Then
        GetEnhancedAPIKey = k
        Exit Function
    End If
    On Error GoTo 0
    
    ' 4. 提示用户输入
    k = InputBox("未检测到 DeepSeek API 密钥。" & vbCrLf & _
                "请从 https://platform.deepseek.com/ 获取 API 密钥" & vbCrLf & _
                "请输入 API 密钥（可留空取消）:", "DeepSeek API 密钥配置")
    
    If Trim(k) <> "" Then
        ' 保存到文档变量
        On Error Resume Next
        If Not ActiveDocument Is Nothing Then
            If DocVariableExists("DeepSeekApiKey") Then
                ActiveDocument.Variables("DeepSeekApiKey").Value = k
            Else
                ActiveDocument.Variables.Add Name:="DeepSeekApiKey", Value:=k
            End If
        End If
        
        ' 尝试保存到注册表
        On Error Resume Next
        wsh.RegWrite "HKEY_CURRENT_USER\Software\DeepSeek\APIKey", k, "REG_SZ"
        On Error GoTo 0
    End If
    
    GetEnhancedAPIKey = k
End Function

' 获取增强版模型配置
Private Function GetEnhancedModel() As String
    If Not ActiveDocument Is Nothing And DocVariableExists("DeepSeekModel") Then
        GetEnhancedModel = Trim(ActiveDocument.Variables("DeepSeekModel").Value)
    Else
        GetEnhancedModel = DEFAULT_MODEL
    End If
End Function

' 获取增强版温度设置
Private Function GetEnhancedTemperature() As Double
    If Not ActiveDocument Is Nothing And DocVariableExists("DeepSeekTemperature") Then
        GetEnhancedTemperature = CDbl(ActiveDocument.Variables("DeepSeekTemperature").Value)
    Else
        GetEnhancedTemperature = DEFAULT_TEMPERATURE
    End If
End Function

' 获取增强版最大令牌数
Private Function GetEnhancedMaxTokens() As Long
    If Not ActiveDocument Is Nothing And DocVariableExists("DeepSeekMaxTokens") Then
        GetEnhancedMaxTokens = CLng(ActiveDocument.Variables("DeepSeekMaxTokens").Value)
    Else
        GetEnhancedMaxTokens = DEFAULT_MAX_TOKENS
    End If
End Function

' 获取增强版 Top-P 设置
Private Function GetEnhancedTopP() As Double
    If Not ActiveDocument Is Nothing And DocVariableExists("DeepSeekTopP") Then
        GetEnhancedTopP = CDbl(ActiveDocument.Variables("DeepSeekTopP").Value)
    Else
        GetEnhancedTopP = DEFAULT_TOP_P
    End If
End Function

' ==================== 智能配置向导 ====================

' 智能配置向导
Private Sub SmartConfigurationWizard()
    On Error GoTo ErrHandler
    
    Dim currentModel As String
    Dim currentTemp As String
    Dim currentTokens As String
    Dim currentTopP As String
    
    ' 获取当前设置
    currentModel = GetEnhancedModel()
    currentTemp = Format(GetEnhancedTemperature(), "0.0#")
    currentTokens = CStr(GetEnhancedMaxTokens())
    currentTopP = Format(GetEnhancedTopP(), "0.0#")
    
    ' 显示配置向导
    Dim configType As String
    configType = InputBox("请选择配置类型：" & vbCrLf & _
                         "1. 快速配置（推荐新手）" & vbCrLf & _
                         "2. 高级配置（自定义参数）" & vbCrLf & _
                         "3. 仅配置 API 密钥", "DeepSeek 配置向导", "1")
    
    Select Case configType
        Case "1"
            QuickConfiguration
        Case "2"
            AdvancedConfiguration
        Case "3"
            APIKeyManager
        Case Else
            Exit Sub
    End Select
    
    Exit Sub

ErrHandler:
    MsgBox "配置向导执行失败: " & Err.Number & " - " & Err.Description, vbExclamation, "DeepSeek"
End Sub

' 快速配置
Private Sub QuickConfiguration()
    Dim useCase As String
    useCase = InputBox("请选择主要使用场景：" & vbCrLf & _
                      "1. 商务文档处理" & vbCrLf & _
                      "2. 学术论文写作" & vbCrLf & _
                      "3. 创意内容创作" & vbCrLf & _
                      "4. 技术文档编写" & vbCrLf & _
                      "5. 多语言翻译", "DeepSeek 快速配置", "1")
    
    Dim model As String, temp As Double, tokens As Long, topP As Double
    
    Select Case useCase
        Case "1" ' 商务文档
            model = "deepseek-chat"
            temp = 0.3
            tokens = 2048
            topP = 0.9
        Case "2" ' 学术论文
            model = "deepseek-chat"
            temp = 0.2
            tokens = 4096
            topP = 0.8
        Case "3" ' 创意内容
            model = "deepseek-chat"
            temp = 0.8
            tokens = 2048
            topP = 0.95
        Case "4" ' 技术文档
            model = "deepseek-coder"
            temp = 0.1
            tokens = 4096
            topP = 0.7
        Case "5" ' 翻译
            model = "deepseek-chat"
            temp = 0.5
            tokens = 2048
            topP = 0.9
        Case Else
            Exit Sub
    End Select
    
    ' 保存配置
    SaveConfiguration model, temp, tokens, topP
    
    MsgBox "快速配置已完成：" & vbCrLf & _
           "模型: " & model & vbCrLf & _
           "温度: " & Format(temp, "0.0#") & vbCrLf & _
           "最大令牌数: " & tokens & vbCrLf & _
           "Top-P: " & Format(topP, "0.0#"), vbInformation, "DeepSeek"
End Sub

' 高级配置
Private Sub AdvancedConfiguration()
    Dim currentModel As String
    Dim currentTemp As String
    Dim currentTokens As String
    Dim currentTopP As String
    
    currentModel = GetEnhancedModel()
    currentTemp = Format(GetEnhancedTemperature(), "0.0#")
    currentTokens = CStr(GetEnhancedMaxTokens())
    currentTopP = Format(GetEnhancedTopP(), "0.0#")
    
    Dim modelInput As String
    Dim tempInput As String
    Dim tokensInput As String
    Dim topPInput As String
    
    modelInput = InputBox("请输入模型名称:", "高级配置 - 模型", currentModel)
    If Trim(modelInput) = "" Then Exit Sub
    
    tempInput = InputBox("请输入温度值 (0.0-2.0):", "高级配置 - 温度", currentTemp)
    If Trim(tempInput) = "" Then Exit Sub
    
    tokensInput = InputBox("请输入最大令牌数 (1-8192):", "高级配置 - 令牌", currentTokens)
    If Trim(tokensInput) = "" Then Exit Sub
    
    topPInput = InputBox("请输入 Top-P 值 (0.0-1.0):", "高级配置 - Top-P", currentTopP)
    If Trim(topPInput) = "" Then Exit Sub
    
    ' 验证输入
    Dim tempValue As Double
    Dim tokensValue As Long
    Dim topPValue As Double
    
    tempValue = CDbl(Replace(tempInput, ",", "."))
    If tempValue < 0 Or tempValue > 2 Then
        MsgBox "温度值必须在 0.0 到 2.0 之间。", vbExclamation, "DeepSeek"
        Exit Sub
    End If
    
    tokensValue = CLng(tokensInput)
    If tokensValue < 1 Or tokensValue > 8192 Then
        MsgBox "最大令牌数必须在 1 到 8192 之间。", vbExclamation, "DeepSeek"
        Exit Sub
    End If
    
    topPValue = CDbl(Replace(topPInput, ",", "."))
    If topPValue < 0 Or topPValue > 1 Then
        MsgBox "Top-P 值必须在 0.0 到 1.0 之间。", vbExclamation, "DeepSeek"
        Exit Sub
    End If
    
    ' 保存配置
    SaveConfiguration modelInput, tempValue, tokensValue, topPValue
    
    MsgBox "高级配置已保存：" & vbCrLf & _
           "模型: " & modelInput & vbCrLf & _
           "温度: " & Format(tempValue, "0.0#") & vbCrLf & _
           "最大令牌数: " & tokensValue & vbCrLf & _
           "Top-P: " & Format(topPValue, "0.0#"), vbInformation, "DeepSeek"
End Sub

' API 密钥管理器
Private Sub APIKeyManager()
    Dim currentKey As String
    currentKey = GetEnhancedAPIKey()
    
    If currentKey <> "" Then
        currentKey = Left(currentKey, 8) & "..." & Right(currentKey, 4)
    End If
    
    Dim options As String
    options = "API 密钥管理：" & vbCrLf & _
              "当前密钥: " & currentKey & vbCrLf & vbCrLf & _
              "请选择操作：" & vbCrLf & _
              "1. 更新 API 密钥" & vbCrLf & _
              "2. 清除 API 密钥" & vbCrLf & _
              "3. 测试 API 连接"
    
    Dim choice As String
    choice = InputBox(options, "DeepSeek API 密钥管理", "1")
    
    Select Case choice
        Case "1"
            UpdateAPIKey
        Case "2"
            ClearAPIKey
        Case "3"
            TestAPIConnection
        Case Else
            Exit Sub
    End Select
End Sub

' 模型配置
Private Sub ModelConfiguration()
    Dim currentModel As String
    currentModel = GetEnhancedModel()
    
    Dim models As String
    models = "请选择模型：" & vbCrLf & _
             "1. deepseek-chat (通用对话)" & vbCrLf & _
             "2. deepseek-coder (代码生成)" & vbCrLf & _
             "3. 自定义模型"
    
    Dim choice As String
    choice = InputBox(models, "DeepSeek 模型配置", "1")
    
    Dim newModel As String
    Select Case choice
        Case "1": newModel = "deepseek-chat"
        Case "2": newModel = "deepseek-coder"
        Case "3": 
            newModel = InputBox("请输入自定义模型名称:", "自定义模型", currentModel)
        Case Else: Exit Sub
    End Select
    
    If Trim(newModel) <> "" Then
        SaveModelConfiguration newModel
        MsgBox "模型配置已更新为: " & newModel, vbInformation, "DeepSeek"
    End If
End Sub

' ==================== 辅助函数 ====================

' 保存配置
Private Sub SaveConfiguration(model As String, temperature As Double, maxTokens As Long, topP As Double)
    On Error Resume Next
    If Not ActiveDocument Is Nothing Then
        ' 保存模型
        If DocVariableExists("DeepSeekModel") Then
            ActiveDocument.Variables("DeepSeekModel").Value = model
        Else
            ActiveDocument.Variables.Add Name:="DeepSeekModel", Value:=model
        End If
        
        ' 保存温度
        If DocVariableExists("DeepSeekTemperature") Then
            ActiveDocument.Variables("DeepSeekTemperature").Value = Format(temperature, "0.0#")
        Else
            ActiveDocument.Variables.Add Name:="DeepSeekTemperature", Value:=Format(temperature, "0.0#")
        End If
        
        ' 保存最大令牌数
        If DocVariableExists("DeepSeekMaxTokens") Then
            ActiveDocument.Variables("DeepSeekMaxTokens").Value = maxTokens
        Else
            ActiveDocument.Variables.Add Name:="DeepSeekMaxTokens", Value:=maxTokens
        End If
        
        ' 保存 Top-P
        If DocVariableExists("DeepSeekTopP") Then
            ActiveDocument.Variables("DeepSeekTopP").Value = Format(topP, "0.0#")
        Else
            ActiveDocument.Variables.Add Name:="DeepSeekTopP", Value:=Format(topP, "0.0#")
        End If
    End If
    
    ' 尝试保存到注册表
    On Error Resume Next
    Dim wsh As Object
    Set wsh = CreateObject("WScript.Shell")
    wsh.RegWrite "HKEY_CURRENT_USER\Software\DeepSeek\Model", model, "REG_SZ"
    wsh.RegWrite "HKEY_CURRENT_USER\Software\DeepSeek\Temperature", Format(temperature, "0.0#"), "REG_SZ"
    wsh.RegWrite "HKEY_CURRENT_USER\Software\DeepSeek\MaxTokens", maxTokens, "REG_DWORD"
    wsh.RegWrite "HKEY_CURRENT_USER\Software\DeepSeek\TopP", Format(topP, "0.0#"), "REG_SZ"
    On Error GoTo 0
End Sub

' 更新 API 密钥
Private Sub UpdateAPIKey()
    Dim newKey As String
    newKey = InputBox("请输入新的 DeepSeek API 密钥:", "更新 API 密钥")
    
    If Trim(newKey) <> "" Then
        On Error Resume Next
        ' 保存到文档变量
        If Not ActiveDocument Is Nothing Then
            If DocVariableExists("DeepSeekApiKey") Then
                ActiveDocument.Variables("DeepSeekApiKey").Value = newKey
            Else
                ActiveDocument.Variables.Add Name:="DeepSeekApiKey", Value:=newKey
            End If
        End If
        
        ' 保存到注册表
        Dim wsh As Object
        Set wsh = CreateObject("WScript.Shell")
        wsh.RegWrite "HKEY_CURRENT_USER\Software\DeepSeek\APIKey", newKey, "REG_SZ"
        On Error GoTo 0
        
        MsgBox "API 密钥已成功更新。", vbInformation, "DeepSeek"
    End If
End Sub

' 清除 API 密钥
Private Sub ClearAPIKey()
    Dim response As VbMsgBoxResult
    response = MsgBox("确定要清除所有存储的 API 密钥吗？", vbYesNo + vbQuestion, "DeepSeek")
    
    If response = vbYes Then
        On Error Resume Next
        ' 清除文档变量
        If Not ActiveDocument Is Nothing Then
            If DocVariableExists("DeepSeekApiKey") Then
                ActiveDocument.Variables("DeepSeekApiKey").Delete
            End If
        End If
        
        ' 清除注册表
        Dim wsh As Object
        Set wsh = CreateObject("WScript.Shell")
        wsh.RegDelete "HKEY_CURRENT_USER\Software\DeepSeek\APIKey"
        On Error GoTo 0
        
        MsgBox "API 密钥已清除。", vbInformation, "DeepSeek"
    End If
End Sub

' 测试 API 连接
Private Sub TestAPIConnection()
    On Error GoTo ErrHandler
    
    Dim apiKey As String
    apiKey = GetEnhancedAPIKey()
    
    If Trim(apiKey) = "" Then
        MsgBox "未配置 API 密钥，请先配置 API 密钥。", vbExclamation, "DeepSeek"
        Exit Sub
    End If
    
    Application.StatusBar = "正在测试 API 连接..."
    
    Dim http As Object
    Set http = CreateObject("MSXML2.XMLHTTP.6.0")
    
    Dim payload As String
    payload = "{""model"": ""deepseek-chat"", ""messages"": [{""role"": ""user"", ""content"": ""Hello""}], ""max_tokens"": 10}"
    
    http.Open "POST", DEEPSEEK_API_URL, False
    http.setRequestHeader "Content-Type", "application/json"
    http.setRequestHeader "Authorization", "Bearer " & apiKey
    
    On Error Resume Next
    http.send payload
    If Err.Number <> 0 Then
        MsgBox "网络连接测试失败：" & Err.Description, vbExclamation, "DeepSeek"
        Application.StatusBar = ""
        Exit Sub
    End If
    On Error GoTo ErrHandler
    
    Application.StatusBar = ""
    
    If http.Status = 200 Then
        MsgBox "API 连接测试成功！", vbInformation, "DeepSeek"
    Else
        MsgBox "API 连接测试失败，状态码：" & http.Status & vbCrLf & _
               "响应：" & Left$(http.responseText, 200), vbExclamation, "DeepSeek"
    End If
    
    Exit Sub
    
ErrHandler:
    Application.StatusBar = ""
    MsgBox "API 连接测试异常：" & Err.Description, vbExclamation, "DeepSeek"
End Sub

' 保存模型配置
Private Sub SaveModelConfiguration(model As String)
    On Error Resume Next
    If Not ActiveDocument Is Nothing Then
        If DocVariableExists("DeepSeekModel") Then
            ActiveDocument.Variables("DeepSeekModel").Value = model
        Else
            ActiveDocument.Variables.Add Name:="DeepSeekModel", Value:=model
        End If
    End If
    
    ' 保存到注册表
    Dim wsh As Object
    Set wsh = CreateObject("WScript.Shell")
    wsh.RegWrite "HKEY_CURRENT_USER\Software\DeepSeek\Model", model, "REG_SZ"
    On Error GoTo 0
End Sub

' ==================== 批量处理功能 ====================

' 批量处理文档
Private Sub BatchProcessDocument(processType As String)
    On Error GoTo ErrHandler
    
    Application.ScreenUpdating = False
    Application.StatusBar = "正在批量处理文档..."
    
    Dim processedCount As Long
    processedCount = 0
    
    Select Case processType
        Case "paragraphs"
            ' 处理所有段落
            Dim para As Paragraph
            For Each para In ActiveDocument.Paragraphs
                If Trim(para.Range.Text) <> vbCr Then
                    para.Range.Select
                    DeepSeek_AdvancedProcess "rewrite:通用"
                    processedCount = processedCount + 1
                End If
            Next para
            
        Case "headings"
            ' 处理所有标题
            Dim heading As Paragraph
            For Each heading In ActiveDocument.Paragraphs
                If heading.Style Like "*标题*" Or heading.Style Like "*Heading*" Then
                    If Trim(heading.Range.Text) <> vbCr Then
                        heading.Range.Select
                        DeepSeek_AdvancedProcess "rewrite:通用"
                        processedCount = processedCount + 1
                    End If
                End If
            Next heading
            
        Case "selected"
            ' 处理选中的多个段落
            If Selection.Type = wdSelectionNormal Then
                Dim selPara As Paragraph
                For Each selPara In Selection.Paragraphs
                    If Trim(selPara.Range.Text) <> vbCr Then
                        selPara.Range.Select
                        DeepSeek_AdvancedProcess "rewrite:通用"
                        processedCount = processedCount + 1
                    End If
                Next selPara
            Else
                MsgBox "请先选择要处理的段落。", vbExclamation, "DeepSeek"
            End If
            
        Case "document"
            ' 处理整个文档
            ActiveDocument.Content.Select
            DeepSeek_AdvancedProcess "rewrite:通用"
            processedCount = 1
    End Select
    
    Application.ScreenUpdating = True
    Application.StatusBar = ""
    
    MsgBox "批量处理完成，共处理 " & processedCount & " 个文本块。", vbInformation, "DeepSeek"
    Exit Sub
    
ErrHandler:
    Application.ScreenUpdating = True
    Application.StatusBar = ""
    MsgBox "批量处理失败：" & Err.Description, vbExclamation, "DeepSeek"
End Sub

' ==================== 格式保留类 ====================

' 增强格式保留类 - 确保原样式不会丢失
Private Class FormatPreservation
    Private savedStyle As Style
    Private savedFont As Font
    Private savedPara As ParagraphFormat
    Private savedText As String
    Private originalRange As Range
    Private savedCharFormats() As Font
    Private savedParaFormats() As ParagraphFormat
    
    ' 保存格式信息 - 增强版本
    Public Sub SaveFormatting(targetRange As Range)
        On Error GoTo ErrHandler
        
        Set originalRange = targetRange.Duplicate
        Set savedStyle = targetRange.Style
        Set savedFont = targetRange.Font.Duplicate
        Set savedPara = targetRange.ParagraphFormat.Duplicate
        savedText = targetRange.Text
        
        ' 保存每个字符的格式信息
        Dim charCount As Long
        charCount = targetRange.Characters.Count
        ReDim savedCharFormats(1 To charCount)
        
        Dim i As Long
        For i = 1 To charCount
            On Error Resume Next
            Set savedCharFormats(i) = targetRange.Characters(i).Font.Duplicate
            On Error GoTo 0
        Next i
        
        ' 保存每个段落的格式信息
        Dim paraCount As Long
        paraCount = targetRange.Paragraphs.Count
        ReDim savedParaFormats(1 To paraCount)
        
        For i = 1 To paraCount
            On Error Resume Next
            Set savedParaFormats(i) = targetRange.Paragraphs(i).Format.Duplicate
            On Error GoTo 0
        Next i
        
        Exit Sub
        
ErrHandler:
        ' 如果详细格式保存失败，至少保存基本格式
        On Error Resume Next
        Set savedStyle = targetRange.Style
        Set savedFont = targetRange.Font.Duplicate
        Set savedPara = targetRange.ParagraphFormat.Duplicate
        savedText = targetRange.Text
        On Error GoTo 0
    End Sub
    
    ' 恢复格式 - 增强版本，确保原样式不会丢失
    Public Sub RestoreFormatting(targetRange As Range, newText As String)
        On Error GoTo ErrHandler
        
        Application.ScreenUpdating = False
        
        ' 方法1：尝试逐字符替换（最精确的格式保留）
        If TryCharacterByCharacterReplacement(targetRange, newText) Then
            Application.ScreenUpdating = True
            Exit Sub
        End If
        
        ' 方法2：如果逐字符替换失败，使用增强的整体格式恢复
        EnhancedBulkFormatRestore targetRange, newText
        
        Application.ScreenUpdating = True
        Exit Sub
        
ErrHandler:
        Application.ScreenUpdating = True
        ' 最终回退：基本格式恢复
        On Error Resume Next
        targetRange.Style = savedStyle
        targetRange.Font.Name = savedFont.Name
        targetRange.Font.Size = savedFont.Size
        targetRange.Font.Bold = savedFont.Bold
        targetRange.Font.Italic = savedFont.Italic
        targetRange.ParagraphFormat.Alignment = savedPara.Alignment
        On Error GoTo 0
    End Sub
    
    ' 尝试逐字符替换（最精确的格式保留方法）
    Private Function TryCharacterByCharacterReplacement(targetRange As Range, newText As String) As Boolean
        On Error GoTo ErrHandler
        
        Dim origCount As Long, newLen As Long
        origCount = originalRange.Characters.Count
        newLen = Len(newText)
        
        ' 如果字符数差异太大，可能不适合逐字符替换
        If Abs(origCount - newLen) > 1000 Then
            TryCharacterByCharacterReplacement = False
            Exit Function
        End If
        
        Dim minCount As Long
        minCount = Min(origCount, newLen)
        
        ' 逐字符替换前 minCount 个字符
        Dim i As Long
        For i = 1 To minCount
            On Error Resume Next
            targetRange.Characters(i).Text = Mid$(newText, i, 1)
            CopyFontProps savedCharFormats(i), targetRange.Characters(i).Font
            On Error GoTo 0
        Next i
        
        ' 处理剩余字符
        If newLen > origCount Then
            ' 插入额外字符，使用最后一个字符的格式
            Dim insertRange As Range
            Set insertRange = targetRange.Duplicate
            insertRange.Collapse wdCollapseEnd
            
            For i = origCount + 1 To newLen
                On Error Resume Next
                insertRange.InsertAfter Mid$(newText, i, 1)
                CopyFontProps savedCharFormats(origCount), insertRange.Font
                Set insertRange = targetRange.Duplicate
                insertRange.Collapse wdCollapseEnd
                On Error GoTo 0
            Next i
        ElseIf newLen < origCount Then
            ' 删除多余字符
            Dim delRange As Range
            Set delRange = targetRange.Duplicate
            delRange.Start = targetRange.Start + newLen
            delRange.End = targetRange.End
            On Error Resume Next
            delRange.Delete
            On Error GoTo 0
        End If
        
        ' 恢复段落格式
        RestoreParagraphFormats targetRange
        
        TryCharacterByCharacterReplacement = True
        Exit Function
        
ErrHandler:
        TryCharacterByCharacterReplacement = False
    End Function
    
    ' 增强的整体格式恢复
    Private Sub EnhancedBulkFormatRestore(targetRange As Range, newText As String)
        On Error GoTo ErrHandler
        
        ' 保存原始位置
        Dim originalStart As Long
        originalStart = targetRange.Start
        
        ' 替换文本
        targetRange.Text = newText
        
        ' 重新设置选区
        Dim newRange As Range
        Set newRange = targetRange.Document.Range(originalStart, originalStart + Len(newText))
        
        ' 全面恢复格式
        On Error Resume Next
        
        ' 恢复样式
        newRange.Style = savedStyle
        
        ' 恢复字体属性
        newRange.Font.Name = savedFont.Name
        newRange.Font.Size = savedFont.Size
        newRange.Font.Bold = savedFont.Bold
        newRange.Font.Italic = savedFont.Italic
        newRange.Font.Underline = savedFont.Underline
        newRange.Font.Color = savedFont.Color
        newRange.Font.Subscript = savedFont.Subscript
        newRange.Font.Superscript = savedFont.Superscript
        newRange.Font.StrikeThrough = savedFont.StrikeThrough
        newRange.Font.Hidden = savedFont.Hidden
        newRange.Font.Spacing = savedFont.Spacing
        newRange.Font.Scaling = savedFont.Scaling
        
        ' 恢复段落格式
        newRange.ParagraphFormat.Alignment = savedPara.Alignment
        newRange.ParagraphFormat.LeftIndent = savedPara.LeftIndent
        newRange.ParagraphFormat.RightIndent = savedPara.RightIndent
        newRange.ParagraphFormat.SpaceBefore = savedPara.SpaceBefore
        newRange.ParagraphFormat.SpaceAfter = savedPara.SpaceAfter
        newRange.ParagraphFormat.LineSpacing = savedPara.LineSpacing
        newRange.ParagraphFormat.FirstLineIndent = savedPara.FirstLineIndent
        newRange.ParagraphFormat.KeepWithNext = savedPara.KeepWithNext
        newRange.ParagraphFormat.KeepTogether = savedPara.KeepTogether
        newRange.ParagraphFormat.WidowControl = savedPara.WidowControl
        
        ' 选择新范围
        newRange.Select
        
        On Error GoTo 0
        Exit Sub
        
ErrHandler:
        ' 如果增强恢复失败，使用基本恢复
        On Error Resume Next
        targetRange.Style = savedStyle
        targetRange.Font.Name = savedFont.Name
        targetRange.Font.Size = savedFont.Size
        On Error GoTo 0
    End Sub
    
    ' 恢复段落格式
    Private Sub RestoreParagraphFormats(targetRange As Range)
        On Error Resume Next
        Dim i As Long
        For i = 1 To targetRange.Paragraphs.Count
            If i <= UBound(savedParaFormats) Then
                With targetRange.Paragraphs(i).Format
                    .Alignment = savedParaFormats(i).Alignment
                    .LeftIndent = savedParaFormats(i).LeftIndent
                    .RightIndent = savedParaFormats(i).RightIndent
                    .SpaceBefore = savedParaFormats(i).SpaceBefore
                    .SpaceAfter = savedParaFormats(i).SpaceAfter
                    .LineSpacing = savedParaFormats(i).LineSpacing
                End With
            End If
        Next i
        On Error GoTo 0
    End Sub
    
    ' 复制字体属性
    Private Sub CopyFontProps(srcFont As Font, dstFont As Font)
        On Error Resume Next
        dstFont.Name = srcFont.Name
        dstFont.Size = srcFont.Size
        dstFont.Bold = srcFont.Bold
        dstFont.Italic = srcFont.Italic
        dstFont.Underline = srcFont.Underline
        dstFont.Color = srcFont.Color
        dstFont.Subscript = srcFont.Subscript
        dstFont.Superscript = srcFont.Superscript
        dstFont.StrikeThrough = srcFont.StrikeThrough
        dstFont.Hidden = srcFont.Hidden
        dstFont.Spacing = srcFont.Spacing
        dstFont.Scaling = srcFont.Scaling
        On Error GoTo 0
    End Sub
    
    ' 返回较小值
    Private Function Min(a As Long, b As Long) As Long
        If a < b Then Min = a Else Min = b
    End Function
End Class

' ==================== 文本分析函数 ====================

' 分析文本特征
Private Function AnalyzeTextCharacteristics(text As String) As String
    Dim analysis As String
    analysis = ""
    
    ' 分析文本长度
    Dim textLength As Long
    textLength = Len(text)
    If textLength > 0 Then
        analysis = analysis & "文本长度: " & textLength & " 字符"
    End If
    
    ' 分析段落数
    Dim paragraphCount As Long
    paragraphCount = Len(text) - Len(Replace(text, vbCr, ""))
    If paragraphCount > 0 Then
        analysis = analysis & ", 段落数: " & paragraphCount + 1
    End If
    
    ' 分析语言特征（简单检测）
    If InStr(text, "。") > 0 Or InStr(text, "，") > 0 Then
        analysis = analysis & ", 语言: 中文"
    ElseIf InStr(text, ".") > 0 Or InStr(text, ",") > 0 Then
        analysis = analysis & ", 语言: 英文"
    End If
    
    ' 分析内容类型（简单检测）
    If InStr(LCase(text), "function") > 0 Or InStr(LCase(text), "sub ") > 0 Then
        analysis = analysis & ", 类型: 代码"
    ElseIf InStr(LCase(text), "http") > 0 Or InStr(LCase(text), "www.") > 0 Then
        analysis = analysis & ", 类型: 包含链接"
    End If
    
    AnalyzeTextCharacteristics = analysis
End Function

' 获取语言名称
Private Function GetLanguageName(langCode As String) As String
    Select Case LCase(langCode)
        Case "zh": GetLanguageName = "中文"
        Case "en": GetLanguageName = "英语"
        Case "ja": GetLanguageName = "日语"
        Case "ko": GetLanguageName = "韩语"
        Case "fr": GetLanguageName = "法语"
        Case "de": GetLanguageName = "德语"
        Case "es": GetLanguageName = "西班牙语"
        Case "ru": GetLanguageName = "俄语"
        Case "ar": GetLanguageName = "阿拉伯语"
        Case Else: GetLanguageName = langCode
    End Select
End Function

' ==================== 错误处理和统计 ====================

' 显示增强版错误信息
Private Sub ShowEnhancedError(errorMsg As String)
    Dim fullMsg As String
    fullMsg = "DeepSeek 处理失败" & vbCrLf & vbCrLf
    
    If Trim(errorMsg) <> "" Then
        fullMsg = fullMsg & "错误详情：" & vbCrLf & errorMsg & vbCrLf & vbCrLf
    End If
    
    fullMsg = fullMsg & "请检查：" & vbCrLf & _
              "• 网络连接是否正常" & vbCrLf & _
              "• API 密钥是否正确配置" & vbCrLf & _
              "• API 调用额度是否充足" & vbCrLf & _
              "• 请求参数是否超出限制"
    
    MsgBox fullMsg, vbExclamation, "DeepSeek"
End Sub

' 显示成功消息
Private Sub ShowSuccessMessage(mode As String, originalLength As Long, newLength As Long)
    Dim modeName As String
    Select Case True
        Case mode Like "rewrite:*": modeName = "智能改写"
        Case mode Like "translate:*": modeName = "翻译"
        Case mode Like "polish:*": modeName = "润色"
        Case mode Like "summarize:*": modeName = "总结"
        Case mode Like "expand:*": modeName = "扩写"
        Case mode Like "style:*": modeName = "风格转换"
        Case mode Like "custom:*": modeName = "自定义处理"
        Case mode = "grammar": modeName = "语法检查"
        Case mode = "code_explain": modeName = "代码解释"
        Case Else: modeName = "处理"
    End Select
    
    Dim lengthInfo As String
    If originalLength > 0 And newLength > 0 Then
        Dim changePercent As Double
        changePercent = ((newLength - originalLength) / originalLength) * 100
        lengthInfo = vbCrLf & "长度变化: " & originalLength & " → " & newLength & _
                    " 字符 (" & Format(changePercent, "+0.0;-0.0") & "%)"
    End If
    
    MsgBox modeName & "完成！" & lengthInfo, vbInformation, "DeepSeek"
End Sub

' 显示使用统计
Private Sub ShowUsageStatistics()
    On Error Resume Next
    
    Dim stats As String
    stats = "DeepSeek for Word - 使用统计" & vbCrLf & vbCrLf
    
    ' 获取配置信息
    stats = stats & "当前配置：" & vbCrLf
    stats = stats & "• 模型: " & GetEnhancedModel() & vbCrLf
    stats = stats & "• 温度: " & Format(GetEnhancedTemperature(), "0.0#") & vbCrLf
    stats = stats & "• 最大令牌数: " & GetEnhancedMaxTokens() & vbCrLf
    stats = stats & "• Top-P: " & Format(GetEnhancedTopP(), "0.0#") & vbCrLf & vbCrLf
    
    ' API 密钥状态
    Dim apiKey As String
    apiKey = GetEnhancedAPIKey()
    If apiKey <> "" Then
        stats = stats & "API 密钥: 已配置 (" & Left(apiKey, 8) & "..." & Right(apiKey, 4) & ")" & vbCrLf
    Else
        stats = stats & "API 密钥: 未配置" & vbCrLf
    End If
    
    ' 文档信息
    If Not ActiveDocument Is Nothing Then
        stats = stats & "当前文档: " & ActiveDocument.Name & vbCrLf
        stats = stats & "文档路径: " & ActiveDocument.Path & vbCrLf
    End If
    
    stats = stats & vbCrLf & "版本信息：" & vbCrLf
    stats = stats & "• DeepSeek for Word v3.0" & vbCrLf
    stats = stats & "• 兼容 Microsoft Word 2010+" & vbCrLf
    stats = stats & "• 最后更新: 2025-11-22" & vbCrLf & vbCrLf
    
    stats = stats & "功能概览：" & vbCrLf
    stats = stats & "• 智能改写 (多种风格)" & vbCrLf
    stats = stats & "• 多语言翻译 (9种语言)" & vbCrLf
    stats = stats & "• 专业润色 (3种强度)" & vbCrLf
    stats = stats & "• 智能总结 (4种类型)" & vbCrLf
    stats = stats & "• 内容扩写 (5种方向)" & vbCrLf
    stats = stats & "• 语法检查" & vbCrLf
    stats = stats & "• 风格转换" & vbCrLf
    stats = stats & "• 代码解释" & vbCrLf
    stats = stats & "• 批量处理" & vbCrLf
    stats = stats & "• 高级自定义"
    
    MsgBox stats, vbInformation, "DeepSeek 使用统计"
End Sub

' 显示全面帮助
Private Sub ShowComprehensiveHelp()
    Dim helpText As String
    helpText = "DeepSeek for Word - 全面使用指南" & vbCrLf & vbCrLf
    
    helpText = helpText & "📋 核心功能命令：" & vbCrLf
    helpText = helpText & "• DS_EnhancedRewrite - 智能改写 (5种风格)" & vbCrLf
    helpText = helpText & "• DS_MultiTranslate - 多语言翻译 (9种语言)" & vbCrLf
    helpText = helpText & "• DS_ProfessionalPolish - 专业润色 (3种强度)" & vbCrLf
    helpText = helpText & "• DS_SmartSummarize - 智能总结 (4种类型)" & vbCrLf
    helpText = helpText & "• DS_ContentExpand - 内容扩写 (5种方向)" & vbCrLf
    helpText = helpText & "• DS_GrammarCheck - 语法检查" & vbCrLf
    helpText = helpText & "• DS_StyleConvert - 风格转换" & vbCrLf
    helpText = helpText & "• DS_CodeExplain - 代码解释" & vbCrLf
    helpText = helpText & "• DS_AdvancedCustom - 高级自定义" & vbCrLf & vbCrLf
    
    helpText = helpText & "🔄 批量处理命令：" & vbCrLf
    helpText = helpText & "• DS_BatchProcess - 批量处理文档" & vbCrLf & vbCrLf
    
    helpText = helpText & "⚙️ 配置管理命令：" & vbCrLf
    helpText = helpText & "• DS_SmartConfigure - 智能配置向导" & vbCrLf
    helpText = helpText & "• DS_KeyManager - API 密钥管理" & vbCrLf
    helpText = helpText & "• DS_ModelSettings - 模型设置" & vbCrLf
    helpText = helpText & "• DS_UsageStats - 使用统计" & vbCrLf
    helpText = helpText & "• DS_ComprehensiveHelp - 全面帮助" & vbCrLf & vbCrLf
    
    helpText = helpText & "🚀 使用步骤：" & vbCrLf
    helpText = helpText & "1. 选中要处理的文本" & vbCrLf
    helpText = helpText & "2. 运行相应的宏命令" & vbCrLf
    helpText = helpText & "3. 首次使用需要配置 API 密钥" & vbCrLf & vbCrLf
    
    helpText = helpText & "🔑 API 密钥获取：" & vbCrLf
    helpText = helpText & "请访问 https://platform.deepseek.com/" & vbCrLf
    helpText = helpText & "注册账号并获取 API 密钥" & vbCrLf & vbCrLf
    
    helpText = helpText & "💡 使用技巧：" & vbCrLf
    helpText = helpText & "• 对于长文档，建议分段处理" & vbCrLf
    helpText = helpText & "• 批量处理功能适合处理多个段落" & vbCrLf
    helpText = helpText & "• 自定义指令可以发挥最大创意" & vbCrLf
    helpText = helpText & "• 格式保留功能确保文档样式不变" & vbCrLf & vbCrLf
    
    helpText = helpText & "📞 技术支持：" & vbCrLf
    helpText = helpText & "• 版本: DeepSeek for Word v3.0" & vbCrLf
    helpText = helpText & "• 兼容: Microsoft Word 2010+" & vbCrLf
    helpText = helpText & "• 更新: 2025-11-22"
    
    MsgBox helpText, vbInformation, "DeepSeek 全面帮助"
End Sub

' ==================== 兼容性函数 ====================

' 文档变量存在性检查
Private Function DocVariableExists(varName As String) As Boolean
    On Error GoTo ErrHandler
    Dim v As Variable
    Dim found As Boolean
    found = False
    If ActiveDocument Is Nothing Then
        DocVariableExists = False
        Exit Function
    End If
    For Each v In ActiveDocument.Variables
        If StrComp(v.Name, varName, vbTextCompare) = 0 Then
            found = True
            Exit For
        End If
    Next v
    DocVariableExists = found
    Exit Function
ErrHandler:
    DocVariableExists = False
End Function

' JSON 字符串转义
Private Function JsonEscape(s As String) As String
    s = Replace(s, "\", "\\")
    s = Replace(s, """", "\""")
    s = Replace(s, vbCrLf, "\n")
    s = Replace(s, vbLf, "\n")
    s = Replace(s, vbCr, "\n")
    s = Replace(s, vbTab, "\t")
    JsonEscape = """" & s & """"
End Function

' ==================== 向后兼容性命令 ====================

' 向后兼容性命令 - 保持与旧版本的兼容性
Public Sub DS_Rewrite()
    DS_EnhancedRewrite
End Sub

Public Sub DS_Translate()
    DS_MultiTranslate
End Sub

Public Sub DS_Polish()
    DS_ProfessionalPolish
End Sub

Public Sub DS_Summarize()
    DS_SmartSummarize
End Sub

Public Sub DS_Expand()
    DS_ContentExpand
End Sub

Public Sub DS_Custom()
    DS_AdvancedCustom
End Sub

Public Sub DS_Configure()
    DS_SmartConfigure
End Sub

Public Sub DS_Help()
    DS_ComprehensiveHelp
End Sub

' ==================== 初始化函数 ====================

' 模块初始化
Private Sub Auto_Open()
    ' 模块加载时执行
    On Error Resume Next
    Application.StatusBar = "DeepSeek for Word v3.0 已加载"
    Application.OnTime Now + TimeValue("00:00:03"), "ClearStatusBar"
End Sub

' 清除状态栏
Public Sub ClearStatusBar()
    Application.StatusBar = ""
End Sub

' ==================== 模块结束 ====================
