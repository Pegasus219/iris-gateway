package middlewares

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/context"
	"iris-gateway/common"
)

func PanicHandle() context.Handler {
	return func(ctx context.Context) {
		defer func() {
			if err := recover(); err != nil {
				if ctx.IsStopped() {
					return
				}
				golog.Error(err)
				rsp := common.Error(999, "网络错误")
				ctx.JSON(rsp)
			}
		}()

		ctx.Next()
	}
}
