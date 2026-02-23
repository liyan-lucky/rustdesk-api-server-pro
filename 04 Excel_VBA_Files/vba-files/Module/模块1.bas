Option Explicit

' 全局变量定义
Private gFileCount As Long, gFolderCount As Long, gStartTime As Double
Private gCancelled As Boolean, gMaxLevel As Integer, gSequenceNumber As Long
Private gCurrentRow As Long, gTotalItemsProcessed As Long, gPermissionDeniedCount As Long
Private gUseHyperlinks As Boolean
Private gGroupStack As Collection, gCurrentLevelStack As Collection

' 常量定义
Private Const COL_SEQUENCE As Integer = 1, COL_FIRST_LEVEL As Integer = 2
Private Const COL_NAME_OFFSET As Integer = 2

' =============================================
' 主程序入口
' =============================================
Sub GetFilePaths()
    On Error GoTo ErrorHandler
    
    If ActiveSheet Is Nothing Then
        MsgBox "无法访问当前工作表", vbExclamation
        Exit Sub
    End If
    
    Call InitializeVariables
    Call ClearWorksheetCompletely
    
    Dim folderPath As String
    folderPath = SelectFolderDialog
    If folderPath = "" Then Exit Sub
    
    Call OptimizeSettings(True)
    
    Dim objFSO As Object, objFolder As Object
    Set objFSO = CreateObject("Scripting.FileSystemObject")
    If objFSO Is Nothing Then
        MsgBox "无法创建文件系统对象", vbExclamation
        GoTo CleanUp
    End If
    
    Set objFolder = GetFolderWithErrorHandling(objFSO, folderPath)
    If objFolder Is Nothing Then GoTo CleanUp
    
    Call SetupHeaders
    Call ProcessFolderWithPermissionHandling(objFolder, 1)
    
    If gCancelled Then
        MsgBox "处理已取消", vbInformation
        GoTo CleanUp
    End If
    
    If gCurrentRow <= 1 Then
        MsgBox "未找到任何文件或文件夹", vbInformation
        GoTo CleanUp
    End If
    
    Call ApplyFinalFormatting
    Call ShowCompletionMessage
    
CleanUp:
    Call OptimizeSettings(False)
    Set objFSO = Nothing
    Set objFolder = Nothing
    Application.StatusBar = False
    Exit Sub
    
ErrorHandler:
    Application.StatusBar = "发生错误，正在恢复..."
    DoEvents
    MsgBox "错误：" & Err.Description, vbExclamation
    Resume CleanUp
End Sub

' =============================================
' 核心功能函数
' =============================================

Sub InitializeVariables()
    gCancelled = False
    gFileCount = 0
    gFolderCount = 0
    gStartTime = Timer
    gMaxLevel = 1
    gSequenceNumber = 0
    gCurrentRow = 1
    gTotalItemsProcessed = 0
    gPermissionDeniedCount = 0
    gUseHyperlinks = True
    Set gGroupStack = New Collection
    Set gCurrentLevelStack = New Collection
End Sub

Sub OptimizeSettings(enable As Boolean)
    With Application
        If enable Then
            .ScreenUpdating = False
            .Calculation = xlCalculationManual
            .EnableEvents = False
            .DisplayAlerts = False
        Else
            .EnableEvents = True
            .Calculation = xlCalculationAutomatic
            .DisplayAlerts = True
            DoEvents
            .ScreenUpdating = True
        End If
    End With
End Sub

Sub SetupHeaders()
    With ActiveSheet
        .Cells.Clear
        .Cells(1, 1).Value = "序号"
        .Cells(1, 2).Value = "选择的文件夹"
        .Cells(1, 3).Value = "名称"
        .Cells(1, 4).Value = "完整路径"
        
        With .Range("A1:D1")
            .Font.Bold = True
            .Interior.Color = RGB(200, 200, 200)
            .HorizontalAlignment = xlCenter
        End With
        
        .Columns(1).ColumnWidth = 8
        .Columns(2).ColumnWidth = 15
        .Columns(3).ColumnWidth = 35
        .Columns(4).ColumnWidth = 50
        .Range("A1").Select
    End With
End Sub

' =============================================
' 文件夹处理核心函数
' =============================================

Sub ProcessFolderWithPermissionHandling(objFolder As Object, currentLevel As Integer)
    If objFolder Is Nothing Or gCancelled Then Exit Sub
    
    gTotalItemsProcessed = gTotalItemsProcessed + 1
    If gTotalItemsProcessed Mod 50 = 0 Then
        Call UpdateProgressStatus(objFolder.Name)
    End If
    
    If currentLevel > gMaxLevel Then
        Call InsertLevelColumns(currentLevel)
        gMaxLevel = currentLevel
    End If
    
    Call StartNewGroup(currentLevel, gCurrentRow + 1)
    
    Dim files() As Object, fileCount As Long
    fileCount = CollectFilesWithPermissionHandling(objFolder, files)
    
    If fileCount > 0 Or objFolder.SubFolders.Count > 0 Then
        Call WriteFolderData(objFolder, currentLevel, files, fileCount)
    End If
    
    If gCancelled Then Exit Sub
    Call ProcessSubFoldersWithPermissionHandling(objFolder, currentLevel)
    Call EndCurrentGroup(gCurrentRow)
End Sub

Function CollectFilesWithPermissionHandling(objFolder As Object, ByRef files() As Object) As Long
    On Error Resume Next
    Dim fileCount As Long
    fileCount = 0
    ReDim files(1 To 200) As Object
    
    Dim objFile As Object
    For Each objFile In objFolder.Files
        If Err.Number <> 0 Then
            Err.Clear
            GoTo NextFile
        End If
        
        If Not objFile Is Nothing Then
            fileCount = fileCount + 1
            If fileCount > UBound(files) Then
                ReDim Preserve files(1 To UBound(files) + 200)
            End If
            Set files(fileCount) = objFile
        End If
        
NextFile:
    Next objFile
    
    If fileCount > 0 Then
        ReDim Preserve files(1 To fileCount) As Object
    Else
        ReDim files(0) As Object
    End If
    
    CollectFilesWithPermissionHandling = fileCount
End Function

Sub ProcessSubFoldersWithPermissionHandling(objFolder As Object, currentLevel As Integer)
    On Error Resume Next
    
    Dim objSubFolder As Object
    For Each objSubFolder In objFolder.SubFolders
        If Err.Number <> 0 Then
            gPermissionDeniedCount = gPermissionDeniedCount + 1
            Err.Clear
            GoTo NextSubFolder
        End If
        
        If Not objSubFolder Is Nothing Then
            If ShouldSkipFolder(objSubFolder.Name) Then
                gPermissionDeniedCount = gPermissionDeniedCount + 1
                GoTo NextSubFolder
            End If
            
            Call ProcessFolderWithPermissionHandling(objSubFolder, currentLevel + 1)
            If gCancelled Then Exit Sub
        End If
        
NextSubFolder:
    Next objSubFolder
End Sub

Function ShouldSkipFolder(folderName As String) As Boolean
    Dim skipFolders As Variant
    skipFolders = Array("System Volume Information", "Recovery", "$Recycle.Bin", _
                       "Config.Msi", "Windows", "Program Files", "Program Files (x86)")
    
    Dim i As Integer
    For i = LBound(skipFolders) To UBound(skipFolders)
        If LCase(folderName) = LCase(skipFolders(i)) Then
            ShouldSkipFolder = True
            Exit Function
        End If
    Next i
End Function

' =============================================
' 数据写入函数
' =============================================

Sub WriteFolderData(objFolder As Object, currentLevel As Integer, files() As Object, fileCount As Long)
    If gCancelled Then Exit Sub
    
    If (gFileCount + gFolderCount) Mod 100 = 0 Then
        Call UpdateProgressStatus(objFolder.Name)
    End If
    
    Call WriteFolderRow(objFolder, currentLevel)
    If gCancelled Then Exit Sub
    
    Dim i As Long
    For i = 1 To fileCount
        If Not files(i) Is Nothing Then
            Call WriteFileRow(files(i))
            If gCancelled Then Exit Sub
        End If
    Next i
    
    If (gFileCount + gFolderCount) Mod 200 = 0 Then
        DoEvents
    End If
End Sub

Sub WriteFolderRow(objFolder As Object, currentLevel As Integer)
    gCurrentRow = gCurrentRow + 1
    gFolderCount = gFolderCount + 1
    
    If gCurrentRow >= 1048576 Then
        MsgBox "已达到Excel最大行数限制", vbExclamation
        gCancelled = True
        Exit Sub
    End If
    
    Dim folderName As String
    folderName = "[" & objFolder.Name & "]"
    
    With ActiveSheet
        .Rows(gCurrentRow).ClearContents
        .Cells(gCurrentRow, currentLevel + 1).Value = folderName
        
        Dim nameCol As Integer, pathCol As Integer
        nameCol = GetNameColumn()
        pathCol = GetPathColumn()
        
        .Cells(gCurrentRow, nameCol).Value = folderName
        .Cells(gCurrentRow, pathCol).Value = objFolder.Path
        
        With .Rows(gCurrentRow)
            .Font.Bold = True
            .Interior.Color = GetLevelColor(currentLevel)
        End With
        
        If gUseHyperlinks Then
            On Error Resume Next
            .Hyperlinks.Add _
                Anchor:=.Cells(gCurrentRow, pathCol), _
                Address:=objFolder.Path, _
                TextToDisplay:=objFolder.Path
        End If
    End With
End Sub

Sub WriteFileRow(objFile As Object)
    gCurrentRow = gCurrentRow + 1
    gFileCount = gFileCount + 1
    gSequenceNumber = gSequenceNumber + 1
    
    If gCurrentRow >= 1048576 Then
        MsgBox "已达到Excel最大行数限制", vbExclamation
        gCancelled = True
        Exit Sub
    End If
    
    With ActiveSheet
        .Rows(gCurrentRow).ClearContents
        .Cells(gCurrentRow, COL_SEQUENCE).Value = gSequenceNumber
        
        Dim nameCol As Integer, pathCol As Integer
        nameCol = GetNameColumn()
        pathCol = GetPathColumn()
        
        .Cells(gCurrentRow, nameCol).Value = objFile.Name
        .Cells(gCurrentRow, pathCol).Value = objFile.Path
        
        If gUseHyperlinks Then
            On Error Resume Next
            .Hyperlinks.Add _
                Anchor:=.Cells(gCurrentRow, pathCol), _
                Address:=objFile.Path, _
                TextToDisplay:=objFile.Path
        End If
    End With
    
    If gFileCount Mod 500 = 0 Then
        Application.ScreenUpdating = True
        Call RealTimeScroll(gCurrentRow)
        Application.ScreenUpdating = False
    End If
End Sub

' =============================================
' 分组管理函数
' =============================================

Sub StartNewGroup(level As Integer, rowIndex As Long)
    While gCurrentLevelStack.Count >= level
        Call EndCurrentGroup(rowIndex - 1)
    Wend
    
    gGroupStack.Add rowIndex
    gCurrentLevelStack.Add level
End Sub

Sub EndCurrentGroup(endRow As Long)
    If gGroupStack.Count = 0 Then Exit Sub
    
    Dim startRow As Long
    startRow = gGroupStack(gGroupStack.Count)
    
    If endRow > startRow Then
        On Error Resume Next
        Rows(startRow + 1 & ":" & endRow).Group
    End If
    
    gGroupStack.Remove gGroupStack.Count
    gCurrentLevelStack.Remove gCurrentLevelStack.Count
End Sub

Sub EndAllGroups()
    While gGroupStack.Count > 0
        Call EndCurrentGroup(gCurrentRow)
    Wend
End Sub

' =============================================
' 工具函数
' =============================================

Sub UpdateProgressStatus(folderName As String)
    Application.StatusBar = "扫描文件夹结构中..." & _
                          " 已处理: " & gTotalItemsProcessed & _
                          " 跳过: " & gPermissionDeniedCount & _
                          " 当前: " & Left(folderName, 25)
End Sub

Function SelectFolderDialog() As String
    On Error Resume Next
    With Application.FileDialog(msoFileDialogFolderPicker)
        .Title = "请选择要遍历的文件夹"
        .InitialFileName = CreateObject("WScript.Shell").SpecialFolders("Desktop")
        If .Show = -1 Then
            SelectFolderDialog = .SelectedItems(1)
        End If
    End With
End Function

Function GetFolderWithErrorHandling(objFSO As Object, folderPath As String) As Object
    On Error Resume Next
    Set GetFolderWithErrorHandling = objFSO.GetFolder(folderPath)
    If Err.Number <> 0 Then
        MsgBox "无法访问文件夹: " & Err.Description, vbExclamation
        Set GetFolderWithErrorHandling = Nothing
    End If
End Function

Function GetLevelColor(level As Integer) As Long
    Select Case level
        Case 1: GetLevelColor = RGB(255, 230, 230)
        Case 2: GetLevelColor = RGB(230, 255, 230)
        Case 3: GetLevelColor = RGB(230, 230, 255)
        Case 4: GetLevelColor = RGB(255, 255, 200)
        Case 5: GetLevelColor = RGB(255, 200, 255)
        Case Else: GetLevelColor = RGB(200, 255, 255)
    End Select
End Function

Sub InsertLevelColumns(targetLevel As Integer)
    Dim columnsToInsert As Integer
    columnsToInsert = targetLevel - gMaxLevel
    
    If columnsToInsert <= 0 Then Exit Sub
    
    Dim insertPosition As Integer
    insertPosition = GetNameColumn()
    
    Columns(insertPosition).Resize(, columnsToInsert).Insert Shift:=xlToRight
    
    Dim i As Integer
    For i = 1 To columnsToInsert
        Dim colIndex As Integer
        colIndex = insertPosition + i - 1
        
        Cells(1, colIndex).Value = "层级" & (gMaxLevel + i)
        
        With Cells(1, colIndex)
            .Font.Bold = True
            .Interior.Color = RGB(200, 200, 200)
            .HorizontalAlignment = xlCenter
        End With
        
        Columns(colIndex).ColumnWidth = 15
    Next i
    
    gMaxLevel = targetLevel
    Call UpdateNamePathHeaders
End Sub

Sub UpdateNamePathHeaders()
    Dim nameCol As Integer, pathCol As Integer
    nameCol = GetNameColumn()
    pathCol = GetPathColumn()
    
    With Cells(1, nameCol)
        .Value = "名称"
        .Font.Bold = True
        .Interior.Color = RGB(200, 200, 200)
        .HorizontalAlignment = xlCenter
    End With
    Columns(nameCol).ColumnWidth = 35
    
    With Cells(1, pathCol)
        .Value = "完整路径"
        .Font.Bold = True
        .Interior.Color = RGB(200, 200, 200)
        .HorizontalAlignment = xlCenter
    End With
    Columns(pathCol).ColumnWidth = 50
End Sub

Function GetNameColumn() As Integer
    GetNameColumn = gMaxLevel + COL_NAME_OFFSET
End Function

Function GetPathColumn() As Integer
    GetPathColumn = gMaxLevel + COL_NAME_OFFSET + 1
End Function

Sub RealTimeScroll(rowIndex As Long)
    If rowIndex < 2 Then Exit Sub
    
    With ActiveWindow
        Dim visibleRows As Long
        visibleRows = .VisibleRange.Rows.Count
        
        Dim currentBottom As Long
        currentBottom = .ScrollRow + visibleRows - 1
        
        If rowIndex >= currentBottom - 2 Then
            Dim newScrollRow As Long
            newScrollRow = rowIndex - visibleRows + 5
            If newScrollRow < 1 Then newScrollRow = 1
            .ScrollRow = newScrollRow
        End If
    End With
End Sub

' =============================================
' 清理和格式化函数
' =============================================

Sub ClearWorksheetCompletely()
    On Error Resume Next
    
    Application.ScreenUpdating = True
    
    With ActiveSheet
        If .AutoFilterMode Then
            .AutoFilterMode = False
        End If
        
        .Cells.Clear
        .Hyperlinks.Delete
        
        Call ForceClearAllGrouping
        
        .Columns.ColumnWidth = 8.38
        .Rows.RowHeight = 13.5
        .Range("A1").Select
    End With
    
    Set gGroupStack = New Collection
    Set gCurrentLevelStack = New Collection
    
    Application.ScreenUpdating = False
End Sub

Sub ForceClearAllGrouping()
    On Error Resume Next
    
    ActiveSheet.Outline.ShowLevels RowLevels:=8, ColumnLevels:=8
    ActiveSheet.Outline.Clear
    
    ' 清理行分组
    Dim maxRowLevel As Integer
    maxRowLevel = GetMaxRowGroupingLevel()
    
    If maxRowLevel > 0 Then
        Dim i As Integer
        For i = maxRowLevel To 1 Step -1
            Rows.Ungroup
        Next i
    Else
        For i = 1 To 10
            Rows.Ungroup
        Next i
    End If
    
    ' 清理列分组
    Dim maxColLevel As Integer
    maxColLevel = GetMaxColumnGroupingLevel()
    
    If maxColLevel > 0 Then
        For i = maxColLevel To 1 Step -1
            Columns.Ungroup
        Next i
    Else
        For i = 1 To 5
            Columns.Ungroup
        Next i
    End If
    
    Cells.Select
    Selection.Ungroup
    ActiveSheet.Outline.Clear
End Sub

Function GetMaxRowGroupingLevel() As Integer
    On Error Resume Next
    GetMaxRowGroupingLevel = ActiveSheet.Outline.SummaryRow
End Function

Function GetMaxColumnGroupingLevel() As Integer
    On Error Resume Next
    GetMaxColumnGroupingLevel = ActiveSheet.Outline.SummaryColumn
End Function

Sub ApplyFinalFormatting()
    If gCurrentRow <= 1 Then Exit Sub
    
    Application.ScreenUpdating = True
    Call EndAllGroups
    
    If gMaxLevel > 1 Then
        Call SafeAddLevelColumnGrouping
        ActiveSheet.Outline.ShowLevels ColumnLevels:=1
    End If
    
    Call ApplyAutoFilter
    Call SmartAdjustColumnWidths
    
    ActiveWindow.FreezePanes = False
    If gCurrentRow >= 2 Then
        ActiveSheet.Rows("2:2").Select
        ActiveWindow.FreezePanes = True
    End If
    
    ActiveSheet.Range("A1").Select
End Sub

Sub SafeAddLevelColumnGrouping()
    If gMaxLevel <= 1 Then Exit Sub
    On Error Resume Next
    ActiveSheet.Columns.Ungroup
    ActiveSheet.Outline.Clear
    
    Dim i As Integer
    For i = 2 To gMaxLevel + 1
        Columns(i).Group
    Next i
End Sub

Sub ApplyAutoFilter()
    On Error Resume Next
    If ActiveSheet.AutoFilterMode Then ActiveSheet.AutoFilterMode = False
    Dim lastCol As Integer
    lastCol = GetPathColumn()
    If lastCol > 0 And gCurrentRow > 1 Then
        ActiveSheet.Range("A1").Resize(1, lastCol).AutoFilter
    End If
End Sub

Sub SmartAdjustColumnWidths()
    On Error Resume Next
    Dim pathCol As Integer
    pathCol = GetPathColumn()
    Columns(pathCol).AutoFit
    If Columns(pathCol).ColumnWidth > 80 Then
        Columns(pathCol).ColumnWidth = 80
    End If
End Sub

Sub ShowCompletionMessage()
    Dim elapsedTime As Double
    elapsedTime = Timer - gStartTime
    
    Dim msg As String
    msg = "文件路径生成完成！" & vbCrLf & _
          "文件夹: " & gFolderCount & vbCrLf & _
          "文件: " & gFileCount & vbCrLf & _
          "总计: " & (gFileCount + gFolderCount) & " 个项目" & vbCrLf & _
          "最大层级: " & gMaxLevel & vbCrLf & _
          "耗时: " & Format(elapsedTime, "0.0") & " 秒"
    
    If gPermissionDeniedCount > 0 Then
        msg = msg & vbCrLf & "跳过 " & gPermissionDeniedCount & " 个无权限文件夹"
    End If
    
    MsgBox msg, vbInformation, "完成"
End Sub

' =============================================
' 辅助函数
' =============================================

Sub CancelProcessing()
    If MsgBox("确定要取消当前操作吗？", vbYesNo + vbQuestion) = vbYes Then
        gCancelled = True
        Application.StatusBar = "正在取消操作..."
    End If
End Sub