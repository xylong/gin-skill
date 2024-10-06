package cmd

import (
	"context"
	"gin-skill/global"
	"gin-skill/pkg/console"
	"github.com/spf13/cobra"
	"time"
)

var CmdPlay = &cobra.Command{
	Use:   "play",
	Short: "Likes the Go Playground, but running at our application context",
	Run:   runPlay,
}

// 调试完成后请记得清除测试代码
func runPlay(cmd *cobra.Command, args []string) {
	var ctx = context.Background()

	// 存进去 redis 中
	global.App.Redis.Set(ctx, "hello", "hi from redis", 10*time.Second)
	// 从 redis 里取出
	console.Success(global.App.Redis.Get(ctx, "hello").String())
}
