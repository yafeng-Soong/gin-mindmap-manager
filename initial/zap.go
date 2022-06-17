package initial

import (
	"fmt"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/yafeng-Soong/gin-mindmap-manager/global"
	"github.com/yafeng-Soong/gin-mindmap-manager/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitZap() {
	// 创建log文件夹
	if ok, _ := utils.PathExists(global.CONFIG.Zap.Dir); !ok {
		_ = os.Mkdir(global.CONFIG.Zap.Dir, os.ModePerm)
	}
	// 只接收debug信息
	debugLevel := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})
	// 接收info和warn的log信息
	infoLevel := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.InfoLevel && lev <= zap.WarnLevel
	})
	// 只接收error以上的log信息
	errorLevel := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})
	level := getLevel(global.CONFIG.Zap.Level)
	cores := []zapcore.Core{}
	// debug级别会输出debug、info、error三个log文件
	// info级别输出info、error两个log文件
	// error及以上级别只输出error.log
	cores = append(cores, newCore("error", errorLevel))
	if level <= zapcore.WarnLevel {
		cores = append(cores, newCore("info", infoLevel))
	}
	if level <= zapcore.DebugLevel {
		cores = append(cores, newCore("debug", debugLevel))
	}
	core := zapcore.NewTee(cores...)
	global.LOG = zap.New(core, zap.AddCaller())
	global.SUGAR_LOG = global.LOG.Sugar()
}

func newCore(level string, enb zapcore.LevelEnabler) zapcore.Core {
	writer := getLogWriter(fmt.Sprintf("./%s/%s.log", global.CONFIG.Zap.Dir, level))
	return zapcore.NewCore(getEncoder(), writer, enb)
}

// 编码器
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = getTimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 大写编码
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 自定义日志输出时间格式
func getTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("[gin-mindmap-manager] " + "2006/01/02 - 15:04:05"))
}

// 输出位置
func getLogWriter(filename string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename, // 日志文件名
		MaxSize:    20,       // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 100,      // 保留旧文件的最大个数
		// MaxAge:     30,       // 保留旧文件的最大天数（不删除）
		Compress: false, // 是否压缩/归档旧文件
	}
	// 开发环境只在控制台输出
	if global.CONFIG.Server.Env == "development" {
		return zapcore.AddSync(os.Stdout)
	}
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
}

func getLevel(level string) (logLevel zapcore.Level) {
	switch level {
	case "debug":
		logLevel = zap.DebugLevel
	case "info":
		logLevel = zap.InfoLevel
	case "warn":
		logLevel = zap.WarnLevel
	case "error":
		logLevel = zap.ErrorLevel
	case "panic":
		logLevel = zap.PanicLevel
	case "fatal":
		logLevel = zap.FatalLevel
	default:
		logLevel = zap.InfoLevel
	}
	return
}
