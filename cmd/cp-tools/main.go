package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	commands "github.com/ministryofjustice/cloud-platform-moj-cp/pkg/commands"
)

func main() {
	cmds := &cobra.Command{
		Use:   "moj-cp",
		Short: "Multi-purpose CLI from the Cloud Platform team",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	commands.AddCommands(cmds)

	if err := cmds.Execute(); err != nil {
		fmt.Printf("Error during command execution: %v", err)
		os.Exit(0)
	}
}
