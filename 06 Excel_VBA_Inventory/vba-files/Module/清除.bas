Attribute VB_Name = "清除"

Public Sub CmdDelete()
    Dim ws As Worksheet    ' 定义工作表对象
    Dim lastRow As Long    ' 定义最后一行变量
    Set ws = ActiveSheet   ' 设置ws为当前活动工作表
    Application.ScreenUpdating = False    ' 关闭屏幕刷新，提高效率
    With ws
        lastRow = .Cells(.Rows.Count, 3).End(xlUp).Row    ' 查找第3列最后有数据的行号
        If lastRow > 2 Then
            '.Range("A3:" & .Cells(lastRow, .Columns.Count).Address).ClearContents '所有单元格的内容（但保留单元格的格式、批注等）
            .Range("A3:" & .Cells(lastRow, .Columns.Count).Address).Clear    ' 清除A3到最后一行所有列的内容和格式
        Else
            'MsgBox "工作表数据不足3行，无需清除"    ' 数据不足3行时不清除
        End If
    End With
End Sub

Public Sub CmdDelete2()
    Dim ws As Worksheet    ' 定义工作表对象
    Dim lastRow As Long    ' 定义最后一行变量
    Dim tbl As ListObject   ' 定义表格对象
    Set ws = ActiveSheet   ' 设置ws为当前活动工作表
    Application.ScreenUpdating = False    ' 关闭屏幕刷新，提高效率
    
    With ws
        ' 检查并删除从第二行开始的所有表格
        On Error Resume Next
        For Each tbl In .ListObjects
            If Not Intersect(tbl.Range, .Range("A2")) Is Nothing Then
                ' 只删除表格的数据区域，保留表头
                If Not tbl.DataBodyRange Is Nothing Then
                    tbl.DataBodyRange.Delete
                End If
            End If
        Next tbl
        On Error GoTo 0
        
        lastRow = .Cells(.Rows.Count, 2).End(xlUp).Row    ' 查找第2列最后有数据的行号
        
        ' 从第二行开始清除（保留第一行表头）
        If lastRow >= 2 Then
            ' 彻底清除所有内容和格式（从第二行开始）
            With .Range("A2:" & .Cells(lastRow, .Columns.Count).Address)
                .Clear    ' 清除内容和格式（包括批注）
                .ClearHyperlinks  ' 清除超链接
                .Borders.LineStyle = xlNone  ' 清除边框
            End With
            
            ' 清除条件格式和数据验证（从第二行开始）
            On Error Resume Next
            .Range("A2:" & .Cells(lastRow, .Columns.Count).Address).FormatConditions.Delete
            .Range("A2:" & .Cells(lastRow, .Columns.Count).Address).Validation.Delete
            On Error GoTo 0
        End If
    End With
    
    Application.ScreenUpdating = True    ' 恢复屏幕刷新
End Sub