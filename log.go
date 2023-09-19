/*
@Time : 2021/4/23 21:09
@Author : LiuKun
@File : log
@Software: GoLand
@Description:
*/

package clog

import (
	"fmt"
	"github.com/gookit/color"
	"log"
	"os"
	"path/filepath"
	"time"
)

// Level 日志等级
type Level int

const (
	DebugLevel           Level = iota //调试
	ReleaseInfoLevel                  //信息
	ReleaseStartLevel                 //开始
	ReleaseCompleteLevel              //完成
	ReleaseSuccessLevel               //成功
	ReleaseErrInfoLevel               //错误信息
	ReleaseFailLevel                  //失败
	ErrorLevel                        //严重错误
	FatalLevel                        //致命错误，程序终止
)

func (l Level) Value() int {
	return int(l)
}

func (l Level) Prefix() string {
	return []string{
		"[debug   ] ",
		"[info    ] ",
		"[start   ] ",
		"[complete] ",
		"[success ] ",
		"[errInfo ] ",
		"[fail    ] ",
		"[error   ] ",
		"[fatal   ] ",
	}[l.Value()]
}

func (l Level) ColorPrinter() color.PrinterFace {
	return []color.PrinterFace{
		color.Gray,     //灰色
		color.Magenta,  //品红，淡紫色
		color.Cyan,     //青色
		color.Blue,     //蓝色
		color.Green,    //绿色
		color.Yellow,   //黄色
		color.LightRed, //浅红
		color.Red,      //红色
		color.White,    //白色
	}[l.Value()]
}

type Logger struct {
	ConsolePrintWhenHasFile bool        //打印到文件时是否打印到控制台
	minLevel                Level       //打印的最小等级
	fileLogger              *log.Logger //文件Logger
	file                    *os.File    //文件
}

func New(level Level, pathname string, flag int) (*Logger, error) {

	// logger
	var fileLogger *log.Logger
	var file *os.File
	if pathname != "" {
		now := time.Now()

		filename := fmt.Sprintf("%d%02d%02d_%02d_%02d_%02d.log",
			now.Year(),
			now.Month(),
			now.Day(),
			now.Hour(),
			now.Minute(),
			now.Second())

		f, err := os.Create(filepath.Join(pathname, filename))
		if err != nil {
			return nil, err
		}

		fileLogger = log.New(f, "", flag)
		file = f
	}

	// new
	logger := new(Logger)
	logger.minLevel = level
	logger.fileLogger = fileLogger
	logger.file = file
	logger.ConsolePrintWhenHasFile = true

	return logger, nil
}

// Close It's dangerous to call the method on logging
func (logger *Logger) Close() {
	if logger.file != nil {
		_ = logger.file.Close()
	}

	logger.fileLogger = nil
	logger.file = nil
}

func (logger *Logger) DoPrintf(printLevel Level, format string, a ...interface{}) {
	if printLevel < logger.minLevel {
		return
	}

	format = printLevel.Prefix() + format
	if logger.fileLogger != nil {
		_ = logger.fileLogger.Output(3, fmt.Sprintf(format, a...))
	}

	if logger.fileLogger == nil || logger.ConsolePrintWhenHasFile {
		format = time.Now().Format("2006/01/02 15:04:05 ") + format
		printLevel.ColorPrinter().Printf(format, a...)
	}

	if printLevel == FatalLevel {
		os.Exit(1)
	}
}

var gLogger, _ = New(DebugLevel, "", log.LstdFlags)

// Export It's dangerous to call the method on logging
func Export(logger *Logger) {
	if logger != nil {
		gLogger = logger
	}
}

// Debug 调试信息， 灰色
func Debug(format string, a ...interface{}) {
	gLogger.DoPrintf(DebugLevel, format, a...)
}

// Info 基本信息，紫色
func Info(format string, a ...interface{}) {
	gLogger.DoPrintf(ReleaseInfoLevel, format, a...)
}

// Start 任务开始信息，青色
func Start(format string, a ...interface{}) {
	gLogger.DoPrintf(ReleaseStartLevel, format, a...)
}

// Complete 任务完成信息，蓝色
func Complete(format string, a ...interface{}) {
	gLogger.DoPrintf(ReleaseCompleteLevel, format, a...)

}

// Success 成功信息，绿色
func Success(format string, a ...interface{}) {
	gLogger.DoPrintf(ReleaseSuccessLevel, format, a...)
}

// ErrInfo 不太重要的错误信息，黄色
func ErrInfo(format string, a ...interface{}) {
	gLogger.DoPrintf(ReleaseErrInfoLevel, format, a...)
}

// Fail 失败信息，浅红色
func Fail(format string, a ...interface{}) {
	gLogger.DoPrintf(ReleaseFailLevel, format, a...)
}

// Error 严重错误，红色
func Error(format string, a ...interface{}) {
	gLogger.DoPrintf(ErrorLevel, format, a...)
}

// Fatal 致命错误, 白色
func Fatal(format string, a ...interface{}) {
	gLogger.DoPrintf(FatalLevel, format, a...)
}

// Close 关闭文件日志输出
func Close() {
	gLogger.Close()
}
