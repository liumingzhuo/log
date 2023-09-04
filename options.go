package log

import "go.uber.org/zap/zapcore"

const (
	flagLevel             = "log.Level"
	flagDisableCaller     = "log.disable-caller"
	flagDisableStacktrace = "log.disable-stacktrace"
	flagFormat            = "log.format"
	flagEnableColor       = "log.enable-color"
	flagOutputPaths       = "log.output-paths"
	flagErrorOutputPaths  = "log.error-output-path"
	flagDevelopment       = "log.development"
	flagName              = "log.name"

	consoleFormat = "console"
	jsonFormat    = "json"
)

// Options log日志的一些配置
type Options struct {
	OutputPaths       []string `json:"output-paths" mapstructure:"output-paths"`
	ErrorOutputPaths  []string `json:"error-output-paths" mapstructure:"error-output-paths"`
	Level             string   `json:"level"mapstructure:"level"`
	Format            string   `json:"format"mapstructure:"format"`
	DisableStacktrace bool     `json:"disable-stacktrace"mapstructure:"disable-stacktrace"`
	DisableCaller     bool     `json:"disable-caller" mapstructure:"disable-caller"`
	EnableColor       bool     `json:"enable-color"mapstructure:"enable-color"`
	Development       bool     `json:"development"mapstructure:"development"`
	Name              string   `json:"name"mapstructure:"name"`
}

func NewOptions() *Options {
	return &Options{
		Level:             zapcore.InfoLevel.String(),
		Format:            consoleFormat,
		DisableStacktrace: false,
		DisableCaller:     false,
		EnableColor:       false,
		Development:       false,
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
	}
}
