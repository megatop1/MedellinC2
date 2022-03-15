/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// launcherCmd represents the launcher command
var launcherCmd = &cobra.Command{
	Use:   "launcher",
	Short: "Generate a Launcher",
	Long: `Launchers are used to generate payloads to be executed
	
on your target machines. Run the launcher on your targt machine
to establish a connection back to the C2 server. Your compromised machine 
will then become an agent`,
}

func init() {
	rootCmd.AddCommand(launcherCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// launcherCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// launcherCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
