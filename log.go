package log

import (
	"fmt"
	"io"
	"time"

	stdlog "log"

	"github.com/fatih/color"
	"github.com/rs/zerolog"
)

// TimeFormat is the default time format used by this package
var TimeFormat = "2006-01-02 15:04:05Z07:00"

// TimestampFunc is the default time func used by this package, useful for testing
var TimestampFunc = time.Now

// Logger is the bare minimum logging interface
type Logger interface {
	Debug(v ...interface{})
	Log(v ...interface{})
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
	Fatal(v ...interface{})
}

// Zero is a basic re-implementation of zerolog to match the logging interface
type Zero struct {
	zerolog.Logger
	output io.Writer
}

// NewZero returns a new Zerolog logger
func NewZero(w io.Writer) *Zero {
	zerolog.TimeFieldFormat = TimeFormat
	return &Zero{
		Logger: zerolog.New(w).With().Timestamp().Logger(),
		output: w,
	}
}
func (z Zero) Pretty() *Zero {
	z.Logger = z.Logger.Output(zerolog.ConsoleWriter{Out: z.output, TimeFormat: TimeFormat})
	return &z
}
func (z Zero) Debug(v ...interface{}) {
	msg := fmt.Sprintln(v...)
	z.Logger.Debug().Msg(msg[0 : len(msg)-1])
}
func (z Zero) Info(v ...interface{}) {
	msg := fmt.Sprintln(v...)
	z.Logger.Info().Msg(msg[0 : len(msg)-1])
}
func (z Zero) Log(v ...interface{}) {
	msg := fmt.Sprintln(v...)
	z.Logger.Log().Msg(msg[0 : len(msg)-1])
}
func (z Zero) Warn(v ...interface{}) {
	msg := fmt.Sprintln(v...)
	z.Logger.Warn().Msg(msg[0 : len(msg)-1])
}
func (z Zero) Error(v ...interface{}) {
	msg := fmt.Sprintln(v...)
	z.Logger.Error().Msg(msg[0 : len(msg)-1])
}
func (z Zero) Fatal(v ...interface{}) {
	msg := fmt.Sprintln(v...)
	z.Logger.Fatal().Msg(msg[0 : len(msg)-1])
	stdlog.Fatal()
}

type std struct {
	tf     func() time.Time
	output io.Writer
}

func (s std) Write(bytes []byte) (int, error) {
	return fmt.Fprint(s.output, s.tf().Format(TimeFormat)+" "+string(bytes))
}

// Std is a basic re-implementation of the standard library log to match the logging interface
type Std struct {
	output io.Writer
	Logger stdlog.Logger
}

// NewStd Returns a new Zerolog logger
func NewStd(w io.Writer) *Std {
	s := std{
		tf:     TimestampFunc,
		output: w,
	}
	return &Std{
		output: s,
		Logger: *stdlog.New(s, "", 0),
	}
}
func (s Std) Debug(v ...interface{}) {
	msg := fmt.Sprintln(v...)
	s.Logger.Println("[Debug]", fmt.Sprint(msg[0:len(msg)-1]))
}
func (s Std) Info(v ...interface{}) {
	msg := fmt.Sprintln(v...)
	s.Logger.Println("[Info]", fmt.Sprint(msg[0:len(msg)-1]))
}
func (s Std) Log(v ...interface{}) {
	msg := fmt.Sprintln(v...)
	s.Logger.Println("[Log]", fmt.Sprint(msg[0:len(msg)-1]))
}
func (s Std) Warn(v ...interface{}) {
	msg := fmt.Sprintln(v...)
	s.Logger.Println("[Warn]", fmt.Sprint(msg[0:len(msg)-1]))
}
func (s Std) Error(v ...interface{}) {
	msg := fmt.Sprintln(v...)
	s.Logger.Println("[Error]", fmt.Sprint(msg[0:len(msg)-1]))
}
func (s Std) Fatal(v ...interface{}) {
	msg := fmt.Sprintln(v...)
	s.Logger.Println("[Fatal]", fmt.Sprint(msg[0:len(msg)-1]))
}

// NewColor returns a new Color logger
func NewColor(w io.Writer) *Color {
	s := std{
		tf:     TimestampFunc,
		output: w,
	}
	return &Color{
		output: s,
		Logger: *stdlog.New(s, "", 0),
	}
}

// Color is a basic re-implementation of the standard library log to match the logging interface
// with added colors
type Color struct {
	output io.Writer
	Logger stdlog.Logger
}

func (c Color) Debug(v ...interface{}) {
	msg := fmt.Sprintln(v...)
	c.Logger.Println(color.GreenString("[Debug]"), fmt.Sprint(msg[0:len(msg)-1]))
}
func (c Color) Info(v ...interface{}) {
	msg := fmt.Sprintln(v...)
	c.Logger.Println(color.HiBlueString("[Info]"), fmt.Sprint(msg[0:len(msg)-1]))
}
func (c Color) Log(v ...interface{}) {
	msg := fmt.Sprintln(v...)
	c.Logger.Println(color.WhiteString("[Log]"), fmt.Sprint(msg[0:len(msg)-1]))
}
func (c Color) Warn(v ...interface{}) {
	msg := fmt.Sprintln(v...)
	c.Logger.Println(color.YellowString("[Warn]"), fmt.Sprint(msg[0:len(msg)-1]))
}
func (c Color) Error(v ...interface{}) {
	msg := fmt.Sprintln(v...)
	c.Logger.Println(color.RedString("[Error]"), fmt.Sprint(msg[0:len(msg)-1]))
}
func (c Color) Fatal(v ...interface{}) {
	msg := fmt.Sprintln(v...)
	c.Logger.Println(color.HiRedString("[Fatal]"), fmt.Sprint(msg[0:len(msg)-1]))
	stdlog.Fatal()
}
