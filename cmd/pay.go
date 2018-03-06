package cmd

import (
	"fmt"
	"github.com/qipaiyouxi/fkmahjongserver/cache"
	"github.com/qipaiyouxi/fkmahjongserver/cron"
	"github.com/qipaiyouxi/fkmahjongserver/db"
	"github.com/qipaiyouxi/fkmahjongserver/log"
	"github.com/qipaiyouxi/fkmahjongserver/notice"
	"github.com/qipaiyouxi/fkmahjongserver/pay"
	"github.com/qipaiyouxi/fkmahjongserver/util"

	"github.com/spf13/cobra"
)

// payCmd represents the pay command
var payCmd = &cobra.Command{
	Use:   "pay",
	Short: "Pay server for fkmahjongserver",
	Long:  `Pay server for fkmahjongserver.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("pay called")
		fmt.Printf("branch: %s\n", branch)
		fmt.Printf("commit: %s\n", commit)
		defer util.Stack()
		log.Config("pay")
		cron.Start("pay")
		db.Start()
		cache.InitShop()
		notice.StartPayNotice()
		pay.Start()
	},
}

func init() {
	RootCmd.AddCommand(payCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// payCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// payCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	payCmd.Flags().StringP("host", "", "0.0.0.0", "The host to listen")
	payCmd.Flags().Int32P("port", "", 5004, "The port to listen")
}
