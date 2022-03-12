/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a listener",
	Long: `A listener listens for connections from our agents.

Listeners are responsible for listeneing for the callbacks from our agents. `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

type promptContent struct {
	errorMsg string
	label    string
}

func init() {
	listenersCmd.AddCommand(createCmd)

}
