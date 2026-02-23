Attribute VB_Name = "模块9"
' 模块名称定义（在VBA项目中显示的名称）
'' 解释：该注释用于解释紧邻代码的业务意图或实现原因。

' 一键流程：按顺序执行导入/清理/生成
'' 解释：说明该行是关键数据流或文件流操作。
Sub 一键处理流程()
    On Error GoTo ErrorHandler

    UpdateStatus "开始执行一键流程..."

    UpdateStatus "清理考勤卡..."
    清理考勤卡

    UpdateStatus "导入企业微信记录..."
    导入企业微信记录

    UpdateStatus "导入考勤机记录..."
    导入考勤机记录

    ' UpdateStatus "清理考勤机数据..."
    '' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
    ' 清理考勤机数据
    '' 解释：该注释用于解释紧邻代码的业务意图或实现原因。

    ' UpdateStatus "清理企业微信数据..."
    '' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
    ' 清理企业微信数据
    '' 解释：该注释用于解释紧邻代码的业务意图或实现原因。

    ' UpdateStatus "清理考勤清单数据..."
    '' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
    ' 清理考勤清单数据
    '' 解释：该注释用于解释紧邻代码的业务意图或实现原因。

    UpdateStatus "生成企业微信考勤卡..."
    获取企业微信记录

    UpdateStatus "生成考勤机考勤卡..."
    获取打卡机记录

    UpdateStatus "生成考勤清单..."
    生成考勤清单

    UpdateStatus "导出考勤卡..."
    导出考勤卡

    UpdateStatus "导出考勤清单..."
    导出考勤清单

    UpdateStatus "一键流程完成", False

    Application.onTime _
        EarliestTime:=Now + timeValue("00:00:05"), _
        Procedure:="恢复状态栏"

    Exit Sub

ErrorHandler:
    ' 异常结束：保留状态栏错误信息，不自动清空
    '' 解释：说明该段用于状态提示或状态栏恢复策略。
    UpdateStatus "一键流程出错：" & Err.Description, False
End Sub
' =================【主程序区】=================
'' 解释：分隔线注释，用于视觉分区。
' 主程序：清理考勤卡工作表的特定区域
'' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
Sub 清理考勤卡()
    ' 错误处理：如果出错，跳转到ErrorHandler1标签
    '' 解释：说明异常捕获与处理路径。
    On Error GoTo ErrorHandler1
    ' 调用核心清理函数，清理"考勤卡"工作表的指定区域
    '' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
    ' Array("C9:H24", "K9:P24", "I21:J24") 指定要清理的三个区域范围
    '' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
    Call 清理工作表("考勤卡", Array("C9:H24", "K9:P24", "I21:J24"))
    Exit Sub  ' 正常执行完后退出，避免执行错误处理代码
    
ErrorHandler1:
    ' 错误处理代码：在状态栏显示错误信息
    '' 解释：说明异常捕获与处理路径。
    UpdateStatus "清理考勤卡时出错：" & Err.Description
    ' 跳转到清理恢复代码
    '' 解释：说明该段用于恢复 Excel 全局状态。
    Resume CleanUp1
    
CleanUp1:
    ' 恢复Excel应用程序设置
    '' 解释：说明该段用于恢复 Excel 全局状态。
    恢复设置
End Sub

' 主程序：清理考勤机工作表的全部内容
'' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
Sub 清理考勤机数据()
    On Error GoTo ErrorHandler2
    ' 调用核心清理函数，不指定区域参数表示清理整个工作表
    '' 解释：说明函数参数的含义与用途。
    Call 清理工作表("考勤机")
    Exit Sub
    
ErrorHandler2:
    UpdateStatus "清理考勤机数据时出错：" & Err.Description
    Resume CleanUp2
    
CleanUp2:
    恢复设置
End Sub

' 主程序：清理企业微信工作表的全部内容
'' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
Sub 清理企业微信数据()
    On Error GoTo ErrorHandler3
    Call 清理工作表("企业微信")
    Exit Sub
    
ErrorHandler3:
    UpdateStatus "清理企业微信数据时出错：" & Err.Description
    Resume CleanUp3
    
CleanUp3:
    恢复设置
End Sub

' 主程序：清理考勤清单工作表的全部内容
'' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
Sub 清理考勤清单数据()
    On Error GoTo ErrorHandler4
    Call 清理考勤清单工作表("考勤清单")
    Exit Sub
    
ErrorHandler4:
    UpdateStatus "清理考勤清单数据时出错：" & Err.Description
    Resume CleanUp4

CleanUp4:
    恢复设置
End Sub

' =================【核心函数区】=================
'' 解释：分隔线注释，用于视觉分区。
' 核心清理函数：清理指定工作表的通用函数
'' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
' 参数说明：
'' 解释：说明函数参数的含义与用途。
'   - sheetName: 要清理的工作表名称
'' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
'   - specialRanges: 可选参数，指定要清理的特定区域数组
'' 解释：说明函数参数的含义与用途。
Private Sub 清理工作表(ByVal sheetName As String, Optional ByVal specialRanges As Variant = Empty)
    ' 初始化Excel应用程序设置（优化性能）
    '' 解释：说明该段用于提升执行速度与稳定性。
    初始化设置
    
    ' 本地错误处理
    '' 解释：说明异常捕获与处理路径。
    On Error GoTo LocalErrorHandler
    
    ' 记录开始时间（用于计算清理耗时）
    '' 解释：说明该行处理日期时间规则与边界。
    Dim startTime As Double
    startTime = Timer
    
    ' 声明工作表对象变量
    '' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
    Dim ws As Worksheet
    ' 获取指定名称的工作表
    '' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
    Set ws = 获取工作表(sheetName)
    
    ' 检查工作表是否存在
    '' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
    If ws Is Nothing Then
        ' 如果工作表不存在，抛出错误代码为1001的自定义错误
        '' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
        Err.Raise 1001, "清理工作表", "未找到名为 """ & sheetName & """ 的工作表！"
    End If
    
    ' 更新状态栏显示清理进度
    '' 解释：说明该段用于状态提示或状态栏恢复策略。
    UpdateStatus "正在清空工作表 """ & sheetName & """..."
    DoEvents  ' 允许处理其他事件，防止界面假死
    
    ' 判断清理方式：完全清空或选择性清空
    '' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
    If IsEmpty(specialRanges) Then
        ' 如果specialRanges参数为空，完全清空工作表
        '' 解释：说明函数参数的含义与用途。
        完全清空工作表 ws
    Else
        ' 否则，只清空指定的区域
        '' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
        选择性清空工作表 ws, specialRanges
    End If
    
    ' 构建完成消息，显示清理用时
    '' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
    Dim msg As String
    msg = "工作表 """ & sheetName & """ 已成功清空！用时 " & _
          Format(Timer - startTime, "0.00") & " 秒"
    
    ' 在状态栏显示完成消息
    '' 解释：说明该段用于状态提示或状态栏恢复策略。
    UpdateStatus msg
    
    ' 使用Application.OnTime定时功能，在5秒后自动恢复状态栏
    '' 解释：说明该段用于恢复 Excel 全局状态。
    ' 这样用户有足够时间看到完成消息，然后状态栏自动恢复正常
    '' 解释：说明该段用于恢复 Excel 全局状态。
    Application.onTime _
        EarliestTime:=Now + timeValue("00:00:05"), _
        Procedure:="恢复状态栏"  ' 调用公共的恢复状态栏函数
    
    Exit Sub  ' 正常退出
    
LocalErrorHandler:
    ' 本地错误处理：将错误重新抛出，让上层错误处理程序捕获
    '' 解释：说明异常捕获与处理路径。
    Err.Raise Err.Number, Err.Source, Err.Description
End Sub

' =================【工具函数区】=================
'' 解释：分隔线注释，用于视觉分区。
' 初始化Excel应用程序设置（优化性能）
'' 解释：说明该段用于提升执行速度与稳定性。
' 更新状态栏并保持界面响应
'' 解释：说明该段用于状态提示或状态栏恢复策略。
Private Sub UpdateStatus(ByVal msg As String, Optional ByVal doEventsFlag As Boolean = True)
    Static lastMsg As String
    Static lastTick As Double

    Dim nowTick As Double
    nowTick = Timer

    If msg = lastMsg Then
        If nowTick >= lastTick Then
            If nowTick - lastTick < 0.2 Then Exit Sub
        End If
    End If

    Application.StatusBar = msg
    lastMsg = msg
    lastTick = nowTick

    If doEventsFlag Then DoEvents
End Sub
Private Sub 初始化设置()
    Application.ScreenUpdating = False  ' 关闭屏幕更新，加快执行速度
    Application.Calculation = xlCalculationManual  ' 手动计算模式，避免自动计算影响性能
    Application.EnableEvents = False  ' 禁用事件触发，防止连锁反应
End Sub

' 恢复Excel应用程序设置
'' 解释：说明该段用于恢复 Excel 全局状态。
Private Sub 恢复设置()
    Application.ScreenUpdating = True   ' 恢复屏幕更新
    Application.Calculation = xlCalculationAutomatic  ' 恢复自动计算
    Application.EnableEvents = True  ' 启用事件触发
End Sub

' 根据名称获取工作表对象
'' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
Private Function 获取工作表(ByVal sheetName As String) As Worksheet
    ' 使用On Error Resume Next避免工作表不存在时程序崩溃
    '' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
    On Error Resume Next
    Set 获取工作表 = ThisWorkbook.Worksheets(sheetName)
    On Error GoTo 0  ' 恢复错误处理
End Function

' 完全清空整个工作表
'' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
Private Sub 完全清空工作表(ByVal ws As Worksheet)
    On Error GoTo ClearError
    ws.Cells.Clear  ' 清除工作表所有单元格的内容和格式
    ws.ScrollArea = ""  ' 清除滚动区域限制
    Exit Sub
    
ClearError:
    ' 如果出错，抛出错误
    '' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
    Err.Raise Err.Number, "完全清空工作表", Err.Description
End Sub

' 选择性清空工作表的指定区域
'' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
Private Sub 选择性清空工作表(ByVal ws As Worksheet, ByVal ranges As Variant)
    On Error GoTo ClearError
    
    Dim i As Long  ' 循环计数器
    
    ' 遍历并清空ranges数组中指定的每个区域
    '' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
    For i = LBound(ranges) To UBound(ranges)
        ws.Range(ranges(i)).ClearContents  ' 只清除内容，保留格式
    Next i
    
    ' 重置表头信息到默认值
    '' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
    ws.Range("E4").value = "姓名："
    ws.Range("H4").value = "工号："
    ws.Range("M4").value = "日期："
    
    ' 如果工作表有26行及以上的数据，仅清空已使用区域中的第26行及以下内容
    '' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
    Dim usedLastCellByRow As Range
    Dim usedLastCellByCol As Range

    Set usedLastCellByRow = ws.Cells.Find(What:="*", SearchOrder:=xlByRows, SearchDirection:=xlPrevious)
    Set usedLastCellByCol = ws.Cells.Find(What:="*", SearchOrder:=xlByColumns, SearchDirection:=xlPrevious)

    If Not usedLastCellByRow Is Nothing And Not usedLastCellByCol Is Nothing Then
        If usedLastCellByRow.Row >= 26 Then
            ws.Range(ws.Cells(26, 1), ws.Cells(usedLastCellByRow.Row, usedLastCellByCol.Column)).Clear
        End If
    End If
    
    Exit Sub
    
ClearError:
    Err.Raise Err.Number, "选择性清空工作表", Err.Description
End Sub

' 专门为考勤清单设计的清理函数：保留第一行数据
'' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
Private Sub 清理考勤清单工作表(ByVal sheetName As String)
    ' 初始化Excel应用程序设置（优化性能）
    '' 解释：说明该段用于提升执行速度与稳定性。
    初始化设置
    
    ' 本地错误处理
    '' 解释：说明异常捕获与处理路径。
    On Error GoTo LocalErrorHandler
    
    ' 记录开始时间（用于计算清理耗时）
    '' 解释：说明该行处理日期时间规则与边界。
    Dim startTime As Double
    startTime = Timer
    
    ' 声明工作表对象变量
    '' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
    Dim ws As Worksheet
    Dim lastRow As Long
    
    ' 获取指定名称的工作表
    '' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
    Set ws = 获取工作表(sheetName)
    
    ' 检查工作表是否存在
    '' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
    If ws Is Nothing Then
        ' 如果工作表不存在，抛出错误代码为1001的自定义错误
        '' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
        Err.Raise 1001, "清理工作表", "未找到名为 """ & sheetName & """ 的工作表！"
    End If
    
    ' 更新状态栏显示清理进度
    '' 解释：说明该段用于状态提示或状态栏恢复策略。
    UpdateStatus "正在清理工作表 """ & sheetName & """（保留第一行）..."
    DoEvents  ' 允许处理其他事件，防止界面假死
    
    ' 获取最后一行（使用列A来判断行数，可以根据需要调整列号）
    '' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
    lastRow = ws.Cells(ws.rows.count, "A").End(xlUp).Row
    
    ' 如果有超过一行的数据，从第二行开始清除
    '' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
    If lastRow > 1 Then
        ' 清除第2行到最后一行（保留表头）
        '' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
        ws.Range("A2:Z" & lastRow).ClearContents
        
        ' 或者如果你不知道最大列，也可以用以下方式：
        '' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
        ' ws.rows("2:" & lastRow).ClearContents
        '' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
        
        ' 如果只需要清除特定列，可以这样：
        '' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
        ' ws.Range("A2:E" & lastRow).ClearContents  ' 只清除A到E列
        '' 解释：说明表格列位与业务字段的映射关系。
    End If
    
    ' 构建完成消息，显示清理用时
    '' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
    Dim msg As String
    msg = "工作表 """ & sheetName & """ 已成功清理（保留第一行）！用时 " & _
          Format(Timer - startTime, "0.00") & " 秒"
    
    ' 在状态栏显示完成消息
    '' 解释：说明该段用于状态提示或状态栏恢复策略。
    UpdateStatus msg
    
    ' 使用Application.OnTime定时功能，在5秒后自动恢复状态栏
    '' 解释：说明该段用于恢复 Excel 全局状态。
    ' 这样用户有足够时间看到完成消息，然后状态栏自动恢复正常
    '' 解释：说明该段用于恢复 Excel 全局状态。
    Application.onTime _
        EarliestTime:=Now + timeValue("00:00:05"), _
        Procedure:="恢复状态栏"  ' 调用公共的恢复状态栏函数
    
    Exit Sub  ' 正常退出
    
LocalErrorHandler:
    ' 本地错误处理：将错误重新抛出，让上层错误处理程序捕获
    '' 解释：说明异常捕获与处理路径。
    Err.Raise Err.Number, Err.Source, Err.Description
End Sub

' 公开的状态栏恢复函数（供Application.OnTime调用）
'' 解释：说明该段用于恢复 Excel 全局状态。
' 必须是Public类型，否则OnTime无法调用
'' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
Public Sub 恢复状态栏()
    Application.StatusBar = False  ' 将状态栏恢复为默认状态
End Sub






