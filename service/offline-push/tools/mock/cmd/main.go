package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/txchat/dtalk/service/offline-push/tools/mock/cmd/push"
)

var rootCmd = &cobra.Command{
	Use:     "tools",
	Short:   "offline push mock tools",
	Example: "  tools auth -d <data>\n",
}

func init() {
	rootCmd.AddCommand(push.SinglePushCmd)
}

// Execute executes the root command and its subcommands.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
