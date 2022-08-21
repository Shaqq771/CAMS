package logger

import (
	"backend-nabati/infrastructure/logger/logrus"
	"backend-nabati/infrastructure/logger/zap"
	"backend-nabati/infrastructure/shared/constant"
)

var useLog string

func InitializeLogger(log string) {
	if log == constant.LOGRUS {
		logrus.InitializeLogrusLogger()
		useLog = constant.LOGRUS
	} else if log == constant.ZAP {
		zap.InitializeZapLogger()
		useLog = constant.ZAP
	} else {
		logrus.InitializeLogrusLogger()
		useLog = constant.LOGRUS
	}
}

func LogInfo(logtype, message string) {
	if useLog == constant.LOGRUS {
		logrus.LogInfo(useLog, logtype, message)
	} else if useLog == constant.ZAP {
		zap.LogInfo(useLog, logtype, message)
	}
}

func LogInfoWithData(data interface{}, logtype, message string) {
	if useLog == constant.LOGRUS {
		logrus.LogInfoWithData(useLog, data, logtype, message)
	} else if useLog == constant.ZAP {
		zap.LogInfoWithData(useLog, data, logtype, message)
	}
}

func LogError(logtype, errtype, message string) {
	if useLog == constant.LOGRUS {
		logrus.LogError(useLog, logtype, errtype, message)
	} else if useLog == constant.ZAP {
		zap.LogInfoWithData(useLog, logtype, errtype, message)
	}
}
