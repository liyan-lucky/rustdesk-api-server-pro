Attribute VB_Name = "模块6"
Option Explicit

' =======================================================
'' 解释：分隔线注释，用于视觉分区。
' 功能：导出考勤卡
'' 解释：说明该过程或函数的核心业务目标。
' 1) 选择需要打开的Excel文件
'' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
' 2) 删除目标文件中名为“考勤卡”的Sheet（若存在）
'' 解释：说明该行是关键数据流或文件流操作。
' 3) 将当前文件的“考勤卡”复制到目标文件
'' 解释：说明该行是关键数据流或文件流操作。
' 4) 保持工作表位置与源文件一致
'' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
' =======================================================
'' 解释：分隔线注释，用于视觉分区。
Sub 导出考勤卡()
    On Error GoTo ErrorHandler

    Dim oldAskToUpdateLinks As Boolean
    Dim oldAutomationSecurity As Long
    Dim shouldResetStatusBar As Boolean  ' 仅正常完成时自动清空状态栏
    oldAskToUpdateLinks = Application.AskToUpdateLinks
    oldAutomationSecurity = Application.AutomationSecurity
    Application.AskToUpdateLinks = False
    Application.AutomationSecurity = 3  ' msoAutomationSecurityForceDisable
    shouldResetStatusBar = False

    Application.ScreenUpdating = False
    Application.Calculation = xlCalculationManual
    Application.EnableEvents = False
    Application.DisplayAlerts = False
    UpdateStatus "初始化中..."

    Dim filePath As Variant
    filePath = Application.GetOpenFilename( _
        FileFilter:="Excel 文件 (*.xlsx;*.xlsm;*.xls;*.xlsb),*.xlsx;*.xlsm;*.xls;*.xlsb", _
        Title:="选择要导出的目标文件")

    If filePath = False Then
        UpdateStatus "已取消选择文件", False
        GoTo CleanExit
    End If

    UpdateStatus "正在打开目标文件..."

    Dim wbSource As Workbook
    Dim wbTarget As Workbook
    Dim wsSource As Worksheet
    Dim pv As ProtectedViewWindow

    Set wbSource = ThisWorkbook

    On Error Resume Next
    Set pv = Application.ProtectedViewWindows.Open(filePath)
    On Error GoTo ErrorHandler

    If Not pv Is Nothing Then
        On Error Resume Next
        Set wbTarget = pv.Edit
        On Error GoTo ErrorHandler
    Else
        Set wbTarget = Workbooks.Open(filePath, UpdateLinks:=0, ReadOnly:=False, IgnoreReadOnlyRecommended:=True, Notify:=False)
    End If

    If wbTarget Is Nothing Then
        UpdateStatus "无法启用编辑或打开目标文件", False
        GoTo CleanExit
    End If

    ' 获取源工作表
    '' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
    Set wsSource = wbSource.Worksheets("考勤卡")

    ' 删除目标文件中已有的“考勤卡”工作表
    '' 解释：说明该行是关键数据流或文件流操作。
    UpdateStatus "正在删除目标文件中的旧表..."
    Dim wsOld As Worksheet
    On Error Resume Next
    Set wsOld = wbTarget.Worksheets("考勤卡")
    On Error GoTo ErrorHandler
    If Not wsOld Is Nothing Then
        wsOld.Delete
    End If

    ' 复制源工作表到目标文件，并保持位置一致
    '' 解释：说明该行是关键数据流或文件流操作。
    UpdateStatus "正在复制工作表..."
    Dim targetCount As Long
    Dim srcIndex As Long
    srcIndex = wsSource.Index
    targetCount = wbTarget.Worksheets.count

    If srcIndex <= targetCount Then
        wsSource.Copy Before:=wbTarget.Worksheets(srcIndex)
    Else
        wsSource.Copy After:=wbTarget.Worksheets(targetCount)
    End If

    ' 保存并关闭目标文件
    '' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
    wbTarget.Save
    wbTarget.Close SaveChanges:=False

    UpdateStatus "导出完成", False

    shouldResetStatusBar = True  ' 成功完成后允许5秒后自动恢复状态栏

CleanExit:
    If shouldResetStatusBar Then ScheduleStatusBarReset
    On Error Resume Next
    If Not pv Is Nothing Then pv.Close
    On Error GoTo 0
    Application.AskToUpdateLinks = oldAskToUpdateLinks
    Application.AutomationSecurity = oldAutomationSecurity
    Application.DisplayAlerts = True
    Application.EnableEvents = True
    Application.Calculation = xlCalculationAutomatic
    Application.ScreenUpdating = True
    Exit Sub

ErrorHandler:
    UpdateStatus "运行错误：" & Err.Description, False
    On Error Resume Next
    If Not wbTarget Is Nothing Then wbTarget.Close SaveChanges:=False
    On Error GoTo 0
    Resume CleanExit
End Sub


Private Sub ScheduleStatusBarReset()
    On Error Resume Next
    Application.onTime _
        EarliestTime:=Now + timeValue("00:00:05"), _
        Procedure:="恢复状态栏"
    On Error GoTo 0
End Sub
