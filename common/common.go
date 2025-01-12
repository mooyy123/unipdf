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

// Package common contains common properties used by the subpackages.
package common ;import (_c "fmt";_ce "io";_g "os";_bd "path/filepath";_b "runtime";_f "time";);

// Trace logs trace message.
func (_cde ConsoleLogger )Trace (format string ,args ...interface{}){if _cde .LogLevel >=LogLevelTrace {_aa :="\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020";_cde .output (_g .Stdout ,_aa ,format ,args ...);};};const (LogLevelTrace LogLevel =5;LogLevelDebug LogLevel =4;
LogLevelInfo LogLevel =3;LogLevelNotice LogLevel =2;LogLevelWarning LogLevel =1;LogLevelError LogLevel =0;);func (_ffd WriterLogger )logToWriter (_gfb _ce .Writer ,_ef string ,_cee string ,_bgc ...interface{}){_fd (_gfb ,_ef ,_cee ,_bgc );};const _egb =30;
var ReleasedAt =_f .Date (_bfa ,_dee ,_fe ,_acg ,_egb ,0,0,_f .UTC );

// Debug logs debug message.
func (_ad ConsoleLogger )Debug (format string ,args ...interface{}){if _ad .LogLevel >=LogLevelDebug {_ac :="\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020";_ad .output (_g .Stdout ,_ac ,format ,args ...);};};

// Error does nothing for dummy logger.
func (DummyLogger )Error (format string ,args ...interface{}){};

// LogLevel is the verbosity level for logging.
type LogLevel int ;

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_gf WriterLogger )IsLogLevel (level LogLevel )bool {return _gf .LogLevel >=level };const _acg =15;

// Notice logs notice message.
func (_cf WriterLogger )Notice (format string ,args ...interface{}){if _cf .LogLevel >=LogLevelNotice {_bdc :="\u005bN\u004f\u0054\u0049\u0043\u0045\u005d ";_cf .logToWriter (_cf .Output ,_bdc ,format ,args ...);};};

// Info logs info message.
func (_ba WriterLogger )Info (format string ,args ...interface{}){if _ba .LogLevel >=LogLevelInfo {_be :="\u005bI\u004e\u0046\u004f\u005d\u0020";_ba .logToWriter (_ba .Output ,_be ,format ,args ...);};};const Version ="\u0033\u002e\u0034\u0039\u002e\u0030";


// NewWriterLogger creates new 'writer' logger.
func NewWriterLogger (logLevel LogLevel ,writer _ce .Writer )*WriterLogger {_bf :=WriterLogger {Output :writer ,LogLevel :logLevel };return &_bf ;};const _ccc ="\u0032\u0020\u004aan\u0075\u0061\u0072\u0079\u0020\u0032\u0030\u0030\u0036\u0020\u0061\u0074\u0020\u0031\u0035\u003a\u0030\u0034";
func _fd (_bc _ce .Writer ,_fde string ,_ggd string ,_de ...interface{}){_ ,_bbb ,_ee ,_fdg :=_b .Caller (3);if !_fdg {_bbb ="\u003f\u003f\u003f";_ee =0;}else {_bbb =_bd .Base (_bbb );};_ggf :=_c .Sprintf ("\u0025s\u0020\u0025\u0073\u003a\u0025\u0064 ",_fde ,_bbb ,_ee )+_ggd +"\u000a";
_c .Fprintf (_bc ,_ggf ,_de ...);};

// Debug does nothing for dummy logger.
func (DummyLogger )Debug (format string ,args ...interface{}){};

// DummyLogger does nothing.
type DummyLogger struct{};func (_ebc ConsoleLogger )output (_acf _ce .Writer ,_acd string ,_afd string ,_cb ...interface{}){_fd (_acf ,_acd ,_afd ,_cb ...);};const _dee =8;

// SetLogger sets 'logger' to be used by the unidoc unipdf library.
func SetLogger (logger Logger ){Log =logger };

// IsLogLevel returns true from dummy logger.
func (DummyLogger )IsLogLevel (level LogLevel )bool {return true };

// Notice does nothing for dummy logger.
func (DummyLogger )Notice (format string ,args ...interface{}){};const _bfa =2023;

// Trace does nothing for dummy logger.
func (DummyLogger )Trace (format string ,args ...interface{}){};

// Notice logs notice message.
func (_cdc ConsoleLogger )Notice (format string ,args ...interface{}){if _cdc .LogLevel >=LogLevelNotice {_eb :="\u005bN\u004f\u0054\u0049\u0043\u0045\u005d ";_cdc .output (_g .Stdout ,_eb ,format ,args ...);};};

// WriterLogger is the logger that writes data to the Output writer
type WriterLogger struct{LogLevel LogLevel ;Output _ce .Writer ;};

// UtcTimeFormat returns a formatted string describing a UTC timestamp.
func UtcTimeFormat (t _f .Time )string {return t .Format (_ccc )+"\u0020\u0055\u0054\u0043"};const _fe =2;

// ConsoleLogger is a logger that writes logs to the 'os.Stdout'
type ConsoleLogger struct{LogLevel LogLevel ;};

// Info does nothing for dummy logger.
func (DummyLogger )Info (format string ,args ...interface{}){};

// Error logs error message.
func (_bb ConsoleLogger )Error (format string ,args ...interface{}){if _bb .LogLevel >=LogLevelError {_ggb :="\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020";_bb .output (_g .Stdout ,_ggb ,format ,args ...);};};

// Warning does nothing for dummy logger.
func (DummyLogger )Warning (format string ,args ...interface{}){};

// Debug logs debug message.
func (_bed WriterLogger )Debug (format string ,args ...interface{}){if _bed .LogLevel >=LogLevelDebug {_ae :="\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020";_bed .logToWriter (_bed .Output ,_ae ,format ,args ...);};};

// Trace logs trace message.
func (_gee WriterLogger )Trace (format string ,args ...interface{}){if _gee .LogLevel >=LogLevelTrace {_bg :="\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020";_gee .logToWriter (_gee .Output ,_bg ,format ,args ...);};};

// Info logs info message.
func (_ec ConsoleLogger )Info (format string ,args ...interface{}){if _ec .LogLevel >=LogLevelInfo {_cg :="\u005bI\u004e\u0046\u004f\u005d\u0020";_ec .output (_g .Stdout ,_cg ,format ,args ...);};};

// Warning logs warning message.
func (_da WriterLogger )Warning (format string ,args ...interface{}){if _da .LogLevel >=LogLevelWarning {_db :="\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020";_da .logToWriter (_da .Output ,_db ,format ,args ...);};};

// Warning logs warning message.
func (_ge ConsoleLogger )Warning (format string ,args ...interface{}){if _ge .LogLevel >=LogLevelWarning {_af :="\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020";_ge .output (_g .Stdout ,_af ,format ,args ...);};};

// Error logs error message.
func (_cc WriterLogger )Error (format string ,args ...interface{}){if _cc .LogLevel >=LogLevelError {_age :="\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020";_cc .logToWriter (_cc .Output ,_age ,format ,args ...);};};

// Logger is the interface used for logging in the unipdf package.
type Logger interface{Error (_d string ,_cd ...interface{});Warning (_fg string ,_ea ...interface{});Notice (_gg string ,_a ...interface{});Info (_ff string ,_eg ...interface{});Debug (_ag string ,_df ...interface{});Trace (_ffe string ,_egd ...interface{});
IsLogLevel (_ab LogLevel )bool ;};

// NewConsoleLogger creates new console logger.
func NewConsoleLogger (logLevel LogLevel )*ConsoleLogger {return &ConsoleLogger {LogLevel :logLevel }};

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_ggc ConsoleLogger )IsLogLevel (level LogLevel )bool {return _ggc .LogLevel >=level };var Log Logger =DummyLogger {};