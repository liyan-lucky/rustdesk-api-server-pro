Attribute VB_Name = "模块5"
Option Explicit

' =======================================================
' 功能：导入考勤机记录
' 1) 选择需要打开的Excel文件
' 2) 删除当前文件中名为"考勤机"的Sheet（若存在）
' 3) 复制外部文件的第一个工作表到当前文件
' 4) 将复制后的工作表命名为"考勤机"
' 5) 若文件处于受保护视图，尝试启用编辑
' =======================================================
Sub 导入考勤机记录()
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
        Title:="选择要导入的考勤机记录文件")

    If filePath = False Then
        UpdateStatus "已取消选择文件", False
        GoTo CleanExit
    End If

    UpdateStatus "正在打开文件..."

    Dim wbSource As Workbook
    Dim wbTarget As Workbook
    Dim pv As ProtectedViewWindow
    Set wbTarget = ThisWorkbook

    On Error Resume Next
    Set pv = Application.ProtectedViewWindows.Open(filePath)
    On Error GoTo ErrorHandler

    If Not pv Is Nothing Then
        On Error Resume Next
        Set wbSource = pv.Edit
        On Error GoTo ErrorHandler
    Else
        Set wbSource = Workbooks.Open(filePath, UpdateLinks:=0, ReadOnly:=True, IgnoreReadOnlyRecommended:=True, Notify:=False)
    End If

    If wbSource Is Nothing Then
        UpdateStatus "无法启用编辑或打开文件", False
        GoTo CleanExit
    End If

    ' 删除当前文件中已有的"考勤机"工作表
    UpdateStatus "正在删除旧的考勤机表..."
    Dim wsOld As Worksheet
    On Error Resume Next
    Set wsOld = wbTarget.Worksheets("考勤机")
    On Error GoTo ErrorHandler
    If Not wsOld Is Nothing Then
        wsOld.Delete
    End If

    ' 复制外部文件的第一个工作表到当前文件
    UpdateStatus "正在复制工作表..."
    Dim wsSource As Worksheet
    Set wsSource = wbSource.Worksheets(1)
    wsSource.Copy After:=wbTarget.Worksheets(wbTarget.Worksheets.count)

    ' 修改新复制的工作表名称为"考勤机"
    Dim wsNew As Worksheet
    Set wsNew = wbTarget.Worksheets(wbTarget.Worksheets.count)
    wsNew.Name = "考勤机"

    wbSource.Close SaveChanges:=False

    UpdateStatus "导入完成", False

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
    If Not wbSource Is Nothing Then
        wbSource.Close SaveChanges:=False
    End If
    Resume CleanExit
End Sub

' 过程说明：ScheduleStatusBarReset
Private Sub ScheduleStatusBarReset()
    On Error Resume Next
    Application.onTime _
        EarliestTime:=Now + timeValue("00:00:05"), _
        Procedure:="恢复状态栏"
    On Error GoTo 0
End Sub

