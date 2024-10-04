package limiter

import (
	"gin-skill/global"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/redis"
	"go.uber.org/zap"
	"strings"
)

// GetKeyRouteWithIP Limitor 的 Key，路由+IP，针对单个路由做限流
func GetKeyRouteWithIP(ctx *gin.Context) string {
	return routeToKeyString(ctx.FullPath()) + ctx.ClientIP()
}

// CheckRate 检测请求是否超额
func CheckRate(ctx *gin.Context, key string, formatted string) (limiter.Context, error) {
	var (
		err     error
		context limiter.Context
		rate    limiter.Rate
		store   limiter.Store
		limit   *limiter.Limiter
	)

	// 实例化依赖的 limiter 包的 limiter.Rate 对象
	if rate, err = limiter.NewRateFromFormatted(formatted); err != nil {
		zap.L().Error("limiter rate error", zap.Error(err))
		return context, err
	}

	// 初始化存储，使用我们程序里共用的 redis.Redis 对象
	if store, err = redis.NewStoreWithOptions(global.App.Redis, limiter.StoreOptions{
		Prefix: global.App.Config.App.AppName + ":limiter",
	}); err != nil {
		zap.L().Error("limiter store error", zap.Error(err))
		return context, err
	}

	// 使用上面的初始化的 limiter.Rate 对象和存储对象
	limit = limiter.New(store, rate)

	// 获取限流的结果
	if ctx.GetBool("limiter-once") {
		// Peek() 取结果，不增加访问次数
		return limit.Peek(ctx, key)
	} else {

		// 确保多个路由组里调用 LimitIP 进行限流时，只增加一次访问次数。
		ctx.Set("limiter-once", true)

		// Get() 取结果且增加访问次数
		return limit.Get(ctx, key)
	}
}

// routeToKeyString 辅助方法，将 URL 中的 / 格式为 -
func routeToKeyString(routeName string) string {
	routeName = strings.ReplaceAll(routeName, "/", "-")
	routeName = strings.ReplaceAll(routeName, ":", "_")
	return routeName
}
