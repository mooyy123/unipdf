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

// Package sighandler implements digital signature handlers for PDF signature validation and signing.
package sighandler ;import (_dd "bytes";_cd "crypto";_de "crypto/rand";_ae "crypto/rsa";_bf "crypto/sha1";_e "crypto/x509";_g "crypto/x509/pkix";_dg "encoding/asn1";_c "errors";_a "fmt";_fg "github.com/unidoc/pkcs7";_da "github.com/unidoc/timestamp";_ac "github.com/unidoc/unipdf/v3/core";
_ag "github.com/unidoc/unipdf/v3/model";_ff "github.com/unidoc/unipdf/v3/model/mdp";_fb "github.com/unidoc/unipdf/v3/model/sigutil";_d "hash";_fe "io/ioutil";_cc "math/big";_fa "net/http";_b "time";);

// DocTimeStampOpts defines options for configuring the timestamp handler.
type DocTimeStampOpts struct{

// SignatureSize is the estimated size of the signature contents in bytes.
// If not provided, a default signature size of 4192 is used.
// The signing process will report the model.ErrSignNotEnoughSpace error
// if the estimated signature size is smaller than the actual size of the
// signature.
SignatureSize int ;

// Client is the timestamp client used to make the signature request.
// If no client is provided, a default one is used.
Client *_fb .TimestampClient ;};

// Validate validates PdfSignature.
func (_gccc *adobePKCS7Detached )Validate (sig *_ag .PdfSignature ,digest _ag .Hasher )(_ag .SignatureValidationResult ,error ){_egc :=sig .Contents .Bytes ();_dfdd ,_fdd :=_fg .Parse (_egc );if _fdd !=nil {return _ag .SignatureValidationResult {},_fdd ;
};_aff ,_bfd :=digest .(*_dd .Buffer );if !_bfd {return _ag .SignatureValidationResult {},_a .Errorf ("c\u0061s\u0074\u0020\u0074\u006f\u0020\u0062\u0075\u0066f\u0065\u0072\u0020\u0066ai\u006c\u0073");};_dfdd .Content =_aff .Bytes ();if _fdd =_dfdd .Verify ();
_fdd !=nil {return _ag .SignatureValidationResult {},_fdd ;};return _ag .SignatureValidationResult {IsSigned :true ,IsVerified :true },nil ;};

// Validate validates PdfSignature.
func (_bab *etsiPAdES )Validate (sig *_ag .PdfSignature ,digest _ag .Hasher )(_ag .SignatureValidationResult ,error ){_fbed :=sig .Contents .Bytes ();_cgc ,_cfe :=_fg .Parse (_fbed );if _cfe !=nil {return _ag .SignatureValidationResult {},_cfe ;};_bd ,_df :=digest .(*_dd .Buffer );
if !_df {return _ag .SignatureValidationResult {},_a .Errorf ("c\u0061s\u0074\u0020\u0074\u006f\u0020\u0062\u0075\u0066f\u0065\u0072\u0020\u0066ai\u006c\u0073");};_cgc .Content =_bd .Bytes ();if _cfe =_cgc .Verify ();_cfe !=nil {return _ag .SignatureValidationResult {},_cfe ;
};for _ ,_ec :=range _cgc .Signers {_gg :=_ec .EncryptedDigest ;for _ ,_dce :=range _ec .AuthenticatedAttributes {if _dce .Type .Equal (_fg .OIDAttributeMessageDigest ){_gg =_dce .Value .Bytes ;};};for _ ,_bdc :=range _ec .UnauthenticatedAttributes {if _bdc .Type .Equal (_fg .OIDAttributeTimeStampToken ){_aga ,_bgd :=_da .Parse (_bdc .Value .Bytes );
if _bgd !=nil {return _ag .SignatureValidationResult {},_bgd ;};_gcda :=_cd .SHA512 .New ();_gcda .Write (_gg );if !_dd .Equal (_gcda .Sum (nil ),_aga .HashedMessage ){return _ag .SignatureValidationResult {},_a .Errorf ("\u0048\u0061\u0073\u0068\u0020i\u006e\u0020\u0074\u0069\u006d\u0065\u0073\u0074\u0061\u006d\u0070\u0020\u0069s\u0020\u0064\u0069\u0066\u0066\u0065\u0072\u0065\u006e\u0074\u0020\u0066\u0072\u006f\u006d\u0020\u0070\u006b\u0063\u0073\u0037");
};break ;};};};_dgf :=_ag .SignatureValidationResult {IsSigned :true ,IsVerified :true };return _dgf ,nil ;};

// NewDigest creates a new digest.
func (_bb *DocMDPHandler )NewDigest (sig *_ag .PdfSignature )(_ag .Hasher ,error ){return _bb ._gd .NewDigest (sig );};

// InitSignature initialises the PdfSignature.
func (_gba *adobePKCS7Detached )InitSignature (sig *_ag .PdfSignature )error {if !_gba ._bac {if _gba ._fcd ==nil {return _c .New ("c\u0065\u0072\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u006d\u0075\u0073\u0074\u0020n\u006f\u0074\u0020b\u0065 \u006e\u0069\u006c");
};if _gba ._egg ==nil {return _c .New ("\u0070\u0072\u0069\u0076\u0061\u0074\u0065\u004b\u0065\u0079\u0020m\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0062\u0065 \u006e\u0069\u006c");};};_ecc :=*_gba ;sig .Handler =&_ecc ;sig .Filter =_ac .MakeName ("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065");
sig .SubFilter =_ac .MakeName ("\u0061\u0064\u0062\u0065.p\u006b\u0063\u0073\u0037\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064");sig .Reference =nil ;_fdc ,_dfd :=_ecc .NewDigest (sig );if _dfd !=nil {return _dfd ;};_fdc .Write ([]byte ("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"));
return _ecc .Sign (sig ,_fdc );};

// NewEtsiPAdESLevelB creates a new Adobe.PPKLite ETSI.CAdES.detached Level T signature handler.
func NewEtsiPAdESLevelT (privateKey *_ae .PrivateKey ,certificate *_e .Certificate ,caCert *_e .Certificate ,certificateTimestampServerURL string )(_ag .SignatureHandler ,error ){return &etsiPAdES {_aa :certificate ,_cba :privateKey ,_fbc :caCert ,_dbd :certificateTimestampServerURL },nil ;
};

// Sign sets the Contents fields for the PdfSignature.
func (_bbe *etsiPAdES )Sign (sig *_ag .PdfSignature ,digest _ag .Hasher )error {_bcc ,_ddc :=digest .(*_dd .Buffer );if !_ddc {return _a .Errorf ("c\u0061s\u0074\u0020\u0074\u006f\u0020\u0062\u0075\u0066f\u0065\u0072\u0020\u0066ai\u006c\u0073");};_cdb ,_bcd :=_fg .NewSignedData (_bcc .Bytes ());
if _bcd !=nil {return _bcd ;};_cdb .SetDigestAlgorithm (_fg .OIDDigestAlgorithmSHA256 );_dag :=_fg .SignerInfoConfig {};_ge :=_cd .SHA256 .New ();_ge .Write (_bbe ._aa .Raw );var _bfc struct{Seq struct{Seq struct{Value []byte ;};};};_bfc .Seq .Seq .Value =_ge .Sum (nil );
_dag .ExtraSignedAttributes =append (_dag .ExtraSignedAttributes ,_fg .Attribute {Type :_fg .OIDAttributeSigningCertificateV2 ,Value :_bfc });var _cfd []*_e .Certificate ;if _bbe ._fbc !=nil {_cfd =[]*_e .Certificate {_bbe ._fbc };};if _feg :=_cdb .AddSignerChainPAdES (_bbe ._aa ,_bbe ._cba ,_cfd ,_dag );
_feg !=nil {return _feg ;};_cdb .Detach ();if len (_bbe ._dbd )> 0{_ceaf :=_cdb .GetSignedData ().SignerInfos [0].EncryptedDigest ;for _ ,_ed :=range _cdb .GetSignedData ().SignerInfos [0].AuthenticatedAttributes {if _ed .Type .Equal (_fg .OIDAttributeMessageDigest ){_ceaf =_ed .Value .Bytes ;
};};_caf ,_gdg :=_bbe .makeTimestampRequest (_bbe ._dbd ,_ceaf );if _gdg !=nil {return _gdg ;};_gdg =_cdb .GetSignedData ().SignerInfos [0].SetUnauthenticatedAttributes ([]_fg .Attribute {{Type :_fg .OIDAttributeTimeStampToken ,Value :_caf }});if _gdg !=nil {return _gdg ;
};};_af ,_bcd :=_cdb .Finish ();if _bcd !=nil {return _bcd ;};_cbb :=make ([]byte ,len (_af )+1024*2);copy (_cbb ,_af );sig .Contents =_ac .MakeHexString (string (_cbb ));if _bbe ._gcd {return nil ;};_ge =_bf .New ();_ge .Write (_cbb );return nil ;};

// InitSignature initialization of the DocMDP signature.
func (_gb *DocMDPHandler )InitSignature (sig *_ag .PdfSignature )error {_eb :=_gb ._gd .InitSignature (sig );if _eb !=nil {return _eb ;};sig .Handler =_gb ;if sig .Reference ==nil {sig .Reference =_ac .MakeArray ();};sig .Reference .Append (_ag .NewPdfSignatureReferenceDocMDP (_ag .NewPdfTransformParamsDocMDP (_gb .Permission )).ToPdfObject ());
return nil ;};

// DocMDPHandler describes handler for the DocMDP realization.
type DocMDPHandler struct{_gd _ag .SignatureHandler ;Permission _ff .DocMDPPermission ;};

// Sign sets the Contents fields.
func (_bgdf *adobePKCS7Detached )Sign (sig *_ag .PdfSignature ,digest _ag .Hasher )error {if _bgdf ._bac {_fad :=_bgdf ._gga ;if _fad <=0{_fad =8192;};sig .Contents =_ac .MakeHexString (string (make ([]byte ,_fad )));return nil ;};_agb ,_gaa :=digest .(*_dd .Buffer );
if !_gaa {return _a .Errorf ("c\u0061s\u0074\u0020\u0074\u006f\u0020\u0062\u0075\u0066f\u0065\u0072\u0020\u0066ai\u006c\u0073");};_ebgc ,_bfa :=_fg .NewSignedData (_agb .Bytes ());if _bfa !=nil {return _bfa ;};if _bbg :=_ebgc .AddSigner (_bgdf ._fcd ,_bgdf ._egg ,_fg .SignerInfoConfig {});
_bbg !=nil {return _bbg ;};_ebgc .Detach ();_bcg ,_bfa :=_ebgc .Finish ();if _bfa !=nil {return _bfa ;};_ada :=make ([]byte ,8192);copy (_ada ,_bcg );sig .Contents =_ac .MakeHexString (string (_ada ));return nil ;};

// NewDocTimeStamp creates a new DocTimeStamp signature handler.
// Both the timestamp server URL and the hash algorithm can be empty for the
// signature validation.
// The following hash algorithms are supported:
// crypto.SHA1, crypto.SHA256, crypto.SHA384, crypto.SHA512.
// NOTE: the handler will do a mock Sign when initializing the signature
// in order to estimate the signature size. Use NewDocTimeStampWithOpts
// for providing the signature size.
func NewDocTimeStamp (timestampServerURL string ,hashAlgorithm _cd .Hash )(_ag .SignatureHandler ,error ){return &docTimeStamp {_gag :timestampServerURL ,_cebb :hashAlgorithm },nil ;};

// Validate implementation of the SignatureHandler interface
// This check is impossible without checking the document's content.
// Please, use ValidateWithOpts with the PdfParser.
func (_ded *DocMDPHandler )Validate (sig *_ag .PdfSignature ,digest _ag .Hasher )(_ag .SignatureValidationResult ,error ){return _ag .SignatureValidationResult {},_c .New ("i\u006d\u0070\u006f\u0073\u0073\u0069b\u006c\u0065\u0020\u0076\u0061\u006ci\u0064\u0061\u0074\u0069\u006f\u006e\u0020w\u0069\u0074\u0068\u006f\u0075\u0074\u0020\u0070\u0061\u0072s\u0065");
};

// NewAdobeX509RSASHA1CustomWithOpts creates a new Adobe.PPKMS/Adobe.PPKLite
// adbe.x509.rsa_sha1 signature handler with a custom signing function. The
// handler is configured based on the provided options. If no options are
// provided, default options will be used. Both the certificate and the sign
// function can be nil for the signature validation.
func NewAdobeX509RSASHA1CustomWithOpts (certificate *_e .Certificate ,signFunc SignFunc ,opts *AdobeX509RSASHA1Opts )(_ag .SignatureHandler ,error ){if opts ==nil {opts =&AdobeX509RSASHA1Opts {};};return &adobeX509RSASHA1 {_cbe :certificate ,_bga :signFunc ,_afc :opts .EstimateSize ,_faf :opts .Algorithm },nil ;
};

// Validate validates PdfSignature.
func (_eae *docTimeStamp )Validate (sig *_ag .PdfSignature ,digest _ag .Hasher )(_ag .SignatureValidationResult ,error ){_bed :=sig .Contents .Bytes ();_gad ,_abaa :=_fg .Parse (_bed );if _abaa !=nil {return _ag .SignatureValidationResult {},_abaa ;};if _abaa =_gad .Verify ();
_abaa !=nil {return _ag .SignatureValidationResult {},_abaa ;};var _gcb timestampInfo ;_ ,_abaa =_dg .Unmarshal (_gad .Content ,&_gcb );if _abaa !=nil {return _ag .SignatureValidationResult {},_abaa ;};_dda ,_abaa :=_feac (_gcb .MessageImprint .HashAlgorithm .Algorithm );
if _abaa !=nil {return _ag .SignatureValidationResult {},_abaa ;};_cdg :=_dda .New ();_bcgd ,_edc :=digest .(*_dd .Buffer );if !_edc {return _ag .SignatureValidationResult {},_a .Errorf ("c\u0061s\u0074\u0020\u0074\u006f\u0020\u0062\u0075\u0066f\u0065\u0072\u0020\u0066ai\u006c\u0073");
};_cdg .Write (_bcgd .Bytes ());_fcdb :=_cdg .Sum (nil );_adgc :=_ag .SignatureValidationResult {IsSigned :true ,IsVerified :_dd .Equal (_fcdb ,_gcb .MessageImprint .HashedMessage ),GeneralizedTime :_gcb .GeneralizedTime };return _adgc ,nil ;};

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_gec *etsiPAdES )IsApplicable (sig *_ag .PdfSignature )bool {if sig ==nil ||sig .Filter ==nil ||sig .SubFilter ==nil {return false ;};return (*sig .Filter =="\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")&&*sig .SubFilter =="\u0045\u0054\u0053\u0049.C\u0041\u0064\u0045\u0053\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064";
};type docTimeStamp struct{_gag string ;_cebb _cd .Hash ;_agbg int ;_gdc *_fb .TimestampClient ;};func _bdg (_bdbd *_ae .PublicKey ,_ee []byte )_cd .Hash {_afb :=_bdbd .Size ();if _afb !=len (_ee ){return 0;};_gfgc :=func (_efd *_cc .Int ,_bcgc *_ae .PublicKey ,_faa *_cc .Int )*_cc .Int {_ffe :=_cc .NewInt (int64 (_bcgc .E ));
_efd .Exp (_faa ,_ffe ,_bcgc .N );return _efd ;};_efc :=new (_cc .Int ).SetBytes (_ee );_bdbdb :=_gfgc (new (_cc .Int ),_bdbd ,_efc );_gbga :=_ffad (_bdbdb .Bytes (),_afb );if _gbga [0]!=0||_gbga [1]!=1{return 0;};_adc :=[]struct{Hash _cd .Hash ;Prefix []byte ;
}{{Hash :_cd .SHA1 ,Prefix :[]byte {0x30,0x21,0x30,0x09,0x06,0x05,0x2b,0x0e,0x03,0x02,0x1a,0x05,0x00,0x04,0x14}},{Hash :_cd .SHA256 ,Prefix :[]byte {0x30,0x31,0x30,0x0d,0x06,0x09,0x60,0x86,0x48,0x01,0x65,0x03,0x04,0x02,0x01,0x05,0x00,0x04,0x20}},{Hash :_cd .SHA384 ,Prefix :[]byte {0x30,0x41,0x30,0x0d,0x06,0x09,0x60,0x86,0x48,0x01,0x65,0x03,0x04,0x02,0x02,0x05,0x00,0x04,0x30}},{Hash :_cd .SHA512 ,Prefix :[]byte {0x30,0x51,0x30,0x0d,0x06,0x09,0x60,0x86,0x48,0x01,0x65,0x03,0x04,0x02,0x03,0x05,0x00,0x04,0x40}},{Hash :_cd .RIPEMD160 ,Prefix :[]byte {0x30,0x20,0x30,0x08,0x06,0x06,0x28,0xcf,0x06,0x03,0x00,0x31,0x04,0x14}}};
for _ ,_dgd :=range _adc {_fbce :=_dgd .Hash .Size ();_ccg :=len (_dgd .Prefix )+_fbce ;if _dd .Equal (_gbga [_afb -_ccg :_afb -_fbce ],_dgd .Prefix ){return _dgd .Hash ;};};return 0;};

// NewEmptyAdobePKCS7Detached creates a new Adobe.PPKMS/Adobe.PPKLite adbe.pkcs7.detached
// signature handler. The generated signature is empty and of size signatureLen.
// The signatureLen parameter can be 0 for the signature validation.
func NewEmptyAdobePKCS7Detached (signatureLen int )(_ag .SignatureHandler ,error ){return &adobePKCS7Detached {_bac :true ,_gga :signatureLen },nil ;};

// NewAdobeX509RSASHA1Custom creates a new Adobe.PPKMS/Adobe.PPKLite
// adbe.x509.rsa_sha1 signature handler with a custom signing function. Both the
// certificate and the sign function can be nil for the signature validation.
// NOTE: the handler will do a mock Sign when initializing the signature in
// order to estimate the signature size. Use NewAdobeX509RSASHA1CustomWithOpts
// for configuring the handler to estimate the signature size.
func NewAdobeX509RSASHA1Custom (certificate *_e .Certificate ,signFunc SignFunc )(_ag .SignatureHandler ,error ){return &adobeX509RSASHA1 {_cbe :certificate ,_bga :signFunc },nil ;};

// NewDigest creates a new digest.
func (_be *adobePKCS7Detached )NewDigest (sig *_ag .PdfSignature )(_ag .Hasher ,error ){return _dd .NewBuffer (nil ),nil ;};

// NewDigest creates a new digest.
func (_gdf *docTimeStamp )NewDigest (sig *_ag .PdfSignature )(_ag .Hasher ,error ){return _dd .NewBuffer (nil ),nil ;};

// NewAdobePKCS7Detached creates a new Adobe.PPKMS/Adobe.PPKLite adbe.pkcs7.detached signature handler.
// Both parameters may be nil for the signature validation.
func NewAdobePKCS7Detached (privateKey *_ae .PrivateKey ,certificate *_e .Certificate )(_ag .SignatureHandler ,error ){return &adobePKCS7Detached {_fcd :certificate ,_egg :privateKey },nil ;};

// NewAdobeX509RSASHA1 creates a new Adobe.PPKMS/Adobe.PPKLite
// adbe.x509.rsa_sha1 signature handler. Both the private key and the
// certificate can be nil for the signature validation.
func NewAdobeX509RSASHA1 (privateKey *_ae .PrivateKey ,certificate *_e .Certificate )(_ag .SignatureHandler ,error ){return &adobeX509RSASHA1 {_cbe :certificate ,_fgf :privateKey },nil ;};

// InitSignature initialises the PdfSignature.
func (_ga *etsiPAdES )InitSignature (sig *_ag .PdfSignature )error {if !_ga ._cea {if _ga ._aa ==nil {return _c .New ("c\u0065\u0072\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u006d\u0075\u0073\u0074\u0020n\u006f\u0074\u0020b\u0065 \u006e\u0069\u006c");
};if _ga ._fbc ==nil {return _c .New ("\u0063\u0061\u0043\u0065rt\u0020\u006d\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006ei\u006c");};if _ga ._cba ==nil {return _c .New ("\u0070\u0072\u0069\u0076\u0061\u0074\u0065\u004b\u0065\u0079\u0020m\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0062\u0065 \u006e\u0069\u006c");
};};_cf :=*_ga ;sig .Handler =&_cf ;sig .Filter =_ac .MakeName ("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065");sig .SubFilter =_ac .MakeName ("\u0045\u0054\u0053\u0049.C\u0041\u0064\u0045\u0053\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064");
sig .Reference =nil ;_abb ,_caa :=_cf .NewDigest (sig );if _caa !=nil {return _caa ;};_ ,_caa =_abb .Write ([]byte ("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"));
if _caa !=nil {return _caa ;};_cf ._gcd =true ;_caa =_cf .Sign (sig ,_abb );_cf ._gcd =false ;return _caa ;};

// InitSignature initialises the PdfSignature.
func (_fbeda *adobeX509RSASHA1 )InitSignature (sig *_ag .PdfSignature )error {if _fbeda ._cbe ==nil {return _c .New ("c\u0065\u0072\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u006d\u0075\u0073\u0074\u0020n\u006f\u0074\u0020b\u0065 \u006e\u0069\u006c");
};if _fbeda ._fgf ==nil &&_fbeda ._bga ==nil {return _c .New ("\u006d\u0075\u0073\u0074\u0020\u0070\u0072o\u0076\u0069\u0064e\u0020\u0065\u0069t\u0068\u0065r\u0020\u0061\u0020\u0070\u0072\u0069v\u0061te\u0020\u006b\u0065\u0079\u0020\u006f\u0072\u0020\u0061\u0020\u0073\u0069\u0067\u006e\u0069\u006e\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e");
};_aba :=*_fbeda ;sig .Handler =&_aba ;sig .Filter =_ac .MakeName ("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065");sig .SubFilter =_ac .MakeName ("\u0061d\u0062e\u002e\u0078\u0035\u0030\u0039.\u0072\u0073a\u005f\u0073\u0068\u0061\u0031");
sig .Cert =_ac .MakeString (string (_aba ._cbe .Raw ));sig .Reference =nil ;_bbge ,_cgcb :=_aba .NewDigest (sig );if _cgcb !=nil {return _cgcb ;};_bbge .Write ([]byte ("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"));
return _aba .sign (sig ,_bbge ,_fbeda ._afc );};

// NewDocMDPHandler returns the new DocMDP handler with the specific DocMDP restriction level.
func NewDocMDPHandler (handler _ag .SignatureHandler ,permission _ff .DocMDPPermission )(_ag .SignatureHandler ,error ){return &DocMDPHandler {_gd :handler ,Permission :permission },nil ;};

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_fbg *DocMDPHandler )IsApplicable (sig *_ag .PdfSignature )bool {_db :=false ;for _ ,_dc :=range sig .Reference .Elements (){if _eg ,_cdf :=_ac .GetDict (_dc );_cdf {if _fd ,_ce :=_ac .GetNameVal (_eg .Get ("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u004d\u0065\u0074\u0068\u006f\u0064"));
_ce {if _fd !="\u0044\u006f\u0063\u004d\u0044\u0050"{return false ;};if _aeg ,_ba :=_ac .GetDict (_eg .Get ("\u0054r\u0061n\u0073\u0066\u006f\u0072\u006d\u0050\u0061\u0072\u0061\u006d\u0073"));_ba {_ ,_ca :=_ac .GetNumberAsInt64 (_aeg .Get ("\u0050"));
if _ca !=nil {return false ;};_db =true ;break ;};};};};return _db &&_fbg ._gd .IsApplicable (sig );};

// NewDigest creates a new digest.
func (_aeb *adobeX509RSASHA1 )NewDigest (sig *_ag .PdfSignature )(_ag .Hasher ,error ){if _dgfa ,_fba :=_aeb .getHashAlgorithm (sig );_dgfa !=0&&_fba ==nil {return _dgfa .New (),nil ;};return _egb .New (),nil ;};

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_babe *adobeX509RSASHA1 )IsApplicable (sig *_ag .PdfSignature )bool {if sig ==nil ||sig .Filter ==nil ||sig .SubFilter ==nil {return false ;};return (*sig .Filter =="A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053"||*sig .Filter =="\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")&&*sig .SubFilter =="\u0061d\u0062e\u002e\u0078\u0035\u0030\u0039.\u0072\u0073a\u005f\u0073\u0068\u0061\u0031";
};

// Validate validates PdfSignature.
func (_dea *adobeX509RSASHA1 )Validate (sig *_ag .PdfSignature ,digest _ag .Hasher )(_ag .SignatureValidationResult ,error ){_agc ,_dced :=_dea .getCertificate (sig );if _dced !=nil {return _ag .SignatureValidationResult {},_dced ;};_gde :=sig .Contents .Bytes ();
var _bdb []byte ;if _ ,_ccb :=_dg .Unmarshal (_gde ,&_bdb );_ccb !=nil {return _ag .SignatureValidationResult {},_ccb ;};_ffd ,_aaa :=digest .(_d .Hash );if !_aaa {return _ag .SignatureValidationResult {},_c .New ("\u0068a\u0073h\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072");
};_cfc ,_ :=_dea .getHashAlgorithm (sig );if _cfc ==0{_cfc =_egb ;};if _gfg :=_ae .VerifyPKCS1v15 (_agc .PublicKey .(*_ae .PublicKey ),_cfc ,_ffd .Sum (nil ),_bdb );_gfg !=nil {return _ag .SignatureValidationResult {},_gfg ;};return _ag .SignatureValidationResult {IsSigned :true ,IsVerified :true },nil ;
};func (_edd *docTimeStamp )getCertificate (_fee *_ag .PdfSignature )(*_e .Certificate ,error ){_gce ,_bcga :=_fee .GetCerts ();if _bcga !=nil {return nil ,_bcga ;};return _gce [0],nil ;};func (_cda *adobeX509RSASHA1 )sign (_fcg *_ag .PdfSignature ,_adad _ag .Hasher ,_gcg bool )error {if !_gcg {return _cda .Sign (_fcg ,_adad );
};_bfaa ,_cfcf :=_cda ._cbe .PublicKey .(*_ae .PublicKey );if !_cfcf {return _a .Errorf ("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0070\u0075\u0062\u006c\u0069\u0063\u0020\u006b\u0065y\u0020\u0074\u0079p\u0065:\u0020\u0025\u0054",_bfaa );};_bfe ,_bbc :=_dg .Marshal (make ([]byte ,_bfaa .Size ()));
if _bbc !=nil {return _bbc ;};_fcg .Contents =_ac .MakeHexString (string (_bfe ));return nil ;};type adobePKCS7Detached struct{_egg *_ae .PrivateKey ;_fcd *_e .Certificate ;_bac bool ;_gga int ;};func _ffad (_bee []byte ,_bec int )(_fafa []byte ){_ddfe :=len (_bee );
if _ddfe > _bec {_ddfe =_bec ;};_fafa =make ([]byte ,_bec );copy (_fafa [len (_fafa )-_ddfe :],_bee );return ;};type etsiPAdES struct{_cba *_ae .PrivateKey ;_aa *_e .Certificate ;_cea bool ;_gcd bool ;_fbc *_e .Certificate ;_dbd string ;};func (_ddd *etsiPAdES )makeTimestampRequest (_bbf string ,_dee []byte )(_dg .RawValue ,error ){_aab :=_cd .SHA512 .New ();
_aab .Write (_dee );_cg :=_aab .Sum (nil );_fc :=_da .Request {HashAlgorithm :_cd .SHA512 ,HashedMessage :_cg ,Certificates :true ,Extensions :nil ,ExtraExtensions :nil };_ceb ,_fed :=_fc .Marshal ();if _fed !=nil {return _dg .RawValue {},_fed ;};_dcf ,_fed :=_fa .NewRequest (_fa .MethodPost ,_bbf ,_dd .NewBuffer (_ceb ));
if _fed !=nil {return _dg .RawValue {},_fed ;};_dcf .Header .Set ("\u0043\u006f\u006et\u0065\u006e\u0074\u002d\u0054\u0079\u0070\u0065","a\u0070\u0070\u006c\u0069\u0063\u0061t\u0069\u006f\u006e\u002f\u0074\u0069\u006d\u0065\u0073t\u0061\u006d\u0070-\u0071u\u0065\u0072\u0079");
_gcf :=_fb .NewTimestampClient ();_ebg ,_fed :=_gcf .HTTPClient .Do (_dcf );if _fed !=nil {return _dg .RawValue {},_fed ;};defer _ebg .Body .Close ();_fgg ,_fed :=_fe .ReadAll (_ebg .Body );if _fed !=nil {return _dg .RawValue {},_fed ;};if _ebg .StatusCode !=_fa .StatusOK {return _dg .RawValue {},_a .Errorf ("\u0068\u0074\u0074\u0070\u0020\u0073\u0074\u0061\u0074\u0075\u0073\u0020\u0063\u006f\u0064e\u0020n\u006f\u0074\u0020\u006f\u006b\u0020\u0028\u0067\u006f\u0074\u0020\u0025\u0064\u0029",_ebg .StatusCode );
};var _cbd struct{Version _dg .RawValue ;Content _dg .RawValue ;};_ ,_fed =_dg .Unmarshal (_fgg ,&_cbd );if _fed !=nil {return _dg .RawValue {},_fed ;};return _cbd .Content ,nil ;};

// InitSignature initialises the PdfSignature.
func (_adf *docTimeStamp )InitSignature (sig *_ag .PdfSignature )error {_egca :=*_adf ;sig .Handler =&_egca ;sig .Filter =_ac .MakeName ("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065");sig .SubFilter =_ac .MakeName ("\u0045\u0054\u0053I\u002e\u0052\u0046\u0043\u0033\u0031\u0036\u0031");
sig .Reference =nil ;if _adf ._agbg > 0{sig .Contents =_ac .MakeHexString (string (make ([]byte ,_adf ._agbg )));}else {_aaae ,_gbe :=_adf .NewDigest (sig );if _gbe !=nil {return _gbe ;};_aaae .Write ([]byte ("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"));
if _gbe =_egca .Sign (sig ,_aaae );_gbe !=nil {return _gbe ;};_adf ._agbg =_egca ._agbg ;};return nil ;};func (_cdbb *adobeX509RSASHA1 )getHashAlgorithm (_ede *_ag .PdfSignature )(_cd .Hash ,error ){_gbg ,_fea :=_cdbb .getCertificate (_ede );if _fea !=nil {if _cdbb ._faf !=0{return _cdbb ._faf ,nil ;
};return _egb ,_fea ;};if _ede .Contents !=nil {_fdb :=_ede .Contents .Bytes ();var _agbf []byte ;if _ ,_cbc :=_dg .Unmarshal (_fdb ,&_agbf );_cbc ==nil {_fdba :=_bdg (_gbg .PublicKey .(*_ae .PublicKey ),_agbf );if _fdba > 0{return _fdba ,nil ;};};};if _cdbb ._faf !=0{return _cdbb ._faf ,nil ;
};return _egb ,nil ;};

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_ggb *docTimeStamp )IsApplicable (sig *_ag .PdfSignature )bool {if sig ==nil ||sig .Filter ==nil ||sig .SubFilter ==nil {return false ;};return (*sig .Filter =="A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053"||*sig .Filter =="\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")&&*sig .SubFilter =="\u0045\u0054\u0053I\u002e\u0052\u0046\u0043\u0033\u0031\u0036\u0031";
};const _egb =_cd .SHA1 ;

// ValidateWithOpts validates a PDF signature by checking PdfReader or PdfParser by the DiffPolicy
// params describes parameters for the DocMDP checks.
func (_gc *DocMDPHandler )ValidateWithOpts (sig *_ag .PdfSignature ,digest _ag .Hasher ,params _ag .SignatureHandlerDocMDPParams )(_ag .SignatureValidationResult ,error ){_cef ,_ffa :=_gc ._gd .Validate (sig ,digest );if _ffa !=nil {return _cef ,_ffa ;
};_cb :=params .Parser ;if _cb ==nil {return _ag .SignatureValidationResult {},_c .New ("p\u0061r\u0073\u0065\u0072\u0020\u0063\u0061\u006e\u0027t\u0020\u0062\u0065\u0020nu\u006c\u006c");};if !_cef .IsVerified {return _cef ,nil ;};_cbf :=params .DiffPolicy ;
if _cbf ==nil {_cbf =_ff .NewDefaultDiffPolicy ();};for _fgd :=0;_fgd <=_cb .GetRevisionNumber ();_fgd ++{_gf ,_deg :=_cb .GetRevision (_fgd );if _deg !=nil {return _ag .SignatureValidationResult {},_deg ;};_acc :=_gf .GetTrailer ();if _acc ==nil {return _ag .SignatureValidationResult {},_c .New ("\u0075\u006e\u0064\u0065f\u0069\u006e\u0065\u0064\u0020\u0074\u0068\u0065\u0020\u0074r\u0061i\u006c\u0065\u0072\u0020\u006f\u0062\u006ae\u0063\u0074");
};_ea ,_gfe :=_ac .GetDict (_acc .Get ("\u0052\u006f\u006f\u0074"));if !_gfe {return _ag .SignatureValidationResult {},_c .New ("\u0075n\u0064\u0065\u0066\u0069n\u0065\u0064\u0020\u0074\u0068e\u0020r\u006fo\u0074\u0020\u006f\u0062\u006a\u0065\u0063t");
};_ad ,_gfe :=_ac .GetDict (_ea .Get ("\u0041\u0063\u0072\u006f\u0046\u006f\u0072\u006d"));if !_gfe {continue ;};_dec ,_gfe :=_ac .GetArray (_ad .Get ("\u0046\u0069\u0065\u006c\u0064\u0073"));if !_gfe {continue ;};for _ ,_ddf :=range _dec .Elements (){_abf ,_fae :=_ac .GetDict (_ddf );
if !_fae {continue ;};_adg ,_fae :=_ac .GetDict (_abf .Get ("\u0056"));if !_fae {continue ;};if _ac .EqualObjects (_adg .Get ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"),sig .Contents ){_cef .DiffResults ,_deg =_cbf .ReviewFile (_gf ,_cb ,&_ff .MDPParameters {DocMDPLevel :_gc .Permission });
if _deg !=nil {return _ag .SignatureValidationResult {},_deg ;};_cef .IsVerified =_cef .DiffResults .IsPermitted ();return _cef ,nil ;};};};return _ag .SignatureValidationResult {},_c .New ("\u0064\u006f\u006e\u0027\u0074\u0020\u0066o\u0075\u006e\u0064 \u0074\u0068\u0069\u0073 \u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e\u0073");
};func _feac (_beg _dg .ObjectIdentifier )(_cd .Hash ,error ){switch {case _beg .Equal (_fg .OIDDigestAlgorithmSHA1 ),_beg .Equal (_fg .OIDDigestAlgorithmECDSASHA1 ),_beg .Equal (_fg .OIDDigestAlgorithmDSA ),_beg .Equal (_fg .OIDDigestAlgorithmDSASHA1 ),_beg .Equal (_fg .OIDEncryptionAlgorithmRSA ):return _cd .SHA1 ,nil ;
case _beg .Equal (_fg .OIDDigestAlgorithmSHA256 ),_beg .Equal (_fg .OIDDigestAlgorithmECDSASHA256 ):return _cd .SHA256 ,nil ;case _beg .Equal (_fg .OIDDigestAlgorithmSHA384 ),_beg .Equal (_fg .OIDDigestAlgorithmECDSASHA384 ):return _cd .SHA384 ,nil ;case _beg .Equal (_fg .OIDDigestAlgorithmSHA512 ),_beg .Equal (_fg .OIDDigestAlgorithmECDSASHA512 ):return _cd .SHA512 ,nil ;
};return _cd .Hash (0),_fg .ErrUnsupportedAlgorithm ;};func (_cbag *adobeX509RSASHA1 )getCertificate (_bff *_ag .PdfSignature )(*_e .Certificate ,error ){if _cbag ._cbe !=nil {return _cbag ._cbe ,nil ;};_bad ,_cdd :=_bff .GetCerts ();if _cdd !=nil {return nil ,_cdd ;
};return _bad [0],nil ;};

// Sign adds a new reference to signature's references array.
func (_fbe *DocMDPHandler )Sign (sig *_ag .PdfSignature ,digest _ag .Hasher )error {return _fbe ._gd .Sign (sig ,digest );};

// NewDocTimeStampWithOpts returns a new DocTimeStamp configured using the
// specified options. If no options are provided, default options will be used.
// Both the timestamp server URL and the hash algorithm can be empty for the
// signature validation.
// The following hash algorithms are supported:
// crypto.SHA1, crypto.SHA256, crypto.SHA384, crypto.SHA512.
func NewDocTimeStampWithOpts (timestampServerURL string ,hashAlgorithm _cd .Hash ,opts *DocTimeStampOpts )(_ag .SignatureHandler ,error ){if opts ==nil {opts =&DocTimeStampOpts {};};if opts .SignatureSize <=0{opts .SignatureSize =4192;};return &docTimeStamp {_gag :timestampServerURL ,_cebb :hashAlgorithm ,_agbg :opts .SignatureSize ,_gdc :opts .Client },nil ;
};

// SignFunc represents a custom signing function. The function should return
// the computed signature.
type SignFunc func (_eac *_ag .PdfSignature ,_dgb _ag .Hasher )([]byte ,error );

// Sign sets the Contents fields for the PdfSignature.
func (_bbee *adobeX509RSASHA1 )Sign (sig *_ag .PdfSignature ,digest _ag .Hasher )error {var _eda []byte ;var _geg error ;if _bbee ._bga !=nil {_eda ,_geg =_bbee ._bga (sig ,digest );if _geg !=nil {return _geg ;};}else {_adaf ,_agf :=digest .(_d .Hash );
if !_agf {return _c .New ("\u0068a\u0073h\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072");};_eab :=_egb ;if _bbee ._faf !=0{_eab =_bbee ._faf ;};_eda ,_geg =_ae .SignPKCS1v15 (_de .Reader ,_bbee ._fgf ,_eab ,_adaf .Sum (nil ));if _geg !=nil {return _geg ;
};};_eda ,_geg =_dg .Marshal (_eda );if _geg !=nil {return _geg ;};sig .Contents =_ac .MakeHexString (string (_eda ));return nil ;};

// NewDigest creates a new digest.
func (_bg *etsiPAdES )NewDigest (_ *_ag .PdfSignature )(_ag .Hasher ,error ){return _dd .NewBuffer (nil ),nil ;};type timestampInfo struct{Version int ;Policy _dg .RawValue ;MessageImprint struct{HashAlgorithm _g .AlgorithmIdentifier ;HashedMessage []byte ;
};SerialNumber _dg .RawValue ;GeneralizedTime _b .Time ;};func (_ceg *adobePKCS7Detached )getCertificate (_edg *_ag .PdfSignature )(*_e .Certificate ,error ){if _ceg ._fcd !=nil {return _ceg ._fcd ,nil ;};_gaf ,_gcc :=_edg .GetCerts ();if _gcc !=nil {return nil ,_gcc ;
};return _gaf [0],nil ;};

// Sign sets the Contents fields for the PdfSignature.
func (_fdf *docTimeStamp )Sign (sig *_ag .PdfSignature ,digest _ag .Hasher )error {_dca ,_ddb :=_fb .NewTimestampRequest (digest .(*_dd .Buffer ),&_da .RequestOptions {Hash :_fdf ._cebb ,Certificates :true });if _ddb !=nil {return _ddb ;};_decf :=_fdf ._gdc ;
if _decf ==nil {_decf =_fb .NewTimestampClient ();};_gfd ,_ddb :=_decf .GetEncodedToken (_fdf ._gag ,_dca );if _ddb !=nil {return _ddb ;};_bbec :=len (_gfd );if _fdf ._agbg > 0&&_bbec > _fdf ._agbg {return _ag .ErrSignNotEnoughSpace ;};if _bbec > 0{_fdf ._agbg =_bbec +128;
};if sig .Contents !=nil {_gfea :=sig .Contents .Bytes ();copy (_gfea ,_gfd );_gfd =_gfea ;};sig .Contents =_ac .MakeHexString (string (_gfd ));return nil ;};

// AdobeX509RSASHA1Opts defines options for configuring the adbe.x509.rsa_sha1
// signature handler.
type AdobeX509RSASHA1Opts struct{

// EstimateSize specifies whether the size of the signature contents
// should be estimated based on the modulus size of the public key
// extracted from the signing certificate. If set to false, a mock Sign
// call is made in order to estimate the size of the signature contents.
EstimateSize bool ;

// Algorithm specifies the algorithm used for performing signing.
// If not specified, defaults to SHA1.
Algorithm _cd .Hash ;};

// NewEtsiPAdESLevelB creates a new Adobe.PPKLite ETSI.CAdES.detached Level B signature handler.
func NewEtsiPAdESLevelB (privateKey *_ae .PrivateKey ,certificate *_e .Certificate ,caCert *_e .Certificate )(_ag .SignatureHandler ,error ){return &etsiPAdES {_aa :certificate ,_cba :privateKey ,_fbc :caCert },nil ;};

// IsApplicable returns true if the signature handler is applicable for the PdfSignature
func (_fbeg *adobePKCS7Detached )IsApplicable (sig *_ag .PdfSignature )bool {if sig ==nil ||sig .Filter ==nil ||sig .SubFilter ==nil {return false ;};return (*sig .Filter =="A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053"||*sig .Filter =="\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")&&*sig .SubFilter =="\u0061\u0064\u0062\u0065.p\u006b\u0063\u0073\u0037\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064";
};type adobeX509RSASHA1 struct{_fgf *_ae .PrivateKey ;_cbe *_e .Certificate ;_bga SignFunc ;_afc bool ;_faf _cd .Hash ;};