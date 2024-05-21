package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	dbHost string
	dbUser string
	dbPass string
	dbName string
)

var rootCmd = &cobra.Command{
	Use:   "zdevs",
	Short: "ZDEVS CLI",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	viper.AutomaticEnv()
	if dbHost = viper.GetString("DB_HOST"); dbHost == "" {
		dbHost = "localhost:5432"
	}
	if dbUser = viper.GetString("DB_USER"); dbUser == "" {
		dbUser = "root"
	}
	if dbPass = viper.GetString("DB_PASS"); dbPass == "" {
		dbPass = "pass"
	}
	if dbName = viper.GetString("DB_NAME"); dbName == "" {
		dbName = "zdevs-db"
	}
}
