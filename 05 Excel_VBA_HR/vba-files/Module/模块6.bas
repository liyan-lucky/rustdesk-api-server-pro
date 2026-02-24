Attribute VB_Name = "模块6"
Option Explicit

' =======================================================
' 功能：导出考勤卡
' 1) 选择需要打开的Excel文件
' 2) 删除目标文件中名为"考勤卡"的Sheet（若存在）
' 3) 将当前文件的"考勤卡"复制到目标文件
' 4) 保持工作表位置与源文件一致
' =======================================================
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

        UpdateStatus "正在检查并解除文件锁定..."
    尝试解除文件锁定 CStr(filePath)

    UpdateStatus "正在打开目标文件..."

    Dim wbSource As Workbook
    Dim wbTarget As Workbook
    Dim wsSource As Worksheet
    Dim pv As ProtectedViewWindow

    Set wbSource = ThisWorkbook

        ' 先按普通方式打开；若仍被拦截，再回退到受保护视图流程
    On Error Resume Next
    Set wbTarget = Workbooks.Open(filePath, UpdateLinks:=0, ReadOnly:=False, IgnoreReadOnlyRecommended:=True, Notify:=False)
    On Error GoTo ErrorHandler

    If wbTarget Is Nothing Then
        On Error Resume Next
        Set pv = Application.ProtectedViewWindows.Open(filePath)
        On Error GoTo ErrorHandler

        If Not pv Is Nothing Then
            On Error Resume Next
            Set wbTarget = pv.Edit
            On Error GoTo ErrorHandler
        End If
    End If

    If wbTarget Is Nothing Then
        UpdateStatus "无法启用编辑或打开目标文件", False
        GoTo CleanExit
    End If

    ' 获取源工作表
    Set wsSource = wbSource.Worksheets("考勤卡")

    ' 若目标文件已存在同名工作表，先重命名为临时备份，避免删除阶段长时间阻塞后影响复制
    UpdateStatus "正在处理目标文件中的旧表..."
    Dim wsOld As Worksheet
    Dim tmpOldSheetName As String
    On Error Resume Next
    Set wsOld = wbTarget.Worksheets("考勤卡")
    On Error GoTo ErrorHandler
    If Not wsOld Is Nothing Then
        tmpOldSheetName = "__bak_kqk_" & Format(Now, "hhmmss")
        UpdateStatus "正在重命名旧表..."
        wsOld.Name = tmpOldSheetName
    End If

    ' 复制源工作表到目标文件，并保持位置一致（改为按内容复制，绕开 Worksheet.Copy 的不稳定问题）
    UpdateStatus "正在复制工作表..."
    Dim targetSheetCount As Long
    Dim srcIndex As Long
    Dim wsNew As Worksheet
    srcIndex = wsSource.Index
    targetSheetCount = wbTarget.Sheets.Count

    按内容复制工作表 wsSource, wbTarget, srcIndex, wsNew

    UpdateStatus "工作表复制完成，准备删除旧表副本..."

    If Not wsOld Is Nothing Then
        UpdateStatus "正在删除旧表副本..."
        wsOld.Delete
        Set wsOld = Nothing
    End If

    ' 保存并关闭目标文件
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
    UpdateStatus "运行错误：" & Err.Number & " - " & Err.Description, False
    On Error Resume Next
    If Not wbTarget Is Nothing Then wbTarget.Close SaveChanges:=False
    On Error GoTo 0
    Resume CleanExit
End Sub



' 过程说明：按内容复制工作表
Private Sub 按内容复制工作表(ByVal wsSource As Worksheet, ByVal wbTarget As Workbook, ByVal srcIndex As Long, ByRef wsNew As Worksheet)
    On Error GoTo CopyFailed

    Dim targetSheetCount As Long
    targetSheetCount = wbTarget.Sheets.Count

    If targetSheetCount <= 0 Then
        Set wsNew = wbTarget.Worksheets.Add
    ElseIf srcIndex <= targetSheetCount Then
        Set wsNew = wbTarget.Worksheets.Add(Before:=wbTarget.Sheets(srcIndex))
    Else
        Set wsNew = wbTarget.Worksheets.Add(After:=wbTarget.Sheets(targetSheetCount))
    End If

    UpdateStatus "正在复制单元格内容与格式..."
    Dim srcUsed As Range
    Set srcUsed = wsSource.UsedRange

    If Not srcUsed Is Nothing Then
        srcUsed.Copy
        wsNew.Range(srcUsed.Cells(1, 1).Address).PasteSpecial xlPasteAll
        wsNew.Range(srcUsed.Cells(1, 1).Address).PasteSpecial xlPasteColumnWidths
    End If
    Application.CutCopyMode = False

    UpdateStatus "正在同步行高..."
    同步行高 wsSource, wsNew

    On Error Resume Next
    wsNew.Name = wsSource.Name
    wsNew.Tab.Color = wsSource.Tab.Color
    On Error GoTo CopyFailed
    Exit Sub

CopyFailed:
    Application.CutCopyMode = False
    Err.Raise Err.Number, "按内容复制工作表", Err.Description
End Sub

' 过程说明：同步行高
Private Sub 同步行高(ByVal wsSource As Worksheet, ByVal wsTarget As Worksheet)
    On Error Resume Next

    Dim usedRg As Range
    Dim firstRow As Long
    Dim lastRow As Long
    Dim r As Long

    Set usedRg = wsSource.UsedRange
    If usedRg Is Nothing Then Exit Sub

    firstRow = usedRg.Row
    lastRow = usedRg.Row + usedRg.Rows.Count - 1

    For r = firstRow To lastRow
        wsTarget.Rows(r).RowHeight = wsSource.Rows(r).RowHeight
    Next r

    On Error GoTo 0
End Sub
' 过程说明：尝试解除文件锁定
Private Sub 尝试解除文件锁定(ByVal filePath As String)
    On Error GoTo SafeExit

    Dim shellObj As Object
    Dim cmd As String
    Dim exitCode As Long
    Dim escapedPath As String

    If Len(Trim$(filePath)) = 0 Then Exit Sub

    Set shellObj = CreateObject("WScript.Shell")

    ' 优先使用 PowerShell 的 Unblock-File 去除网络来源标记（Mark of the Web）
    escapedPath = Replace(filePath, "'", "''")
    cmd = "powershell -NoProfile -ExecutionPolicy Bypass -Command ""try { Unblock-File -LiteralPath '" & escapedPath & "' -ErrorAction Stop; exit 0 } catch { exit 1 }"""
    exitCode = shellObj.Run(cmd, 0, True)

    If exitCode <> 0 Then
        ' 兼容方案：直接删除 Zone.Identifier 附加数据流（若存在）
        cmd = "cmd /c del /f /q """ & filePath & ":Zone.Identifier"" >nul 2>nul"
        shellObj.Run cmd, 0, True
    End If

SafeExit:
    Set shellObj = Nothing
    On Error GoTo 0
End Sub
' 过程说明：ScheduleStatusBarReset
Private Sub ScheduleStatusBarReset()
    On Error Resume Next
    Application.onTime _
        EarliestTime:=Now + timeValue("00:00:05"), _
        Procedure:="恢复状态栏"
    On Error GoTo 0
End Sub


















