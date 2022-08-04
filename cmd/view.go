/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/megatop1/MedellinC2/data"
	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View a list of active agents",
	Long:  `View a list of agents that are still currently active and connected to the C2 server`,
	Run: func(cmd *cobra.Command, args []string) {
		viewAliveAgents()
	},
}

func viewAliveAgents() {
	data.GetAliveAgents()
}

func init() {
	agentCmd.AddCommand(viewCmd)
}
