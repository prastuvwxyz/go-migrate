package log

import (
	"fmt"
	"io"

	"github.com/fatih/color"
)

var (
	DiagnosticWriter = color.Error
	OutputWriter     = color.Output
)

var (
	successSprintf = color.HiGreenString
	errorSprintf   = color.HiRedString
	warningSprintf = color.YellowString
	debugSprintf   = color.New(color.Faint).Sprintf
)

const (
	successPrefix = "✔"
	errorPrefix   = "✘"
	warningPrefix = "Note:"
)

func Success(args ...interface{}) {
	success(DiagnosticWriter, args...)
}

func Successln(args ...interface{}) {
	successln(DiagnosticWriter, args...)
}

func Successf(format string, args ...interface{}) {
	successf(DiagnosticWriter, format, args...)
}

func Ssuccess(args ...interface{}) string {
	return fmt.Sprintf("%s %s", successSprintf(successPrefix), fmt.Sprint(args...))
}

func Ssuccessln(args ...interface{}) string {
	msg := fmt.Sprintf("%s %s", successSprintf(successPrefix), fmt.Sprint(args...))
	return fmt.Sprintln(msg)
}

func Ssuccessf(format string, args ...interface{}) string {
	wrappedFormat := fmt.Sprintf("%s %s", successSprintf(successPrefix), format)
	return fmt.Sprintf(wrappedFormat, args...)
}

func Error(args ...interface{}) {
	err(DiagnosticWriter, args...)
}

func Errorln(args ...interface{}) {
	errln(DiagnosticWriter, args...)
}

func Errorf(format string, args ...interface{}) {
	errf(DiagnosticWriter, format, args...)
}

func Serror(args ...interface{}) string {
	return fmt.Sprintf("%s %s", errorSprintf(errorPrefix), fmt.Sprint(args...))
}

func Serrorln(args ...interface{}) string {
	msg := fmt.Sprintf("%s %s", errorSprintf(errorPrefix), fmt.Sprint(args...))
	return fmt.Sprintln(msg)
}

func Serrorf(format string, args ...interface{}) string {
	wrappedFormat := fmt.Sprintf("%s %s", errorSprintf(errorPrefix), format)
	return fmt.Sprintf(wrappedFormat, args...)
}

func Warning(args ...interface{}) {
	warning(DiagnosticWriter, args...)
}

func Warningln(args ...interface{}) {
	warningln(DiagnosticWriter, args...)
}

func Warningf(format string, args ...interface{}) {
	warningf(DiagnosticWriter, format, args...)
}

func Info(args ...interface{}) {
	info(DiagnosticWriter, args...)
}

func Infoln(args ...interface{}) {
	infoln(DiagnosticWriter, args...)
}

func Infof(format string, args ...interface{}) {
	infof(DiagnosticWriter, format, args...)
}

func Debug(args ...interface{}) {
	debug(DiagnosticWriter, args...)
}

func Debugln(args ...interface{}) {
	debugln(DiagnosticWriter, args...)
}

func Debugf(format string, args ...interface{}) {
	debugf(DiagnosticWriter, format, args...)
}

func success(w io.Writer, args ...interface{}) {
	msg := fmt.Sprintf("%s %s", successSprintf(successPrefix), fmt.Sprint(args...))
	fmt.Fprint(w, msg)
}

func successln(w io.Writer, args ...interface{}) {
	msg := fmt.Sprintf("%s %s", successSprintf(successPrefix), fmt.Sprint(args...))
	fmt.Fprintln(w, msg)
}

func successf(w io.Writer, format string, args ...interface{}) {
	wrappedFormat := fmt.Sprintf("%s %s", successSprintf(successPrefix), format)
	fmt.Fprintf(w, wrappedFormat, args...)
}

func err(w io.Writer, args ...interface{}) {
	msg := fmt.Sprintf("%s %s", errorSprintf(errorPrefix), fmt.Sprint(args...))
	fmt.Fprint(w, msg)
}

func errln(w io.Writer, args ...interface{}) {
	msg := fmt.Sprintf("%s %s", errorSprintf(errorPrefix), fmt.Sprint(args...))
	fmt.Fprintln(w, msg)
}

func errf(w io.Writer, format string, args ...interface{}) {
	wrappedFormat := fmt.Sprintf("%s %s", errorSprintf(errorPrefix), format)
	fmt.Fprintf(w, wrappedFormat, args...)
}

func warning(w io.Writer, args ...interface{}) {
	msg := fmt.Sprint(args...)
	fmt.Fprint(w, warningSprintf(fmt.Sprintf("%s %s", warningPrefix, msg)))
}

func warningln(w io.Writer, args ...interface{}) {
	msg := fmt.Sprint(args...)
	fmt.Fprintln(w, warningSprintf(fmt.Sprintf("%s %s", warningPrefix, msg)))
}

func warningf(w io.Writer, format string, args ...interface{}) {
	wrappedFormat := fmt.Sprintf("%s %s", warningPrefix, format)
	fmt.Fprint(w, warningSprintf(wrappedFormat, args...))
}

func info(w io.Writer, args ...interface{}) {
	fmt.Fprint(w, args...)
}

func infoln(w io.Writer, args ...interface{}) {
	fmt.Fprintln(w, args...)
}

func infof(w io.Writer, format string, args ...interface{}) {
	fmt.Fprintf(w, format, args...)
}

func debug(w io.Writer, args ...interface{}) {
	fmt.Fprint(w, debugSprintf(fmt.Sprint(args...)))
}

func debugln(w io.Writer, args ...interface{}) {
	fmt.Fprintln(w, debugSprintf(fmt.Sprint(args...)))
}

func debugf(w io.Writer, format string, args ...interface{}) {
	fmt.Fprint(w, debugSprintf(format, args...))
}
