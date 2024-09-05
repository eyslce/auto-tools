package cmd

import (
	"auto-tools/tools"
	"github.com/spf13/cobra"
)

var oacmd = &cobra.Command{
	Use:   "oa",
	Short: "OA command line tool",
	Long:  `OA command line tool`,
	Run: func(cmd *cobra.Command, args []string) {
		tool := new(tools.OaTools)
		tool.Run()
	},
}

var bingcmd = &cobra.Command{
	Use:   "bing",
	Short: "Bing command line tool",
	Long:  `Bing command line tool`,
	Run: func(cmd *cobra.Command, args []string) {
		tool := new(tools.BingTools)
		if len(args) > 0 {
			tool.RunE(true)
			return
		}
		tool.RunE(false)
	},
}
