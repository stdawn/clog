/**
 * @Time: 2023/9/19 16:55
 * @Author: LiuKun
 * @File: log_test.go
 * @Description:
 */

package clog

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestAll(t *testing.T) {

	p, err := os.Executable()
	if err != nil {
		Error(err.Error())
		return
	}
	dir := filepath.Dir(p)

	l, _ := New(DebugLevel, dir, log.LstdFlags)
	l.ConsoleNotPrintWhenHasFile = true
	Export(l)

	Debug(DebugLevel.Prefix())
	Info(ReleaseInfoLevel.Prefix())
	Start(ReleaseStartLevel.Prefix())
	l.SetLoggerFlag(log.LstdFlags | log.Llongfile)
	Complete(ReleaseCompleteLevel.Prefix() + "\n\n")
	Success(ReleaseSuccessLevel.Prefix() + "\n")
	ErrInfo(ReleaseErrInfoLevel.Prefix() + "\n")
	Fail(ReleaseFailLevel.Prefix() + "\n")
	Error(ErrorLevel.Prefix() + "\n")
	//Fatal(FatalLevel.Prefix() + "\n")
}
