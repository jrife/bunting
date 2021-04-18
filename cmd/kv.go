package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	kvCmd.AddCommand(kvSchedulerCmd)
	kvCmd.AddCommand(kvServerCmd)
	kvCmd.AddCommand(kvShardHostCmd)
}

var kvCmd = &cobra.Command{
	Use:   "kv",
	Short: "Interact with or run buntingkv",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var kvSchedulerCmd = &cobra.Command{
	Use:   "scheduler",
	Short: "Run a scheduler instance",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var kvServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Run an API server instance",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var kvShardHostCmd = &cobra.Command{
	Use:   "processor",
	Short: "Run a processor instance",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
