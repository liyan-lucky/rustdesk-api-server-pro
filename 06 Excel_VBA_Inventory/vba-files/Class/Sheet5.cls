VERSION 1.0 CLASS
BEGIN
MultiUse = -1  'True
END
Attribute VB_Name = "Sheet5"
Attribute VB_GlobalNameSpace = False
Attribute VB_Creatable = False
Attribute VB_PredeclaredId = True
Attribute VB_Exposed = True
' 类对外公开

Private Sub cmdbtnConfirmItem_Click()
    nCurrentRow = ActiveCell.Row    ' 获取当前活动单元格的行号

    If Sheets("仓库").Range("D1").Value = 1 Then
        ' 如果D1为1，将仓库表当前行的第1、2列内容写入清单制作表H1、H2
        Sheets("清单制作").Range("H1").Value = Sheets("仓库").Cells(nCurrentRow, 1).Value
        Sheets("清单制作").Range("H2").Value = Sheets("仓库").Cells(nCurrentRow, 2).Value
        Sheets("清单制作").Select    ' 选中清单制作表
        Sheets("清单制作").Range("H2").Select    ' 选中H2单元格
    End If

    If Sheets("仓库").Range("D1").Value = 2 Then
        ' 如果D1为2，将仓库表当前行的第1、2列内容写入清单制作表I1、I2
        Sheets("清单制作").Range("I1").Value = Sheets("仓库").Cells(nCurrentRow, 1).Value
        Sheets("清单制作").Range("I2").Value = Sheets("仓库").Cells(nCurrentRow, 2).Value
        Sheets("清单制作").Select    ' 选中清单制作表
        Sheets("清单制作").Range("I2").Select    ' 选中I2单元格
    End If

    If Sheets("仓库").Range("D1").Value = 3 Then
        ' 如果D1为3，将仓库表当前行的第1、2列内容写入清单制作表J1、J2
        Sheets("清单制作").Range("J1").Value = Sheets("仓库").Cells(nCurrentRow, 1).Value
        Sheets("清单制作").Range("J2").Value = Sheets("仓库").Cells(nCurrentRow, 2).Value
        Sheets("清单制作").Select    ' 选中清单制作表
        Sheets("清单制作").Range("J2").Select    ' 选中J2单元格
    End If

End Sub
