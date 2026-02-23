Attribute VB_Name = "模块3"
Sub Macro1()
Attribute Macro1.VB_ProcData.VB_Invoke_Func = " \n14"
'
' Macro1 宏
'
' 空宏，无操作
'
End Sub
Sub Macro2()
Attribute Macro2.VB_ProcData.VB_Invoke_Func = " \n14"
'
' Macro2 宏
'
'
    ActiveCell.FormulaR1C1 = _
        "=IF(表_ERP_最低库存汇总表_1[[#This Row],[库存量]]<表_ERP_最低库存汇总表_1[[#This Row],[最低库存]],表_ERP_最低库存汇总表_1[[#This Row],[最高库存]]-表_ERP_最低库存汇总表_1[[#This Row],[库存量]],"""")"
        ' 在当前单元格插入IF公式，判断库存量是否低于最低库存，若低则计算补货量，否则为空
    Range("K3").Select    ' 选中K3单元格
End Sub
