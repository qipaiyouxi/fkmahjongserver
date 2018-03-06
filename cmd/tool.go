package cmd

import (
	"fmt"
	"github.com/qipaiyouxi/fkmahjongserver/cache"
	"github.com/qipaiyouxi/fkmahjongserver/cron"
	"github.com/qipaiyouxi/fkmahjongserver/db"
	"github.com/qipaiyouxi/fkmahjongserver/log"
	"github.com/qipaiyouxi/fkmahjongserver/notice"
	"github.com/qipaiyouxi/fkmahjongserver/tool"
	"github.com/qipaiyouxi/fkmahjongserver/util"

	"github.com/spf13/cobra"
)

// toolCmd represents the tool command
var toolCmd = &cobra.Command{
	Use:   "tool",
	Short: "Tool server for fkmahjongserver",
	Long:  `Tool server for fkmahjongserver.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("tool called")
		fmt.Printf("branch: %s\n", branch)
		fmt.Printf("commit: %s\n", commit)
		defer util.Stack()
		log.Config("tool")
		cron.Start("tool")
		db.Start()
		cache.InitAccount()
		cache.InitCost()
		cache.InitGame()
		cache.InitModule()
		cache.InitPermission()
		cache.InitRole()
		cache.InitShop()
		notice.StartToolNotice()
		tool.Start()
	},
}

func init() {
	RootCmd.AddCommand(toolCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// toolCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// toolCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	toolCmd.Flags().StringP("host", "", "0.0.0.0", "The host to listen")
	toolCmd.Flags().Int32P("port", "", 5003, "The port to listen")
}
