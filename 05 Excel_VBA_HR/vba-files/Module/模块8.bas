Attribute VB_Name = "模块8"
Option Explicit

' =======================================================
'' 解释：分隔线注释，用于视觉分区。
' 功能：根据“职员”表补全“考勤清单”中的工号(C列)与部门(E列)
'' 解释：说明该过程或函数的核心业务目标。
' 规则：
'' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
' 1) 通过姓名(F列)匹配职员表B列，回填工号到C列
'' 解释：说明表格列位与业务字段的映射关系。
' 2) 通过工号(C列)匹配职员表A列，回填部门到E列
'' 解释：说明表格列位与业务字段的映射关系。
' =======================================================
'' 解释：分隔线注释，用于视觉分区。
Public Sub 获取工号和部门信息()
    On Error GoTo ErrorHandler

    Dim startTime As Double
    startTime = Timer

    Dim calcMode As XlCalculation
    Dim screenUpdating As Boolean
    Dim enableEvents As Boolean
    Dim shouldResetStatusBar As Boolean  ' 仅正常完成时自动清空状态栏

    calcMode = Application.Calculation
    screenUpdating = Application.ScreenUpdating
    enableEvents = Application.EnableEvents

    Application.ScreenUpdating = False
    Application.Calculation = xlCalculationManual
    Application.EnableEvents = False

    shouldResetStatusBar = False

    Dim wsAttend As Worksheet
    Dim wsStaff As Worksheet

    Set wsAttend = GetAttendSheet(ThisWorkbook)
    If wsAttend Is Nothing Then
        UpdateStatus "错误：找不到考勤清单工作表", False
        GoTo CleanUp
    End If

    On Error Resume Next
    Set wsStaff = ThisWorkbook.Sheets("职员")
    On Error GoTo ErrorHandler
    If wsStaff Is Nothing Then
        UpdateStatus "错误：找不到职员工作表", False
        GoTo CleanUp
    End If

    UpdateStatus "正在读取数据..."

    Dim lastRow As Long
    lastRow = wsAttend.Cells(wsAttend.Rows.Count, "F").End(xlUp).Row
    If lastRow < 2 Then
        UpdateStatus "考勤清单中没有可处理的数据", False
        GoTo CleanUp
    End If

    Dim staffLastRow As Long
    staffLastRow = wsStaff.Cells(wsStaff.Rows.Count, "A").End(xlUp).Row
    If staffLastRow < 2 Then
        UpdateStatus "职员表中没有可用数据", False
        GoTo CleanUp
    End If

    ' 职员表：A=工号，B=姓名，C=部门
    '' 解释：该注释用于解释紧邻代码的业务意图或实现原因。
    Dim staffData As Variant
    staffData = wsStaff.Range("A2:C" & staffLastRow).Value

    Dim dictNameToID As Object
    Dim dictIDToDept As Object
    Set dictNameToID = CreateObject("Scripting.Dictionary")
    Set dictIDToDept = CreateObject("Scripting.Dictionary")
    dictNameToID.CompareMode = vbTextCompare
    dictIDToDept.CompareMode = vbTextCompare

    Dim i As Long
    Dim empID As String, empName As String, dept As String

    For i = 1 To UBound(staffData, 1)
        empID = Trim(CStr(staffData(i, 1) & ""))
        empName = Trim(CStr(staffData(i, 2) & ""))
        dept = Trim(CStr(staffData(i, 3) & ""))

        If empName <> "" And empID <> "" Then
            If Not dictNameToID.Exists(empName) Then dictNameToID.Add empName, empID
        End If

        If empID <> "" Then
            If Not dictIDToDept.Exists(empID) Then dictIDToDept.Add empID, dept
        End If
    Next i

    UpdateStatus "正在匹配并回填..."

        Dim nameArr As Variant
    Dim idArr As Variant
    Dim deptArr As Variant
    
    nameArr = wsAttend.Range("F2:F" & lastRow).Value
    idArr = wsAttend.Range("C2:C" & lastRow).Value
    deptArr = wsAttend.Range("E2:E" & lastRow).Value

    Dim processedCount As Long
    Dim matchedIDCount As Long
    Dim matchedDeptCount As Long
    Dim skippedCompletedCount As Long

    For i = 1 To UBound(nameArr, 1)
        processedCount = processedCount + 1

        empName = Trim(CStr(nameArr(i, 1) & ""))
        empID = Trim(CStr(idArr(i, 1) & ""))
        dept = Trim(CStr(deptArr(i, 1) & ""))

        ' 续接模式：C/E都已有值则跳过
        '' 解释：说明满足条件时会跳过当前分支处理。
        If empID <> "" And dept <> "" Then
            skippedCompletedCount = skippedCompletedCount + 1
            GoTo NextRow
        End If

        ' C列为空时按姓名补工号
        '' 解释：说明表格列位与业务字段的映射关系。
        If empID = "" And empName <> "" Then
            If dictNameToID.Exists(empName) Then
                empID = CStr(dictNameToID(empName))
                idArr(i, 1) = empID
                matchedIDCount = matchedIDCount + 1
            End If
        End If

        ' E列为空时按工号补部门
        '' 解释：说明表格列位与业务字段的映射关系。
        If dept = "" And empID <> "" Then
            If dictIDToDept.Exists(empID) Then
                dept = CStr(dictIDToDept(empID))
                deptArr(i, 1) = dept
                If dept <> "" Then matchedDeptCount = matchedDeptCount + 1
            End If
        End If

NextRow:
    Next i

    ' 批量写入，避免逐单元格写入
    '' 解释：说明该行是关键数据流或文件流操作。
    wsAttend.Range("C2:C" & lastRow).NumberFormatLocal = "@"
    wsAttend.Range("C2:C" & lastRow).Value = idArr
    wsAttend.Range("E2:E" & lastRow).Value = deptArr

    Dim elapsed As Double
    elapsed = Timer - startTime

    UpdateStatus "完成：工号补全 " & matchedIDCount & " 条，部门补全 " & matchedDeptCount & " 条，跳过已完整 " & skippedCompletedCount & " 条，耗时 " & Format(elapsed, "0.0") & " 秒", False
    shouldResetStatusBar = True  ' 成功完成后允许5秒后自动恢复状态栏


CleanUp:
    If shouldResetStatusBar Then ScheduleStatusBarReset
    Set dictNameToID = Nothing
    Set dictIDToDept = Nothing

    Application.EnableEvents = enableEvents
    Application.Calculation = calcMode
    Application.ScreenUpdating = screenUpdating
    Exit Sub

ErrorHandler:
    UpdateStatus "运行错误：" & Err.Description, False
    Resume CleanUp
End Sub

Private Function GetAttendSheet(ByVal wb As Workbook) As Worksheet
    On Error Resume Next
    Set GetAttendSheet = wb.Sheets("考勤清单")
    If GetAttendSheet Is Nothing Then
        Set GetAttendSheet = wb.Sheets("Sheet5")
    End If
    On Error GoTo 0
End Function

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








Private Sub ScheduleStatusBarReset()
    On Error Resume Next
    Application.onTime _
        EarliestTime:=Now + timeValue("00:00:05"), _
        Procedure:="恢复状态栏"
    On Error GoTo 0
End Sub
