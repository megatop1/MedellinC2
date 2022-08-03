/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// agentCmd represents the agent command
var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Interact with agents",
	Long: `Agents are compromised machines that have an established callback to the C2 server. Use the various agents commands to interact with your agents, and run commands. 
	You are in charge of your "houses" in your city of medellin!
	`,
}

func init() {
	rootCmd.AddCommand(agentCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// launcherCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// launcherCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
