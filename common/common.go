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

// Package common contains wrapper types and utilities common to all of the
// OOXML document formats.
//
// Package common contains common properties used by the subpackages.
package common ;import (_cb "archive/zip";_ce "bytes";_cbc "encoding/xml";_e "fmt";_bb "github.com/unidoc/unioffice";_fb "github.com/unidoc/unioffice/common/tempstorage";_ff "github.com/unidoc/unioffice/common/tempstorage/diskstore";_fd "github.com/unidoc/unioffice/measurement";_fa "github.com/unidoc/unioffice/schema/soo/dml";_be "github.com/unidoc/unioffice/schema/soo/ofc/custom_properties";_ag "github.com/unidoc/unioffice/schema/soo/ofc/docPropsVTypes";_ca "github.com/unidoc/unioffice/schema/soo/ofc/extended_properties";_ced "github.com/unidoc/unioffice/schema/soo/pkg/content_types";_ed "github.com/unidoc/unioffice/schema/soo/pkg/metadata/core_properties";_agf "github.com/unidoc/unioffice/schema/soo/pkg/relationships";_feg "github.com/unidoc/unioffice/zippkg";_d "image";_ "image/gif";_ "image/jpeg";_ "image/png";_cf "os";_fc "regexp";_fe "strconv";_f "strings";_a "time";);func (_ega CustomProperties )SetPropertyAsStream (name string ,stream string ){_ceed :=_ega .getNewProperty (name );_ceed .Stream =&stream ;_ega .setProperty (_ceed );};func (_gcc CustomProperties )SetPropertyAsUi8 (name string ,ui8 uint64 ){_gga :=_gcc .getNewProperty (name );_gga .Ui8 =&ui8 ;_gcc .setProperty (_gga );};func (_aae CustomProperties )SetPropertyAsLpstr (name string ,lpstr string ){_afd :=_aae .getNewProperty (name );_afd .Lpstr =&lpstr ;_aae .setProperty (_afd );};

// SetCompany sets the name of the company that created the document.
func (_ae AppProperties )SetCompany (s string ){_ae ._efa .Company =&s };

// Type returns the type of a relationship.
func (_fdff Relationship )Type ()string {return _fdff ._gbd .TypeAttr };

// CoreProperties contains document specific properties.
type CoreProperties struct{_af *_ed .CoreProperties };

// RelativeHeight returns the relative height of an image given a fixed width.
// This is used when setting image to a fixed width to calculate the height
// required to keep the same image aspect ratio.
func (_cga ImageRef )RelativeHeight (w _fd .Distance )_fd .Distance {_debd :=float64 (_cga .Size ().Y )/float64 (_cga .Size ().X );return w *_fd .Distance (_debd );};

// ApplicationVersion returns the version of the application that created the
// document.
func (_efac AppProperties )ApplicationVersion ()string {if _efac ._efa .AppVersion !=nil {return *_efac ._efa .AppVersion ;};return "";};

// PropertiesList returns the list of all custom properties of the document.
func (_ac CustomProperties )PropertiesList ()[]*_be .CT_Property {return _ac ._bdf .Property };

// SetApplicationVersion sets the version of the application that created the
// document.  Per MS, the verison string mut be in the form 'XX.YYYY'.
func (_dc AppProperties )SetApplicationVersion (s string ){_dc ._efa .AppVersion =&s };

// WriteExtraFiles writes the extra files to the zip package.
func (_gfa *DocBase )WriteExtraFiles (z *_cb .Writer )error {for _ ,_feb :=range _gfa .ExtraFiles {if _bee :=_feg .AddFileFromDisk (z ,_feb .ZipPath ,_feb .DiskPath );_bee !=nil {return _bee ;};};return nil ;};

// CustomProperties contains document specific properties.
type CustomProperties struct{_bdf *_be .Properties };

// ImageRef is a reference to an image within a document.
type ImageRef struct{_dae *DocBase ;_gccd Relationships ;_abdf Image ;_edd string ;};const _eacb ="\u0032\u0020\u004aan\u0075\u0061\u0072\u0079\u0020\u0032\u0030\u0030\u0036\u0020\u0061\u0074\u0020\u0031\u0035\u003a\u0030\u0034";

// TableStyles contains document specific properties.
type TableStyles struct{_gab *_fa .TblStyleLst };func (_dbd CustomProperties )SetPropertyAsR8 (name string ,r8 float64 ){_ccc :=_dbd .getNewProperty (name );_ccc .R8 =&r8 ;_dbd .setProperty (_ccc );};func (_bbc CustomProperties )SetPropertyAsOstream (name string ,ostream string ){_bac :=_bbc .getNewProperty (name );_bac .Ostream =&ostream ;_bbc .setProperty (_bac );};

// FindRIDForN returns the relationship ID for the i'th relationship of type t.
func (_ebfg Relationships )FindRIDForN (i int ,t string )string {for _ ,_gda :=range _ebfg ._fgf .CT_Relationships .Relationship {if _gda .TypeAttr ==t {if i ==0{return _gda .IdAttr ;};i --;};};return "";};func (_agc CustomProperties )SetPropertyAsInt (name string ,i int ){_def :=_agc .getNewProperty (name );_ge :=int32 (i );_def .Int =&_ge ;_agc .setProperty (_def );};

// SetModified sets the time that the document was modified.
func (_fee CoreProperties )SetModified (t _a .Time ){_fee ._af .Modified =_bbf (t ,"\u0064\u0063t\u0065\u0072\u006ds\u003a\u006d\u006f\u0064\u0069\u0066\u0069\u0065\u0064");};var ReleasedAt =_a .Date (_gee ,_bba ,_aggc ,_fgff ,_cce ,0,0,_a .UTC );

// AddOverride adds an override content type for a given path name.
func (_cfc ContentTypes )AddOverride (path ,contentType string ){if !_f .HasPrefix (path ,"\u002f"){path ="\u002f"+path ;};if _f .HasPrefix (contentType ,"\u0068\u0074\u0074\u0070"){_bb .Log ("\u0063\u006f\u006e\u0074\u0065\u006et\u0020\u0074\u0079p\u0065\u0020\u0027%\u0073\u0027\u0020\u0069\u0073\u0020\u0069\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u002c m\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0068\u0074\u0074\u0070",contentType );};_fdf :=_ced .NewOverride ();_fdf .PartNameAttr =path ;_fdf .ContentTypeAttr =contentType ;_cfc ._ceee .Override =append (_cfc ._ceee .Override ,_fdf );};

// Author returns the author of the document
func (_ccg CoreProperties )Author ()string {if _ccg ._af .Creator !=nil {return string (_ccg ._af .Creator .Data );};return "";};func UtcTimeFormat (t _a .Time )string {return t .Format (_eacb )+"\u0020\u0055\u0054\u0043"};

// SetDescription records the description of the document.
func (_aad CoreProperties )SetDescription (s string ){if _aad ._af .Description ==nil {_aad ._af .Description =&_bb .XSDAny {XMLName :_cbc .Name {Local :"\u0064\u0063\u003a\u0064\u0065\u0073\u0063\u0072\u0069p\u0074\u0069\u006f\u006e"}};};_aad ._af .Description .Data =[]byte (s );};func (_gcg CustomProperties )SetPropertyAsI2 (name string ,i2 int16 ){_aca :=_gcg .getNewProperty (name );_aca .I2 =&i2 ;_gcg .setProperty (_aca );};

// Application returns the name of the application that created the document.
// For unioffice created documents, it defaults to github.com/unidoc/unioffice
func (_bga AppProperties )Application ()string {if _bga ._efa .Application !=nil {return *_bga ._efa .Application ;};return "";};

// X returns the inner wrapped XML type.
func (_ebc CustomProperties )X ()*_be .Properties {return _ebc ._bdf };func (_dgc CustomProperties )SetPropertyAsFiletime (name string ,filetime _a .Time ){_ebfc :=_dgc .getNewProperty (name );_ebfc .Filetime =&filetime ;_dgc .setProperty (_ebfc );};

// ImageFromStorage reads an image using the currently set
// temporary storage mechanism (see tempstorage). You can also
// construct an Image directly if the file and size are known.
func ImageFromStorage (path string )(Image ,error ){_dca :=Image {};_gde ,_efef :=_fb .Open (path );if _efef !=nil {return _dca ,_e .Errorf ("\u0065\u0072\u0072or\u0020\u0072\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073",_efef );};defer _gde .Close ();_babc ,_dda ,_efef :=_d .Decode (_gde );if _efef !=nil {return _dca ,_e .Errorf ("\u0075n\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0070\u0061r\u0073e\u0020i\u006d\u0061\u0067\u0065\u003a\u0020\u0025s",_efef );};_dca .Path =path ;_dca .Format =_dda ;_dca .Size =_babc .Bounds ().Size ();return _dca ,nil ;};

// SetApplication sets the name of the application that created the document.
func (_cd AppProperties )SetApplication (s string ){_cd ._efa .Application =&s };

// AddExtraFileFromZip is used when reading an unsupported file from an OOXML
// file. This ensures that unsupported file content will at least round-trip
// correctly.
func (_cbd *DocBase )AddExtraFileFromZip (f *_cb .File )error {_bbff ,_acg :=_feg .ExtractToDiskTmp (f ,_cbd .TmpPath );if _acg !=nil {return _e .Errorf ("\u0065\u0072r\u006f\u0072\u0020\u0065x\u0074\u0072a\u0063\u0074\u0069\u006e\u0067\u0020\u0075\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0066\u0069\u006ce\u003a\u0020\u0025\u0073",_acg );};_cbd .ExtraFiles =append (_cbd .ExtraFiles ,ExtraFile {ZipPath :f .Name ,DiskPath :_bbff });return nil ;};

// Path returns the path to an image file, if any.
func (_gdc ImageRef )Path ()string {return _gdc ._abdf .Path };func (_add CustomProperties )SetPropertyAsUi1 (name string ,ui1 uint8 ){_ffd :=_add .getNewProperty (name );_ffd .Ui1 =&ui1 ;_add .setProperty (_ffd );};

// NewRelationshipsCopy creates a new relationships wrapper as a copy of passed in instance.
func NewRelationshipsCopy (rels Relationships )Relationships {_cece :=*rels ._fgf ;return Relationships {_fgf :&_cece };};

// RemoveOverride removes an override given a path.
func (_g ContentTypes )RemoveOverride (path string ){if !_f .HasPrefix (path ,"\u002f"){path ="\u002f"+path ;};for _abd ,_eg :=range _g ._ceee .Override {if _eg .PartNameAttr ==path {copy (_g ._ceee .Override [_abd :],_g ._ceee .Override [_abd +1:]);_g ._ceee .Override =_g ._ceee .Override [0:len (_g ._ceee .Override )-1];};};};func (_face CustomProperties )SetPropertyAsArray (name string ,array *_ag .Array ){_abg :=_face .getNewProperty (name );_abg .Array =array ;_face .setProperty (_abg );};

// Target returns the target (path) of a relationship.
func (_gbe Relationship )Target ()string {return _gbe ._gbd .TargetAttr };func (_afac CustomProperties )SetPropertyAsR4 (name string ,r4 float32 ){_eba :=_afac .getNewProperty (name );_eba .R4 =&r4 ;_afac .setProperty (_eba );};

// RelID returns the relationship ID.
func (_gce ImageRef )RelID ()string {return _gce ._edd };func (_da CustomProperties )SetPropertyAsNull (name string ){_dfcg :=_da .getNewProperty (name );_dfcg .Null =_ag .NewNull ();_da .setProperty (_dfcg );};

// X returns the inner raw content types.
func (_agfd ContentTypes )X ()*_ced .Types {return _agfd ._ceee };

// Description returns the description of the document
func (_dea CoreProperties )Description ()string {if _dea ._af .Description !=nil {return string (_dea ._af .Description .Data );};return "";};

// ID returns the ID of a relationship.
func (_eea Relationship )ID ()string {return _eea ._gbd .IdAttr };

// Company returns the name of the company that created the document.
// For unioffice created documents, it defaults to github.com/unidoc/unioffice
func (_bef AppProperties )Company ()string {if _bef ._efa .Company !=nil {return *_bef ._efa .Company ;};return "";};

// SetAuthor records the author of the document.
func (_ba CoreProperties )SetAuthor (s string ){if _ba ._af .Creator ==nil {_ba ._af .Creator =&_bb .XSDAny {XMLName :_cbc .Name {Local :"\u0064\u0063\u003a\u0063\u0072\u0065\u0061\u0074\u006f\u0072"}};};_ba ._af .Creator .Data =[]byte (s );};

// AddImageToZip adds an image (either from bytes or from disk) and adds it to the zip file.
func AddImageToZip (z *_cb .Writer ,img ImageRef ,imageNum int ,dt _bb .DocType )error {_eag :=_bb .AbsoluteImageFilename (dt ,imageNum ,_f .ToLower (img .Format ()));if img .Data ()!=nil &&len (*img .Data ())> 0{if _efe :=_feg .AddFileFromBytes (z ,_eag ,*img .Data ());_efe !=nil {return _efe ;};}else if img .Path ()!=""{if _fgg :=_feg .AddFileFromDisk (z ,_eag ,img .Path ());_fgg !=nil {return _fgg ;};}else {return _e .Errorf ("\u0075\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0073\u006f\u0075\u0072\u0063\u0065\u003a %\u002b\u0076",img );};return nil ;};const _aggc =31;func (_ede CustomProperties )SetPropertyAsBool (name string ,b bool ){_acf :=_ede .getNewProperty (name );_acf .Bool =&b ;_ede .setProperty (_acf );};

// CopyOverride copies override content type for a given `path` and puts it with a path `newPath`.
func (_fac ContentTypes )CopyOverride (path ,newPath string ){if !_f .HasPrefix (path ,"\u002f"){path ="\u002f"+path ;};if !_f .HasPrefix (newPath ,"\u002f"){newPath ="\u002f"+newPath ;};for _bgb :=range _fac ._ceee .Override {if _fac ._ceee .Override [_bgb ].PartNameAttr ==path {_eff :=*_fac ._ceee .Override [_bgb ];_eff .PartNameAttr =newPath ;_fac ._ceee .Override =append (_fac ._ceee .Override ,&_eff );};};};

// Created returns the time that the document was created.
func (_aa CoreProperties )Created ()_a .Time {return _bc (_aa ._af .Created )};func (_gb CustomProperties )SetPropertyAsI1 (name string ,i1 int8 ){_bfa :=_gb .getNewProperty (name );_bfa .I1 =&i1 ;_gb .setProperty (_bfa );};

// RelativeWidth returns the relative width of an image given a fixed height.
// This is used when setting image to a fixed height to calculate the width
// required to keep the same image aspect ratio.
func (_dbb ImageRef )RelativeWidth (h _fd .Distance )_fd .Distance {_afc :=float64 (_dbb .Size ().X )/float64 (_dbb .Size ().Y );return h *_fd .Distance (_afc );};

// Remove removes an existing relationship.
func (_bfc Relationships )Remove (rel Relationship )bool {for _aegg ,_bce :=range _bfc ._fgf .Relationship {if _bce ==rel ._gbd {copy (_bfc ._fgf .Relationship [_aegg :],_bfc ._fgf .Relationship [_aegg +1:]);_bfc ._fgf .Relationship =_bfc ._fgf .Relationship [0:len (_bfc ._fgf .Relationship )-1];return true ;};};return false ;};func (_aac CustomProperties )SetPropertyAsUi2 (name string ,ui2 uint16 ){_afa :=_aac .getNewProperty (name );_afa .Ui2 =&ui2 ;_aac .setProperty (_afa );};

// X returns the inner wrapped XML type.
func (_ecgd TableStyles )X ()*_fa .TblStyleLst {return _ecgd ._gab };

// Data returns the data of an image file, if any.
func (_gdg ImageRef )Data ()*[]byte {return _gdg ._abdf .Data };

// X returns the underlying raw XML data.
func (_ccd Relationships )X ()*_agf .Relationships {return _ccd ._fgf };func (_ebcg CustomProperties )SetPropertyAsError (name string ,error string ){_daf :=_ebcg .getNewProperty (name );_daf .Error =&error ;_ebcg .setProperty (_daf );};

// LastModifiedBy returns the name of the last person to modify the document
func (_cff CoreProperties )LastModifiedBy ()string {if _cff ._af .LastModifiedBy !=nil {return *_cff ._af .LastModifiedBy ;};return "";};

// ContentTypes is the top level "[Content_Types].xml" in a zip package.
type ContentTypes struct{_ceee *_ced .Types };

// MakeImageRef constructs an image reference which is a reference to a
// particular image file inside a document.  The same image can be used multiple
// times in a document by re-use the ImageRef.
func MakeImageRef (img Image ,d *DocBase ,rels Relationships )ImageRef {return ImageRef {_abdf :img ,_dae :d ,_gccd :rels };};func (_ebbd CustomProperties )SetPropertyAsBstr (name string ,bstr string ){_effg :=_ebbd .getNewProperty (name );_effg .Bstr =&bstr ;_ebbd .setProperty (_effg );};

// NewCoreProperties constructs a new CoreProperties.
func NewCoreProperties ()CoreProperties {return CoreProperties {_af :_ed .NewCoreProperties ()}};

// SetDocSecurity sets the document security flag.
func (_cee AppProperties )SetDocSecurity (v int32 ){_cee ._efa .DocSecurity =_bb .Int32 (v )};

// SetTarget set the target (path) of a relationship.
func (_efec Relationship )SetTarget (s string ){_efec ._gbd .TargetAttr =s };

// X returns the inner wrapped XML type.
func (_dfd AppProperties )X ()*_ca .Properties {return _dfd ._efa };

// Pages returns total number of pages which are saved by the text editor which produced the document.
// For unioffice created documents, it is 0.
func (_bea AppProperties )Pages ()int32 {if _bea ._efa .Pages !=nil {return *_bea ._efa .Pages ;};return 0;};

// Relationships represents a .rels file.
type Relationships struct{_fgf *_agf .Relationships };func (_dfdd CustomProperties )SetPropertyAsStorage (name string ,storage string ){_cdf :=_dfdd .getNewProperty (name );_cdf .Storage =&storage ;_dfdd .setProperty (_cdf );};

// Image is a container for image information. It's used as we need format and
// and size information to use images.
// It contains either the filesystem path to the image, or the image itself.
type Image struct{Size _d .Point ;Format string ;Path string ;Data *[]byte ;};func (_gaf CustomProperties )setProperty (_dcg *_be .CT_Property ){_dee :=_gaf .GetPropertyByName (*_dcg .NameAttr );if (_dee ==CustomProperty {}){_gaf ._bdf .Property =append (_gaf ._bdf .Property ,_dcg );}else {_dcg .FmtidAttr =_dee ._bda .FmtidAttr ;if _dee ._bda .PidAttr ==0{_dcg .PidAttr =_dee ._bda .PidAttr ;};_dcg .LinkTargetAttr =_dee ._bda .LinkTargetAttr ;*_dee ._bda =*_dcg ;};};

// X returns the inner wrapped XML type.
func (_dga CoreProperties )X ()*_ed .CoreProperties {return _dga ._af };func (_daa CustomProperties )SetPropertyAsCy (name string ,cy string ){_ebff :=_daa .getNewProperty (name );_ebff .Cy =&cy ;_daa .setProperty (_ebff );};func (_cbg CustomProperties )SetPropertyAsUi4 (name string ,ui4 uint32 ){_ffc :=_cbg .getNewProperty (name );_ffc .Ui4 =&ui4 ;_cbg .setProperty (_ffc );};const _ee ="2\u00300\u0036\u002d\u0030\u0031\u002d\u0030\u0032\u00541\u0035\u003a\u0030\u0034:0\u0035\u005a";

// SetContentStatus records the content status of the document.
func (_afg CoreProperties )SetContentStatus (s string ){_afg ._af .ContentStatus =&s };

// EnsureDefault esnures that an extension and default content type exist,
// adding it if necessary.
func (_bgc ContentTypes )EnsureDefault (ext ,contentType string ){for _ ,_fbg :=range _bgc ._ceee .Default {if _fbg .ExtensionAttr ==ext {_fbg .ContentTypeAttr =contentType ;return ;};};_aeg :=&_ced .Default {};_aeg .ContentTypeAttr =contentType ;_aeg .ExtensionAttr =ext ;_bgc ._ceee .Default =append (_bgc ._ceee .Default ,_aeg );};

// IsEmpty returns true if there are no relationships.
func (_gdea Relationships )IsEmpty ()bool {return _gdea ._fgf ==nil ||len (_gdea ._fgf .Relationship )==0;};

// NewTheme constructs a new theme.
func NewTheme ()Theme {return Theme {_fa .NewTheme ()}};

// NewRelationship constructs a new relationship.
func NewRelationship ()Relationship {return Relationship {_agf .NewRelationship ()}};

// Theme is a drawingml theme.
type Theme struct{_agcg *_fa .Theme };

// Relationship is a relationship within a .rels file.
type Relationship struct{_gbd *_agf .Relationship };const _fgff =20;

// NewTableStyles constructs a new TableStyles.
func NewTableStyles ()TableStyles {return TableStyles {_gab :_fa .NewTblStyleLst ()}};

// AddHyperlink adds an external hyperlink relationship.
func (_adc Relationships )AddHyperlink (target string )Hyperlink {_bcb :=_adc .AddRelationship (target ,_bb .HyperLinkType );_bcb ._gbd .TargetModeAttr =_agf .ST_TargetModeExternal ;return Hyperlink (_bcb );};const _cce =3;

// AddAutoRelationship adds a relationship with an automatically generated
// filename based off of the type. It should be preferred over AddRelationship
// to ensure consistent filenames are maintained.
func (_cbee Relationships )AddAutoRelationship (dt _bb .DocType ,src string ,idx int ,ctype string )Relationship {return _cbee .AddRelationship (_bb .RelativeFilename (dt ,src ,ctype ,idx ),ctype );};

// SetCategory records the category of the document.
func (_abc CoreProperties )SetCategory (s string ){_abc ._af .Category =&s };

// CustomProperty contains document specific property.
// Using of this type is deprecated.
type CustomProperty struct{_bda *_be .CT_Property };func (_aaff CustomProperties )SetPropertyAsOstorage (name string ,ostorage string ){_cab :=_aaff .getNewProperty (name );_cab .Ostorage =&ostorage ;_aaff .setProperty (_cab );};const _gee =2020;func (_egf CustomProperties )SetPropertyAsBlob (name ,blob string ){_fef :=_egf .getNewProperty (name );_fef .Blob =&blob ;_egf .setProperty (_fef );};

// Relationships returns a slice of all of the relationships.
func (_cgb Relationships )Relationships ()[]Relationship {_fdg :=[]Relationship {};for _ ,_acd :=range _cgb ._fgf .Relationship {_fdg =append (_fdg ,Relationship {_acd });};return _fdg ;};

// Clear removes any existing relationships.
func (_dfce Relationships )Clear (){_dfce ._fgf .Relationship =nil };func (_fbga Relationship )String ()string {return _e .Sprintf ("\u007b\u0049\u0044\u003a \u0025\u0073\u0020\u0054\u0061\u0072\u0067\u0065\u0074\u003a \u0025s\u0020\u0054\u0079\u0070\u0065\u003a\u0020%\u0073\u007d",_fbga .ID (),_fbga .Target (),_fbga .Type ());};

// NewAppProperties constructs a new AppProperties.
func NewAppProperties ()AppProperties {_fcf :=AppProperties {_efa :_ca .NewProperties ()};_fcf .SetCompany ("\u0046\u006f\u0078\u0079\u0055\u0074\u0069\u006c\u0073\u0020\u0065\u0068\u0066");_fcf .SetApplication ("g\u0069\u0074\u0068\u0075\u0062\u002ec\u006f\u006d\u002f\u0075\u006e\u0069\u0064\u006f\u0063/\u0075\u006e\u0069o\u0066f\u0069\u0063\u0065");_fcf .SetDocSecurity (0);_fcf .SetLinksUpToDate (false );var _ab ,_cedg ,_df int64 ;_e .Sscanf (Version ,"\u0025\u0064\u002e\u0025\u0064\u002e\u0025\u0064",&_ab ,&_cedg ,&_df );_bg :=float64 (_ab )+float64 (_cedg )/10000.0;_fcf .SetApplicationVersion (_e .Sprintf ("\u0025\u0030\u0037\u002e\u0034\u0066",_bg ));return _fcf ;};

// ImageFromBytes returns an Image struct for an in-memory image. You can also
// construct an Image directly if the file and size are known.
func ImageFromBytes (data []byte )(Image ,error ){_gced :=Image {};_ccf ,_gdge ,_dfg :=_d .Decode (_ce .NewReader (data ));if _dfg !=nil {return _gced ,_e .Errorf ("\u0075n\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0070\u0061r\u0073e\u0020i\u006d\u0061\u0067\u0065\u003a\u0020\u0025s",_dfg );};_gced .Data =&data ;_gced .Format =_gdge ;_gced .Size =_ccf .Bounds ().Size ();return _gced ,nil ;};

// SetLinksUpToDate sets the links up to date flag.
func (_bgg AppProperties )SetLinksUpToDate (v bool ){_bgg ._efa .LinksUpToDate =_bb .Bool (v )};func (_abca CustomProperties )SetPropertyAsClsid (name string ,clsid string ){_geb :=_abca .getNewProperty (name );_geb .Clsid =&clsid ;_abca .setProperty (_geb );};

// SetTitle records the title of the document.
func (_cae CoreProperties )SetTitle (s string ){if _cae ._af .Title ==nil {_cae ._af .Title =&_bb .XSDAny {XMLName :_cbc .Name {Local :"\u0064\u0063\u003a\u0074\u0069\u0074\u006c\u0065"}};};_cae ._af .Title .Data =[]byte (s );};

// Title returns the Title of the document
func (_ebd CoreProperties )Title ()string {if _ebd ._af .Title !=nil {return string (_ebd ._af .Title .Data );};return "";};func (_gba CustomProperties )SetPropertyAsI4 (name string ,i4 int32 ){_edc :=_gba .getNewProperty (name );_edc .I4 =&i4 ;_gba .setProperty (_edc );};

// EnsureOverride ensures that an override for the given path exists, adding it if necessary
func (_eaa ContentTypes )EnsureOverride (path ,contentType string ){for _ ,_dg :=range _eaa ._ceee .Override {if _dg .PartNameAttr ==path {if _f .HasPrefix (contentType ,"\u0068\u0074\u0074\u0070"){_bb .Log ("\u0063\u006f\u006e\u0074\u0065\u006et\u0020\u0074\u0079p\u0065\u0020\u0027%\u0073\u0027\u0020\u0069\u0073\u0020\u0069\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u002c m\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0068\u0074\u0074\u0070",contentType );};_dg .ContentTypeAttr =contentType ;return ;};};_eaa .AddOverride (path ,contentType );};

// NewContentTypes returns a wrapper around a newly constructed content-types.
func NewContentTypes ()ContentTypes {_cbe :=ContentTypes {_ceee :_ced .NewTypes ()};_cbe .AddDefault ("\u0078\u006d\u006c","\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c");_cbe .AddDefault ("\u0072\u0065\u006c\u0073","\u0061\u0070\u0070\u006c\u0069\u0063a\u0074\u0069\u006fn\u002f\u0076\u006ed\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006fr\u006d\u0061\u0074\u0073\u002dpa\u0063\u006b\u0061\u0067\u0065\u002e\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073\u002b\u0078\u006d\u006c");_cbe .AddDefault ("\u0070\u006e\u0067","\u0069m\u0061\u0067\u0065\u002f\u0070\u006eg");_cbe .AddDefault ("\u006a\u0070\u0065\u0067","\u0069\u006d\u0061\u0067\u0065\u002f\u006a\u0070\u0065\u0067");_cbe .AddDefault ("\u006a\u0070\u0067","\u0069m\u0061\u0067\u0065\u002f\u006a\u0070g");_cbe .AddDefault ("\u0077\u006d\u0066","i\u006d\u0061\u0067\u0065\u002f\u0078\u002d\u0077\u006d\u0066");_cbe .AddOverride ("\u002fd\u006fc\u0050\u0072\u006f\u0070\u0073/\u0063\u006fr\u0065\u002e\u0078\u006d\u006c","\u0061\u0070\u0070\u006c\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0076\u006e\u0064\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073-\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002e\u0063\u006f\u0072\u0065\u002dp\u0072\u006f\u0070\u0065\u0072\u0074i\u0065\u0073\u002bx\u006d\u006c");_cbe .AddOverride ("\u002f\u0064\u006f\u0063\u0050\u0072\u006f\u0070\u0073\u002f\u0061\u0070p\u002e\u0078\u006d\u006c","a\u0070\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0076\u006e\u0064\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066o\u0072\u006d\u0061\u0074\u0073\u002d\u006f\u0066\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075m\u0065\u006e\u0074\u002e\u0065\u0078\u0074\u0065\u006e\u0064\u0065\u0064\u002dp\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073\u002b\u0078m\u006c");return _cbe ;};

// Modified returns the time that the document was modified.
func (_fbd CoreProperties )Modified ()_a .Time {return _bc (_fbd ._af .Modified )};

// SetLanguage records the language of the document.
func (_ege CoreProperties )SetLanguage (s string ){_ege ._af .Language =&_bb .XSDAny {XMLName :_cbc .Name {Local :"d\u0063\u003a\u006c\u0061\u006e\u0067\u0075\u0061\u0067\u0065"}};_ege ._af .Language .Data =[]byte (s );};

// DocBase is the type embedded in in the Document/Workbook/Presentation types
// that contains members common to all.
type DocBase struct{ContentTypes ContentTypes ;AppProperties AppProperties ;Rels Relationships ;CoreProperties CoreProperties ;CustomProperties CustomProperties ;Thumbnail _d .Image ;Images []ImageRef ;ExtraFiles []ExtraFile ;TmpPath string ;};func (_cde CustomProperties )SetPropertyAsUint (name string ,ui uint ){_fbgc :=_cde .getNewProperty (name );_egc :=uint32 (ui );_fbgc .Uint =&_egc ;_cde .setProperty (_fbgc );};func (_ace CustomProperties )getNewProperty (_dfe string )*_be .CT_Property {_ebf :=_ace ._bdf .Property ;_cdde :=int32 (1);for _ ,_dfc :=range _ebf {if _dfc .PidAttr > _cdde {_cdde =_dfc .PidAttr ;};};_befb :=_be .NewCT_Property ();_befb .NameAttr =&_dfe ;_befb .PidAttr =_cdde +1;_befb .FmtidAttr ="\u007b\u0044\u0035\u0043\u0044\u0044\u0035\u0030\u0035\u002d\u0032\u0045\u0039\u0043\u002d\u0031\u0030\u0031\u0042\u002d\u0039\u0033\u0039\u0037-\u0030\u0038\u0030\u0030\u0032B\u0032\u0043F\u0039\u0041\u0045\u007d";return _befb ;};func _bbf (_ec _a .Time ,_dcf string )*_bb .XSDAny {_cdd :=&_bb .XSDAny {XMLName :_cbc .Name {Local :_dcf }};_cdd .Attrs =append (_cdd .Attrs ,_cbc .Attr {Name :_cbc .Name {Local :"\u0078\u0073\u0069\u003a\u0074\u0079\u0070\u0065"},Value :"\u0064\u0063\u0074\u0065\u0072\u006d\u0073\u003a\u00573\u0043\u0044\u0054\u0046"});_cdd .Attrs =append (_cdd .Attrs ,_cbc .Attr {Name :_cbc .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u0073i"},Value :"\u0068\u0074\u0074\u0070\u003a/\u002f\u0077\u0077\u0077\u002e\u0077\u0033\u002e\u006f\u0072\u0067\u002f\u00320\u0030\u0031\u002f\u0058\u004d\u004c\u0053\u0063\u0068\u0065\u006d\u0061\u002d\u0069\u006e\u0073\u0074\u0061\u006e\u0063\u0065"});_cdd .Attrs =append (_cdd .Attrs ,_cbc .Attr {Name :_cbc .Name {Local :"\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0063\u0074\u0065\u0072\u006d\u0073"},Value :"\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/"});_cdd .Data =[]byte (_ec .Format (_ee ));return _cdd ;};func (_gf CustomProperties )SetPropertyAsEmpty (name string ){_deee :=_gf .getNewProperty (name );_deee .Empty =_ag .NewEmpty ();_gf .setProperty (_deee );};

// GetPropertyByName returns a custom property selected by it's name.
func (_aaf CustomProperties )GetPropertyByName (name string )CustomProperty {_fea :=_aaf ._bdf .Property ;for _ ,_gc :=range _fea {if *_gc .NameAttr ==name {return CustomProperty {_bda :_gc };};};return CustomProperty {};};

// X returns the inner wrapped XML type of CustomProperty.
func (_bab CustomProperty )X ()*_be .CT_Property {return _bab ._bda };

// ContentStatus returns the content status of the document (e.g. "Final", "Draft")
func (_ebb CoreProperties )ContentStatus ()string {if _ebb ._af .ContentStatus !=nil {return *_ebb ._af .ContentStatus ;};return "";};

// CopyRelationship copies the relationship.
func (_baa Relationships )CopyRelationship (idAttr string )(Relationship ,bool ){for _bae :=range _baa ._fgf .Relationship {if _baa ._fgf .Relationship [_bae ].IdAttr ==idAttr {_fbgf :=*_baa ._fgf .Relationship [_bae ];_dgf :=len (_baa ._fgf .Relationship )+1;_cgf :=map[string ]struct{}{};for _ ,_abcad :=range _baa ._fgf .Relationship {_cgf [_abcad .IdAttr ]=struct{}{};};for _ ,_aea :=_cgf [_e .Sprintf ("\u0072\u0049\u0064%\u0064",_dgf )];_aea ;_ ,_aea =_cgf [_e .Sprintf ("\u0072\u0049\u0064%\u0064",_dgf )]{_dgf ++;};_fbgf .IdAttr =_e .Sprintf ("\u0072\u0049\u0064%\u0064",_dgf );_baa ._fgf .Relationship =append (_baa ._fgf .Relationship ,&_fbgf );return Relationship {&_fbgf },true ;};};return Relationship {},false ;};

// ExtraFile is an unsupported file type extracted from, or to be written to a
// zip package
type ExtraFile struct{ZipPath string ;DiskPath string ;};func (_bgac CustomProperties )SetPropertyAsVector (name string ,vector *_ag .Vector ){_ad :=_bgac .getNewProperty (name );_ad .Vector =vector ;_bgac .setProperty (_ad );};func _bc (_fdd *_bb .XSDAny )_a .Time {if _fdd ==nil {return _a .Time {};};_egb ,_cge :=_a .Parse (_ee ,string (_fdd .Data ));if _cge !=nil {_bb .Log ("\u0065\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073i\u006e\u0067\u0020\u0074\u0069\u006d\u0065 \u0066\u0072\u006f\u006d\u0020\u0025\u0073\u003a\u0020\u0025\u0073",string (_fdd .Data ),_cge );};return _egb ;};

// ImageFromFile reads an image from a file on disk. It doesn't keep the image
// in memory and only reads it to determine the format and size. You can also
// construct an Image directly if the file and size are known.
// NOTE: See also ImageFromStorage.
func ImageFromFile (path string )(Image ,error ){_befg ,_afga :=_cf .Open (path );_cgeb :=Image {};if _afga !=nil {return _cgeb ,_e .Errorf ("\u0065\u0072\u0072or\u0020\u0072\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073",_afga );};defer _befg .Close ();_fga ,_adf ,_afga :=_d .Decode (_befg );if _afga !=nil {return _cgeb ,_e .Errorf ("\u0075n\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0070\u0061r\u0073e\u0020i\u006d\u0061\u0067\u0065\u003a\u0020\u0025s",_afga );};_cgeb .Path =path ;_cgeb .Format =_adf ;_cgeb .Size =_fga .Bounds ().Size ();return _cgeb ,nil ;};

// SetCreated sets the time that the document was created.
func (_cac CoreProperties )SetCreated (t _a .Time ){_cac ._af .Created =_bbf (t ,"\u0064c\u0074e\u0072\u006d\u0073\u003a\u0063\u0072\u0065\u0061\u0074\u0065\u0064");};

// AddRelationship adds a relationship.
func (_eed Relationships )AddRelationship (target ,ctype string )Relationship {if !_f .HasPrefix (ctype ,"\u0068t\u0074\u0070\u003a\u002f\u002f"){_bb .Log ("\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006es\u0068\u0069\u0070 t\u0079\u0070\u0065\u0020\u0025\u0073 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u0077\u0069t\u0068\u0020\u0027\u0068\u0074\u0074\u0070\u003a/\u002f\u0027",ctype );};_ddg :=_agf .NewRelationship ();_ccff :=len (_eed ._fgf .Relationship )+1;_gdf :=map[string ]struct{}{};for _ ,_egbd :=range _eed ._fgf .Relationship {_gdf [_egbd .IdAttr ]=struct{}{};};for _ ,_aeeb :=_gdf [_e .Sprintf ("\u0072\u0049\u0064%\u0064",_ccff )];_aeeb ;_ ,_aeeb =_gdf [_e .Sprintf ("\u0072\u0049\u0064%\u0064",_ccff )]{_ccff ++;};_ddg .IdAttr =_e .Sprintf ("\u0072\u0049\u0064%\u0064",_ccff );_ddg .TargetAttr =target ;_ddg .TypeAttr =ctype ;_eed ._fgf .Relationship =append (_eed ._fgf .Relationship ,_ddg );return Relationship {_ddg };};

// Category returns the category of the document
func (_bf CoreProperties )Category ()string {if _bf ._af .Category !=nil {return *_bf ._af .Category ;};return "";};

// X returns the inner wrapped XML type.
func (_ceg Relationship )X ()*_agf .Relationship {return _ceg ._gbd };

// AppProperties contains properties specific to the document and the
// application that created it.
type AppProperties struct{_efa *_ca .Properties };const _bba =8;

// X returns the inner wrapped XML type.
func (_ada Theme )X ()*_fa .Theme {return _ada ._agcg };

// NewRelationships creates a new relationship wrapper.
func NewRelationships ()Relationships {return Relationships {_fgf :_agf .NewRelationships ()}};func init (){_ff .SetAsStorage ()};

// RemoveOverrideByIndex removes an override given a path and override index.
func (_dd ContentTypes )RemoveOverrideByIndex (path string ,indexToFind int )error {_fdc :=path [0:len (path )-5];if !_f .HasPrefix (_fdc ,"\u002f"){_fdc ="\u002f"+_fdc ;};_cad ,_bd :=_fc .Compile (_fdc +"\u0028\u005b\u0030-\u0039\u005d\u002b\u0029\u002e\u0078\u006d\u006c");if _bd !=nil {return _bd ;};_de :=0;_eb :=-1;for _efb ,_aee :=range _dd ._ceee .Override {if _db :=_cad .FindStringSubmatch (_aee .PartNameAttr );len (_db )> 1{if _de ==indexToFind {_eb =_efb ;}else if _de > indexToFind {_cg ,_ :=_fe .Atoi (_db [1]);_cg --;_aee .PartNameAttr =_e .Sprintf ("\u0025\u0073\u0025\u0064\u002e\u0078\u006d\u006c",_fdc ,_cg );};_de ++;};};if _eb > -1{copy (_dd ._ceee .Override [_eb :],_dd ._ceee .Override [_eb +1:]);_dd ._ceee .Override =_dd ._ceee .Override [0:len (_dd ._ceee .Override )-1];};return nil ;};const Version ="\u0031\u002e\u0035.\u0031";func (_feea CustomProperties )SetPropertyAsLpwstr (name string ,lpwstr string ){_dbc :=_feea .getNewProperty (name );_dbc .Lpwstr =&lpwstr ;_feea .setProperty (_dbc );};

// TblStyle returns the TblStyle property.
func (_gdae TableStyles )TblStyle ()[]*_fa .CT_TableStyle {return _gdae ._gab .TblStyle };

// DefAttr returns the DefAttr property.
func (_caef TableStyles )DefAttr ()string {return _caef ._gab .DefAttr };func (_fae CustomProperties )SetPropertyAsDate (name string ,date _a .Time ){date =date .UTC ();_bcf ,_bad ,_eaf :=date .Date ();_deb ,_cba ,_caa :=date .Clock ();_cbf :=_a .Date (_bcf ,_bad ,_eaf ,_deb ,_cba ,_caa ,0,_a .UTC );_eac :=_fae .getNewProperty (name );_eac .Filetime =&_cbf ;_fae .setProperty (_eac );};

// SetLastModifiedBy records the last person to modify the document.
func (_gg CoreProperties )SetLastModifiedBy (s string ){_gg ._af .LastModifiedBy =&s };func (_ebe CustomProperties )SetPropertyAsI8 (name string ,i8 int64 ){_dcc :=_ebe .getNewProperty (name );_dcc .I8 =&i8 ;_ebe .setProperty (_dcc );};

// AddDefault registers a default content type for a given file extension.
func (_aec ContentTypes )AddDefault (fileExtension string ,contentType string ){_cec :=_ced .NewDefault ();_cec .ExtensionAttr =fileExtension ;_cec .ContentTypeAttr =contentType ;_aec ._ceee .Default =append (_aec ._ceee .Default ,_cec );};

// Format returns the format of the underlying image
func (_ecg ImageRef )Format ()string {return _ecg ._abdf .Format };func (_agg CustomProperties )SetPropertyAsOblob (name ,oblob string ){_gcb :=_agg .getNewProperty (name );_gcb .Oblob =&oblob ;_agg .setProperty (_gcb );};func (_dff *ImageRef )SetRelID (id string ){_dff ._edd =id };func (_ddd CustomProperties )SetPropertyAsVstream (name string ,vstream *_ag .Vstream ){_dba :=_ddd .getNewProperty (name );_dba .Vstream =vstream ;_ddd .setProperty (_dba );};func (_aef CustomProperties )SetPropertyAsDecimal (name string ,decimal float64 ){_afaca :=_aef .getNewProperty (name );_afaca .Decimal =&decimal ;_aef .setProperty (_afaca );};

// Hyperlink is just an appropriately configured relationship.
type Hyperlink Relationship ;

// NewCustomProperties constructs a new CustomProperties.
func NewCustomProperties ()CustomProperties {return CustomProperties {_bdf :_be .NewProperties ()}};

// Size returns the size of an image
func (_dafc ImageRef )Size ()_d .Point {return _dafc ._abdf .Size };