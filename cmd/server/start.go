package server

import (
	server "github.com/rugggger/go-fiber/pkg/server"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:     "start",
	Aliases: []string{"s"},
	Short:   "Starts a server",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		server.Start(args[0])
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
