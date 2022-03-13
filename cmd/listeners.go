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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listenersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listenersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
