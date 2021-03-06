package cmd

import (
	"fmt"
	"github.com/qipaiyouxi/fkmahjongserver/cache"
	"github.com/qipaiyouxi/fkmahjongserver/cron"
	"github.com/qipaiyouxi/fkmahjongserver/db"
	"github.com/qipaiyouxi/fkmahjongserver/log"
	"github.com/qipaiyouxi/fkmahjongserver/notice"
	"github.com/qipaiyouxi/fkmahjongserver/sale"
	"github.com/qipaiyouxi/fkmahjongserver/util"

	"github.com/spf13/cobra"
)

// saleCmd represents the sale command
var saleCmd = &cobra.Command{
	Use:   "sale",
	Short: "Sale server for fkmahjongserver",
	Long:  `Sale server for fkmahjongserver.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("sale called")
		fmt.Printf("branch: %s\n", branch)
		fmt.Printf("commit: %s\n", commit)
		defer util.Stack()
		log.Config("sale")
		cron.StartForSale("sale")
		db.Start()
		cache.InitAgAccount()
		cache.InitAgAuth()
		cache.InitAgBill()
		notice.StartSaleNotice()
		sale.Start()
	},
}

func init() {
	RootCmd.AddCommand(saleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// saleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// saleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	saleCmd.Flags().StringP("host", "", "0.0.0.0", "The host to listen")
	saleCmd.Flags().Int32P("port", "", 5005, "The port to listen")
}
