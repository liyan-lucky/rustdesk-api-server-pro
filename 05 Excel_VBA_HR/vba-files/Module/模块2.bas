Attribute VB_Name = "模块2"
' 过程说明：获取打卡机记录
Sub 获取打卡机记录()
    ' =======================================================
    ' 功能：从考勤机数据生成员工考勤卡
    ' 输入：人员信息表（B2:C2为日期范围，B5起为员工姓名，C列为状态标记）
    ' 输出：考勤卡表（按模板格式生成每位员工的考勤记录）
    ' 修改：参考企业微信代码优化续接逻辑
    ' =======================================================

    ' ================= 性能设置 =================
    Application.ScreenUpdating = False
    Application.Calculation = xlCalculationManual
    Application.EnableEvents = False
    UpdateStatus "初始化中..."

    On Error GoTo ErrorHandler

    Dim startTime As Double
    startTime = Timer

    ' ================= 工作表对象定义 =================
    Dim wsStaff As Worksheet
    Dim wsAttendance As Worksheet
    Dim wsCard As Worksheet

    Set wsStaff = ThisWorkbook.Sheets("人员信息")
    Set wsAttendance = ThisWorkbook.Sheets("考勤机")
    Set wsCard = ThisWorkbook.Sheets("考勤卡")

    ' ================= 日期范围验证 =================
    If Not IsDate(wsStaff.Range("B2").value) Or Not IsDate(wsStaff.Range("C2").value) Then
        UpdateStatus "人员信息 B2 / C2 日期无效", False
        GoTo CleanUp
    End If

    Dim startDate As Date
    Dim endDate As Date

    startDate = CDate(wsStaff.Range("B2").value)
    endDate = CDate(wsStaff.Range("C2").value)

    Dim yearM As Long, monthM As Long
    yearM = Year(startDate)
    monthM = Month(startDate)

    ' ================= 模板参数定义 =================
    Const TPL_START_ROW As Long = 1
    Const BLOCK_ROWS As Long = 25
    Const DATE_ROW_BASE As Long = 8

    ' ================= 员工数据预读 =================
    Dim staffLastRow As Long
    staffLastRow = wsStaff.Cells(wsStaff.rows.count, "B").End(xlUp).Row

    If staffLastRow < 5 Then
        UpdateStatus "人员信息表中无员工数据", False
        GoTo CleanUp
    End If

    Dim staffArr As Variant
    staffArr = wsStaff.Range("B5:C" & staffLastRow).value

    Dim staffCount As Long
    staffCount = UBound(staffArr, 1)

    Dim staffEligible As Object
    Set staffEligible = CreateObject("Scripting.Dictionary")

    Dim si As Long
    For si = 1 To staffCount
        Dim sName As String
        sName = Trim(staffArr(si, 1) & "")
        If sName <> "" Then
            ' 仅处理布尔False的员工；非布尔值直接跳过
            If IsFalsyStatus(staffArr(si, 2)) Then
                staffEligible(NormalizeName(sName)) = True
            End If
        End If
    Next si

    ' ================= 读取考勤机数据到数组 =================
    UpdateStatus "正在加载考勤机数据到内存..."

    Dim attLastRow As Long
    attLastRow = wsAttendance.Cells(wsAttendance.rows.count, "A").End(xlUp).Row

    Dim attArr As Variant
    attArr = wsAttendance.Range("A2:G" & attLastRow).value

    ' ================= 预处理考勤机数据（按人名+日期分组） =================
    UpdateStatus "正在整理考勤机数据..."

    Dim attNameDict As Object
    Dim attRecords As Object
    Dim attEmpNoDict As Object
    Set attNameDict = CreateObject("Scripting.Dictionary")
    Set attRecords = CreateObject("Scripting.Dictionary")
    Set attEmpNoDict = CreateObject("Scripting.Dictionary")

    Dim totalAttRows As Long
    totalAttRows = UBound(attArr, 1)

    Dim ar As Long
    For ar = 1 To totalAttRows
        If ar Mod 1000 = 0 Then UpdateStatus "正在整理考勤机数据... " & ar & "/" & totalAttRows

        Dim nameKey As String
        nameKey = NormalizeName(attArr(ar, 2) & "")
        If nameKey <> "" Then
            If staffEligible.Exists(nameKey) Then
                If Not attNameDict.Exists(nameKey) Then attNameDict.Add nameKey, True

                Dim sourceEmpNo As String
                sourceEmpNo = NormalizeEmpNo(attArr(ar, 1))
                If sourceEmpNo <> "" Then
                    If Not attEmpNoDict.Exists(nameKey) Then
                        attEmpNoDict.Add nameKey, sourceEmpNo
                    End If
                End If

                Dim curDate As Date
                If TryParseDateValue(attArr(ar, 3), curDate) Then
                    If curDate >= startDate And curDate <= endDate Then
                        Dim dateKey As String
                        dateKey = Format(curDate, "yyyy-mm-dd")

                        Dim tin As String, tout As String
                        tin = Trim(CStr(attArr(ar, 4)))
                        tout = Trim(CStr(attArr(ar, 7)))

                        Dim empDict As Object
                        If Not attRecords.Exists(nameKey) Then
                            Set empDict = CreateObject("Scripting.Dictionary")
                            attRecords.Add nameKey, empDict
                        Else
                            Set empDict = attRecords(nameKey)
                        End If

                        If Not empDict.Exists(dateKey) Then
                            empDict.Add dateKey, Array(tin, tout)
                        Else
                            Dim t()
                            t = empDict(dateKey)
                            If tin <> "" And (t(0) = "" Or tin < t(0)) Then
                                t(0) = tin
                            End If
                            If tout <> "" And (t(1) = "" Or tout > t(1)) Then
                                t(1) = tout
                            End If
                            empDict(dateKey) = t
                        End If
                    End If
                End If
            End If
        End If
    Next ar

    ' ================= 查找考勤卡中最后一个已生成的员工 =================
    Dim r As Long
    Dim lastNameRow As Long

    lastNameRow = 0

    For r = wsCard.Cells(wsCard.rows.count, "E").End(xlUp).Row To 1 Step -1
        If InStr(wsCard.Cells(r, "E").MergeArea.Cells(1, 1).value, "姓名：") > 0 Then
            Dim oldName As String
            oldName = Trim(Replace(wsCard.Cells(r, "E").MergeArea.Cells(1, 1).value, "姓名：", ""))
            If oldName <> "" Then
                lastNameRow = r
                Exit For
            End If
        End If
    Next r

    ' ================= 确定起始位置 =================
    Dim cardRow As Long
    Dim needCopyTemplate As Boolean
    Dim empNoText As String

    If lastNameRow = 0 Then
        ' 第一次运行
        cardRow = TPL_START_ROW
        needCopyTemplate = False
    Else
        ' 参考企业微信代码：使用函数查找块起始行
        Dim lastBlockStartRow As Long
        lastBlockStartRow = FindLastBlockStartRow(wsCard, lastNameRow, BLOCK_ROWS, TPL_START_ROW)

        ' 计算下一个员工的起始行
        cardRow = lastBlockStartRow + BLOCK_ROWS

        ' 检查该行是否为空，如果不为空则继续向下查找空行
        Do While cardRow <= wsCard.rows.count And _
               Not IsEmptyRow(wsCard, cardRow, 5)  ' 检查前5列是否为空
            cardRow = cardRow + 1
        Loop

        needCopyTemplate = True
    End If

    ' ================= 建立已生成员工字典 =================
    Dim existsDict As Object
    Set existsDict = CreateObject("Scripting.Dictionary")

    For r = 1 To wsCard.Cells(wsCard.rows.count, "E").End(xlUp).Row
        If InStr(wsCard.Cells(r, "E").MergeArea.Cells(1, 1).value, "姓名：") > 0 Then
            Dim oldName2 As String
            oldName2 = Trim(Replace(wsCard.Cells(r, "E").MergeArea.Cells(1, 1).value, "姓名：", ""))
            If oldName2 <> "" Then existsDict(NormalizeName(oldName2)) = True
        End If
    Next r

    ' ================= 员工循环处理 =================
    Dim i As Long
    Dim processedCount As Long
    processedCount = 0
    Dim templateCopied As Boolean
    templateCopied = False

    ' 记录是否有新员工需要处理
    Dim hasNewEmployee As Boolean
    hasNewEmployee = False

    ' 记录是否所有员工都不在考勤机中
    Dim allEmployeesNotFound As Boolean
    allEmployeesNotFound = True

    For i = 1 To staffCount
        Dim empName As String
        empName = Trim(staffArr(i, 1) & "")
        If empName = "" Then
            UpdateStatus "第" & (i + 4) & "行员工姓名为空，处理结束", False
            GoTo CleanUp
        End If

        ' ========= 检查人员状态标记 =========
        Dim statusMark As Variant
        statusMark = staffArr(i, 2)

        ' 如果状态是True（企业微信员工），跳过此员工
        If IsTruthyStatus(statusMark) Then
            UpdateStatus "跳过员工：" & empName & "（企业微信员工）"
            GoTo ContinueNextEmployee
        End If

        Dim empKey As String
        empKey = NormalizeName(empName)

        ' 跳过已生成的员工
        If existsDict.Exists(empKey) Then
            UpdateStatus "跳过已存在员工：" & empName
            GoTo ContinueNextEmployee
        End If

        ' ========= 检查员工是否在考勤机中存在 =========
        If Not attNameDict.Exists(empKey) Then
            UpdateStatus "跳过员工：" & empName & "（考勤机中不存在）"
            GoTo ContinueNextEmployee
        End If

        ' 找到需要处理的新员工
        hasNewEmployee = True
        allEmployeesNotFound = False

        UpdateStatus "正在处理：" & empName & _
            "（" & i & "/" & staffCount & "）"

        ' ========= 拷贝模板（如果需要） =========
        If needCopyTemplate Then
            ' 参考企业微信代码：使用函数查找模板源行
            Dim copySourceRow As Long

            If cardRow <= TPL_START_ROW + BLOCK_ROWS Then
                ' 前两个员工使用原始模板
                copySourceRow = TPL_START_ROW
            Else
                ' 查找上一个完整的员工块
                copySourceRow = FindTemplateSourceRow(wsCard, cardRow, BLOCK_ROWS, TPL_START_ROW)
                If copySourceRow < TPL_START_ROW Then
                    copySourceRow = TPL_START_ROW
                End If
            End If

            ' 确保源行是完整块的起始行
            If (copySourceRow - TPL_START_ROW) Mod BLOCK_ROWS <> 0 Then
                copySourceRow = TPL_START_ROW + Int((copySourceRow - TPL_START_ROW) / BLOCK_ROWS) * BLOCK_ROWS
            End If

            ' 复制模板
            If copySourceRow >= TPL_START_ROW And copySourceRow <= wsCard.rows.count Then
                wsCard.rows(copySourceRow & ":" & copySourceRow + BLOCK_ROWS - 1).Copy
                wsCard.rows(cardRow).Insert Shift:=xlDown
                Application.CutCopyMode = False
                templateCopied = True

                ClearTemplateBlock wsCard, cardRow, DATE_ROW_BASE
            Else
                copySourceRow = TPL_START_ROW
                wsCard.rows(copySourceRow & ":" & copySourceRow + BLOCK_ROWS - 1).Copy
                wsCard.rows(cardRow).Insert Shift:=xlDown
                Application.CutCopyMode = False
                templateCopied = True
                ClearTemplateBlock wsCard, cardRow, DATE_ROW_BASE
            End If
        End If

        ' ========= 检查是否有日期范围内的记录 =========
        Dim empRecords As Object
        If attRecords.Exists(empKey) Then
            Set empRecords = attRecords(empKey)
        Else
            Set empRecords = Nothing
        End If

        If empRecords Is Nothing Then
            ' 删除已拷贝的模板（如果有）
            If templateCopied Then
                wsCard.rows(cardRow & ":" & cardRow + BLOCK_ROWS - 1).Delete Shift:=xlUp
                templateCopied = False
            End If
            UpdateStatus empName & " 不在日期范围内", False
            GoTo CleanUp
        End If

        ' ========= 有匹配记录，继续写入信息 =========
        ' 填写表头信息
        SafeSet wsCard.Cells(cardRow + 3, "E"), "姓名：" & empName
        empNoText = ""
        If attEmpNoDict.Exists(empKey) Then empNoText = attEmpNoDict(empKey)
        SafeSet wsCard.Cells(cardRow + 3, "H"), "工号：" & empNoText
        SafeSet wsCard.Cells(cardRow + 3, "M"), _
            "日期:" & Format(startDate, "yy.mm.dd") & "～" & Format(endDate, "yy.mm.dd")

        ' ========= 填写日期轴 =========
        Dim d As Long, rr As Long
        For d = 1 To 16
            rr = cardRow + DATE_ROW_BASE + d - 1

            ' 填写上半月日期（1-16日）
            curDate = DateSerial(yearM, monthM, d)
            If Month(curDate) = monthM Then
                SafeSet wsCard.Cells(rr, "A"), curDate, "mm.dd"
                SafeSet wsCard.Cells(rr, "B"), CNWeek1(curDate)
            Else
                SafeSet wsCard.Cells(rr, "A"), ""
                SafeSet wsCard.Cells(rr, "B"), ""
            End If

            ' 填写下半月日期（17-31日）
            curDate = DateSerial(yearM, monthM, d + 16)
            If Month(curDate) = monthM Then
                SafeSet wsCard.Cells(rr, "I"), curDate, "mm.dd"
                SafeSet wsCard.Cells(rr, "J"), CNWeek1(curDate)
            Else
                SafeSet wsCard.Cells(rr, "I"), ""
                SafeSet wsCard.Cells(rr, "J"), ""
            End If
        Next d

        ' ========= 写入上班和下班时间到考勤卡 =========
        Dim dayNo As Long, rowW As Long, cIn As String, cOut As String
        Dim k As String
        For dayNo = 1 To 31
            curDate = DateSerial(yearM, monthM, dayNo)

            If Month(curDate) <> monthM Then GoTo SkipDay
            If curDate < startDate Or curDate > endDate Then GoTo SkipDay

            If dayNo <= 16 Then
                rowW = cardRow + DATE_ROW_BASE + dayNo - 1
                cIn = "C"
                cOut = "H"
            Else
                rowW = cardRow + DATE_ROW_BASE + dayNo - 17
                cIn = "K"
                cOut = "P"
            End If

            k = Format(curDate, "yyyy-mm-dd")
            If empRecords.Exists(k) Then
                Dim dayRecord As Variant
                dayRecord = empRecords(k)
                If dayRecord(0) <> "" Then
                    SafeSet wsCard.Cells(rowW, cIn), dayRecord(0), "hh:mm"
                Else
                    SafeSet wsCard.Cells(rowW, cIn), ""
                End If

                If dayRecord(1) <> "" Then
                    SafeSet wsCard.Cells(rowW, cOut), dayRecord(1), "hh:mm"
                Else
                    SafeSet wsCard.Cells(rowW, cOut), ""
                End If
            Else
                ' 没有记录，写入空白
                SafeSet wsCard.Cells(rowW, cIn), ""
                SafeSet wsCard.Cells(rowW, cOut), ""
            End If

SkipDay:
        Next dayNo

        ' ========= 更新变量，准备下一个员工 =========
        ApplyCardBlockFont wsCard, cardRow, BLOCK_ROWS
        existsDict(empKey) = True
        processedCount = processedCount + 1
        UpdateStatus "完成：" & empName & "（" & processedCount & "/" & staffCount & "）"

        ' 工号来源于考勤机A列，不做自增
        cardRow = cardRow + BLOCK_ROWS
        needCopyTemplate = True
        templateCopied = False

ContinueNextEmployee:
    Next i

    ' ================= 检查是否有新员工需要处理 =================
    If Not hasNewEmployee Then
        UpdateStatus "所有员工考勤卡已存在，无需生成新记录", False
        processedCount = -1
    ElseIf allEmployeesNotFound Then
        UpdateStatus "所有员工都不在考勤机中", False
        processedCount = -2
    End If

CleanUp:
    ' ================= 清理和恢复设置 =================
    Dim elapsedTime As Double
    elapsedTime = Timer - startTime

    Application.ScreenUpdating = True
    Application.Calculation = xlCalculationAutomatic
    Application.EnableEvents = True

    If i > staffCount Then
        If processedCount >= 0 Then
            Dim completionMessage As String
            If processedCount > 0 Then
                completionMessage = "处理完成，用时 " & Format(elapsedTime, "0.00") & " 秒，处理了 " & processedCount & " 个员工"
            Else
                completionMessage = "所有员工都已存在，无需处理"
            End If
            UpdateStatus completionMessage, False
        Else
            ' 异常结束：所有员工已存在或都不在考勤机中
        End If

        ' 仅正常完成时5秒后恢复状态栏；异常/取消保持当前文本
        If processedCount >= 0 Then
            Application.onTime _
                EarliestTime:=Now + timeValue("00:00:05"), _
                Procedure:="恢复状态栏"
        End If
    Else
        UpdateStatus Application.StatusBar & "，用时 " & Format(elapsedTime, "0.00") & " 秒", False
    End If

    Exit Sub

ErrorHandler:
    UpdateStatus "运行错误：" & Err.Description, False
    Resume CleanUp
End Sub

' ================= 新增辅助函数（参考企业微信代码） =================
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

' 函数说明：NormalizeName
Private Function NormalizeName(ByVal nm As String) As String
    NormalizeName = Replace(Trim(nm & ""), " ", "")
End Function

' 函数说明：IsTruthyStatus
Private Function IsTruthyStatus(ByVal statusMark As Variant) As Boolean
    IsTruthyStatus = (VarType(statusMark) = vbBoolean And statusMark = True)
End Function

' 函数说明：IsFalsyStatus
Private Function IsFalsyStatus(ByVal statusMark As Variant) As Boolean
    IsFalsyStatus = (VarType(statusMark) = vbBoolean And statusMark = False)
End Function

' 函数说明：NormalizeEmpNo
Private Function NormalizeEmpNo(ByVal v As Variant) As String
    If IsEmpty(v) Or IsNull(v) Then Exit Function

    If IsNumeric(v) Then
        NormalizeEmpNo = Format(CLng(v), "00000")
    Else
        NormalizeEmpNo = Trim(CStr(v))
    End If
End Function
' 函数说明：TryParseDateValue
Private Function TryParseDateValue(ByVal v As Variant, ByRef d As Date) As Boolean
    On Error Resume Next
    If IsDate(v) Then d = CDate(v)
    TryParseDateValue = (Err.Number = 0 And d <> 0)
    Err.Clear
    On Error GoTo 0
End Function

' 过程说明：ClearTemplateBlock
Private Sub ClearTemplateBlock(ByVal ws As Worksheet, ByVal cardRow As Long, ByVal dateRowBase As Long)
    Dim dr As Long
    For dr = 0 To 15
        SafeSet ws.Cells(cardRow + dateRowBase + dr, "C"), ""
        SafeSet ws.Cells(cardRow + dateRowBase + dr, "H"), ""
        SafeSet ws.Cells(cardRow + dateRowBase + dr, "K"), ""
        SafeSet ws.Cells(cardRow + dateRowBase + dr, "P"), ""
    Next dr

    SafeSet ws.Cells(cardRow + 3, "E"), ""
    SafeSet ws.Cells(cardRow + 3, "H"), ""
    SafeSet ws.Cells(cardRow + 3, "M"), ""
End Sub

' 功能：查找最后一个员工块的起始行
Private Function FindLastBlockStartRow(ws As Worksheet, lastNameRow As Long, blockRows As Long, startRow As Long) As Long
    Dim rowNum As Long
    rowNum = lastNameRow

    While rowNum > startRow And (rowNum - startRow) Mod blockRows <> 0
        rowNum = rowNum - 1
    Wend

    If rowNum < startRow Then
        FindLastBlockStartRow = startRow
    Else
        FindLastBlockStartRow = rowNum
    End If
End Function

' 功能：查找模板源行
Private Function FindTemplateSourceRow(ws As Worksheet, currentRow As Long, blockRows As Long, startRow As Long) As Long
    Dim rowNum As Long
    rowNum = currentRow - 1

    While rowNum >= startRow
        If (rowNum - startRow) Mod blockRows = 0 Then
            Dim cellValue As String
            cellValue = ws.Cells(rowNum + 3, "E").MergeArea.Cells(1, 1).value & ""
            If InStr(cellValue, "姓名：") > 0 Then
                FindTemplateSourceRow = rowNum
                Exit Function
            End If
        End If
        rowNum = rowNum - 1
    Wend

    FindTemplateSourceRow = startRow
End Function

' 功能：检查行是否为空
Private Function IsEmptyRow(ws As Worksheet, rowNum As Long, Optional colCount As Long = 5) As Boolean
    Dim col As Long
    IsEmptyRow = True

    For col = 1 To colCount
        If Trim(ws.Cells(rowNum, col).value & "") <> "" Then
            IsEmptyRow = False
            Exit Function
        End If
    Next col
End Function

' ================= 原有工具函数 =================
Private Sub ApplyCardBlockFont(ByVal ws As Worksheet, ByVal startRow As Long, ByVal blockRows As Long)
    With ws.Rows(startRow & ":" & startRow + blockRows - 1).Font
        .Name = "宋体"
        .Size = 10
    End With
End Sub
' 过程说明：SafeSet
Private Sub SafeSet(c As Range, v As Variant, Optional fmt As String = "")
    If c.MergeCells Then
        c.MergeArea.value = v
        If fmt <> "" Then c.MergeArea.NumberFormat = fmt
    Else
        c.value = v
        If fmt <> "" Then c.NumberFormat = fmt
    End If
End Sub

' 函数说明：CNWeek1
Private Function CNWeek1(d As Date) As String
    CNWeek1 = Array("", "日", "一", "二", "三", "四", "五", "六")(Weekday(d, vbSunday))
End Function

