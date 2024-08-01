package cmd

import (
	"auto-tools/config"
	"auto-tools/logger"
	"flag"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "auto-tools",
	Short: "A CLI tool for automating common tasks",
	Long:  `A CLI tool for automating common tasks such as database migrations, server setup, and data generation.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Default().Printf("run service error: %v\n", err)
		os.Exit(1)
	}
}

func initApp() {
	//初始化配置
	configFile := flag.String("c", "config.json", "config file path")
	flag.Parse()
	config.InitConfig(*configFile)
	//初始化日志
	logfile := config.GetLogFile()
	logger.InitLogger(logfile)
}

func init() {
	cobra.OnInitialize(initApp)
}
