VERSION 1.0 CLASS
BEGIN
  MultiUse = -1  'True
END
Attribute VB_Name = "Sheet2"
Attribute VB_GlobalNameSpace = False
Attribute VB_Creatable = False
Attribute VB_PredeclaredId = True
Attribute VB_Exposed = True

' 下面的代码是之前的版本，注释掉了。现在改成了只画行列的外框，去掉了单元格内的边框。
' Private Sub Worksheet_SelectionChange(ByVal Target As Range)
'     Cells.Interior.ColorIndex = xlNone
'     Target.EntireRow.Interior.Color = RGB(255, 255, 200)
'     Target.EntireColumn.Interior.Color = RGB(255, 255, 200)
' End Sub

' ' 下面的代码是新的版本，画了行列的外框，去掉了单元格内的边框。
' '========================
' ' 选择一种范围模式（只改这里）
' ' 1 = 自动UsedRange
' ' 2 = 固定范围（手动写地址）
' ' 3 = 整张表（不建议大表）
' '========================
' Private Const RANGE_MODE As Integer = 2

' ' 固定范围模式用（RANGE_MODE=2时生效）
' Private Const FIXED_RANGE_ADDR As String = "A1:Z1000"

' ' 线条名称（别改，除非你知道自己在干啥）
' Private Const HLINE_NAME As String = "__CrossHair_H"
' Private Const VLINE_NAME As String = "__CrossHair_V"

' Private Sub Worksheet_SelectionChange(ByVal Target As Range)
'     On Error GoTo SafeExit
'     If Target Is Nothing Then Exit Sub

'     Application.ScreenUpdating = False

'     Dim rng As Range
'     Set rng = GetCrossRange()

'     ' 只在范围内才画
'     If Intersect(Target, rng) Is Nothing Then
'         RemoveCrossHair
'         GoTo SafeExit
'     End If

'     DrawCrossHair Target, rng

' SafeExit:
'     Application.ScreenUpdating = True
' End Sub

' Private Function GetCrossRange() As Range
'     Select Case RANGE_MODE
'         Case 1 ' 自动UsedRange（加一点缓冲，避免UsedRange太小）
'             Dim ur As Range
'             Set ur = Me.UsedRange
'             If ur Is Nothing Then
'                 Set GetCrossRange = Me.Range("A1")
'             Else
'                 Set GetCrossRange = Me.Range(ur.Cells(1, 1), ur.Cells(ur.Rows.Count, ur.Columns.Count))
'             End If

'         Case 2 ' 固定范围
'             Set GetCrossRange = Me.Range(FIXED_RANGE_ADDR)

'         Case 3 ' 整表（注意：巨大）
'             Set GetCrossRange = Me.Cells

'         Case Else
'             Set GetCrossRange = Me.UsedRange
'     End Select
' End Function

' Private Sub DrawCrossHair(ByVal Target As Range, ByVal rng As Range)
'     Dim x1 As Double, x2 As Double, y1 As Double, y2 As Double
'     Dim tCell As Range

'     ' 横线：在 Target 这一行，跨 rng 的左右边界
'     Set tCell = Intersect(Me.Rows(Target.Row), rng).Cells(1, 1)
'     If tCell Is Nothing Then Exit Sub

'     x1 = rng.Cells(1, 1).Left
'     x2 = rng.Cells(1, rng.Columns.Count).Left + rng.Cells(1, rng.Columns.Count).Width
'     y1 = Target.Top
'     y2 = Target.Top + Target.Height

'     ' 竖线：在 Target 这一列，跨 rng 的上下边界
'     Dim topCell As Range, bottomCell As Range
'     Set topCell = Intersect(Me.Columns(Target.Column), rng).Cells(1, 1)
'     Set bottomCell = Intersect(Me.Columns(Target.Column), rng).Cells(rng.Rows.Count, 1)

'     Dim vx As Double, vy1 As Double, vy2 As Double
'     vx = Target.Left
'     vy1 = rng.Cells(1, 1).Top
'     vy2 = rng.Cells(rng.Rows.Count, 1).Top + rng.Cells(rng.Rows.Count, 1).Height

'     ' 画两条矩形"框线"（上+下表示横框，左+右表示竖框）
'     ' 用两条线更接近"外框"，且不会出现一格一格

'     ' 横框：上边+下边
'     UpsertLine HLINE_NAME & "_T", x1, y1, x2, y1
'     UpsertLine HLINE_NAME & "_B", x1, y2, x2, y2

'     ' 竖框：左边+右边
'     UpsertLine VLINE_NAME & "_L", vx, vy1, vx, vy2
'     UpsertLine VLINE_NAME & "_R", vx + Target.Width, vy1, vx + Target.Width, vy2
' End Sub

' Private Sub UpsertLine(ByVal nm As String, ByVal x1 As Double, ByVal y1 As Double, ByVal x2 As Double, ByVal y2 As Double)
'     Dim shp As Shape
'     On Error Resume Next
'     Set shp = Me.Shapes(nm)
'     On Error GoTo 0

'     If shp Is Nothing Then
'         Set shp = Me.Shapes.AddLine(x1, y1, x2, y2)
'         shp.Name = nm
'         shp.Placement = xlMove ' 跟随单元格移动（不随缩放也可改 xlMoveAndSize）
'         shp.Line.Weight = 1
'         shp.Line.Transparency = 0
'     Else
'         shp.Left = 0 ' 防止某些情况下Left/Top残留影响
'         shp.Top = 0
'         shp.Width = 0
'         shp.Height = 0
'         shp.SetShapesDefaultProperties
'         shp.Delete ' 保险起见删了重建（更稳定）
'         Set shp = Me.Shapes.AddLine(x1, y1, x2, y2)
'         shp.Name = nm
'         shp.Placement = xlMove
'         shp.Line.Weight = 1
'         shp.Line.Transparency = 0
'     End If
' End Sub

' Private Sub RemoveCrossHair()
'     DeleteShapeIfExists HLINE_NAME & "_T"
'     DeleteShapeIfExists HLINE_NAME & "_B"
'     DeleteShapeIfExists VLINE_NAME & "_L"
'     DeleteShapeIfExists VLINE_NAME & "_R"
' End Sub

' Private Sub DeleteShapeIfExists(ByVal nm As String)
'     On Error Resume Next
'     Me.Shapes(nm).Delete
'     On Error GoTo 0
' End Sub

' 下面的代码是之前的版本，注释掉了。现在改成了只画行列的外框，去掉了单元格内的边框。
'   Private Sub Worksheet_SelectionChange(ByVal Target As Range)
'     Cells.Borders.LineStyle = xlNone

'     Target.EntireRow.Borders(xlEdgeTop).LineStyle = xlContinuous
'     Target.EntireColumn.Borders(xlEdgeLeft).LineStyle = xlContinuous
' End Sub
' ' 下面的代码是之前的版本，注释掉了。现在改成了只画行列的外框，去掉了单元格内的边框。

' Private Const H_NAME As String = "__cross_h"
' Private Const V_NAME As String = "__cross_v"

' Private Sub Worksheet_SelectionChange(ByVal Target As Range)
'     On Error GoTo SafeExit

'     If Target Is Nothing Then Exit Sub
'     Dim t As Range: Set t = Target.Cells(1, 1)

'     Application.ScreenUpdating = False

'     Dim vr As Range
'     Set vr = ActiveWindow.VisibleRange   ' 只画在可见区域，避免超范围

'     Dim h As Shape, v As Shape
'     Set h = GetOrCreateLine(H_NAME)
'     Set v = GetOrCreateLine(V_NAME)

'     ' 横线：Left/Top/Width/Height 控制
'     With h
'         .Visible = msoTrue
'         .Line.Visible = msoTrue
'         .ZOrder msoBringToFront
'         .Left = vr.Left
'         .Top = t.Top
'         .Width = vr.Width
'         .Height = 0.1        ' 用0可能不显示，给个极小值更稳
'         .Line.Weight = 1
'     End With

'     ' 竖线
'     With v
'         .Visible = msoTrue
'         .Line.Visible = msoTrue
'         .ZOrder msoBringToFront
'         .Left = t.Left
'         .Top = vr.Top
'         .Width = 0.1
'         .Height = vr.Height
'         .Line.Weight = 1
'     End With

' SafeExit:
'     Application.ScreenUpdating = True
' End Sub

' Private Function GetOrCreateLine(ByVal nm As String) As Shape
'     On Error Resume Next
'     Set GetOrCreateLine = Me.Shapes(nm)
'     On Error GoTo 0

'     If GetOrCreateLine Is Nothing Then
'         Set GetOrCreateLine = Me.Shapes.AddLine(0, 0, 10, 0)
'         GetOrCreateLine.Name = nm
'         GetOrCreateLine.Placement = xlMove
'     End If
' End Function