package main

import (
	"github.com/go-kit/kit/log"
	"time"
)

// 构造一个函数，然后返回函数类型
func loggingMiddleware (logger log.Logger) ServiceMiddleware {
	return func(next StringService) StringService {
		return logmw{logger, next}
	}
}

// 日志中间件，采用静态代理模式
type logmw struct {
	logger log.Logger
	next StringService
}

func (mw logmw) Uppercase(s string) (output string, err error) {
	// 延迟返回，日志记录返回值和错误信息
	defer func(begin time.Time) {
		mw.logger.Log("method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin))
	}(time.Now())

	output, err = mw.next.Uppercase(s)
	return
}

func(mw logmw) Count(s string) (output int){
	defer func(begin time.Time) {
		mw.logger.Log("method", "count",
			"input", s,
			"output", output,
			"took", time.Since(begin))
	}(time.Now())

	output = mw.next.Count(s)
	return
}