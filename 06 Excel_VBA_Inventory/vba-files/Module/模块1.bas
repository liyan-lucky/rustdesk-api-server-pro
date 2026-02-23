Attribute VB_Name = "模块1"
Public nFindRow As Long    ' 定义全局变量，查找的行号
Public nFindCol As Long    ' 定义全局变量，查找的列号




Public Sub GetDictDate()
    ' 获取字典数据并填充到各工作表

    Sheets("存货目录").Activate    ' 激活“存货目录”工作表
    Sheets("存货目录").Range("A1").Select    ' 选中A1单元格

    ' 判断是否有筛选，若有则取消筛选
    '   If Sheets("存货目录").ListObjects("表_Inv").AutoFilter.FilterMode Then
    If Not Sheets("存货目录").AutoFilter Is Nothing Then
        Selection.AutoFilter    ' 取消筛选
    End If

    With Selection.Interior
        .Pattern = xlNone    ' 清除单元格填充图案
        .TintAndShade = 0    ' 清除色调
        .PatternTintAndShade = 0    ' 清除图案色调
    End With

    ActiveSheet.UsedRange.Select    ' 选中当前工作表的已用区域
    Selection.ClearContents         ' 清除内容


    ' 查询条件：新编号和停用日为空
    ' WHERE (Inventory.cInvDefine1 Is Null) And (Inventory.dEDate Is Null)

    ' 从SQL Server数据库获取存货目录数据，并填充到表格
    With ActiveSheet.ListObjects.Add(SourceType:=0, Source:= _
        "ODBC;DRIVER=SQL Server;SERVER=ERPServer;UID=sa;PWD=1QAZ2wsx;APP=2007 Microsoft Office system;WSID=HUANGGUOPING;DATABASE=UFDATA_002_2015" _
        , Destination:=Range("$A$1")).QueryTable
        .CommandText = Array( _
        "Select Inventory.cInvCode 编码, Inventory.cInvName 名称, Inventory.cInvStd 规格, ComputationUnit.cComUnitName 主单位, '' 辅助单位," _
        , _
        " ''  换算率,  Inventory.cPosition 货架 ,Inventory.cInvCCode 大类,Inventory.cInvDefine1 新编号 , Inventory.dEDate 停用日, Inventory.dSDate 启用日 " _
        , _
        "FROM UFDATA_002_2015.dbo.ComputationUnit ComputationUnit, UFDATA_002_2015.dbo.Inventory Inventory " _
        , _
        "WHERE (ComputationUnit.cComunitCode = Inventory.cComUnitCode) And (Inventory.dEDate IS NULL) ORDER BY Inventory.cInvCode")
        .RowNumbers = False    ' 不显示行号
        .FillAdjacentFormulas = False    ' 不填充相邻公式
        .PreserveFormatting = True       ' 保留格式
        .RefreshOnFileOpen = False       ' 打开文件时不刷新
        .BackgroundQuery = True          ' 后台查询
        .RefreshStyle = xlInsertDeleteCells    ' 刷新方式：插入或删除单元格
        .SavePassword = False            ' 不保存密码
        .SaveData = True                 ' 保存数据
        .AdjustColumnWidth = True        ' 自动调整列宽
        .RefreshPeriod = 0               ' 刷新周期为0
        .PreserveColumnInfo = True       ' 保留列信息
        .ListObject.DisplayName = "表_Inv"    ' 设置表名
        .Refresh BackgroundQuery:=False       ' 刷新数据
    End With
    ActiveSheet.ListObjects("表_Inv").TableStyle = ""    ' 清除表格样式


    Rows("1:1").RowHeight = 47.25    ' 设置第一行行高

    Columns("B:B").ColumnWidth = 19.83    ' 设置B列宽度
    Columns("C:C").ColumnWidth = 21.83    ' 设置C列宽度
    Columns("D:D").ColumnWidth = 6.5      ' 设置D列宽度
    Columns("I:I").ColumnWidth = 12.5     ' 设置I列宽度
    Range("A2").Select                    ' 选中A2单元格





    Sheets("项目").Activate    ' 激活“项目”工作表
    ActiveSheet.UsedRange.Select    ' 选中已用区域
    Selection.ClearContents         ' 清除内容

    ' 从SQL Server数据库获取项目信息，并填充到表格
    With ActiveSheet.ListObjects.Add(SourceType:=0, Source:= _
        "ODBC;DRIVER=SQL Server;SERVER=ERPServer;UID=sa;PWD=1QAZ2wsx;APP=2007 Microsoft Office system;WSID=HUANGGUOPING;DATABASE=UFDATA_002_2015" _
        , Destination:=Range("项目!$A$1")).QueryTable
        .CommandText = Array( _
        "Select  fitemss00.citemcode 项目编号, fitemss00.citemname 项目名称 , fitemss00.客户单位名称, fitemss00.项目序号 " _
        , _
        " FROM  UFDATA_002_2015.dbo.fitemss00 fitemss00 ORDER BY fitemss00.citemcode ")   ' ,UFDATA_002_2015.dbo.fitem fitem,UFDATA_002_2015.dbo.fitemss00Class fitemss00Class")
        .RowNumbers = False    ' 不显示行号
        .FillAdjacentFormulas = False    ' 不填充相邻公式
        .PreserveFormatting = True       ' 保留格式
        .RefreshOnFileOpen = False       ' 打开文件时不刷新
        .BackgroundQuery = True          ' 后台查询
        .RefreshStyle = xlInsertDeleteCells    ' 刷新方式：插入或删除单元格
        .SavePassword = False            ' 不保存密码
        .SaveData = True                 ' 保存数据
        .AdjustColumnWidth = True        ' 自动调整列宽
        .RefreshPeriod = 0               ' 刷新周期为0
        .PreserveColumnInfo = True       ' 保留列信息
        .ListObject.DisplayName = "表_ExternalData_1"    ' 设置表名
        .Refresh BackgroundQuery:=False       ' 刷新数据
    End With
    ActiveSheet.ListObjects("表_ExternalData_1").TableStyle = ""    ' 清除表格样式
    Range("A1").Select    ' 选中A1单元格


    Sheets("存货类别").Activate    ' 激活“存货类别”工作表
    ActiveSheet.UsedRange.Select    ' 选中已用区域
    Selection.ClearContents         ' 清除内容
    Selection.ClearContents         ' 再次清除内容（重复）

    ' 从SQL Server数据库获取存货类别信息，并填充到表格
    With ActiveSheet.ListObjects.Add(SourceType:=0, Source:= _
        "ODBC;DRIVER=SQL Server;SERVER=ERPServer;UID=sa;PWD=1QAZ2wsx;APP=2007 Microsoft Office system;WSID=HUANGGUOPING;DATABASE=UFDATA_002_2015" _
        , Destination:=Range("存货类别!$A$1")).QueryTable
        .CommandText = Array("Select cInvCName As 类别名, cInvCCode As 类别编号, iInvCGrade, bInvCEnd FROM UFDATA_002_2015.dbo.InventoryClass")
        .RowNumbers = False    ' 不显示行号
        .FillAdjacentFormulas = False    ' 不填充相邻公式
        .PreserveFormatting = True       ' 保留格式
        .RefreshOnFileOpen = False       ' 打开文件时不刷新
        .BackgroundQuery = True          ' 后台查询
        .RefreshStyle = xlInsertDeleteCells    ' 刷新方式：插入或删除单元格
        .SavePassword = False            ' 不保存密码
        .SaveData = True                 ' 保存数据
        .AdjustColumnWidth = True        ' 自动调整列宽
        .RefreshPeriod = 0               ' 刷新周期为0
        .PreserveColumnInfo = True       ' 保留列信息
        .ListObject.DisplayName = "表_存货类别"    ' 设置表名
        .Refresh BackgroundQuery:=False       ' 刷新数据
    End With
    ActiveSheet.ListObjects("表_存货类别").TableStyle = ""    ' 清除表格样式
    Range("A1").Select    ' 选中A1单元格




End Sub
