package cmd

import (
	"fmt"
	"github.com/aminzdev/auth"
	"github.com/aminzdev/authserver"
	"github.com/aminzdev/db"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var (
	grpcHost    string
	grpcWebHost string
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "starts auth server",
	Run: func(cmd *cobra.Command, args []string) {
		database, err := db.NewDB(
			&db.Postgres{
				Host:     dbHost,
				User:     dbUser,
				Pass:     dbPass,
				DBName:   dbName,
				SSL:      "disable",
				TimeZone: "Asia/Tehran",
			},
			&auth.User{},
			&auth.Session{},
		)
		cobra.CheckErr(err)
		authserver.RunAuthServer(grpcHost, grpcWebHost, database)
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig
		fmt.Println("\nAuth Server Terminated.")
	},
}

func init() {
	serverCmd.AddCommand(authCmd)
	authCmd.Flags().StringVarP(&grpcHost, "grpc", "g", ":50051", "gRPC Host Address")
	authCmd.Flags().StringVarP(&grpcWebHost, "web", "w", ":80", "gRPC-Web Host Address")
}
