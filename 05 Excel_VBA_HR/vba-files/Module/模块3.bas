Attribute VB_Name = "模块3"
Option Explicit

' =======================================================
' 功能：从考勤卡数据生成考勤清单（横排样式）
' 输入：考勤卡数据表、人员信息表
' 输出：考勤清单表（横排格式，每日一行，数据从第2行开始）
' 注意：考勤清单表必须已存在，且表头在第1行
' =======================================================

Sub 生成考勤清单()
    ' 错误处理
    On Error GoTo ErrorHandler

    ' 初始化设置
    初始化设置

    Dim startTime As Double
    startTime = Timer

    ' ================= 工作表对象定义 =================
    Dim wsStaff As Worksheet        ' 人员信息表
    Dim wsCard As Worksheet         ' 考勤卡数据表
    Dim wsOutput As Worksheet       ' 考勤清单输出表
    Dim shouldResetStatusBar As Boolean  ' 仅正常完成时自动清空状态栏

    ' 获取工作表
    Set wsStaff = ThisWorkbook.Sheets("人员信息")
    Set wsCard = ThisWorkbook.Sheets("考勤卡")
    Set wsOutput = ThisWorkbook.Sheets("考勤清单")
    shouldResetStatusBar = False

    ' ================= 检查表中现有数据（支持续接） =================
    Dim lastDataRow As Long
    lastDataRow = wsOutput.Cells(wsOutput.rows.count, "F").End(xlUp).Row

    ' ================= 日期范围验证 =================
    Dim startDate As Date, endDate As Date
    startDate = CDate(wsStaff.Range("B2").value)
    endDate = CDate(wsStaff.Range("C2").value)

    ' ================= 从人员信息表获取时间参数（包括格式） =================
    Dim STANDARD_START As Variant, STANDARD_END As Variant, MEAL_TIME As Variant
    Dim OVERTIME_START As Variant, NIGHT_MEAL_TIME As Variant

    ' 获取时间参数值和格式
    STANDARD_START = wsStaff.Range("F2").value
    STANDARD_END = wsStaff.Range("G2").value
    MEAL_TIME = wsStaff.Range("H2").value
    OVERTIME_START = wsStaff.Range("I2").value
    NIGHT_MEAL_TIME = wsStaff.Range("J2").value

    ' 验证时间参数是否为空
    If IsEmpty(STANDARD_START) Or IsEmpty(STANDARD_END) Or IsEmpty(MEAL_TIME) Or _
       IsEmpty(OVERTIME_START) Or IsEmpty(NIGHT_MEAL_TIME) Then
        UpdateStatus "时间参数错误，请检查人员信息表F2-J2单元格是否填写"
        GoTo CleanExit
    End If

    ' ================= 读取考勤卡数据 =================
    UpdateStatus "读取考勤卡数据..."
    DoEvents  ' 允许用户操作

    Dim cardLastRow As Long
    cardLastRow = wsCard.Cells(wsCard.rows.count, "A").End(xlUp).Row

    If cardLastRow < 5 Then
        UpdateStatus "考勤卡表中没有数据"
        GoTo CleanExit
    End If

    ' 一次性读取考勤卡数据到数组
    Dim cardData As Variant
    cardData = wsCard.Range("A1:P" & cardLastRow).value

    ' 创建字典存储员工考勤数据
    Dim cardDict As Object
    Set cardDict = CreateObject("Scripting.Dictionary")

    ' 存储员工基础信息（姓名 -> 工号、部门）
    Dim empInfoDict As Object
    Set empInfoDict = CreateObject("Scripting.Dictionary")

    Dim cardDataLastRow As Long
    cardDataLastRow = UBound(cardData, 1)

    Dim i As Long
    i = 1
    Dim employeeCount As Long
    employeeCount = 0
    Dim totalRecords As Long
    totalRecords = 0

    ' 遍历考勤卡数据，每25行处理一个员工
    Do While i <= cardDataLastRow
        ' 检查是否到达员工信息行（第4行开始是第一个员工的信息）
        If i + 3 <= cardDataLastRow Then
            Dim unitCell As String, nameCell As String, empNoCell As String
            unitCell = Trim(cardData(i + 3, 1) & "")       ' A列：单位
            nameCell = Trim(cardData(i + 3, 5) & "")      ' E列：姓名
            empNoCell = Trim(cardData(i + 3, 8) & "")     ' H列：工号

            ' 检查是否是员工信息行（包含"姓名："）
            If InStr(nameCell, "姓名：") > 0 Then
                employeeCount = employeeCount + 1

                ' 提取员工信息
                Dim employeeName As String
                Dim cardEmployeeKey As String
                employeeName = Trim(Replace(nameCell, "姓名：", ""))
                cardEmployeeKey = NormalizeName(employeeName)

                If employeeName <> "" Then
                    ' 提取部门信息
                    Dim department As String
                    If InStr(unitCell, "单位：") > 0 Then
                        department = Trim(Replace(unitCell, "单位：", ""))
                    Else
                        department = ""
                    End If

                    ' 提取工号
                    Dim employeeNo As String
                    If InStr(empNoCell, "工号：") > 0 Then
                        employeeNo = Trim(Replace(empNoCell, "工号：", ""))
                    Else
                        employeeNo = ""
                    End If

                    If cardEmployeeKey <> "" Then
                        If Not empInfoDict.Exists(cardEmployeeKey) Then
                            empInfoDict(cardEmployeeKey) = Array(employeeNo, department)
                        End If
                    End If

                    ' 每个员工数据块分左右两栏处理（每半栏含标题共17行，实际16天数据）
                    Dim startRow As Long, endRow As Long
                    startRow = i + 7  ' 打卡数据开始行
                    endRow = i + 23   ' 半栏结束行（扣除标题行后对应16天数据）

                    If startRow <= cardDataLastRow And endRow <= cardDataLastRow Then
                        Dim dayRow As Long

                        ' 处理左半部分数据（A到H列）
                        For dayRow = startRow To endRow
                            ' 获取日期（A列）
                            Dim dateStr As String
                            dateStr = Trim(cardData(dayRow, 1) & "")  ' A列：日期

                            ' 跳过"日期"标题行
                            If dateStr = "日期" Then
                                GoTo NextDayRow
                            End If

                            ' 尝试解析日期
                            Dim recordDate As Date
                            Dim dateValid As Boolean
                            dateValid = False

                            ' 尝试解析日期（兼容 MM.DD 格式）
                            dateValid = TryParseCardDateByRange(dateStr, startDate, endDate, recordDate)

                            If dateValid Then
                                ' 检查日期是否在指定范围内
                                If recordDate >= startDate And recordDate <= endDate Then
                                    Dim dateKey As String
                                    dateKey = Format(recordDate, "yyyy-mm-dd")

                                    ' 获取打卡时间（处理数字格式）
                                    Dim workStart As String, workEnd As String
                                    ' C列是上午上班时间（作为上班时间）
                                    workStart = FormatExcelTime(cardData(dayRow, 3))
                                    ' H列是加班下班时间（作为下班时间）
                                    workEnd = FormatExcelTime(cardData(dayRow, 8))

                                    If workStart <> "" Or workEnd <> "" Then
                                        ' 存储到字典
                                        Dim dictKey As String
                                        dictKey = cardEmployeeKey & "|" & dateKey

                                        If Not cardDict.Exists(dictKey) Then
                                            cardDict(dictKey) = Array(workStart, workEnd, 1, employeeNo, department)
                                            totalRecords = totalRecords + 1
                                        End If
                                    End If
                                End If
                            End If
NextDayRow:
                        Next dayRow

                        ' 处理右半部分数据（I到P列）
                        For dayRow = startRow To endRow
                            ' 获取日期（I列）
                            dateStr = Trim(cardData(dayRow, 9) & "")  ' I列：日期

                            ' 跳过"日期"标题行
                            If dateStr = "日期" Then
                                GoTo NextDayRow2
                            End If

                            ' 尝试解析日期
                            dateValid = False

                            ' 尝试解析日期（兼容 MM.DD 格式）
                            dateValid = TryParseCardDateByRange(dateStr, startDate, endDate, recordDate)

                            If dateValid Then
                                ' 检查日期是否在指定范围内
                                If recordDate >= startDate And recordDate <= endDate Then
                                    dateKey = Format(recordDate, "yyyy-mm-dd")

                                    ' 获取打卡时间（处理数字格式）
                                    ' K列是上午上班时间（作为上班时间）
                                    workStart = FormatExcelTime(cardData(dayRow, 11))
                                    ' P列是加班下班时间（作为下班时间）
                                    workEnd = FormatExcelTime(cardData(dayRow, 16))

                                    If workStart <> "" Or workEnd <> "" Then
                                        ' 存储到字典
                                        dictKey = cardEmployeeKey & "|" & dateKey

                                        If cardDict.Exists(dictKey) Then
                                            ' 合并同一日期的时间
                                            Dim existingData As Variant
                                            existingData = cardDict(dictKey)
                                            Dim existingStart As String, existingEnd As String
                                            existingStart = existingData(0)
                                            existingEnd = existingData(1)

                                            ' 取更早的上班时间
                                            If workStart <> "" Then
                                                If existingStart = "" Or workStart < existingStart Then
                                                    existingStart = workStart
                                                End If
                                            End If

                                            ' 取更晚的下班时间
                                            If workEnd <> "" Then
                                                If existingEnd = "" Or workEnd > existingEnd Then
                                                    existingEnd = workEnd
                                                End If
                                            End If

                                            cardDict(dictKey) = Array(existingStart, existingEnd, existingData(2) + 1, employeeNo, department)
                                        Else
                                            cardDict(dictKey) = Array(workStart, workEnd, 1, employeeNo, department)
                                            totalRecords = totalRecords + 1
                                        End If
                                    End If
                                End If
                            End If
NextDayRow2:
                        Next dayRow
                    End If
                End If
            End If
        End If

        i = i + 25  ' 每个员工占25行

        ' 更新状态
        If i Mod 250 = 0 Then  ' 每处理10个员工更新一次状态
            UpdateStatus "正在读取考勤卡数据，已处理 " & i & " 行"
            DoEvents  ' 允许用户操作
        End If
    Loop

    ' ================= 读取员工信息 =================
    Dim staffLastRow As Long
    staffLastRow = wsStaff.Cells(wsStaff.rows.count, "B").End(xlUp).Row

    If staffLastRow < 5 Then
        UpdateStatus "人员信息表中无员工数据"
        GoTo CleanExit
    End If

    ' 读取员工信息到数组
    Dim staffData As Variant
    staffData = wsStaff.Range("B5:D" & staffLastRow).value

    ' ================= 批量写入数据 =================
    ' 使用Collection代替ArrayList
    Dim outputCollection As Collection
    Set outputCollection = New Collection

    ' 已存在记录去重（姓名+日期）
    Dim existingRecordDict As Object
    Set existingRecordDict = CreateObject("Scripting.Dictionary")

    If lastDataRow > 1 Then
        Dim existingArr As Variant
        existingArr = wsOutput.Range("F2:G" & lastDataRow).value

        Dim er As Long
        For er = 1 To UBound(existingArr, 1)
            Dim exNameKey As String
            exNameKey = NormalizeName(existingArr(er, 1) & "")
            If exNameKey <> "" And IsDate(existingArr(er, 2)) Then
                existingRecordDict(exNameKey & "|" & Format(CDate(existingArr(er, 2)), "yyyy-mm-dd")) = True
            End If
        Next er
    End If

    Dim totalEmployees As Long
    totalEmployees = staffLastRow - 4
    Dim processedCount As Long
    processedCount = 0

    UpdateStatus "生成考勤数据..."
    DoEvents  ' 允许用户操作

    ' 定义变量
    Dim j As Long
    Dim matchCount As Long
    Dim addedCount As Long
    Dim skippedExistingCount As Long
    matchCount = 0
    addedCount = 0
    skippedExistingCount = 0

    ' 检查是否有标记为布尔True的员工（D列）
    Dim hasValidEmployees As Boolean
    hasValidEmployees = False

    For j = 1 To UBound(staffData, 1)
        Dim statusMark As Variant
        statusMark = staffData(j, 3)
        If IsMarkedAsTrue(statusMark) Then
            hasValidEmployees = True
            Exit For
        End If
    Next j

    If Not hasValidEmployees Then
        UpdateStatus "没有标记为布尔True的员工（D列）"
        GoTo CleanExit
    End If

    For j = 1 To UBound(staffData, 1)
        processedCount = processedCount + 1
        If processedCount Mod 20 = 0 Then
            UpdateStatus "处理进度：" & processedCount & "/" & totalEmployees
            DoEvents  ' 允许用户操作
        End If

        employeeName = Trim(staffData(j, 1) & "")
        Dim employeeKey As String
        employeeKey = NormalizeName(employeeName)

        If employeeName = "" Then
            GoTo NextEmployee
        End If

        ' 检查员工状态
        statusMark = staffData(j, 3)
        If Not IsMarkedAsTrue(statusMark) Then
            GoTo NextEmployee
        End If

                ' 获取员工基础信息（工号/部门）
        Dim defaultEmployeeNo As String, defaultDepartment As String
        If empInfoDict.Exists(employeeKey) Then
            Dim empInfo As Variant
            empInfo = empInfoDict(employeeKey)
            defaultEmployeeNo = empInfo(0)
            defaultDepartment = empInfo(1)
        Else
            ' 考勤卡中不存在该员工，不生成清单
            GoTo NextEmployee
        End If

        ' 为每个日期生成记录
        Dim currentDate As Date
        For currentDate = startDate To endDate
            Dim dateKey2 As String
            dateKey2 = Format(currentDate, "yyyy-mm-dd")

            ' 查找打卡记录
            Dim searchKey As String
            searchKey = employeeKey & "|" & dateKey2
            If existingRecordDict.Exists(searchKey) Then
                skippedExistingCount = skippedExistingCount + 1
                GoTo NextCurrentDate
            End If

            Dim workIn As String, workOut As String
            Dim foundEmployeeNo As String, foundDepartment As String

            If cardDict.Exists(searchKey) Then
                matchCount = matchCount + 1
                Dim record As Variant
                record = cardDict(searchKey)
                workIn = record(0)
                workOut = record(1)
                foundEmployeeNo = record(3)
                foundDepartment = record(4)
                If foundEmployeeNo = "" Then foundEmployeeNo = defaultEmployeeNo
                If foundDepartment = "" Then foundDepartment = defaultDepartment
            Else
                workIn = ""
                workOut = ""
                foundEmployeeNo = defaultEmployeeNo
                foundDepartment = defaultDepartment
            End If

            ' 创建一行数据数组
            Dim rowData() As Variant
            ReDim rowData(1 To 19)

            ' 填充行数据
            rowData(1) = ""  ' A列：项目编号
            rowData(2) = ""  ' B列：项目名称
            rowData(3) = ""  ' C列：公司工号

            ' D列：考勤机卡号（写入工号）
            rowData(4) = defaultEmployeeNo

            ' E列：部门
            If foundDepartment <> "" Then
                rowData(5) = foundDepartment
            Else
                rowData(5) = ""
            End If

            ' F列：姓名
            rowData(6) = employeeName

            ' G列：日期
            rowData(7) = currentDate

            ' H列：上班时间 - 如果考勤卡中上班时间为空，则考勤清单中也保持为空
            If workIn <> "" Then
                On Error Resume Next
                Dim timeParts() As String
                timeParts = Split(workIn, ":")
                If UBound(timeParts) >= 1 Then
                    Dim hh As Integer, mm As Integer
                    hh = CInt(timeParts(0))
                    mm = CInt(timeParts(1))

                    ' 创建时间值
                    Dim timeIn As Date
                    timeIn = TimeSerial(hh, mm, 0)

                    ' 上班时间：日期+时间
                    rowData(8) = currentDate + timeIn
                Else
                    ' 时间格式不正确，留空
                    rowData(8) = ""
                End If
                On Error GoTo 0
            Else
                ' 考勤卡中上班时间为空，考勤清单中也保持为空
                rowData(8) = ""
            End If

            ' I列：下班时间 - 如果考勤卡中下班时间为空，则考勤清单中也保持为空
            If workOut <> "" Then
                On Error Resume Next
                timeParts = Split(workOut, ":")
                If UBound(timeParts) >= 1 Then
                    hh = CInt(timeParts(0))
                    mm = CInt(timeParts(1))

                    Dim timeOut As Date
                    timeOut = TimeSerial(hh, mm, 0)

                    Dim outTime As Date

                    ' 检查是否需要跨天：如果下班时间小于上班时间
                    If rowData(8) <> "" And IsDate(rowData(8)) Then
                        ' 有上班时间，需要比较
                        Dim checkInTime As Date
                        checkInTime = CDate(rowData(8))

                        If timeOut < timeValue(checkInTime) Then
                            ' 下班时间小于上班时间，说明跨天了
                            outTime = currentDate + 1 + timeOut
                        Else
                            ' 同一天
                            outTime = currentDate + timeOut
                        End If
                    Else
                        ' 没有上班时间，只保存下班时间（同一天）
                        outTime = currentDate + timeOut
                    End If

                    rowData(9) = outTime
                Else
                    ' 时间格式不正确，留空
                    rowData(9) = ""
                End If
                On Error GoTo 0
            Else
                ' 考勤卡中下班时间为空，考勤清单中也保持为空
                rowData(9) = ""
            End If

            ' J到N列：基准时间（直接使用原始值，包括格式）
            rowData(10) = STANDARD_START
            rowData(11) = STANDARD_END
            rowData(12) = MEAL_TIME
            rowData(13) = OVERTIME_START
            rowData(14) = NIGHT_MEAL_TIME

            ' O到S列：公式（将在写入后设置）
            Dim k As Long
            For k = 15 To 19
                rowData(k) = ""
            Next k

            ' 添加到集合
            outputCollection.Add rowData
            existingRecordDict(searchKey) = True
            addedCount = addedCount + 1
NextCurrentDate:
        Next currentDate

NextEmployee:
    Next j

        ' ================= 将集合转换为数组 =================
    If outputCollection.count = 0 Then
        UpdateStatus "没有可续接的新增记录（跳过已存在 " & skippedExistingCount & " 条）", False
        GoTo CleanExit
    End If

        If outputCollection.count > 0 Then
        UpdateStatus "写入数据到工作表..."
        DoEvents

        Dim writeLastRow As Long
        writeLastRow = lastDataRow

        ' 先构建员工顺序（按本次生成顺序）
        Dim empOrder As Collection
        Set empOrder = New Collection
        Dim empSeen As Object
        Set empSeen = CreateObject("Scripting.Dictionary")
        Dim empRowsDict As Object
        Set empRowsDict = CreateObject("Scripting.Dictionary")

        Dim rowIndex As Long
        For rowIndex = 1 To outputCollection.count
            Dim tempRow As Variant
            tempRow = outputCollection(rowIndex)
            Dim kEmp As String
            kEmp = NormalizeName(tempRow(6) & "")
            If kEmp <> "" Then
                If Not empSeen.Exists(kEmp) Then
                    empSeen(kEmp) = True
                    empOrder.Add kEmp
                    Set empRowsDict(kEmp) = New Collection
                End If
                empRowsDict(kEmp).Add tempRow
            End If
        Next rowIndex

        Dim deptSearchArrayLiteral As String
        deptSearchArrayLiteral = BuildDeptPrefixArrayLiteral()

        Dim oi As Long
        For oi = 1 To empOrder.count
            Dim currentEmpKey As String
            currentEmpKey = CStr(empOrder(oi))

            ' 收集该员工本次新增行（预分组，避免重复扫描）
            Dim empRows As Collection
            Set empRows = empRowsDict(currentEmpKey)

            If empRows.count > 0 Then
                Dim blockData() As Variant
                ReDim blockData(1 To empRows.count, 1 To 19)

                Dim bi As Long, colIndex As Long
                For bi = 1 To empRows.count
                    tempRow = empRows(bi)
                    For colIndex = 1 To 19
                        blockData(bi, colIndex) = tempRow(colIndex)
                    Next colIndex
                Next bi

                Dim startWriteRow As Long
                Dim lastEmpRow As Long
                lastEmpRow = FindLastEmployeeRowByKey(wsOutput, currentEmpKey, writeLastRow)

                If lastEmpRow > 0 Then
                    startWriteRow = lastEmpRow + 1
                    wsOutput.Rows(startWriteRow & ":" & startWriteRow + empRows.count - 1).Insert Shift:=xlDown
                Else
                    startWriteRow = writeLastRow + 1
                End If

                Dim targetRange As Range
                Set targetRange = wsOutput.Range("A" & startWriteRow).Resize(empRows.count, 19)
                targetRange.Value = blockData

                ApplyOutputBlockFormat wsOutput, wsStaff, targetRange, startWriteRow, startWriteRow + empRows.count - 1, deptSearchArrayLiteral

                writeLastRow = writeLastRow + empRows.count
            End If
        Next oi

        ' 显示完成消息
        Dim elapsedTime As Double
        elapsedTime = Timer - startTime

        Dim completionMsg As String
        completionMsg = "完成！新增 " & addedCount & " 条，跳过已存在 " & skippedExistingCount & " 条，用时 " & Format(elapsedTime, "0.00") & " 秒"

        UpdateStatus completionMsg
        shouldResetStatusBar = True  ' 仅完整成功时自动恢复状态栏
    End If

CleanExit:
    ' ================= 清理和恢复设置 =================
    恢复设置

    ' 清理对象 - 只在对象存在时才清理
    If Not cardDict Is Nothing Then
        Set cardDict = Nothing
    End If
    If Not empInfoDict Is Nothing Then
        Set empInfoDict = Nothing
    End If

    If Not outputCollection Is Nothing Then
        Set outputCollection = Nothing
    End If

    If shouldResetStatusBar Then
        Application.onTime _
            EarliestTime:=Now + timeValue("00:00:05"), _
            Procedure:="恢复状态栏"
    End If

    Exit Sub

ErrorHandler:
    ' ================= 错误处理 =================
    恢复设置

    ' 清理对象 - 只在对象存在时才清理
    If Not cardDict Is Nothing Then
        Set cardDict = Nothing
    End If
    If Not empInfoDict Is Nothing Then
        Set empInfoDict = Nothing
    End If

    If Not outputCollection Is Nothing Then
        Set outputCollection = Nothing
    End If

    UpdateStatus "错误：" & Err.Number & " - " & Err.Description
End Sub

' 函数说明：TryParseCardDateByRange
Private Function TryParseCardDateByRange(ByVal dateText As String, ByVal startDate As Date, _
    ByVal endDate As Date, ByRef parsedDate As Date) As Boolean
    Dim txt As String
    txt = Trim(dateText & "")
    If txt = "" Then Exit Function
    ' 优先使用 Excel/系统内置日期解析
    On Error Resume Next
    If IsDate(txt) Then
        parsedDate = CDate(txt)
        If Err.Number = 0 Then
            TryParseCardDateByRange = True
            On Error GoTo 0
            Exit Function
        End If
    End If
    Err.Clear
    On Error GoTo 0
    ' 若内置解析失败，再按 "MM.DD" 格式结合查询区间年份补全年份
    If Len(txt) <> 5 Or InStr(txt, ".") = 0 Then Exit Function
    Dim mdText As String
    mdText = Replace(txt, ".", "-")
    Dim y1 As Long, y2 As Long, y As Variant
    y1 = Year(startDate)
    y2 = Year(endDate)
    Dim yearsToTry As Collection
    Set yearsToTry = New Collection
    On Error Resume Next
    yearsToTry.Add y1, CStr(y1)
    yearsToTry.Add y2, CStr(y2)
    yearsToTry.Add y1 - 1, CStr(y1 - 1)
    yearsToTry.Add y2 + 1, CStr(y2 + 1)
    On Error GoTo 0
    Dim candidate As Date
    For Each y In yearsToTry
        On Error Resume Next
        candidate = CDate(CStr(y) & "-" & mdText)
        If Err.Number = 0 Then
            If candidate >= startDate And candidate <= endDate Then
                parsedDate = candidate
                TryParseCardDateByRange = True
                On Error GoTo 0
                Exit Function
            End If
        End If
        Err.Clear
        On Error GoTo 0
    Next y
    ' 兜底处理：未命中区间时按起始年份尝试一次解析
    On Error Resume Next
    parsedDate = CDate(CStr(y1) & "-" & mdText)
    If Err.Number = 0 Then TryParseCardDateByRange = True
    On Error GoTo 0
End Function
' 函数说明：FindLastEmployeeRowByKey
Private Function FindLastEmployeeRowByKey(ByVal ws As Worksheet, ByVal empKey As String, ByVal lastRow As Long) As Long
    Dim r As Long
    For r = lastRow To 2 Step -1
        If NormalizeName(ws.Cells(r, "F").Value & "") = empKey Then
            FindLastEmployeeRowByKey = r
            Exit Function
        End If
    Next r
End Function

' 过程说明：ApplyOutputBlockFormat
Private Sub ApplyOutputBlockFormat(ByVal wsOutput As Worksheet, ByVal wsStaff As Worksheet, _
    ByVal targetRange As Range, ByVal startRow As Long, ByVal lastRow As Long, ByVal deptSearchArrayLiteral As String)

    Dim formatStandardStart As String, formatStandardEnd As String
    Dim formatMealTime As String, formatOvertimeStart As String, formatNightMealTime As String

    formatStandardStart = wsStaff.Range("F2").NumberFormat
    formatStandardEnd = wsStaff.Range("G2").NumberFormat
    formatMealTime = wsStaff.Range("H2").NumberFormat
    formatOvertimeStart = wsStaff.Range("I2").NumberFormat
    formatNightMealTime = wsStaff.Range("J2").NumberFormat

    With targetRange
        .Font.Name = "宋体"
        .Font.Size = 10
        .Columns(10).NumberFormat = formatStandardStart
        .Columns(11).NumberFormat = formatStandardEnd
        .Columns(12).NumberFormat = formatMealTime
        .Columns(13).NumberFormat = formatOvertimeStart
        .Columns(14).NumberFormat = formatNightMealTime
    End With

    targetRange.Columns(7).NumberFormat = "yyyy-mm-dd"
    targetRange.Columns(8).NumberFormat = "hh:mm"
    targetRange.Columns(9).NumberFormat = "hh:mm"

    Dim c As Range
    For Each c In targetRange.Columns(8).Cells
        If c.Value = "" Then
            c.ClearFormats
            c.NumberFormat = "General"
        End If
    Next c
    For Each c In targetRange.Columns(9).Cells
        If c.Value = "" Then
            c.ClearFormats
            c.NumberFormat = "General"
        End If
    Next c

    wsOutput.Range("O" & startRow & ":O" & lastRow).Formula = _
        "=IF(OR(H" & startRow & "="""",I" & startRow & "=""""),"""",(MIN(I" & startRow & ",G" & startRow & "+K" & startRow & ")-MAX(H" & startRow & ",G" & startRow & "+J" & startRow & ")-L" & startRow & ")*24)"

    wsOutput.Range("P" & startRow & ":P" & lastRow).Formula = _
        "=IF(OR(G" & startRow & "="""",I" & startRow & "=""""),0,MAX(I" & startRow & "-(G" & startRow & "+M" & startRow & ")-N" & startRow & ",0)*24)"

    wsOutput.Range("Q" & startRow & ":Q" & lastRow).Formula = _
        "=IF(O" & startRow & "="""","""",INT(O" & startRow & ")+IF((O" & startRow & "-INT(O" & startRow & "))>=0.5,0.5,0))"

    wsOutput.Range("R" & startRow & ":R" & lastRow).Formula = _
        "=IF(OR(P" & startRow & "="""",ISNUMBER(SEARCH({""设计"",""工艺"",""项目"",""电气"",""上海""},E" & startRow & ")),E" & startRow & "=""""),"""",MROUND(P" & startRow & ",0.5))"
        ' "=IF(OR(P" & startRow & "="""",ISNUMBER(SEARCH(" & deptSearchArrayLiteral & ",E" & startRow & ")),E" & startRow & "=""""),"""",MROUND(P" & startRow & ",0.5))"
    wsOutput.Range("S" & startRow & ":S" & lastRow).Formula = _
        "=IF(OR(H" & startRow & "="""",J" & startRow & "=""""),"""",MAX((H" & startRow & "-(G" & startRow & "+J" & startRow & "))*24*60,0))"

    wsOutput.Range("O" & startRow & ":S" & lastRow).Calculate
End Sub
' 函数说明：BuildDeptPrefixArrayLiteral
Private Function BuildDeptPrefixArrayLiteral() As String
    On Error GoTo Fallback

    Dim wsEmp As Worksheet
    Set wsEmp = ThisWorkbook.Sheets("职员")

    Dim lastRow As Long
    lastRow = wsEmp.Cells(wsEmp.Rows.Count, "C").End(xlUp).Row
    If lastRow < 2 Then GoTo Fallback

    Dim arr As Variant
    arr = wsEmp.Range("C2:C" & lastRow).Value

    Dim dict As Object
    Set dict = CreateObject("Scripting.Dictionary")
    dict.CompareMode = vbTextCompare

    Dim i As Long
    For i = 1 To UBound(arr, 1)
        Dim deptText As String
        deptText = Trim(CStr(arr(i, 1) & ""))
        If deptText <> "" Then
            Dim token As String
            If Len(deptText) >= 2 Then
                token = Left(deptText, 2)
            Else
                token = deptText
            End If
            If token <> "" Then
                If Not dict.Exists(token) Then dict.Add token, True
            End If
        End If
    Next i

    If dict.Count = 0 Then GoTo Fallback

    Dim q As String
    q = Chr(34)

    Dim result As String
    result = "{"

    Dim k As Variant
    For Each k In dict.Keys
        If result <> "{" Then result = result & ","
        result = result & q & CStr(k) & q
    Next k

    result = result & "}"
    BuildDeptPrefixArrayLiteral = result
    Exit Function

Fallback:
    BuildDeptPrefixArrayLiteral = "{""__""}"
End Function
' ================= 辅助函数 =================
' 初始化Excel应用程序设置（优化性能）
' 更新状态栏并保持界面响应
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
' 过程说明：初始化设置
Private Sub 初始化设置()
    Application.ScreenUpdating = False  ' 关闭屏幕更新，加快执行速度
    Application.Calculation = xlCalculationManual  ' 手动计算模式，避免自动计算影响性能
    Application.EnableEvents = False  ' 禁用事件触发，防止连锁反应
End Sub

' 恢复Excel应用程序设置
Private Sub 恢复设置()
    Application.ScreenUpdating = True   ' 恢复屏幕更新
    Application.Calculation = xlCalculationAutomatic  ' 恢复自动计算
    Application.EnableEvents = True  ' 启用事件触发
End Sub

' 检查状态标记是否为真
Function IsMarkedAsTrue(mark As Variant) As Boolean
    IsMarkedAsTrue = (VarType(mark) = vbBoolean And mark = True)
End Function

' 规范化姓名键（去空格）
Function NormalizeName(ByVal nm As String) As String
    NormalizeName = Replace(Trim(nm & ""), " ", "")
End Function

' 将Excel时间数值转换为"HH:MM"字符串

Function FormatExcelTime(excelTime As Variant) As String
    If IsNumeric(excelTime) And excelTime <> 0 Then
        Dim timeValue As Double
        timeValue = CDbl(excelTime)

        ' 确保是合法的时间值（0到1之间）
        If timeValue >= 0 And timeValue < 1 Then
            Dim totalMinutes As Long
            totalMinutes = CLng(timeValue * 24 * 60)

            Dim h As Long, m As Long
            h = totalMinutes \ 60
            m = totalMinutes Mod 60

            FormatExcelTime = Format(h, "00") & ":" & Format(m, "00")
        Else
            FormatExcelTime = ""
        End If
    ElseIf VarType(excelTime) = vbString Then
        ' 如果是字符串格式的时间，直接返回
        Dim timeStr As String
        timeStr = Trim(CStr(excelTime))
        If Len(timeStr) = 5 And Mid(timeStr, 3, 1) = ":" Then
            ' 检查是否是有效的HH:MM格式
            Dim hh As Long, mm As Long
            hh = Val(Left(timeStr, 2))
            mm = Val(Right(timeStr, 2))
            If hh <= 23 And mm <= 59 Then
                FormatExcelTime = timeStr
            Else
                FormatExcelTime = ""
            End If
        Else
            FormatExcelTime = ""
        End If
    Else
        FormatExcelTime = ""
    End If
End Function

