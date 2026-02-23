Attribute VB_Name = "模块2"
Sub 宏1()
Attribute 宏1.VB_ProcData.VB_Invoke_Func = " \n14"
    '
    ' 宏1 宏
    '
    ActiveCell.FormulaR1C1 = "=COUNTIFS(存货目录!B:B,清单制作!C67,存货目录!C:C,清单制作!D67)"    ' 在当前单元格插入COUNTIFS公式，统计符合条件的数量
    With ActiveCell.Characters(Start:=1, Length:=10).Font    ' 设置公式前10个字符的字体
        .Name = "Times New Roman"    ' 字体为Times New Roman
        .FontStyle = "常规"          ' 字体样式为常规
        .Size = 10                  ' 字号为10
        .Strikethrough = False      ' 不加删除线
        .Superscript = False        ' 不为上标
        .Subscript = False          ' 不为下标
        .OutlineFont = False        ' 不为轮廓字体
        .Shadow = False             ' 不加阴影
        .Underline = xlUnderlineStyleNone    ' 不加下划线
        .ThemeColor = xlThemeColorLight1     ' 主题颜色
        .TintAndShade = 0                   ' 色调
        .ThemeFont = xlThemeFontNone         ' 主题字体
    End With
    With ActiveCell.Characters(Start:=11, Length:=4).Font    ' 设置第11到14个字符的字体
        .Name = "宋体"
        .FontStyle = "常规"
        .Size = 10
        .Strikethrough = False
        .Superscript = False
        .Subscript = False
        .OutlineFont = False
        .Shadow = False
        .Underline = xlUnderlineStyleNone
        .ThemeColor = xlThemeColorLight1
        .TintAndShade = 0
        .ThemeFont = xlThemeFontNone
    End With
    With ActiveCell.Characters(Start:=15, Length:=14).Font    ' 设置第15到28个字符的字体
        .Name = "Times New Roman"
        .FontStyle = "常规"
        .Size = 10
        .Strikethrough = False
        .Superscript = False
        .Subscript = False
        .OutlineFont = False
        .Shadow = False
        .Underline = xlUnderlineStyleNone
        .ThemeColor = xlThemeColorLight1
        .TintAndShade = 0
        .ThemeFont = xlThemeFontNone
    End With
    With ActiveCell.Characters(Start:=29, Length:=4).Font    ' 设置第29到32个字符的字体
        .Name = "宋体"
        .FontStyle = "常规"
        .Size = 10
        .Strikethrough = False
        .Superscript = False
        .Subscript = False
        .OutlineFont = False
        .Shadow = False
        .Underline = xlUnderlineStyleNone
        .ThemeColor = xlThemeColorLight1
        .TintAndShade = 0
        .ThemeFont = xlThemeFontNone
    End With
    With ActiveCell.Characters(Start:=33, Length:=14).Font    ' 设置第33到46个字符的字体
        .Name = "Times New Roman"
        .FontStyle = "常规"
        .Size = 10
        .Strikethrough = False
        .Superscript = False
        .Subscript = False
        .OutlineFont = False
        .Shadow = False
        .Underline = xlUnderlineStyleNone
        .ThemeColor = xlThemeColorLight1
        .TintAndShade = 0
        .ThemeFont = xlThemeFontNone
    End With
    ActiveCell.Select    ' 选中当前单元格
    Selection.NumberFormatLocal = "G/通用格式"    ' 设置单元格格式为通用格式
    ActiveCell.FormulaR1C1 = "=COUNTIFS(存货目录!C,清单制作!RC[1],存货目录!C[1],清单制作!RC[2])"    ' 再次插入COUNTIFS公式，使用相对引用
    ActiveCell.Select    ' 选中当前单元格
End Sub
Sub 宏3()
Attribute 宏3.VB_ProcData.VB_Invoke_Func = " \n14"
    '
    ' 宏3 宏
    '
    Selection.FormulaArray = "=INDEX(表_Inv[编码],MATCH(""*""&RC[1]&""*""&RC[2]&""*"",表_Inv[名称]&表_Inv[规格],0),1)"    ' 在选中区域插入数组公式，查找编码
End Sub
Sub 宏4()
Attribute 宏4.VB_ProcData.VB_Invoke_Func = " \n14"
    '
    ' 宏4 宏
    '
    ActiveCell.Offset(0, 1).Range("A1:B2").Select    ' 选中当前单元格右侧的A1:B2区域
    Selection.Copy                                    ' 复制选中区域
    ActiveCell.Offset(0, 5).Range("A1").Select        ' 选中当前单元格右移5列的A1单元格
    ActiveSheet.Paste                                 ' 粘贴内容
End Sub
Sub 宏5()
Attribute 宏5.VB_ProcData.VB_Invoke_Func = " \n14"
    '
    ' 宏5 宏
    '
    ActiveCell.Offset(3, -1).Range("A1:A3").Select    ' 选中当前单元格下移3行左移1列的A1:A3区域
    Selection.Copy                                    ' 复制选中区域
    ActiveCell.Offset(0, 5).Range("A1").Select        ' 选中当前单元格右移5列的A1单元格
    ActiveSheet.Paste                                 ' 粘贴内容
    Application.CutCopyMode = False                   ' 取消剪贴板模式
End Sub
Sub 宏6()
Attribute 宏6.VB_ProcData.VB_Invoke_Func = " \n14"
    '
    ' 宏6 宏
    '
    ActiveCell.Offset(-1, 1).Range("A1:A2").Select    ' 选中当前单元格上移1行右移1列的A1:A2区域
    Selection.ClearContents                           ' 清除选中区域内容
End Sub
    '
    ' 宏6 宏
    '

    '
    ActiveCell.Offset(-1, 1).Range("A1:A2").Select
    Selection.ClearContents
End Sub
