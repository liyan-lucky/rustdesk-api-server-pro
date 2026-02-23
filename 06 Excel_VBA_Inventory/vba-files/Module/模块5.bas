Attribute VB_Name = "模块5"

' 本模块包含两个主要的公共子程序：
' 1. Cmdliushui：用于查询项目的流水账，并将结果输出到Excel工作表中。
' 2. WarehouseSummary：用于生成库存汇总表，并将其保存到桌面。

'----------------------------------------------------------

Public Sub Cmdliushui()
    '定义公共子程序：项目流水账查询

    '========== 变量声明 ==========
    Dim ws As Worksheet    ' 工作表对象
    Dim cnnConnect As ADODB.Connection    ' 数据库连接对象
    Dim rstRecordset As ADODB.Recordset   ' 数据库记录集对象
    Dim strSQL As String                  ' SQL语句
    Dim strRdRecord() As Variant          ' 存放表名的数组
    Dim userInput As String               ' 用户输入
    Dim strFindcInvCode As String         ' 查找的项目编号
    Dim cNextRowNumber As Long            ' 下一个写入行号
    Dim tmpQueryName As String            ' 临时查询名称
    Dim nFieldsCount As Integer           ' 字段数量
    Dim lastRow As Long                   ' 最后一行
    Dim i As Integer, Q As Integer        ' 循环变量
    Dim hasData As Boolean                ' 是否有数据

    '========== 检查工作表是否存在 ==========
    On Error Resume Next
    Set ws = ThisWorkbook.Sheets("流水账")    ' 获取名为“流水账”的工作表
    If Err.Number <> 0 Then
        On Error GoTo 0
        CreateObject("WScript.Shell").popup _
            "找不到名为'流水账'的工作表，请先创建该工作表！", _
            1, "错误", 16 + 4096
        Exit Sub
    End If
    On Error GoTo 0

    '========== 获取用户输入 ==========
    userInput = InputBox("请输入6位项目编号:", "项目编号验证", "888888")    ' 弹出输入框获取项目编号

    '检查用户是否取消输入
    If userInput = "" Then
        CreateObject("WScript.Shell").popup _
            "项目流水账查询已取消！", _
            1, "输入已取消", 64 + 4096
        Exit Sub
    End If

    '验证输入是否为6位数字
    If Len(userInput) = 6 And IsNumeric(userInput) Then
        strFindcInvCode = userInput    ' 合法则赋值
    Else
        CreateObject("WScript.Shell").popup _
            "输入无效！必须为6位数字。", _
            1, "错误", 16 + 4096
        Exit Sub
    End If

    '========== 准备工作表 ==========
    Application.ScreenUpdating = False    ' 关闭屏幕刷新
    Application.DisplayAlerts = False     ' 关闭警告提示

    With ws
        .Activate
        .Cells.ClearContents '清空所有内容

        '删除历史查询名称
        Dim nm As Name
        For Each nm In .Names
            If Left(nm.Name, 14) = "流水账!ContactList" Then
                nm.Delete
            End If
        Next nm

        '删除所有查询表
        For Q = .QueryTables.Count To 1 Step -1
            .QueryTables(Q).Delete
        Next Q
    End With

    '========== 数据库查询配置 ==========
    strRdRecord = Array("RdRecord34", "RdRecords34", "RdRecord01", "RdRecords01", _
        "RdRecord11", "RdRecords11", "RdRecord08", "RdRecords08")    ' 相关表名数组

    cNextRowNumber = 1    ' 初始化写入行号
    hasData = False       ' 初始化数据标志

    '========== 执行多个SQL查询 ==========
    For i = 0 To 3    ' 循环4次，分别查询不同的表
        '构建SQL语句 - 修正SQL语法
        strSQL = "Select RdRecord.dDate As 日期, RdRecord.cWhCode As 仓库, RdRecord.cCode As 单据号, " & _
            "rdrecords.cInvCode As 货号, Inventory.cInvName As 名称, Inventory.cInvStd As 规格, " & _
            "ComputationUnit.cComUnitName As 单位, " & _
            "rdrecords.iQuantity * (RdRecord.bRdFlag * 2 - 1) As 数量, " & _
            "RdRecord.cMemo As 备注, " & _
            "rdrecords.cItemCode As 项目编号, rdrecords.cName As 项目名称, " & _
            "RdRecord.cBusType As 进出类别, RdRecord.cRdCode As 单别, " & _
            "RdRecord.cDepCode As 部门, RdRecord.cMaker As 制单 " & _
            "FROM UFDATA_002_2015.dbo.ComputationUnit ComputationUnit, " & _
            "UFDATA_002_2015.dbo.Inventory Inventory, " & _
            "UFDATA_002_2015.dbo." & strRdRecord(i * 2) & " RdRecord, " & _
            "UFDATA_002_2015.dbo." & strRdRecord(i * 2 + 1) & " rdrecords " & _
            "WHERE RdRecord.ID = rdrecords.ID " & _
            "And rdrecords.cInvCode = Inventory.cInvCode " & _
            "And Inventory.cComUnitCode = ComputationUnit.cComunitCode " & _
            "And rdrecords.cItemCode = '" & strFindcInvCode & "' "

        '添加日期条件
        If i = 0 Then
            strSQL = strSQL & "And (RdRecord.dDate < '2015-01-01')"    ' 第一组查早于2015年
        Else
            strSQL = strSQL & "And (RdRecord.dDate >= '2015-01-01')"   ' 其他查2015年及以后
        End If

        Debug.Print "SQL " & i + 1 & ": " & strSQL '调试输出SQL语句

        '执行查询
        Set cnnConnect = New ADODB.Connection
        Set rstRecordset = New ADODB.Recordset

        On Error GoTo SqlConnectError
        cnnConnect.Open "DRIVER=SQL Server;SERVER=ERPServer;UID=sa;PWD=1QAZ2wsx;DATABASE=UFDATA_002_2015"    ' 打开数据库连接

        On Error GoTo QueryError
        rstRecordset.Open Source:=strSQL, ActiveConnection:=cnnConnect, _
            CursorType:=adOpenStatic, LockType:=adLockReadOnly, _
            Options:=adCmdText    ' 执行SQL查询

        '检查是否有数据
        If Not rstRecordset.EOF Then
            hasData = True    ' 查询到数据

            '将结果直接复制到工作表（更可靠的方法）
            If cNextRowNumber = 1 Then
                '写入标题
                For col = 0 To rstRecordset.Fields.Count - 1
                    ws.Cells(1, col + 1).Value = rstRecordset.Fields(col).Name
                Next col
                cNextRowNumber = 2
            End If

            '写入数据
            ws.Range("A" & cNextRowNumber).CopyFromRecordset rstRecordset

            '更新行号
            cNextRowNumber = ws.Cells(ws.Rows.Count, "A").End(xlUp).Row + 1
        End If

QueryContinue:
        '清理对象
        If rstRecordset.State = adStateOpen Then rstRecordset.Close
        If cnnConnect.State = adStateOpen Then cnnConnect.Close
        Set rstRecordset = Nothing
        Set cnnConnect = Nothing

    Next i

    '========== 数据后处理 ==========
    Application.DisplayAlerts = True    ' 恢复警告提示

    With ws
        lastRow = .Cells(.Rows.Count, "A").End(xlUp).Row    ' 获取最后一行

        If lastRow > 1 And hasData Then
            '数据排序
            .Range("A1:S" & lastRow).Sort Key1:=.Range("D2"), Order1:=xlAscending, _
                Key2:=.Range("B2"), Order2:=xlAscending, _
                Key3:=.Range("A2"), Order3:=xlAscending, _
                Header:=xlYes

            '添加筛选
            .Range("A1:S" & lastRow).AutoFilter

            '调整列宽
            .Columns("A:S").AutoFit
        End If

        .Range("A2").Select    ' 选中A2单元格
    End With

    '========== 完成提示 ==========
    Application.ScreenUpdating = True    ' 恢复屏幕刷新

    Dim dataRows As Long
    dataRows = lastRow - 1
    If dataRows < 0 Then dataRows = 0

    If dataRows > 0 Then
        CreateObject("WScript.Shell").popup _
            "项目流水账查询完成！" & vbCrLf & _
            "项目编号: " & strFindcInvCode & vbCrLf & _
            "共查询到 " & dataRows & " 条记录", _
            3, "查询成功", 64
    Else
        CreateObject("WScript.Shell").popup _
            "未找到项目编号 " & strFindcInvCode & " 的相关记录！" & vbCrLf & _
            "请检查：1.项目编号是否正确 2.数据库中是否存在该项目的流水数据", _
            3, "查询结果", 48
    End If

    Exit Sub

    '========== 错误处理 ==========
SqlConnectError:
    CreateObject("WScript.Shell").popup _
        "数据库连接失败！" & vbCrLf & _
        "请检查：1.网络连接 2.数据库服务器状态 3.连接字符串配置", _
        3, "数据库错误", 16
    Application.ScreenUpdating = True
    Exit Sub

QueryError:
    CreateObject("WScript.Shell").popup _
        "第 " & i + 1 & " 个查询执行失败！" & vbCrLf & _
        "错误信息: " & Err.Description, _
        1, "查询错误", 16
    Resume QueryContinue
End Sub

'----------------------------------------------------------

Sub WarehouseSummary()
    Dim wsOriginal As Worksheet    ' 原始工作表
    Dim wbNew As Workbook          ' 新工作簿
    Dim wsCopy As Worksheet        ' 复制的工作表
    Dim wsPivot As Worksheet       ' 数据透视表工作表
    Dim pivotCache As pivotCache   ' 数据透视表缓存
    Dim pivotTable As pivotTable   ' 数据透视表对象
    Dim lastRow As Long            ' 最后一行
    Dim lastCol As Long            ' 最后一列
    Dim dataRange As Range         ' 数据范围
    Dim ws As Worksheet            ' 工作表对象
    Dim pf As PivotField           ' 数据透视表字段
    Dim projectNumber As String    ' 项目编号
    Dim fileName As String         ' 文件名
    Dim desktopPath As String      ' 桌面路径
    Dim projectNumberCol As String ' 项目编号列
    Dim i As Integer               ' 循环变量
    Dim foundProjectNumber As Boolean ' 是否找到项目编号列
    Dim btn As Object              ' 声明按钮对象变量
    Dim popup As Object            ' 用于创建弹窗对象

    ' 设置原始工作表
    On Error Resume Next
    Set wsOriginal = ThisWorkbook.Worksheets("流水账")    ' 获取“流水账”工作表
    On Error GoTo 0

    ' 检查原始工作表是否存在
    If wsOriginal Is Nothing Then
        Set popup = CreateObject("WScript.Shell")
        popup.popup "名为'流水账'的工作表不存在！", 1, "错误", 16 + 4096
        Exit Sub
    End If

    ' 查找项目编号列
    foundProjectNumber = False
    projectNumberCol = "J" ' 从J列开始

    For i = 0 To 4 ' 检查J列及后面4列
        If wsOriginal.Cells(1, Asc(projectNumberCol) - 64 + i).Value = "项目编号" Then
            projectNumberCol = Chr(Asc(projectNumberCol) + i) ' 计算列字母
            foundProjectNumber = True
            Exit For
        End If
    Next i

    ' 如果未找到项目编号列，停止运行
    If Not foundProjectNumber Then
        Set popup = CreateObject("WScript.Shell")
        popup.popup "未找到'项目编号'列！请确保在J列至N列中存在'项目编号'标题。", 1, "错误", 16 + 4096
        Exit Sub
    End If

    ' 获取项目编号
    With wsOriginal
        lastRow = .Cells(.Rows.Count, projectNumberCol).End(xlUp).Row
        If lastRow > 1 Then
            projectNumber = .Range(projectNumberCol & "2").Value
        Else
            projectNumber = "未知项目"
        End If
    End With

    ' 创建新工作簿
    Set wbNew = Workbooks.Add
    Application.DisplayAlerts = False ' 防止删除工作表时出现提示

    ' 添加新工作表用于数据透视表，表名称为"库存汇总"
    Set wsPivot = wbNew.Worksheets.Add(Before:=wbNew.Sheets(1))
    wsPivot.Name = "库存汇总"

    ' 复制流水账工作表到新工作簿
    wsOriginal.Copy After:=wbNew.Sheets(1)
    Set wsCopy = wbNew.Sheets(2)
    ' 直接命名为"流水账"，不带副本二字
    On Error Resume Next
    wsCopy.Name = "流水账"
    ' 如果名称已存在，添加后缀
    If Err.Number <> 0 Then
        Err.Clear
        wsCopy.Name = "流水账_" & Format(Now(), "hhmmss")
    End If
    On Error GoTo 0

    ' 删除新工作簿中的默认空白工作表
    For Each ws In wbNew.Worksheets
        If ws.Name <> "库存汇总" And ws.Name <> wsCopy.Name Then
            Application.DisplayAlerts = False
            ws.Delete
            Application.DisplayAlerts = True
        End If
    Next ws

    ' 确定数据范围
    With wsCopy
        lastRow = .Cells(.Rows.Count, 1).End(xlUp).Row
        lastCol = .Cells(1, .Columns.Count).End(xlToLeft).Column

        If lastRow < 2 Then
            Set popup = CreateObject("WScript.Shell")
            popup.popup "流水账工作表没有数据！", 1, "错误", 16 + 4096
            wbNew.Close SaveChanges:=False
            Application.DisplayAlerts = True
            Exit Sub
        End If

        Set dataRange = .Range("A1").Resize(lastRow, lastCol)
    End With

    ' 创建数据透视表缓存
    Set pivotCache = wbNew.PivotCaches.Create( _
        SourceType:=xlDatabase, _
        SourceData:=dataRange.Address)

    ' 创建数据透视表
    Set pivotTable = pivotCache.CreatePivotTable( _
        TableDestination:=wsPivot.Range("A3"), _
        TableName:="库存透视表")

    ' 设置数据透视表字段
    With pivotTable
        ' 添加行字段 - 包括规格字段
        .PivotFields("货号").Orientation = xlRowField
        .PivotFields("名称").Orientation = xlRowField
        .PivotFields("规格").Orientation = xlRowField
        .PivotFields("单位").Orientation = xlRowField

        ' 添加值字段（数量求和），并命名为"库存数量"
        .AddDataField .PivotFields("数量"), "库存数量", xlSum

        ' 不显示总计
        .RowGrand = False
        .ColumnGrand = False

        ' 不显示+/-按钮
        .ShowDrillIndicators = False

        ' 设置表格布局为表格形式
        .RowAxisLayout xlTabularRow

        ' 关闭数据透视表字段列表
        .EnableFieldList = False
    End With

    ' 为每个行字段设置不显示分类汇总
    For Each pf In pivotTable.PivotFields
        If pf.Orientation = xlRowField Then
            pf.Subtotals(1) = False ' 禁用自动分类汇总
            pf.Subtotals = Array(False, False, False, False, False, False, False, False, False, False, False, False)
        End If
    Next pf

    ' 重复所有项目标签
    pivotTable.RepeatAllLabels xlRepeatLabels

    ' 设置标题和格式
    With wsPivot
        .Range("A1").Value = "库存汇总表"
        .Range("A1").Font.Bold = True
        .Range("A1").Font.Size = 14

        ' 设置指定列宽
        .Columns("A:A").ColumnWidth = 15 ' 货号
        .Columns("B:B").ColumnWidth = 30 ' 名称
        .Columns("C:C").ColumnWidth = 30 ' 规格
        .Columns("D:D").ColumnWidth = 10 ' 单位
        .Columns("E:E").ColumnWidth = 10 ' 库存数量
    End With

    ' 确保库存汇总表是第一个工作表
    wsPivot.Move Before:=wbNew.Sheets(1)

    ' 恢复显示警告
    Application.DisplayAlerts = True

    ' 激活数据透视表工作表
    wsPivot.Activate

    ' 构建文件名并保存到桌面
    desktopPath = CreateObject("WScript.Shell").SpecialFolders("Desktop")

    ' 清理项目编号中的非法字符
    projectNumber = CleanFileName(projectNumber)

    ' 构建文件名：项目编号_当前年月日（格式为00000000）
    fileName = projectNumber & "_" & Format(Date, "yyyymmdd") & ".xlsx"

    ' 完整文件路径
    Dim fullPath As String
    fullPath = desktopPath & "\" & fileName

    ' 防止覆盖提示
    Application.DisplayAlerts = False

    ' 只删除新工作簿中名为"导出桌面"的按钮
    For Each ws In wbNew.Worksheets
        On Error Resume Next ' 忽略错误，继续执行

        ' 删除表单控件按钮（名称为"导出桌面"的按钮）
        For Each btn In ws.Buttons
            If btn.Name = "导出桌面" Or btn.Caption = "导出桌面" Then
                btn.Delete
            End If
        Next btn

        ' 删除ActiveX控件按钮（名称为"导出桌面"的按钮）
        For Each btn In ws.OLEObjects
            If TypeName(btn.Object) = "CommandButton" Then
                If btn.Name = "导出桌面" Or btn.Object.Caption = "导出桌面" Then
                    btn.Delete
                End If
            End If
        Next btn

        On Error GoTo 0
    Next ws

    ' 保存工作簿 - 直接覆盖已存在的文件
    On Error Resume Next
    ' 如果文件已存在，先删除
    If Dir(fullPath) <> "" Then
        Kill fullPath
    End If
    On Error GoTo 0

    wbNew.SaveAs fileName:=fullPath, FileFormat:=xlOpenXMLWorkbook    ' 保存到桌面
    ' 关闭工作簿
    wbNew.Close SaveChanges:=False

    ' 恢复显示警告
    Application.DisplayAlerts = True

    ' 删除原工作簿中"流水账"表第一行以外的所有内容（使用.Clear方法）
    With wsOriginal
        ' 获取最后一行
        lastRow = .Cells(.Rows.Count, 1).End(xlUp).Row

        ' 如果有超过一行的数据，清除第二行及以后的所有内容
        If lastRow > 1 Then
            .Range("2:" & lastRow).Clear
        End If
    End With

    ' 可选：激活原工作簿的流水账工作表
    wsOriginal.Activate
    wsOriginal.Range("A2").Select ' 选择A2单元格，方便继续输入

    Set popup = CreateObject("WScript.Shell")
    popup.popup "库存汇总表已生成并保存到桌面，文件名为: " & fileName & vbCrLf & _
        "流水账表已清空，只保留标题行。", _
        1, "操作完成", 64 + 4096

End Sub

' 辅助函数：清理文件名中的非法字符
Function CleanFileName(strFileName As String) As String
    Dim strInvalidChars As String    ' 非法字符集合
    Dim i As Integer                 ' 循环变量

    strInvalidChars = "\/:*?""<>|"

    For i = 1 To Len(strInvalidChars)
        strFileName = Replace(strFileName, Mid(strInvalidChars, i, 1), "")
    Next i

    CleanFileName = Trim(strFileName)    ' 返回清理后的文件名
End Function
