/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// listenersCmd represents the listeners command
var listenersCmd = &cobra.Command{
	Use:   "listeners",
	Short: "Please Configure your listener",
	Long: `Listeners listen for connections from agents:

Please specify the [Listener Name] [Port] [Protocol]`,
}

func init() {
	rootCmd.AddCommand(listenersCmd)

}
