package main

import (
	"fmt"
	"gin-skill/app/cmd"
	"gin-skill/app/dao"
	"gin-skill/bootstrap"
	"gin-skill/global"
	"gin-skill/pkg/console"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	defer bootstrap.CloseDB()

	var rootCmd = &cobra.Command{
		Use:   "Gin",
		Short: "A simple forum project",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,

		// rootCmd 的所有子命令都会执行以下代码
		PersistentPreRun: func(command *cobra.Command, args []string) {

			// 配置初始化，依赖命令行 --env 参数
			bootstrap.InitConfig()

			// 初始化 Logger
			bootstrap.InitLog()

			// 初始化数据库
			bootstrap.Migrate()
			dao.SetDefault(bootstrap.DB())

			// 初始化 Redis
			global.App.Redis = bootstrap.InitRedis()

			// 初始化验证器
			bootstrap.InitValidator()
		},
	}

	// 注册子命令
	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdKey,
		cmd.CmdPlay,
		cmd.CmdMake,
	)

	// 配置默认运行 Web 服务
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	// 注册全局参数，--env
	cmd.RegisterGlobalFlags(rootCmd)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}
