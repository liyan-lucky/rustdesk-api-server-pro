Attribute VB_Name = "模块1"
' 过程说明：获取企业微信记录
Sub 获取企业微信记录()
    ' =======================================================
    ' 功能：从企业微信数据生成员工考勤卡
    ' 输入：人员信息表（B2:C2为日期范围，B5起为员工姓名，C列为状态标记）
    ' 输出：考勤卡表（按模板格式生成每位员工的考勤记录）
    ' =======================================================

    ' ================= 性能设置 =================
    Application.ScreenUpdating = False      ' 关闭屏幕更新，提高运行速度
    Application.Calculation = xlCalculationManual  ' 手动计算模式
    Application.EnableEvents = False        ' 禁用事件触发
    UpdateStatus "初始化中..."              ' 设置状态栏提示

    On Error GoTo ErrorHandler             ' 错误处理跳转

    Dim startTime As Double
    startTime = Timer                       ' 记录开始时间，用于计算运行时长

    ' ================= 工作表对象定义 =================
    Dim wsStaff As Worksheet        ' 人员信息表
    Dim wsWX As Worksheet           ' 企业微信数据表
    Dim wsCard As Worksheet         ' 考勤卡输出表

    Set wsStaff = ThisWorkbook.Sheets("人员信息")
    Set wsWX = ThisWorkbook.Sheets("企业微信")
    Set wsCard = ThisWorkbook.Sheets("考勤卡")

    ' ================= 日期范围验证 =================
    ' 检查人员信息表中B2和C2单元格是否为有效日期
    If Not IsDate(wsStaff.Range("B2").value) Or Not IsDate(wsStaff.Range("C2").value) Then
        UpdateStatus "人员信息 B2 / C2 日期无效", False
        GoTo CleanUp
    End If

    Dim startDate As Date      ' 考勤开始日期
    Dim endDate As Date        ' 考勤结束日期

    startDate = CDate(wsStaff.Range("B2").value)
    endDate = CDate(wsStaff.Range("C2").value)

    Dim yearM As Long, monthM As Long
    yearM = Year(startDate)
    monthM = Month(startDate)

    ' ================= 计算当月首末日期 =================
    Dim monthFirst As Date, monthLast As Date
    monthFirst = DateSerial(Year(startDate), Month(startDate), 1)  ' 当月第一天
    monthLast = DateSerial(Year(monthFirst), Month(monthFirst) + 1, 0)  ' 当月最后一天
    ' ====================================================

    ' ================= 模板参数定义 =================
    Const TPL_START_ROW As Long = 1      ' 模板起始行
    Const BLOCK_ROWS As Long = 25        ' 每个员工考勤卡占用的行数
    Const DATE_ROW_BASE As Long = 8      ' 日期行的基准行号

    ' ================= 员工数据预读 =================
    Dim staffLastRow As Long
    staffLastRow = wsStaff.Cells(wsStaff.rows.count, "B").End(xlUp).Row  ' 人员信息最后行

    ' 检查是否有员工数据
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
            If IsTruthyStatus(staffArr(si, 2)) Then
                staffEligible(NormalizeName(sName)) = True
            End If
        End If
    Next si

    ' ================= 读取企业微信数据到数组（提高处理速度） =================
    UpdateStatus "正在加载企业微信数据到内存..."

    Dim wxLastRow As Long
    wxLastRow = wsWX.Cells(wsWX.rows.count, "A").End(xlUp).Row  ' 获取企业微信数据最后行

    ' 企业微信表列说明：
    ' A列: 日期 (格式: 2025/12/31 星期三)
    ' B列: 姓名
    ' C列: 工号
    ' I列: 上班时间
    ' J列: 下班时间
    Dim wxArr As Variant
    wxArr = wsWX.Range("A2:J" & wxLastRow).value  ' 将数据读入数组，提高处理速度

    ' ================= 预处理企业微信数据（按人名+日期分组） =================
    UpdateStatus "正在整理企业微信数据..."

    Dim wxNameDict As Object
    Dim wxRecords As Object
    Dim wxEmpNoDict As Object
    Set wxNameDict = CreateObject("Scripting.Dictionary")  ' 存储企业微信中的姓名
    Set wxRecords = CreateObject("Scripting.Dictionary")   ' 存储姓名对应的日期记录
    Set wxEmpNoDict = CreateObject("Scripting.Dictionary") ' 存储姓名对应工号

    Dim totalWxRows As Long
    totalWxRows = UBound(wxArr, 1)

    Dim ar As Long
    For ar = 1 To totalWxRows
        If ar Mod 1000 = 0 Then UpdateStatus "正在整理企业微信数据... " & ar & "/" & totalWxRows

        Dim nameKey As String
        nameKey = NormalizeName(wxArr(ar, 2) & "")
        If nameKey <> "" Then
            If staffEligible.Exists(nameKey) Then
                If Not wxNameDict.Exists(nameKey) Then wxNameDict.Add nameKey, True

                Dim sourceEmpNo As String
                sourceEmpNo = NormalizeWxEmpNo(wxArr(ar, 3))
                If sourceEmpNo <> "" Then
                    If Not wxEmpNoDict.Exists(nameKey) Then
                        wxEmpNoDict.Add nameKey, sourceEmpNo
                    End If
                End If

                Dim dtText As String
                dtText = Trim(wxArr(ar, 1) & "")
                If dtText <> "" Then
                    Dim recordDate As Date
                    If TryParseRecordDate(dtText, recordDate) Then
                        If recordDate >= startDate And recordDate <= endDate Then
                            Dim dateKey As String
                            dateKey = Format(recordDate, "yyyy-mm-dd")

                            Dim timeIn As String, timeOut As String
                            timeIn = NormalizeTimeText(ExtractTimeFromText(CStr(wxArr(ar, 9))))   ' I列：上班时间
                            timeOut = NormalizeTimeText(ExtractTimeFromText(CStr(wxArr(ar, 10)))) ' J列：下班时间

                            Dim empDict As Object
                            If Not wxRecords.Exists(nameKey) Then
                                Set empDict = CreateObject("Scripting.Dictionary")
                                wxRecords.Add nameKey, empDict
                            Else
                                Set empDict = wxRecords(nameKey)
                            End If

                            If Not empDict.Exists(dateKey) Then
                                empDict.Add dateKey, Array(timeIn, timeOut)
                            Else
                                Dim t()
                                t = empDict(dateKey)
                                If timeIn <> "" And (t(0) = "" Or timeIn < t(0)) Then
                                    t(0) = timeIn
                                End If
                                If timeOut <> "" And (t(1) = "" Or timeOut > t(1)) Then
                                    t(1) = timeOut
                                End If
                                empDict(dateKey) = t
                            End If
                        End If
                    End If
                End If
            End If
        End If
    Next ar

    ' ================= 员工循环处理 =================

    Dim r As Long
    Dim lastNameRow As Long

    lastNameRow = 0

    ' ================= 查找考勤卡中最后一个已生成的员工 =================
    ' 从下往上查找最后一个包含"姓名："的合并单元格
    For r = wsCard.Cells(wsCard.rows.count, "E").End(xlUp).Row To 1 Step -1
        If InStr(wsCard.Cells(r, "E").MergeArea.Cells(1, 1).value, "姓名：") > 0 Then
            Dim oldName As String
            oldName = Trim(Replace(wsCard.Cells(r, "E").MergeArea.Cells(1, 1).value, "姓名：", ""))
            If oldName <> "" Then
                lastNameRow = r  ' 记录行号
                Exit For
            End If
        End If
    Next r

    ' ================= 确定起始位置 =================
    Dim cardRow As Long        ' 当前写入行
    Dim needCopyTemplate As Boolean  ' 是否需要复制模板
    Dim empNoText As String

    If lastNameRow = 0 Then
        ' 第一次运行，从模板起始行开始
        cardRow = TPL_START_ROW
        needCopyTemplate = False  ' 第一行不需要复制
    Else
        ' 修正部分：正确计算下一个员工的起始行
        ' 先找到最后一个员工块的第一行
        Dim lastBlockStartRow As Long
        lastBlockStartRow = FindLastBlockStartRow(wsCard, lastNameRow, BLOCK_ROWS, TPL_START_ROW)

        ' 计算下一个员工的起始行
        cardRow = lastBlockStartRow + BLOCK_ROWS

        ' 检查该行是否为空，如果不为空则继续向下查找空行
        Do While cardRow <= wsCard.rows.count And _
               Not IsEmptyRow(wsCard, cardRow, 1)
            cardRow = cardRow + 1
        Loop

        needCopyTemplate = True   ' 需要复制模板
    End If

    ' ================= 建立已生成员工字典（避免重复生成） =================
    Dim existsDict As Object
    Set existsDict = CreateObject("Scripting.Dictionary")  ' 存储已生成的员工姓名

    ' 扫描考勤卡表中所有已生成的员工
    For r = 1 To wsCard.Cells(wsCard.rows.count, "E").End(xlUp).Row
        If InStr(wsCard.Cells(r, "E").MergeArea.Cells(1, 1).value, "姓名：") > 0 Then
            Dim oldName2 As String
            oldName2 = Trim(Replace(wsCard.Cells(r, "E").MergeArea.Cells(1, 1).value, "姓名：", ""))
            If oldName2 <> "" Then existsDict(NormalizeName(oldName2)) = True  ' 去除空格后添加到字典
        End If
    Next r

    Dim i As Long
    Dim processedCount As Long
    processedCount = 0
    Dim templateCopied As Boolean  ' 标记是否已拷贝模板
    templateCopied = False

    ' ================= 新增：记录是否有新的员工需要处理 =================
    Dim hasNewEmployee As Boolean
    hasNewEmployee = False  ' 初始化为没有新员工

    For i = 1 To staffCount  ' 从第5行开始（跳过表头）
        Dim empName As String
        empName = Trim(staffArr(i, 1) & "")  ' 获取员工姓名
        If empName = "" Then
            UpdateStatus "第" & (i + 4) & "行员工姓名为空，处理结束", False
            GoTo CleanUp
        End If

        ' ========= 检查人员状态标记（C列） =========
        ' 检查C列的状态标记，如果是True才从企业微信表处理
        Dim statusMark As Variant
        statusMark = staffArr(i, 2)

        ' 判断状态是否为TRUE（仅接受布尔TRUE，非布尔值直接跳过）
        Dim isWXEmployee As Boolean
        isWXEmployee = IsTruthyStatus(statusMark)

        ' 如果状态不是True，跳过此员工
        If Not isWXEmployee Then
            UpdateStatus "跳过员工：" & empName & "（非企业微信员工）"
            GoTo NextEmployee  ' 跳过此员工，继续下一个
        End If

        Dim empKey As String
        empKey = NormalizeName(empName)  ' 去除空格作为键值

        ' 跳过已生成的员工
        If existsDict.Exists(empKey) Then
            UpdateStatus "跳过已存在员工：" & empName
            GoTo NextEmployee
        End If

        ' ========= 检查员工是否在企业微信中存在 =========
        If Not wxNameDict.Exists(empKey) Then
            ' 员工在企业微信中不存在，跳过继续处理下一个
            UpdateStatus "跳过员工：" & empName & "（企业微信中不存在）"
            GoTo NextEmployee  ' 跳过此员工，继续下一个
        End If

        ' ========= 新增：标记有新员工需要处理 =========
        hasNewEmployee = True

        ' 更新状态栏显示进度
        UpdateStatus "正在处理：" & empName & _
            "（" & i & "/" & staffCount & "）"

        ' ========= 拷贝模板（如果需要） =========
        If needCopyTemplate Then
            ' 修正部分：更可靠的模板复制逻辑
            Dim copySourceRow As Long

            ' 找到可用的模板源行（从已存在的员工块中复制）
            If cardRow <= TPL_START_ROW + BLOCK_ROWS Then
                ' 如果是前两个员工，使用原始模板
                copySourceRow = TPL_START_ROW
            Else
                ' 查找上一个完整的员工块
                ' 从当前行向上查找最近的完整块
                copySourceRow = FindTemplateSourceRow(wsCard, cardRow, BLOCK_ROWS, TPL_START_ROW)

                If copySourceRow < TPL_START_ROW Then
                    ' 如果找不到合适的源行，使用原始模板
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
                templateCopied = True  ' 标记已拷贝模板

                ClearTemplateBlock wsCard, cardRow, DATE_ROW_BASE
            Else
                ' 无法找到合适的模板源，使用原始模板
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
        If wxRecords.Exists(empKey) Then
            Set empRecords = wxRecords(empKey)
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
        SafeSet wsCard.Cells(cardRow + 3, "E"), "姓名：" & empName  ' 姓名
        empNoText = ""
        If wxEmpNoDict.Exists(empKey) Then empNoText = wxEmpNoDict(empKey)
        SafeSet wsCard.Cells(cardRow + 3, "H"), "工号：" & empNoText
        SafeSet wsCard.Cells(cardRow + 3, "M"), _
            "日期:" & Format(startDate, "yy.mm.dd") & "～" & Format(endDate, "yy.mm.dd")  ' 考勤日期范围

        ' ========= 填写日期轴 =========
        Dim d As Long, rr As Long
        For d = 1 To 16  ' 每月最多31天，分两列显示（1-16，17-31）
            rr = cardRow + DATE_ROW_BASE + d - 1  ' 计算行号

            ' 填写上半月日期（1-16日）
            Dim curDate As Date
            curDate = DateSerial(yearM, monthM, d)
            If Month(curDate) = monthM Then
                SafeSet wsCard.Cells(rr, "A"), curDate, "mm.dd"   ' 日期格式：月.日
                SafeSet wsCard.Cells(rr, "B"), CNWeek1(curDate)   ' 星期几
            Else
                SafeSet wsCard.Cells(rr, "A"), ""   ' 非本月日期留空
                SafeSet wsCard.Cells(rr, "B"), ""
            End If

            ' 填写下半月日期（17-31日）
            curDate = DateSerial(yearM, monthM, d + 16)
            If Month(curDate) = monthM Then
                SafeSet wsCard.Cells(rr, "I"), curDate, "mm.dd"   ' 日期格式：月.日
                SafeSet wsCard.Cells(rr, "J"), CNWeek1(curDate)   ' 星期几
            Else
                SafeSet wsCard.Cells(rr, "I"), ""   ' 非本月日期留空
                SafeSet wsCard.Cells(rr, "J"), ""
            End If
        Next d

        ' ========= 写入上班和下班时间到考勤卡 =========
        Dim dayNo As Long, rowW As Long, cIn As String, cOut As String
        For dayNo = 1 To 31  ' 遍历31天
            curDate = DateSerial(yearM, monthM, dayNo)

            ' 跳过非本月日期和不在考勤范围内的日期
            If Month(curDate) <> monthM Then GoTo SkipDay
            If curDate < startDate Or curDate > endDate Then GoTo SkipDay

            ' 根据日期判断写入位置（上半月或下半月）
            If dayNo <= 16 Then
                rowW = cardRow + DATE_ROW_BASE + dayNo - 1  ' 上半月行号
                cIn = "C"   ' 上半月上班时间列
                cOut = "H"  ' 上半月下班时间列
            Else
                rowW = cardRow + DATE_ROW_BASE + dayNo - 17  ' 下半月行号
                cIn = "K"   ' 下半月上班时间列
                cOut = "P"  ' 下半月下班时间列
            End If

            Dim k As String
            k = Format(curDate, "yyyy-mm-dd")
            If empRecords.Exists(k) Then
                Dim dayRecord As Variant
                dayRecord = empRecords(k)
                ' 写入上班时间
                If dayRecord(0) <> "" Then
                    SafeSet wsCard.Cells(rowW, cIn), dayRecord(0), "hh:mm"
                Else
                    SafeSet wsCard.Cells(rowW, cIn), ""
                End If

                ' 写入下班时间
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

        ' 更新字典和变量，准备下一个员工
        ApplyCardBlockFont wsCard, cardRow, BLOCK_ROWS
        existsDict(empKey) = True  ' 标记该员工已处理
        cardRow = cardRow + BLOCK_ROWS  ' 行号递增
        needCopyTemplate = True    ' 后续都需要复制模板
        processedCount = processedCount + 1
        templateCopied = False     ' 重置模板拷贝标记
        UpdateStatus "完成：" & empName & "（" & processedCount & "/" & staffCount & "）"

NextEmployee:
    Next i

    ' ================= 新增：检查是否有新员工需要处理 =================
    ' 如果循环完成且hasNewEmployee为False，说明所有员工都已存在
    If Not hasNewEmployee Then
        UpdateStatus "所有员工考勤卡已存在，无需生成新记录", False
        ' 设置异常结束标志
        processedCount = -1  ' 使用-1表示异常结束
    End If

CleanUp:
    ' ================= 清理和恢复设置 =================
    Dim elapsedTime As Double
    elapsedTime = Timer - startTime  ' 计算运行时间

    ' 恢复应用程序设置
    Application.ScreenUpdating = True
    Application.Calculation = xlCalculationAutomatic
    Application.EnableEvents = True

    ' 检查是否正常完成所有员工处理
    If i > staffCount Then  ' 循环完成
        If processedCount >= 0 Then  ' 正常处理完成（有处理新员工）
            Dim completionMessage As String
            If processedCount > 0 Then
                completionMessage = "处理完成，用时 " & Format(elapsedTime, "0.00") & " 秒，处理了 " & processedCount & " 个员工"
            Else
                ' 这种情况理论上不会发生，因为processedCount=0且i>staffLastRow
                ' 意味着遍历了所有员工但没有处理任何新员工
                completionMessage = "所有员工都已存在，无需处理"
            End If
            UpdateStatus completionMessage, False
        Else
            ' 异常结束：所有员工都已存在
            ' 保持异常状态，不设置自动清除
            ' Application.StatusBar 已在上面设置为"所有员工考勤卡已存在，无需生成新记录"
        End If

        ' 仅正常完成时5秒后恢复状态栏；异常/取消保持当前文本
        If processedCount >= 0 Then
            安排状态栏恢复 5
        End If
    Else
        ' 非正常结束（包括日期范围不匹配等情况）
        ' 状态栏信息已经在相应位置设置，这里只添加用时信息
        If VarType(Application.StatusBar) = vbBoolean Then
            UpdateStatus "用时 " & Format(elapsedTime, "0.00") & " 秒", False
        Else
            UpdateStatus CStr(Application.StatusBar) & "，用时 " & Format(elapsedTime, "0.00") & " 秒", False
        End If
    End If

    Exit Sub

ErrorHandler:
    ' ================= 错误处理 =================
    ' 记录错误信息并跳转到清理环节
    UpdateStatus "运行错误：" & Err.Description, False
    Resume CleanUp
End Sub

' ================= 新增辅助函数 =================

' 函数说明：NormalizeName
Function NormalizeName(ByVal nm As String) As String
    NormalizeName = Replace(Trim(nm & ""), " ", "")
End Function

' 函数说明：IsTruthyStatus
Function IsTruthyStatus(ByVal statusMark As Variant) As Boolean
    IsTruthyStatus = (VarType(statusMark) = vbBoolean And statusMark = True)
End Function

' 函数说明：NormalizeWxEmpNo
Function NormalizeWxEmpNo(ByVal v As Variant) As String
    If IsEmpty(v) Or IsNull(v) Then Exit Function

    Dim s As String
    s = Trim(CStr(v))
    If s = "" Then Exit Function

    If Len(s) > 5 Then
        s = Right(s, 5)
    End If

    NormalizeWxEmpNo = s
End Function

' 函数说明：TryParseRecordDate
Function TryParseRecordDate(ByVal dtText As String, ByRef recordDate As Date) As Boolean
    Dim dateParts() As String
    On Error Resume Next
    dateParts = Split(Trim(dtText), " ")
    If UBound(dateParts) >= 0 Then
        recordDate = DateValue(dateParts(0))
        If Err.Number <> 0 Then
            Err.Clear
            recordDate = DateValue(Replace(dateParts(0), "/", "-"))
        End If
    End If
    TryParseRecordDate = (Err.Number = 0 And recordDate <> 0)
    Err.Clear
    On Error GoTo 0
End Function

' 函数说明：NormalizeTimeText
Function NormalizeTimeText(ByVal t As String) As String
    If t = "" Then Exit Function
    If t = "未打卡" Then Exit Function
    If Left(t, 2) = "次日" Then
        t = Mid(t, 3)
        If Not IsValidTimeFormat(t) Then Exit Function
    End If
    NormalizeTimeText = t
End Function

' 过程说明：ClearTemplateBlock
Sub ClearTemplateBlock(ByVal ws As Worksheet, ByVal cardRow As Long, ByVal dateRowBase As Long)
    Dim dr As Long
    For dr = 0 To 15
        SafeSet ws.Cells(cardRow + dateRowBase + dr, "C"), ""   ' 上半月上班时间
        SafeSet ws.Cells(cardRow + dateRowBase + dr, "H"), ""   ' 上半月下班时间
        SafeSet ws.Cells(cardRow + dateRowBase + dr, "K"), ""   ' 下半月上班时间
        SafeSet ws.Cells(cardRow + dateRowBase + dr, "P"), ""   ' 下半月下班时间
    Next dr

    ' 清空姓名和工号
    SafeSet ws.Cells(cardRow + 3, "E"), ""
    SafeSet ws.Cells(cardRow + 3, "H"), ""
    SafeSet ws.Cells(cardRow + 3, "M"), ""
End Sub

' ================= 原有工具函数 =================

' 功能：查找最后一个员工块的起始行
' 参数：ws - 工作表，lastNameRow - 找到的最后一个员工行，blockRows - 块大小，startRow - 起始行
Function FindLastBlockStartRow(ws As Worksheet, lastNameRow As Long, blockRows As Long, startRow As Long) As Long
    ' 从找到的员工行向上查找块的起始行
    Dim rowNum As Long
    rowNum = lastNameRow

    ' 向上查找，直到找到块起始行
    While rowNum > startRow And (rowNum - startRow) Mod blockRows <> 0
        rowNum = rowNum - 1
    Wend

    ' 确保不会低于起始行
    If rowNum < startRow Then
        FindLastBlockStartRow = startRow
    Else
        FindLastBlockStartRow = rowNum
    End If
End Function

' 功能：查找模板源行
' 参数：ws - 工作表，currentRow - 当前行，blockRows - 块大小，startRow - 起始行
Function FindTemplateSourceRow(ws As Worksheet, currentRow As Long, blockRows As Long, startRow As Long) As Long
    ' 从当前行向上查找最近的完整块
    Dim rowNum As Long
    rowNum = currentRow - 1

    ' 向上查找，直到找到包含"姓名："的完整块
    While rowNum >= startRow
        ' 检查是否是完整块的起始行
        If (rowNum - startRow) Mod blockRows = 0 Then
            ' 检查该行是否有姓名信息
            Dim cellValue As String
            cellValue = ws.Cells(rowNum + 3, "E").MergeArea.Cells(1, 1).value & ""
            If InStr(cellValue, "姓名：") > 0 Then
                FindTemplateSourceRow = rowNum
                Exit Function
            End If
        End If
        rowNum = rowNum - 1
    Wend

    ' 如果找不到，返回起始行
    FindTemplateSourceRow = startRow
End Function

' 功能：检查行是否为空
' 参数：ws - 工作表，rowNum - 行号，colCount - 检查的列数
Function IsEmptyRow(ws As Worksheet, rowNum As Long, Optional colCount As Long = 10) As Boolean
    Dim col As Long
    IsEmptyRow = True

    ' 检查指定数量的列是否都为空
    For col = 1 To colCount
        If Trim(ws.Cells(rowNum, col).value & "") <> "" Then
            IsEmptyRow = False
            Exit Function
        End If
    Next col
End Function

' 功能：安全设置单元格值（处理合并单元格）
' 参数：c - 单元格对象，v - 要设置的值，fmt - 数字格式（可选）
Private Sub ApplyCardBlockFont(ByVal ws As Worksheet, ByVal startRow As Long, ByVal blockRows As Long)
    With ws.Rows(startRow & ":" & startRow + blockRows - 1).Font
        .Name = "宋体"
        .Size = 10
    End With
End Sub
' 过程说明：SafeSet
Sub SafeSet(c As Range, v As Variant, Optional fmt As String = "")
    If c.MergeCells Then
        ' 如果是合并单元格，设置整个合并区域的值
        c.MergeArea.value = v
        If fmt <> "" Then c.MergeArea.NumberFormat = fmt
    Else
        ' 普通单元格直接设置
        c.value = v
        If fmt <> "" Then c.NumberFormat = fmt
    End If
End Sub

' 功能：将日期转换为中文星期（返回"日"、"一"..."六"）
' 参数：d - 日期
Function CNWeek1(d As Date) As String
    ' 数组索引：0-空，1-日，2-一，...，7-六
    CNWeek1 = Array("", "日", "一", "二", "三", "四", "五", "六")(Weekday(d, vbSunday))
End Function

' 函数：ExtractTimeFromText
' 功能：从文本中提取时间
' 参数：txt - 包含时间的文本
' 返回：提取的时间字符串（格式：HH:MM 或 "未打卡" 或 "次日HH:MM"）
Function ExtractTimeFromText(txt As String) As String
    If IsNull(txt) Or txt = "" Then
        ExtractTimeFromText = ""
        Exit Function
    End If

    txt = Trim(CStr(txt))

    ' 检查是否是"未打卡"
    If txt = "未打卡" Or InStr(txt, "未打卡") > 0 Then
        ExtractTimeFromText = "未打卡"
        Exit Function
    End If

    ' 检查是否是"次日XX:XX"
    If InStr(txt, "次日") > 0 Then
        Dim re As Object
        Set re = CreateObject("VBScript.RegExp")
        re.Global = True
        re.Pattern = "次日\s*(\d{1,2}):(\d{2})"  ' 匹配"次日"后的时间
        If re.Test(txt) Then
            With re.Execute(txt)(0)
                ExtractTimeFromText = "次日" & Format(.SubMatches(0), "00") & ":" & .SubMatches(1)
            End With
            Exit Function
        End If
    End If

    ' 检查普通时间格式 HH:MM 或 H:MM
    Dim re2 As Object
    Set re2 = CreateObject("VBScript.RegExp")
    re2.Global = True
    re2.Pattern = "(\d{1,2}):(\d{2})"  ' 匹配时间格式
    If re2.Test(txt) Then
        With re2.Execute(txt)(0)
            ExtractTimeFromText = Format(.SubMatches(0), "00") & ":" & .SubMatches(1)
        End With
        Exit Function
    End If

    ' 如果没有匹配到任何格式，返回空字符串
    ExtractTimeFromText = ""
End Function

' 函数：IsValidTimeFormat
' 功能：验证时间格式是否为有效的 HH:MM
' 参数：t - 时间字符串
' 返回：Boolean，是否有效
Function IsValidTimeFormat(t As String) As Boolean
    ' 检查长度是否为5位（HH:MM）
    If Len(t) <> 5 Then Exit Function

    ' 检查第3位是否为冒号
    If Mid(t, 3, 1) <> ":" Then Exit Function

    ' 提取小时和分钟
    Dim h As Integer, m As Integer
    h = Val(Left(t, 2))
    m = Val(Right(t, 2))

    ' 验证小时和分钟是否在有效范围内
    IsValidTimeFormat = (h >= 0 And h <= 23 And m >= 0 And m <= 59)
End Function


