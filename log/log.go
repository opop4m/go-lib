// Copyright 2014 mqant Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package log 日志初始化
package log

import (
	beegolog "github.com/opop4m/go-lib/log/beego"
)

var beego *beegolog.BeeLogger
var bi *beegolog.BeeLogger

// InitLog 初始化日志
func InitLog(debug bool, ProcessID string, Logdir string, settings map[string]interface{}) {
	beego = NewBeegoLogger(debug, ProcessID, Logdir, settings)
}

// InitBI 初始化BI日志
func InitBI(debug bool, ProcessID string, Logdir string, settings map[string]interface{}) {
	bi = NewBeegoLogger(debug, ProcessID, Logdir, settings)
}

// LogBeego LogBeego
func LogBeego() *beegolog.BeeLogger {
	if beego == nil {
		beego = beegolog.NewLogger()
	}
	return beego
}

// BiBeego BiBeego
func BiBeego() *beegolog.BeeLogger {
	return bi
}

// CreateRootTrace CreateRootTrace
func CreateRootTrace() TraceSpan {
	return &TraceSpanImp{
		Trace: GenerateID().String(),
		Span:  GenerateID().String(),
	}
}

// CreateTrace CreateTrace
func CreateTrace(trace, span string) TraceSpan {
	return &TraceSpanImp{
		Trace: trace,
		Span:  span,
	}
}

// BiReport BiReport
func BiReport(msg string) {
	//gLogger.doPrintf(debugLevel, printDebugLevel, format, a...)
	l := BiBeego()
	if l != nil {
		l.BiReport(msg)
	}
}

// Debug Debug
func Debug(format string, a ...interface{}) {
	//gLogger.doPrintf(debugLevel, printDebugLevel, format, a...)
	LogBeego().Debug(nil, format, a...)
}

// Info Info
func Info(format string, a ...interface{}) {
	//gLogger.doPrintf(releaseLevel, printReleaseLevel, format, a...)
	LogBeego().Info(nil, format, a...)
}

// Error Error
func Error(format string, a ...interface{}) {
	//gLogger.doPrintf(errorLevel, printErrorLevel, format, a...)
	LogBeego().Error(nil, format, a...)
}

// Warning Warning
func Warning(format string, a ...interface{}) {
	//gLogger.doPrintf(fatalLevel, printFatalLevel, format, a...)
	LogBeego().Warning(nil, format, a...)
}

// TDebug TDebug
func TDebug(span TraceSpan, format string, a ...interface{}) {
	if span != nil {
		LogBeego().Debug(
			&beegolog.BeegoTraceSpan{
				Trace: span.TraceId(),
				Span:  span.SpanId(),
			}, format, a...)
	} else {
		LogBeego().Debug(nil, format, a...)
	}
}

// TInfo TInfo
func TInfo(span TraceSpan, format string, a ...interface{}) {
	if span != nil {
		LogBeego().Info(
			&beegolog.BeegoTraceSpan{
				Trace: span.TraceId(),
				Span:  span.SpanId(),
			}, format, a...)
	} else {
		LogBeego().Info(nil, format, a...)
	}
}

// TError TError
func TError(span TraceSpan, format string, a ...interface{}) {
	if span != nil {
		LogBeego().Error(
			&beegolog.BeegoTraceSpan{
				Trace: span.TraceId(),
				Span:  span.SpanId(),
			}, format, a...)
	} else {
		LogBeego().Error(nil, format, a...)
	}
}

// TWarning TWarning
func TWarning(span TraceSpan, format string, a ...interface{}) {
	if span != nil {
		LogBeego().Warning(
			&beegolog.BeegoTraceSpan{
				Trace: span.TraceId(),
				Span:  span.SpanId(),
			}, format, a...)
	} else {
		LogBeego().Warning(nil, format, a...)
	}
}

// Close Close
func Close() {
	LogBeego().Close()
}

type LogConf struct {
	Debug       bool   `mapstructure:"debug" json:"debug" yaml:"debug"`
	Contenttype string `mapstructure:"contenttype" json:"contenttype" yaml:"env"`
	Dir         string `mapstructure:"dir" json:"dir" yaml:"dir"`
	BiDir       string `mapstructure:"bi-dir" json:"bi-dir" yaml:"bi-dir"`

	File      LogFile  `mapstructure:"file" yaml:"file" json:"file" `
	Multifile LogFile  `mapstructure:"multifile" yaml:"multifile" json:"multifile"`
	Slack     LogSlack `mapstructure:"slack" yaml:"slack" json:"slack"`
	Smtp      LogSmtp  `mapstructure:"smtp" yaml:"smtp" json:"smtp"`
	Conn      LogConn  `mapstructure:"conn" yaml:"conn" json:"conn"`
}
type LogFile struct {
	Prefix string `mapstructure:"prefix" yaml:"prefix" json:"prefix"`
	Suffix string `mapstructure:"suffix" yaml:"suffix" json:"suffix"`
	Daily  bool   `mapstructure:"daily" yaml:"daily" json:"daily"`
	Level  int    `mapstructure:"level" yaml:"level" json:"level"`
}
type LogConn struct {
	ReconnectOnMsg bool   `mapstructure:"reconnectOnMsg" yaml:"enreconnectOnMsgv" json:"reconnectOnMsg"`
	Reconnect      bool   `mapstructure:"reconnect" yaml:"reconnect" json:"reconnect"`
	Net            string `mapstructure:"net" yaml:"net" json:"net"`
	Addr           string `mapstructure:"addr" yaml:"addr" json:"addr"`
	Level          int    `mapstructure:"level" yaml:"level" json:"level"`
}
type LogSlack struct {
	WebhookURL string `mapstructure:"webhookurl" yaml:"webhookurl" json:"webhookurl"`
	Level      int    `mapstructure:"level" yaml:"level" json:"level"`
}
type LogSmtp struct {
	Username           string   `mapstructure:"username" yaml:"username" json:"username"`
	Password           string   `mapstructure:"password" yaml:"password" json:"password"`
	Host               string   `mapstructure:"host" yaml:"host" json:"host"`
	Subject            string   `mapstructure:"subject" yaml:"subject" json:"subject"`
	FromAddress        string   `mapstructure:"fromAddress" yaml:"fromAddress" json:"fromAddress"`
	RecipientAddresses []string `mapstructure:"sendTos" yaml:"sendTos" json:"sendTos"`
	Level              int      `mapstructure:"level" yaml:"level" json:"level"`
}
