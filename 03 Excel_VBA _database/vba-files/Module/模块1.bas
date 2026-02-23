Sub 选择输入Excel表格()
    ' 弹出文件选择对话框，选择要输入的Excel表格
    Dim fd As FileDialog
    Set fd = Application.FileDialog(msoFileDialogFilePicker)
    fd.Title = "请选择要输入的Excel表格"
    fd.Filters.Clear
    fd.Filters.Add "Excel 文件", "*.xls; *.xlsx; *.xlsm"
    fd.AllowMultiSelect = True ' 允许多选

    If fd.Show = -1 Then ' 用户点击了“打开”
        Dim selectedFile As Variant
        Dim sheetName As String
        'sheetName = InputBox("请输入要操作的工作表名称（如：Integrator Offer Supplier）：", "工作表名称")
        sheetName = "Integrator Offer Supplier"
        If sheetName = "" Then
            MsgBox "未输入工作表名称，程序结束。", vbExclamation
            End
        End If
        For Each selectedFile In fd.SelectedItems
            ' 打开选中的Excel文件
            Dim wb As Workbook
            Set wb = Workbooks.Open(selectedFile)
            ' 判断工作簿是否包含指定名称的工作表
            Dim ws As Worksheet
            On Error Resume Next
            Set ws = wb.Sheets(sheetName)
            On Error GoTo 0
            If Not ws Is Nothing Then
                ' 打开database.xlsm文件
                Dim dbWb As Workbook
                Dim dbWs As Worksheet
                Dim lookupValue As Variant
                Dim foundCell As Range

                On Error Resume Next
                Set dbWb = Workbooks("database.xlsm")
                On Error GoTo 0
                If dbWb Is Nothing Then
                    Set dbWb = Workbooks.Open("e:\Vsc_VBA - database\database.xlsm")
                End If

                Set dbWs = dbWb.Sheets("price")

                Dim rowIdx As Long
                rowIdx = 10 ' 从M10开始

                Dim emptyCount As Integer
                emptyCount = 0

                ' 主循环：分段处理数据
                Do While True
                    ' 查找第一个文本类型单元格，作为求和起始行
                    Do While Not (VarType(ws.Cells(rowIdx, "M").Value) = vbString And ws.Cells(rowIdx, "M").Value <> "")
                        emptyCount = emptyCount + 1
                        If emptyCount >= 10 Then Exit Do ' 连续10行为空则结束
                        rowIdx = rowIdx + 1
                    Loop
                    If emptyCount >= 10 Then Exit Do
                    Dim sumStartRow As Long
                    Dim sumRow As Long
                    sumRow = rowIdx - 1 ' 求和结果输出行（段前一行）
                    sumStartRow = rowIdx ' 求和起始行

                    ' 向下查找，直到遇到第一个非文本类型或空单元格，作为求和结束行
                    Do While VarType(ws.Cells(rowIdx, "M").Value) = vbString And ws.Cells(rowIdx, "M").Value <> ""
                        emptyCount = 0
                        lookupValue = ws.Cells(rowIdx, "M").Value
                        ' 在price表A列查找M列的值
                        Set foundCell = dbWs.Range("A:A").Find(lookupValue, LookIn:=xlValues, LookAt:=xlWhole)
                        If Not foundCell Is Nothing Then
                            ' 获取E到L列，输出到S到Z列
                            Dim i As Integer
                            For i = 0 To 7 ' E=5,S=19; L=12,Z=26
                                ws.Cells(rowIdx, 19 + i).Value = dbWs.Cells(foundCell.Row, 5 + i).Value
                            Next i
                        End If
                        rowIdx = rowIdx + 1
                    Loop
                    Dim sumEndRow As Long
                    sumEndRow = rowIdx - 1 ' 求和结束行

                    ' 在S到Z列的sumRow行直接输出SUM的求和结果，并清除前面查询输出
                    Dim k As Integer
                    For k = 0 To 7
                        Dim colIdx As Integer
                        colIdx = 19 + k ' S=19, Z=26
                        If colIdx <= ws.Columns.Count Then
                            ' 计算并输出求和结果
                            ws.Cells(sumRow, colIdx).Value = Application.WorksheetFunction.Sum(ws.Range(ws.Cells(sumStartRow, colIdx), ws.Cells(sumEndRow, colIdx)))
                            ' 清除本段查询输出
                            ws.Range(ws.Cells(sumStartRow, colIdx), ws.Cells(sumEndRow, colIdx)).ClearContents
                        End If
                    Next k

                    ' 继续查找下一个段
                Loop

                ' 保存并关闭当前工作簿
                wb.Save
                wb.Close
            Else
                MsgBox "文件 " & selectedFile & " 不包含名为'" & sheetName & "'的工作表，程序将跳过该文件。", vbExclamation
            End If
        Next selectedFile
    Else
        MsgBox "未选择文件。", vbExclamation
        End ' 未选择文件时结束程序
    End If

    Set fd = Nothing
End Sub