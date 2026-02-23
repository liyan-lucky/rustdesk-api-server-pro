Attribute VB_Name = "模块4"

' 定义一个名为"宏2"的子过程（也就是宏）
Sub 宏2()
    ' 这是VBA自动生成的属性声明，用于指定宏的调用方式，通常不需要手动修改
    ' 宏2 宏

    ' 核心功能：显示当前活动工作表中的所有数据
    ' 作用：取消任何筛选状态，显示被隐藏的行
    ActiveSheet.ShowAllData    ' 显示所有数据，取消筛选隐藏

    ' 结束子过程的定义
End Sub
Sub 宏7()
Attribute 宏7.VB_ProcData.VB_Invoke_Func = " \n14"
    ' 宏7
    Range("D2").Select    ' 选中D2单元格
    Selection.AutoFilter  ' 对选中区域应用自动筛选
    Range("表_ERP_最低库存汇总表_1[[#Headers],[仓库编号]]").Select    ' 选中表头的仓库编号列
    Selection.AutoFilter  ' 对选中区域应用自动筛选
    Selection.AutoFilter  ' 再次应用自动筛选（重复）
End Sub
