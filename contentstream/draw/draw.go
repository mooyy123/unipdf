//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

// Package draw has handy features for defining paths which can be used to draw content on a PDF page.  Handles
// defining paths as points, vector calculations and conversion to PDF content stream data which can be used in
// page content streams and XObject forms and thus also in annotation appearance streams.
//
// Also defines utility functions for drawing common shapes such as rectangles, lines and circles (ovals).
package draw ;import (_d "fmt";_ge "github.com/unidoc/unipdf/v3/contentstream";_b "github.com/unidoc/unipdf/v3/core";_e "github.com/unidoc/unipdf/v3/internal/transform";_g "github.com/unidoc/unipdf/v3/model";_ac "math";);

// GetBoundingBox returns the bounding box of the path.
func (_ee Path )GetBoundingBox ()BoundingBox {_cgg :=BoundingBox {};_cc :=0.0;_ded :=0.0;_geg :=0.0;_eee :=0.0;for _fa ,_fdc :=range _ee .Points {if _fa ==0{_cc =_fdc .X ;_ded =_fdc .X ;_geg =_fdc .Y ;_eee =_fdc .Y ;continue ;};if _fdc .X < _cc {_cc =_fdc .X ;
};if _fdc .X > _ded {_ded =_fdc .X ;};if _fdc .Y < _geg {_geg =_fdc .Y ;};if _fdc .Y > _eee {_eee =_fdc .Y ;};};_cgg .X =_cc ;_cgg .Y =_geg ;_cgg .Width =_ded -_cc ;_cgg .Height =_eee -_geg ;return _cgg ;};

// Length returns the number of points in the path.
func (_gg Path )Length ()int {return len (_gg .Points )};

// RemovePoint removes the point at the index specified by number from the
// path. The index is 1-based.
func (_ad Path )RemovePoint (number int )Path {if number < 1||number > len (_ad .Points ){return _ad ;};_bae :=number -1;_ad .Points =append (_ad .Points [:_bae ],_ad .Points [_bae +1:]...);return _ad ;};

// FlipX flips the sign of the Dx component of the vector.
func (_ae Vector )FlipX ()Vector {_ae .Dx =-_ae .Dx ;return _ae };

// FlipY flips the sign of the Dy component of the vector.
func (_add Vector )FlipY ()Vector {_add .Dy =-_add .Dy ;return _add };

// Draw draws the composite Bezier curve. A graphics state name can be
// specified for setting the curve properties (e.g. setting the opacity).
// Otherwise leave empty (""). Returns the content stream as a byte array and
// the curve bounding box.
func (_cd PolyBezierCurve )Draw (gsName string )([]byte ,*_g .PdfRectangle ,error ){if _cd .BorderColor ==nil {_cd .BorderColor =_g .NewPdfColorDeviceRGB (0,0,0);};_dae :=NewCubicBezierPath ();for _ ,_eafg :=range _cd .Curves {_dae =_dae .AppendCurve (_eafg );
};_eb :=_ge .NewContentCreator ();_eb .Add_q ();_cd .FillEnabled =_cd .FillEnabled &&_cd .FillColor !=nil ;if _cd .FillEnabled {_eb .SetNonStrokingColor (_cd .FillColor );};_eb .SetStrokingColor (_cd .BorderColor );_eb .Add_w (_cd .BorderWidth );if len (gsName )> 1{_eb .Add_gs (_b .PdfObjectName (gsName ));
};for _caa ,_ged :=range _dae .Curves {if _caa ==0{_eb .Add_m (_ged .P0 .X ,_ged .P0 .Y );}else {_eb .Add_l (_ged .P0 .X ,_ged .P0 .Y );};_eb .Add_c (_ged .P1 .X ,_ged .P1 .Y ,_ged .P2 .X ,_ged .P2 .Y ,_ged .P3 .X ,_ged .P3 .Y );};if _cd .FillEnabled {_eb .Add_h ();
_eb .Add_B ();}else {_eb .Add_S ();};_eb .Add_Q ();return _eb .Bytes (),_dae .GetBoundingBox ().ToPdfRectangle (),nil ;};

// Draw draws the polyline. A graphics state name can be specified for
// setting the polyline properties (e.g. setting the opacity). Otherwise leave
// empty (""). Returns the content stream as a byte array and the polyline
// bounding box.
func (_baee Polyline )Draw (gsName string )([]byte ,*_g .PdfRectangle ,error ){if _baee .LineColor ==nil {_baee .LineColor =_g .NewPdfColorDeviceRGB (0,0,0);};_bge :=NewPath ();for _ ,_fcdd :=range _baee .Points {_bge =_bge .AppendPoint (_fcdd );};_cfb :=_ge .NewContentCreator ();
_cfb .Add_q ().SetStrokingColor (_baee .LineColor ).Add_w (_baee .LineWidth );if len (gsName )> 1{_cfb .Add_gs (_b .PdfObjectName (gsName ));};DrawPathWithCreator (_bge ,_cfb );_cfb .Add_S ();_cfb .Add_Q ();return _cfb .Bytes (),_bge .GetBoundingBox ().ToPdfRectangle (),nil ;
};

// Draw draws the composite curve polygon. A graphics state name can be
// specified for setting the curve properties (e.g. setting the opacity).
// Otherwise leave empty (""). Returns the content stream as a byte array
// and the bounding box of the polygon.
func (_bc CurvePolygon )Draw (gsName string )([]byte ,*_g .PdfRectangle ,error ){_egd :=_ge .NewContentCreator ();_egd .Add_q ();_bc .FillEnabled =_bc .FillEnabled &&_bc .FillColor !=nil ;if _bc .FillEnabled {_egd .SetNonStrokingColor (_bc .FillColor );
};_bc .BorderEnabled =_bc .BorderEnabled &&_bc .BorderColor !=nil ;if _bc .BorderEnabled {_egd .SetStrokingColor (_bc .BorderColor );_egd .Add_w (_bc .BorderWidth );};if len (gsName )> 1{_egd .Add_gs (_b .PdfObjectName (gsName ));};_ccg :=NewCubicBezierPath ();
for _ ,_bf :=range _bc .Rings {for _gd ,_bace :=range _bf {if _gd ==0{_egd .Add_m (_bace .P0 .X ,_bace .P0 .Y );}else {_egd .Add_l (_bace .P0 .X ,_bace .P0 .Y );};_egd .Add_c (_bace .P1 .X ,_bace .P1 .Y ,_bace .P2 .X ,_bace .P2 .Y ,_bace .P3 .X ,_bace .P3 .Y );
_ccg =_ccg .AppendCurve (_bace );};_egd .Add_h ();};if _bc .FillEnabled &&_bc .BorderEnabled {_egd .Add_B ();}else if _bc .FillEnabled {_egd .Add_f ();}else if _bc .BorderEnabled {_egd .Add_S ();};_egd .Add_Q ();return _egd .Bytes (),_ccg .GetBoundingBox ().ToPdfRectangle (),nil ;
};

// PolyBezierCurve represents a composite curve that is the result of
// joining multiple cubic Bezier curves.
type PolyBezierCurve struct{Curves []CubicBezierCurve ;BorderWidth float64 ;BorderColor _g .PdfColor ;FillEnabled bool ;FillColor _g .PdfColor ;};

// CubicBezierPath represents a collection of cubic Bezier curves.
type CubicBezierPath struct{Curves []CubicBezierCurve ;};

// LineStyle refers to how the line will be created.
type LineStyle int ;

// NewCubicBezierCurve returns a new cubic Bezier curve.
func NewCubicBezierCurve (x0 ,y0 ,x1 ,y1 ,x2 ,y2 ,x3 ,y3 float64 )CubicBezierCurve {_f :=CubicBezierCurve {};_f .P0 =NewPoint (x0 ,y0 );_f .P1 =NewPoint (x1 ,y1 );_f .P2 =NewPoint (x2 ,y2 );_f .P3 =NewPoint (x3 ,y3 );return _f ;};

// GetBoundingBox returns the bounding box of the Bezier path.
func (_efe CubicBezierPath )GetBoundingBox ()Rectangle {_af :=Rectangle {};_bab :=0.0;_bd :=0.0;_efc :=0.0;_efa :=0.0;for _dee ,_fb :=range _efe .Curves {_cf :=_fb .GetBounds ();if _dee ==0{_bab =_cf .Llx ;_bd =_cf .Urx ;_efc =_cf .Lly ;_efa =_cf .Ury ;
continue ;};if _cf .Llx < _bab {_bab =_cf .Llx ;};if _cf .Urx > _bd {_bd =_cf .Urx ;};if _cf .Lly < _efc {_efc =_cf .Lly ;};if _cf .Ury > _efa {_efa =_cf .Ury ;};};_af .X =_bab ;_af .Y =_efc ;_af .Width =_bd -_bab ;_af .Height =_efa -_efc ;return _af ;
};

// Add shifts the coordinates of the point with dx, dy and returns the result.
func (_bac Point )Add (dx ,dy float64 )Point {_bac .X +=dx ;_bac .Y +=dy ;return _bac };

// Circle represents a circle shape with fill and border properties that can be drawn to a PDF content stream.
type Circle struct{X float64 ;Y float64 ;Width float64 ;Height float64 ;FillEnabled bool ;FillColor _g .PdfColor ;BorderEnabled bool ;BorderWidth float64 ;BorderColor _g .PdfColor ;Opacity float64 ;};

// Rotate rotates the vector by the specified angle.
func (_daab Vector )Rotate (phi float64 )Vector {_gfg :=_daab .Magnitude ();_eegb :=_daab .GetPolarAngle ();return NewVectorPolar (_gfg ,_eegb +phi );};

// Add adds the specified vector to the current one and returns the result.
func (_fbd Vector )Add (other Vector )Vector {_fbd .Dx +=other .Dx ;_fbd .Dy +=other .Dy ;return _fbd };

// Draw draws the polygon. A graphics state name can be specified for
// setting the polygon properties (e.g. setting the opacity). Otherwise leave
// empty (""). Returns the content stream as a byte array and the polygon
// bounding box.
func (_cagc Polygon )Draw (gsName string )([]byte ,*_g .PdfRectangle ,error ){_dba :=_ge .NewContentCreator ();_dba .Add_q ();_cagc .FillEnabled =_cagc .FillEnabled &&_cagc .FillColor !=nil ;if _cagc .FillEnabled {_dba .SetNonStrokingColor (_cagc .FillColor );
};_cagc .BorderEnabled =_cagc .BorderEnabled &&_cagc .BorderColor !=nil ;if _cagc .BorderEnabled {_dba .SetStrokingColor (_cagc .BorderColor );_dba .Add_w (_cagc .BorderWidth );};if len (gsName )> 1{_dba .Add_gs (_b .PdfObjectName (gsName ));};_fc :=NewPath ();
for _ ,_daa :=range _cagc .Points {for _ccbe ,_fdfe :=range _daa {_fc =_fc .AppendPoint (_fdfe );if _ccbe ==0{_dba .Add_m (_fdfe .X ,_fdfe .Y );}else {_dba .Add_l (_fdfe .X ,_fdfe .Y );};};_dba .Add_h ();};if _cagc .FillEnabled &&_cagc .BorderEnabled {_dba .Add_B ();
}else if _cagc .FillEnabled {_dba .Add_f ();}else if _cagc .BorderEnabled {_dba .Add_S ();};_dba .Add_Q ();return _dba .Bytes (),_fc .GetBoundingBox ().ToPdfRectangle (),nil ;};

// ToPdfRectangle returns the rectangle as a PDF rectangle.
func (_cfdf Rectangle )ToPdfRectangle ()*_g .PdfRectangle {return &_g .PdfRectangle {Llx :_cfdf .X ,Lly :_cfdf .Y ,Urx :_cfdf .X +_cfdf .Width ,Ury :_cfdf .Y +_cfdf .Height };};const (LineStyleSolid LineStyle =0;LineStyleDashed LineStyle =1;);

// NewVector returns a new vector with the direction specified by dx and dy.
func NewVector (dx ,dy float64 )Vector {_bdc :=Vector {};_bdc .Dx =dx ;_bdc .Dy =dy ;return _bdc };

// DrawBezierPathWithCreator makes the bezier path with the content creator.
// Adds the PDF commands to draw the path to the creator instance.
func DrawBezierPathWithCreator (bpath CubicBezierPath ,creator *_ge .ContentCreator ){for _dag ,_gea :=range bpath .Curves {if _dag ==0{creator .Add_m (_gea .P0 .X ,_gea .P0 .Y );};creator .Add_c (_gea .P1 .X ,_gea .P1 .Y ,_gea .P2 .X ,_gea .P2 .Y ,_gea .P3 .X ,_gea .P3 .Y );
};};

// LineEndingStyle defines the line ending style for lines.
// The currently supported line ending styles are None, Arrow (ClosedArrow) and Butt.
type LineEndingStyle int ;

// NewPath returns a new empty path.
func NewPath ()Path {return Path {}};

// CurvePolygon is a multi-point shape with rings containing curves that can be
// drawn to a PDF content stream.
type CurvePolygon struct{Rings [][]CubicBezierCurve ;FillEnabled bool ;FillColor _g .PdfColor ;BorderEnabled bool ;BorderColor _g .PdfColor ;BorderWidth float64 ;};

// Draw draws the rectangle. A graphics state can be specified for
// setting additional properties (e.g. opacity). Otherwise pass an empty string
// for the `gsName` parameter. The method returns the content stream as a byte
// array and the bounding box of the shape.
func (_dad Rectangle )Draw (gsName string )([]byte ,*_g .PdfRectangle ,error ){_ade :=_ge .NewContentCreator ();_ade .Add_q ();if _dad .FillEnabled {_ade .SetNonStrokingColor (_dad .FillColor );};if _dad .BorderEnabled {_ade .SetStrokingColor (_dad .BorderColor );
_ade .Add_w (_dad .BorderWidth );};if len (gsName )> 1{_ade .Add_gs (_b .PdfObjectName (gsName ));};var (_cac ,_bce =_dad .X ,_dad .Y ;_fcf ,_ga =_dad .Width ,_dad .Height ;_aae =_ac .Abs (_dad .BorderRadiusTopLeft );_ec =_ac .Abs (_dad .BorderRadiusTopRight );
_edf =_ac .Abs (_dad .BorderRadiusBottomLeft );_adb =_ac .Abs (_dad .BorderRadiusBottomRight );_cgc =0.4477;);_gae :=Path {Points :[]Point {{X :_cac +_fcf -_adb ,Y :_bce },{X :_cac +_fcf ,Y :_bce +_ga -_ec },{X :_cac +_aae ,Y :_bce +_ga },{X :_cac ,Y :_bce +_edf }}};
_ebc :=[][7]float64 {{_adb ,_cac +_fcf -_adb *_cgc ,_bce ,_cac +_fcf ,_bce +_adb *_cgc ,_cac +_fcf ,_bce +_adb },{_ec ,_cac +_fcf ,_bce +_ga -_ec *_cgc ,_cac +_fcf -_ec *_cgc ,_bce +_ga ,_cac +_fcf -_ec ,_bce +_ga },{_aae ,_cac +_aae *_cgc ,_bce +_ga ,_cac ,_bce +_ga -_aae *_cgc ,_cac ,_bce +_ga -_aae },{_edf ,_cac ,_bce +_edf *_cgc ,_cac +_edf *_cgc ,_bce ,_cac +_edf ,_bce }};
_ade .Add_m (_cac +_edf ,_bce );for _df :=0;_df < 4;_df ++{_eegg :=_gae .Points [_df ];_ade .Add_l (_eegg .X ,_eegg .Y );_cfd :=_ebc [_df ];if _dcd :=_cfd [0];_dcd !=0{_ade .Add_c (_cfd [1],_cfd [2],_cfd [3],_cfd [4],_cfd [5],_cfd [6]);};};_ade .Add_h ();
if _dad .FillEnabled &&_dad .BorderEnabled {_ade .Add_B ();}else if _dad .FillEnabled {_ade .Add_f ();}else if _dad .BorderEnabled {_ade .Add_S ();};_ade .Add_Q ();return _ade .Bytes (),_gae .GetBoundingBox ().ToPdfRectangle (),nil ;};

// GetPolarAngle returns the angle the magnitude of the vector forms with the
// positive X-axis going counterclockwise.
func (_ced Vector )GetPolarAngle ()float64 {return _ac .Atan2 (_ced .Dy ,_ced .Dx )};const (LineEndingStyleNone LineEndingStyle =0;LineEndingStyleArrow LineEndingStyle =1;LineEndingStyleButt LineEndingStyle =2;);

// Polygon is a multi-point shape that can be drawn to a PDF content stream.
type Polygon struct{Points [][]Point ;FillEnabled bool ;FillColor _g .PdfColor ;BorderEnabled bool ;BorderColor _g .PdfColor ;BorderWidth float64 ;};

// Path consists of straight line connections between each point defined in an array of points.
type Path struct{Points []Point ;};

// NewPoint returns a new point with the coordinates x, y.
func NewPoint (x ,y float64 )Point {return Point {X :x ,Y :y }};

// Point represents a two-dimensional point.
type Point struct{X float64 ;Y float64 ;};

// Rectangle is a shape with a specified Width and Height and a lower left corner at (X,Y) that can be
// drawn to a PDF content stream.  The rectangle can optionally have a border and a filling color.
// The Width/Height includes the border (if any specified), i.e. is positioned inside.
type Rectangle struct{

// Position and size properties.
X float64 ;Y float64 ;Width float64 ;Height float64 ;

// Fill properties.
FillEnabled bool ;FillColor _g .PdfColor ;

// Border properties.
BorderEnabled bool ;BorderColor _g .PdfColor ;BorderWidth float64 ;BorderRadiusTopLeft float64 ;BorderRadiusTopRight float64 ;BorderRadiusBottomLeft float64 ;BorderRadiusBottomRight float64 ;

// Shape opacity (0-1 interval).
Opacity float64 ;};

// CubicBezierCurve is defined by:
// R(t) = P0*(1-t)^3 + P1*3*t*(1-t)^2 + P2*3*t^2*(1-t) + P3*t^3
// where P0 is the current point, P1, P2 control points and P3 the final point.
type CubicBezierCurve struct{P0 Point ;P1 Point ;P2 Point ;P3 Point ;};

// Draw draws the line to PDF contentstream. Generates the content stream which can be used in page contents or
// appearance stream of annotation. Returns the stream content, XForm bounding box (local), bounding box and an error
// if one occurred.
func (_bfe Line )Draw (gsName string )([]byte ,*_g .PdfRectangle ,error ){_ag ,_babe :=_bfe .X1 ,_bfe .X2 ;_gca ,_gge :=_bfe .Y1 ,_bfe .Y2 ;_ccf :=_gge -_gca ;_gcc :=_babe -_ag ;_eba :=_ac .Atan2 (_ccf ,_gcc );L :=_ac .Sqrt (_ac .Pow (_gcc ,2.0)+_ac .Pow (_ccf ,2.0));
_bbg :=_bfe .LineWidth ;_ce :=_ac .Pi ;_aadc :=1.0;if _gcc < 0{_aadc *=-1.0;};if _ccf < 0{_aadc *=-1.0;};VsX :=_aadc *(-_bbg /2*_ac .Cos (_eba +_ce /2));VsY :=_aadc *(-_bbg /2*_ac .Sin (_eba +_ce /2)+_bbg *_ac .Sin (_eba +_ce /2));V1X :=VsX +_bbg /2*_ac .Cos (_eba +_ce /2);
V1Y :=VsY +_bbg /2*_ac .Sin (_eba +_ce /2);V2X :=VsX +_bbg /2*_ac .Cos (_eba +_ce /2)+L *_ac .Cos (_eba );V2Y :=VsY +_bbg /2*_ac .Sin (_eba +_ce /2)+L *_ac .Sin (_eba );V3X :=VsX +_bbg /2*_ac .Cos (_eba +_ce /2)+L *_ac .Cos (_eba )+_bbg *_ac .Cos (_eba -_ce /2);
V3Y :=VsY +_bbg /2*_ac .Sin (_eba +_ce /2)+L *_ac .Sin (_eba )+_bbg *_ac .Sin (_eba -_ce /2);V4X :=VsX +_bbg /2*_ac .Cos (_eba -_ce /2);V4Y :=VsY +_bbg /2*_ac .Sin (_eba -_ce /2);_egf :=NewPath ();_egf =_egf .AppendPoint (NewPoint (V1X ,V1Y ));_egf =_egf .AppendPoint (NewPoint (V2X ,V2Y ));
_egf =_egf .AppendPoint (NewPoint (V3X ,V3Y ));_egf =_egf .AppendPoint (NewPoint (V4X ,V4Y ));_dfe :=_bfe .LineEndingStyle1 ;_gbg :=_bfe .LineEndingStyle2 ;_ab :=3*_bbg ;_cfa :=3*_bbg ;_cdg :=(_cfa -_bbg )/2;if _gbg ==LineEndingStyleArrow {_fba :=_egf .GetPointNumber (2);
_dbae :=NewVectorPolar (_ab ,_eba +_ce );_gdf :=_fba .AddVector (_dbae );_edd :=NewVectorPolar (_cfa /2,_eba +_ce /2);_abe :=NewVectorPolar (_ab ,_eba );_gegf :=NewVectorPolar (_cdg ,_eba +_ce /2);_cgf :=_gdf .AddVector (_gegf );_gda :=_abe .Add (_edd .Flip ());
_dfg :=_cgf .AddVector (_gda );_cfdb :=_edd .Scale (2).Flip ().Add (_gda .Flip ());_egff :=_dfg .AddVector (_cfdb );_gee :=_gdf .AddVector (NewVectorPolar (_bbg ,_eba -_ce /2));_bg :=NewPath ();_bg =_bg .AppendPoint (_egf .GetPointNumber (1));_bg =_bg .AppendPoint (_gdf );
_bg =_bg .AppendPoint (_cgf );_bg =_bg .AppendPoint (_dfg );_bg =_bg .AppendPoint (_egff );_bg =_bg .AppendPoint (_gee );_bg =_bg .AppendPoint (_egf .GetPointNumber (4));_egf =_bg ;};if _dfe ==LineEndingStyleArrow {_gege :=_egf .GetPointNumber (1);_fdg :=_egf .GetPointNumber (_egf .Length ());
_fe :=NewVectorPolar (_bbg /2,_eba +_ce +_ce /2);_dgg :=_gege .AddVector (_fe );_ead :=NewVectorPolar (_ab ,_eba ).Add (NewVectorPolar (_cfa /2,_eba +_ce /2));_dd :=_dgg .AddVector (_ead );_gdfc :=NewVectorPolar (_cdg ,_eba -_ce /2);_bbb :=_dd .AddVector (_gdfc );
_agg :=NewVectorPolar (_ab ,_eba );_gaf :=_fdg .AddVector (_agg );_gcf :=NewVectorPolar (_cdg ,_eba +_ce +_ce /2);_fcd :=_gaf .AddVector (_gcf );_efg :=_dgg ;_edb :=NewPath ();_edb =_edb .AppendPoint (_dgg );_edb =_edb .AppendPoint (_dd );_edb =_edb .AppendPoint (_bbb );
for _ ,_gcd :=range _egf .Points [1:len (_egf .Points )-1]{_edb =_edb .AppendPoint (_gcd );};_edb =_edb .AppendPoint (_gaf );_edb =_edb .AppendPoint (_fcd );_edb =_edb .AppendPoint (_efg );_egf =_edb ;};_cb :=_ge .NewContentCreator ();_cb .Add_q ().SetNonStrokingColor (_bfe .LineColor );
if len (gsName )> 1{_cb .Add_gs (_b .PdfObjectName (gsName ));};_egf =_egf .Offset (_bfe .X1 ,_bfe .Y1 );_ffec :=_egf .GetBoundingBox ();DrawPathWithCreator (_egf ,_cb );if _bfe .LineStyle ==LineStyleDashed {_cb .Add_d ([]int64 {1,1},0).Add_S ().Add_f ().Add_Q ();
}else {_cb .Add_f ().Add_Q ();};return _cb .Bytes (),_ffec .ToPdfRectangle (),nil ;};

// ToPdfRectangle returns the bounding box as a PDF rectangle.
func (_ffb BoundingBox )ToPdfRectangle ()*_g .PdfRectangle {return &_g .PdfRectangle {Llx :_ffb .X ,Lly :_ffb .Y ,Urx :_ffb .X +_ffb .Width ,Ury :_ffb .Y +_ffb .Height };};

// BoundingBox represents the smallest rectangular area that encapsulates an object.
type BoundingBox struct{X float64 ;Y float64 ;Width float64 ;Height float64 ;};

// GetPointNumber returns the path point at the index specified by number.
// The index is 1-based.
func (_ca Path )GetPointNumber (number int )Point {if number < 1||number > len (_ca .Points ){return Point {};};return _ca .Points [number -1];};func (_gcg Point )String ()string {return _d .Sprintf ("(\u0025\u002e\u0031\u0066\u002c\u0025\u002e\u0031\u0066\u0029",_gcg .X ,_gcg .Y );
};

// Offset shifts the path with the specified offsets.
func (_bb Path )Offset (offX ,offY float64 )Path {for _bdb ,_gb :=range _bb .Points {_bb .Points [_bdb ]=_gb .Add (offX ,offY );};return _bb ;};

// Polyline defines a slice of points that are connected as straight lines.
type Polyline struct{Points []Point ;LineColor _g .PdfColor ;LineWidth float64 ;};

// AppendCurve appends the specified Bezier curve to the path.
func (_ff CubicBezierPath )AppendCurve (curve CubicBezierCurve )CubicBezierPath {_ff .Curves =append (_ff .Curves ,curve );return _ff ;};

// AppendPoint adds the specified point to the path.
func (_eaf Path )AppendPoint (point Point )Path {_eaf .Points =append (_eaf .Points ,point );return _eaf };

// Offset shifts the Bezier path with the specified offsets.
func (_ba CubicBezierPath )Offset (offX ,offY float64 )CubicBezierPath {for _dc ,_da :=range _ba .Curves {_ba .Curves [_dc ]=_da .AddOffsetXY (offX ,offY );};return _ba ;};

// NewVectorBetween returns a new vector with the direction specified by
// the subtraction of point a from point b (b-a).
func NewVectorBetween (a Point ,b Point )Vector {_gff :=Vector {};_gff .Dx =b .X -a .X ;_gff .Dy =b .Y -a .Y ;return _gff ;};

// Draw draws the circle. Can specify a graphics state (gsName) for setting opacity etc.  Otherwise leave empty ("").
// Returns the content stream as a byte array, the bounding box and an error on failure.
func (_eeg Circle )Draw (gsName string )([]byte ,*_g .PdfRectangle ,error ){_eeb :=_eeg .Width /2;_db :=_eeg .Height /2;if _eeg .BorderEnabled {_eeb -=_eeg .BorderWidth /2;_db -=_eeg .BorderWidth /2;};_ggg :=0.551784;_bbe :=_eeb *_ggg ;_ffe :=_db *_ggg ;
_gf :=NewCubicBezierPath ();_gf =_gf .AppendCurve (NewCubicBezierCurve (-_eeb ,0,-_eeb ,_ffe ,-_bbe ,_db ,0,_db ));_gf =_gf .AppendCurve (NewCubicBezierCurve (0,_db ,_bbe ,_db ,_eeb ,_ffe ,_eeb ,0));_gf =_gf .AppendCurve (NewCubicBezierCurve (_eeb ,0,_eeb ,-_ffe ,_bbe ,-_db ,0,-_db ));
_gf =_gf .AppendCurve (NewCubicBezierCurve (0,-_db ,-_bbe ,-_db ,-_eeb ,-_ffe ,-_eeb ,0));_gf =_gf .Offset (_eeb ,_db );if _eeg .BorderEnabled {_gf =_gf .Offset (_eeg .BorderWidth /2,_eeg .BorderWidth /2);};if _eeg .X !=0||_eeg .Y !=0{_gf =_gf .Offset (_eeg .X ,_eeg .Y );
};_cag :=_ge .NewContentCreator ();_cag .Add_q ();if _eeg .FillEnabled {_cag .SetNonStrokingColor (_eeg .FillColor );};if _eeg .BorderEnabled {_cag .SetStrokingColor (_eeg .BorderColor );_cag .Add_w (_eeg .BorderWidth );};if len (gsName )> 1{_cag .Add_gs (_b .PdfObjectName (gsName ));
};DrawBezierPathWithCreator (_gf ,_cag );_cag .Add_h ();if _eeg .FillEnabled &&_eeg .BorderEnabled {_cag .Add_B ();}else if _eeg .FillEnabled {_cag .Add_f ();}else if _eeg .BorderEnabled {_cag .Add_S ();};_cag .Add_Q ();_be :=_gf .GetBoundingBox ();if _eeg .BorderEnabled {_be .Height +=_eeg .BorderWidth ;
_be .Width +=_eeg .BorderWidth ;_be .X -=_eeg .BorderWidth /2;_be .Y -=_eeg .BorderWidth /2;};return _cag .Bytes (),_be .ToPdfRectangle (),nil ;};

// Copy returns a clone of the Bezier path.
func (_dg CubicBezierPath )Copy ()CubicBezierPath {_fd :=CubicBezierPath {};_fd .Curves =append (_fd .Curves ,_dg .Curves ...);return _fd ;};

// AddOffsetXY adds X,Y offset to all points on a curve.
func (_ea CubicBezierCurve )AddOffsetXY (offX ,offY float64 )CubicBezierCurve {_ea .P0 .X +=offX ;_ea .P1 .X +=offX ;_ea .P2 .X +=offX ;_ea .P3 .X +=offX ;_ea .P0 .Y +=offY ;_ea .P1 .Y +=offY ;_ea .P2 .Y +=offY ;_ea .P3 .Y +=offY ;return _ea ;};

// Flip changes the sign of the vector: -vector.
func (_gcac Vector )Flip ()Vector {_dgc :=_gcac .Magnitude ();_gcab :=_gcac .GetPolarAngle ();_gcac .Dx =_dgc *_ac .Cos (_gcab +_ac .Pi );_gcac .Dy =_dgc *_ac .Sin (_gcab +_ac .Pi );return _gcac ;};

// GetBounds returns the bounding box of the Bezier curve.
func (_ef CubicBezierCurve )GetBounds ()_g .PdfRectangle {_c :=_ef .P0 .X ;_cg :=_ef .P0 .X ;_fg :=_ef .P0 .Y ;_de :=_ef .P0 .Y ;for _aa :=0.0;_aa <=1.0;_aa +=0.001{Rx :=_ef .P0 .X *_ac .Pow (1-_aa ,3)+_ef .P1 .X *3*_aa *_ac .Pow (1-_aa ,2)+_ef .P2 .X *3*_ac .Pow (_aa ,2)*(1-_aa )+_ef .P3 .X *_ac .Pow (_aa ,3);
Ry :=_ef .P0 .Y *_ac .Pow (1-_aa ,3)+_ef .P1 .Y *3*_aa *_ac .Pow (1-_aa ,2)+_ef .P2 .Y *3*_ac .Pow (_aa ,2)*(1-_aa )+_ef .P3 .Y *_ac .Pow (_aa ,3);if Rx < _c {_c =Rx ;};if Rx > _cg {_cg =Rx ;};if Ry < _fg {_fg =Ry ;};if Ry > _de {_de =Ry ;};};_eg :=_g .PdfRectangle {};
_eg .Llx =_c ;_eg .Lly =_fg ;_eg .Urx =_cg ;_eg .Ury =_de ;return _eg ;};

// AddVector adds vector to a point.
func (_gc Point )AddVector (v Vector )Point {_gc .X +=v .Dx ;_gc .Y +=v .Dy ;return _gc };

// Line defines a line shape between point 1 (X1,Y1) and point 2 (X2,Y2).  The line ending styles can be none (regular line),
// or arrows at either end.  The line also has a specified width, color and opacity.
type Line struct{X1 float64 ;Y1 float64 ;X2 float64 ;Y2 float64 ;LineColor _g .PdfColor ;Opacity float64 ;LineWidth float64 ;LineEndingStyle1 LineEndingStyle ;LineEndingStyle2 LineEndingStyle ;LineStyle LineStyle ;};

// Rotate returns a new Point at `p` rotated by `theta` degrees.
func (_fad Point )Rotate (theta float64 )Point {_cff :=_e .NewPoint (_fad .X ,_fad .Y ).Rotate (theta );return NewPoint (_cff .X ,_cff .Y );};

// Vector represents a two-dimensional vector.
type Vector struct{Dx float64 ;Dy float64 ;};

// NewVectorPolar returns a new vector calculated from the specified
// magnitude and angle.
func NewVectorPolar (length float64 ,theta float64 )Vector {_eecc :=Vector {};_eecc .Dx =length *_ac .Cos (theta );_eecc .Dy =length *_ac .Sin (theta );return _eecc ;};

// Copy returns a clone of the path.
func (_fdf Path )Copy ()Path {_cad :=Path {};_cad .Points =append (_cad .Points ,_fdf .Points ...);return _cad ;};

// Magnitude returns the magnitude of the vector.
func (_acg Vector )Magnitude ()float64 {return _ac .Sqrt (_ac .Pow (_acg .Dx ,2.0)+_ac .Pow (_acg .Dy ,2.0));};

// BasicLine defines a line between point 1 (X1,Y1) and point 2 (X2,Y2). The line has a specified width, color and opacity.
type BasicLine struct{X1 float64 ;Y1 float64 ;X2 float64 ;Y2 float64 ;LineColor _g .PdfColor ;Opacity float64 ;LineWidth float64 ;LineStyle LineStyle ;DashArray []int64 ;DashPhase int64 ;};

// Draw draws the basic line to PDF. Generates the content stream which can be used in page contents or appearance
// stream of annotation. Returns the stream content, XForm bounding box (local), bounding box and an error if
// one occurred.
func (_baea BasicLine )Draw (gsName string )([]byte ,*_g .PdfRectangle ,error ){_cgfc :=NewPath ();_cgfc =_cgfc .AppendPoint (NewPoint (_baea .X1 ,_baea .Y1 ));_cgfc =_cgfc .AppendPoint (NewPoint (_baea .X2 ,_baea .Y2 ));_dgf :=_ge .NewContentCreator ();
_dgf .Add_q ().Add_w (_baea .LineWidth ).SetStrokingColor (_baea .LineColor );if _baea .LineStyle ==LineStyleDashed {if _baea .DashArray ==nil {_baea .DashArray =[]int64 {1,1};};_dgf .Add_d (_baea .DashArray ,_baea .DashPhase );};if len (gsName )> 1{_dgf .Add_gs (_b .PdfObjectName (gsName ));
};DrawPathWithCreator (_cgfc ,_dgf );_dgf .Add_S ().Add_Q ();return _dgf .Bytes (),_cgfc .GetBoundingBox ().ToPdfRectangle (),nil ;};

// NewCubicBezierPath returns a new empty cubic Bezier path.
func NewCubicBezierPath ()CubicBezierPath {_aad :=CubicBezierPath {};_aad .Curves =[]CubicBezierCurve {};return _aad ;};

// DrawPathWithCreator makes the path with the content creator.
// Adds the PDF commands to draw the path to the creator instance.
func DrawPathWithCreator (path Path ,creator *_ge .ContentCreator ){for _eddc ,_ece :=range path .Points {if _eddc ==0{creator .Add_m (_ece .X ,_ece .Y );}else {creator .Add_l (_ece .X ,_ece .Y );};};};

// Scale scales the vector by the specified factor.
func (_ggef Vector )Scale (factor float64 )Vector {_fbc :=_ggef .Magnitude ();_gcaa :=_ggef .GetPolarAngle ();_ggef .Dx =factor *_fbc *_ac .Cos (_gcaa );_ggef .Dy =factor *_fbc *_ac .Sin (_gcaa );return _ggef ;};