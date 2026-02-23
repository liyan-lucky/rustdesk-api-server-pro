VERSION 1.0 CLASS
BEGIN
MultiUse = -1  'True
END
Attribute VB_Name = "Sheet1"
Attribute VB_GlobalNameSpace = False
Attribute VB_Creatable = False
Attribute VB_PredeclaredId = True
Attribute VB_Exposed = True

' 查找数量按钮事件
Public Sub cmdFindQty()
    Dim ws As Worksheet    ' 工作表对象
    Dim lastRow As Long    ' 最后一行变量
    Set ws = ActiveSheet   ' 设置为当前活动工作表

    ' 获取F列最后一个非空单元格的行号（从下往上找）
    lastRow = ws.Cells(ws.Rows.Count, "F").End(xlUp).Row

    ' 如果F列有数据（行号大于等于2），则清空F3到最后一个单元格
    If lastRow >= 2 Then
        If lastRow > 2 Then
            ws.Range("F3:F" & lastRow).ClearContents    ' 清空F3到最后一行内容
        End If
    Else
        ' 如果F列为空或只有标题，确保F2有内容
        ws.Range("F2").Value = "总库存"    ' 设置F2为标题
    End If
    ' 确保F2单元格有标题
    If ws.Range("F2").Value <> "总库存" Then
        ws.Range("F2").Value = "总库存"
    End If

    Set ws = ActiveSheet
    On Error Resume Next ' 忽略错误
    lastRow = ws.Cells(ws.Rows.Count, "B").End(xlUp).Row    ' 获取B列最后一行
    On Error Goto 0 ' 恢复错误处理

        ' 检查是否找到有效行号（避免返回第一行）
        If lastRow = ws.Rows.Count And ws.Cells(ws.Rows.Count, "B").Value = "" Then
            ' 如果返回的是最后一行且该单元格为空，说明整列为空
            lastRow = 0
        Elseif lastRow = 1 And ws.Cells(1, "B").Value = "" Then
            ' 如果返回的是第一行且该单元格为空，说明整列为空
            lastRow = 0
        End If

        If lastRow = 2 Then
            MsgBox "存货编号为空"    ' 提示存货编号为空
        Else
            ' 在F3:F区域填充SUMIFS公式，统计现存量
            Range("F3:F" & lastRow).FormulaR1C1 = "=SUMIFS(现存量!C[1],现存量!C[-3],清单制作!RC[-4] )"
        End If
End Sub

' 回车切换方向按钮事件
Private Sub cmdEnter_Click()
    ' 如果当前回车方向为向下，则切换为向右，否则切换为向下
    If Application.MoveAfterReturnDirection = xlDown Then
        Application.MoveAfterReturnDirection = xlToRight
    Else
        Application.MoveAfterReturnDirection = xlDown
    End If
End Sub
'-------------------------------------------------------------------------------

Public Sub CmdBtnDetail()
    '定义公共子程序：存货编号明细账查询

    '========== 变量声明 ==========
    Dim wsList As Worksheet, wsDetail As Worksheet    ' 工作表对象
    Dim cnnConnect As ADODB.Connection               ' 数据库连接对象
    Dim rstRecordset As ADODB.Recordset              ' 数据库记录集对象
    Dim strSQL As String                             ' SQL语句
    Dim strRdRecord() As Variant                     ' 表名数组
    Dim userInput As String                          ' 用户输入
    Dim strFindcInvCode As String                    ' 查找的存货编号
    Dim cNextRowNumber As Long                       ' 下一个写入行号
    Dim tmpQueryName As String                       ' 临时查询名称
    Dim nFieldsCount As Integer                      ' 字段数量
    Dim lastRow As Long                              ' 最后一行
    Dim i As Integer, Q As Integer, col As Integer   ' 循环变量
    Dim rngSelected As Range                         ' 选中单元格
    Dim hasData As Boolean                           ' 是否有数据
    Dim wshShell As Object                           ' 弹窗对象

    On Error Goto ErrorHandler

        '创建WScript.Shell对象用于弹出窗口
        Set wshShell = CreateObject("WScript.Shell")

        '========== 检查工作表是否存在 ==========
        On Error Resume Next
        Set wsList = ThisWorkbook.Sheets("清单制作")
        Set wsDetail = ThisWorkbook.Sheets("明细账")
        If Err.Number <> 0 Then
            On Error Goto ErrorHandler
                wshShell.popup "找不到必要的工作表，请确保存在'清单制作'和'明细账'工作表！", 1, "错误", 16
             Exit Sub
            End If
            On Error Goto ErrorHandler

                '========== 获取当前选中的单元格内容 ==========
                Set rngSelected = ActiveCell ' 使用ActiveCell而不是Selection

                '检查是否选中了单元格
                If rngSelected Is Nothing Then
                    wshShell.popup "请先选中包含存货编号的单元格！", 1, "提示", 48
                 Exit Sub
                End If

                '获取选中单元格所在行的第2列内容（存货编号）
                strFindcInvCode = Trim(wsList.Cells(rngSelected.Row, 2).Value)

                '检查单元格是否为空
                If strFindcInvCode = "" Then
                    wshShell.popup "选中的行中存货编号为空，请选择包含存货编号的行！", 1, "提示", 48
                 Exit Sub
                End If

                '验证输入是否为大于等于4个字符的字符串
                If Len(strFindcInvCode) < 4 Then
                    wshShell.popup "存货编号不合法！" & vbCrLf & _
                    "存货编号必须大于等于4个字符。" & vbCrLf & _
                    "当前选中的内容: '" & strFindcInvCode & "' (长度: " & Len(strFindcInvCode) & ")", 1, "错误", 48
                 Exit Sub
                End If

                '========== 准备工作表 ==========
                Application.ScreenUpdating = False    ' 关闭屏幕刷新
                Application.DisplayAlerts = False     ' 关闭警告提示

                With wsDetail
                    .Activate
                    .Cells.ClearContents '清空所有内容

                    '删除历史查询名称
                    Dim nm As Name
                    For Each nm In .Names
                        If Left(nm.Name, 14) = "明细账!ContactList" Then
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
                "RdRecord11", "RdRecords11", "RdRecord08", "RdRecords08", _
                "RdRecord09", "RdRecords09")

                cNextRowNumber = 1    ' 初始化写入行号
                hasData = False       ' 初始化数据标志

                '========== 执行多个SQL查询 ==========
                For i = 0 To 4
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
                    "And rdrecords.cInvCode = '" & strFindcInvCode & "' "

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

                    On Error Goto SqlConnectError
                        cnnConnect.Open "DRIVER=SQL Server;SERVER=ERPServer;UID=sa;PWD=1QAZ2wsx;DATABASE=UFDATA_002_2015"    ' 打开数据库连接

                        On Error Goto QueryError
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
                                        wsDetail.Cells(1, col + 1).Value = rstRecordset.Fields(col).Name
                                    Next col
                                    cNextRowNumber = 2
                                End If

                                '写入数据
                                wsDetail.Range("A" & cNextRowNumber).CopyFromRecordset rstRecordset

                                '更新行号
                                cNextRowNumber = wsDetail.Cells(wsDetail.Rows.Count, "A").End(xlUp).Row + 1
                            End If

 QueryContinue:
                            '清理对象
                            If Not rstRecordset Is Nothing Then
                                If rstRecordset.State = adStateOpen Then rstRecordset.Close
                                    Set rstRecordset = Nothing
                                End If
                                If Not cnnConnect Is Nothing Then
                                    If cnnConnect.State = adStateOpen Then cnnConnect.Close
                                        Set cnnConnect = Nothing
                                    End If
                                Next i

                                '========== 数据后处理 ==========
                                Application.DisplayAlerts = True    ' 恢复警告提示

                                With wsDetail
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
                                        wshShell.popup "存货明细账查询完成！" & vbCrLf & _
                                        "存货编号: " & strFindcInvCode & vbCrLf & _
                                        "共查询到 " & dataRows & " 条记录", 1, "查询成功", 64
                                    Else
                                        wshShell.popup "未找到存货编号 " & strFindcInvCode & " 的相关记录！" & vbCrLf & _
                                        "请检查：1.存货编号是否正确 2.数据库中是否存在该存货的流水数据", 1, "查询结果", 48
                                    End If

                                 Exit Sub

                                    '========== 错误处理 ==========
 SqlConnectError:
                                    wshShell.popup "数据库连接失败！" & vbCrLf & _
                                    "请检查：1.网络连接 2.数据库服务器状态 3.连接字符串配置", 1, "数据库错误", 16
                                    Application.ScreenUpdating = True
                                 Exit Sub

 QueryError:
                                    wshShell.popup "第 " & i + 1 & " 个查询执行失败！" & vbCrLf & _
                                    "错误信息: " & Err.Description, 1, "查询错误", 16
                                    Resume QueryContinue

 ErrorHandler:
                                        wshShell.popup "发生未知错误：" & Err.Description & vbCrLf & _
                                        "错误代码：" & Err.Number, 1, "系统错误", 16
                                        Application.ScreenUpdating = True
End Sub

'-------------------------------------------------------------------------------


Public Sub cmdbtnRefreshData()

    Application.ScreenUpdating = False
    On Error Goto ErrorHandler

        ' 刷新所有外部数据连接
        ThisWorkbook.RefreshAll
        ' 或者分别刷新不同类型的数据连接
        ' ActiveWorkbook.UpdateLink Name:=ActiveWorkbook.LinkSources
        ' MsgBox "所有链接已刷新完成！", vbInformation, "刷新完成"

     ExitSub:
        Application.ScreenUpdating = True
     Exit Sub

 ErrorHandler:
        MsgBox "刷新链接时出现错误：" & Err.Description, vbExclamation, "错误"
        Resume ExitSub

End Sub

'-------------------------------------------------------------------------------
' 确认按钮事件
Public Sub ComBtnConfirm()
    ' 在G列填充公式，校验存货目录信息
    Dim lastRow As Long
    With ActiveSheet
        ' 找到B列最后一个非空单元格的行号
        lastRow = .Cells(.Rows.Count, "B").End(xlUp).Row

        ' 确保至少从第3行开始
        If lastRow < 3 Then lastRow = 3

            ' 应用公式到G3到最后一个非空单元格
            .Range("G3:G" & lastRow).FormulaR1C1 = _
            "=If(ISBLANK(RC[-5]), """"," & _
            "IFERROR(" & _
            "And(" & _
            "VLOOKUP(RC[-5],存货目录!C[-6]:C[-4],2,False)=RC[-4]," & _
            "VLOOKUP(RC[-5],存货目录!C[-6]:C[-4],3,False)=RC[-3]," & _
            "UPPER(VLOOKUP(RC[-5],存货目录!C[-6]:C[-3],4,False))=UPPER(RC[-2])" & _
            ")," & _
            "False)" & _
            ")"
        End With
End Sub

' 批量查找存货按钮事件
Public Sub BatchFindInv()
    ' 变量声明
    Dim wsList As Worksheet      ' 清单制作工作表
    Dim wsInv As Worksheet       ' 存货目录工作表
    Dim nBegRow As Long          ' 起始行号
    Dim nRows As Long            ' 选中行数
    Dim i As Long                ' 循环计数器
    Dim n As Long                ' 匹配数量计数器
    Dim nInvFindRows As Long     ' 匹配行号
    Dim searchRange As Range     ' 搜索范围

    ' 设置工作表对象（提高代码可读性和性能）
    Set wsList = Sheets("清单制作")
    Set wsInv = Sheets("存货目录")

    ' 验证选择区域是否连续
    If Selection.Areas.Count > 1 Then
        MsgBox "请选择连续的区域！", vbExclamation
     Exit Sub
    End If

    ' 获取选择区域信息
    nBegRow = Selection.Rows.Row
    nRows = Selection.Rows.Count

    ' 清空B列选定区域内容
    With wsList
        .Range("B" & nBegRow & ":B" & nBegRow + nRows - 1).ClearContents

        ' 复制C、D列内容到H列（备份原始数据）
        .Range("C" & nBegRow & ":D" & nBegRow + nRows - 1).Copy _
        Destination:=.Range("H" & nBegRow)
        Application.CutCopyMode = False

        ' 循环处理每一行
        For i = nBegRow To nBegRow + nRows - 1
            ' 使用CountIfs进行模糊匹配查询
            ' 原来的模糊匹配（包含通配符*）
            'n = WorksheetFunction.CountIfs(wsInv.Range("B:B"), "*" & .Cells(i, 3) & "*", _
            'wsInv.Range("C:C"), "*" & .Cells(i, 4) & "*")

            ' 修改后的精准匹配（移除通配符*）
            n = WorksheetFunction.CountIfs(wsInv.Range("B:B"), .Cells(i, 3), _
            wsInv.Range("C:C"), .Cells(i, 4))

            ' 如果找到唯一匹配项
            If n = 1 Then
                ' 使用数组公式查找匹配行号
                .Range("J" & i).FormulaArray = "=MATCH(""*""&RC[-2]&""*""&RC[-1]&""*"",表_Inv[名称]&表_Inv[规格],0)"
                nInvFindRows = .Range("J" & i).Value

                ' 从存货目录中提取对应信息并填充
                .Range("B" & i).Value = WorksheetFunction.Index(wsInv.Range("表_Inv[编码]"), nInvFindRows, 1)    ' 存货编码
                .Range("C" & i).Value = WorksheetFunction.Index(wsInv.Range("表_Inv[名称]"), nInvFindRows, 1)    ' 存货名称
                .Range("D" & i).Value = WorksheetFunction.Index(wsInv.Range("表_Inv[规格]"), nInvFindRows, 1)    ' 存货规格
                .Range("E" & i).Value = WorksheetFunction.Index(wsInv.Range("表_Inv[主单位]"), nInvFindRows, 1)  ' 计量单位
            Elseif n > 1 Then
                ' 多个匹配项时的处理（可选：添加提示信息）
                .Range("B" & i).Value = "多个匹配项"
            Else
                ' 无匹配项时的处理（可选：添加提示信息）
                .Range("B" & i).Value = "未找到匹配项"
            End If
        Next i

        ' 清理辅助列J列
        .Range("J" & nBegRow & ":J" & nBegRow + nRows - 1).ClearContents

        ' 将焦点返回到B列起始位置
        .Range("B" & nBegRow).Select
    End With

    ' 释放对象变量
    Set wsList = Nothing
    Set wsInv = Nothing
    'MsgBox "批量查找完成！", vbInformation
End Sub

' 查找存货按钮事件
Public Sub FindInv()
    ' 过程说明：根据当前工作表（单位目录）中选中的存货名称和规格型号，
    '          在"存货目录"工作表中进行模糊筛选匹配

    ' 声明变量
    Dim cUnitSheetName As String     ' 当前工作表名称
    Dim nFindRow As Long             ' 当前选中行号
    Dim nFindCol As Long             ' 当前选中列号
    Dim cFindInvName As String       ' 要查找的存货名称
    Dim cFindStdName As String       ' 要查找的规格型号
    Dim wsInv As Worksheet           ' 存货目录工作表对象（新增变量提高可读性）

    ' 设置工作表对象引用，避免重复使用Sheets()提高效率
    Set wsInv = Sheets("存货目录")
    cUnitSheetName = Me.Name

    ' 获取当前选中的单元格位置
    nFindRow = ActiveCell.Row
    nFindCol = ActiveCell.Column

    ' 从当前工作表的第3列（C列）和第4列（D列）获取查找内容
    With Sheets(cUnitSheetName)
        cFindInvName = Trim(.Cells(nFindRow, 3).Value)  ' C列：存货名称
        cFindStdName = Trim(.Cells(nFindRow, 4).Value)  ' D列：规格型号
    End With

    ' 将空格替换为通配符*，支持模糊查找
    ' 使用Replace函数替代WorksheetFunction.Substitute提高效率
    cFindInvName = Replace(cFindInvName, " ", "*")
    cFindStdName = Replace(cFindStdName, " ", "*")

    ' 切换到存货目录工作表
    wsInv.Activate

    ' 如果当前存在筛选，则先清除筛选
    If wsInv.AutoFilterMode Then
        wsInv.AutoFilterMode = False
    End If

    ' 根据查找条件进行筛选
    If Len(Trim(cFindInvName)) = 0 And Len(Trim(cFindStdName)) = 0 Then
        ' 如果两个查找条件都为空，不进行任何筛选
        Debug.Print "查找条件为空，未进行筛选"
    Else
        ' 设置筛选条件
        With wsInv.ListObjects("表_Inv")
            ' 筛选第10字段为空的内容（根据实际需求调整字段序号）
            .Range.AutoFilter Field:=10, Criteria1:="="

            ' 筛选存货名称包含指定内容
            If Len(Trim(cFindInvName)) > 0 Then
                .Range.AutoFilter Field:=2, Criteria1:="=*" & cFindInvName & "*"
            End If

            ' 筛选规格型号包含指定内容（如果有的话）
            If Len(Trim(cFindStdName)) > 0 Then
                .Range.AutoFilter Field:=3, Criteria1:="=*" & cFindStdName & "*"
            End If
        End With
    End If

    ' 导航到筛选结果的最后一行（优化导航逻辑）
    With wsInv
        ' 选择B列的第一个单元格
        .Cells(1, 2).Select

        ' 如果有数据，跳转到最后一行；如果数据量很大（超过10万行），则向上查找
        If .Cells(.Rows.Count, 2).End(xlUp).Row > 1 Then
            .Cells(.Rows.Count, 2).End(xlUp).Select
        Else
            .Cells(1, 2).Select
        End If
    End With

    ' 清理对象引用
    Set wsInv = Nothing

    Debug.Print "查找完成：存货名称=" & cFindInvName & ", 规格型号=" & cFindStdName
End Sub
















