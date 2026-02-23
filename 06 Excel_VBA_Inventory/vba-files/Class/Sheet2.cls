VERSION 1.0 CLASS
BEGIN
  MultiUse = -1  'True
END
Attribute VB_Name = "Sheet2"
Attribute VB_GlobalNameSpace = False
Attribute VB_Creatable = False
Attribute VB_PredeclaredId = True
Attribute VB_Exposed = True

Public Sub cmdConfirm()
    nCurrentRow = ActiveCell.Row    ' 获取当前活动单元格的行号

    If nCurrentRow = 1 Then
        response = MsgBox("请选择具体要查询存货编码所在行！")    ' 如果在第一行，弹窗提示
     Exit Sub
    End If

    If Date > #1/1/2019# Then
     Exit Sub    ' 日期大于2019年则退出
    End If

    If Not IsEmpty(Sheets("存货目录").Cells(nCurrentRow, 10).Value) Then
        tmpRow = ActiveCell.Row
        tmpCol = ActiveCell.Column
        response = MsgBox("该货号已停用，请另选货号")    ' 如果第10列有值，提示货号已停用
        Sheets("存货目录").Cells(tmpRow, tmpCol).Select
     Exit Sub
    End If

    If IsEmpty(Sheets("存货目录").Cells(nCurrentRow, 1).Value) Then
        tmpRow = ActiveCell.Row
        tmpCol = ActiveCell.Column
        response = MsgBox("请选择有存货编号行！")    ' 如果第1列为空，提示必须选择有存货编号的行
        Sheets("存货目录").Cells(tmpRow, tmpCol).Select
     Exit Sub
    End If

    Sheets("清单制作").Cells(nFindRow, 1).Value = nFindRow - 2    ' 填写序号
    Sheets("清单制作").Cells(nFindRow, 2).Value = Sheets("存货目录").Cells(nCurrentRow, 1).Value    ' 填写存货编码
    Sheets("清单制作").Cells(nFindRow, 3).Value = Sheets("存货目录").Cells(nCurrentRow, 2).Value    ' 填写存货名称
    Sheets("清单制作").Cells(nFindRow, 4).Value = Sheets("存货目录").Cells(nCurrentRow, 3).Value    ' 填写规格
    Sheets("清单制作").Cells(nFindRow, 5).Value = Sheets("存货目录").Cells(nCurrentRow, 4).Value    ' 填写单位

    Sheets("清单制作").Cells(nFindRow, 6).FormulaR1C1 = "=SUMIFS(现存量!C[1],现存量!C[-3],清单制作!RC[-4] )"    ' 填写库存量公式

    'Sheets("清单制作").Cells(nFindRow, 7).FormulaR1C1 = "=SUMIFS(现存量!C7,现存量!C3,清单制作!RC2,现存量!C1,清单制作!R1C)"
    'Sheets("清单制作").Cells(nFindRow, 8).FormulaR1C1 = "=SUMIFS(现存量!C7,现存量!C3,清单制作!RC2,现存量!C1,清单制作!R1C)"
    'Sheets("清单制作").Cells(nFindRow, 9).FormulaR1C1 = "=SUMIFS(现存量!C7,现存量!C3,清单制作!RC2,现存量!C1,清单制作!R1C)"
    'Sheets("清单制作").Cells(nFindRow, 10).FormulaR1C1 = "=SUMIFS(现存量!C7,现存量!C3,清单制作!RC2,现存量!C1,清单制作!R1C)"



    Sheets("清单制作").Select ' 选中清单制作表
    'Sheets("清单制作").Cells(nFindRow, 7).Select
    '    With Selection.Interior
    '        .Pattern = xlSolid
    '        .PatternColorIndex = xlAutomatic
    '        .ThemeColor = xlThemeColorDark1
    '        .TintAndShade = -0.149998474074526
    '        .PatternTintAndShade = 0
    '    End With

    Sheets("清单制作").Cells(nFindRow + 1, 4).Select    ' 选中下一行的第4列
End Sub

Public Sub cmdRefresh()
    GetDictDate    ' 调用获取数据过程

    Sheets("存货目录").Select
    Columns("J:K").Select
    Selection.NumberFormatLocal = "yy-mm-dd"    ' 设置J、K列日期格式
    Columns("L:L").Select
    Selection.NumberFormatLocal = "#,##0.0000_ ;[红色]-#,##0.0000 "    ' 设置L列数字格式
    Columns("B:B").ColumnWidth = 28    ' 设置B列宽度
    Columns("C:C").ColumnWidth = 24    ' 设置C列宽度
    Rows("1:1").RowHeight = 54.75      ' 设置第一行行高
    Range("B2").Select                 ' 选中B2单元格
End Sub

Public Sub cmdReturn()
    Sheets("清单制作").Select    ' 切换到清单制作表
    Sheets("清单制作").Cells(nFindRow, nFindCol).Select    ' 选中指定单元格
End Sub

Public Sub cmdStop()
    ' 判断第10字段筛选状态，切换筛选
    If ActiveSheet.ListObjects("表_Inv").AutoFilter.Filters(10).On Then
        ActiveSheet.ListObjects("表_Inv").Range.AutoFilter Field:=10    ' 取消筛选
    Else
        ActiveSheet.ListObjects("表_Inv").Range.AutoFilter Field:=10, Criteria1:="="    ' 筛选为空
    End If
End Sub

