package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	kvCmd.AddCommand(kvQueryCmd)
	kvCmd.AddCommand(kvCreateCmd)
	kvCmd.AddCommand(kvUpdateCmd)
	kvCmd.AddCommand(kvDeleteCmd)
	kvCmd.AddCommand(kvServerCmd)
	kvCmd.AddCommand(kvProcessorCmd)
	kvCmd.AddCommand(kvShardHostCmd)
}

var kvCmd = &cobra.Command{
	Use:   "kv",
	Short: "Interact with or run buntingkv",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var kvQueryCmd = &cobra.Command{
	Use:   "query",
	Short: "Submit a query request to the server",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var kvCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Submit a create request to the server",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var kvUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Submit a update request to the server",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var kvDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Submit a delete request to the server",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// Maybe there needs to be a designated group of "primaries"
// with well known addresses that maintain the cluster state.
// These could just be specially marked shard hosts.
//
// apiserver  -> aware of shard host addresses.
// processor  -> aware of apiserver address(es).
// shard host -> aware of apiserver address(es).
var kvServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Run an API server instance",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var kvProcessorCmd = &cobra.Command{
	Use:   "processor",
	Short: "Run a processor instance",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var kvShardHostCmd = &cobra.Command{
	Use:   "shard-host",
	Short: "Run a shard host instance",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
