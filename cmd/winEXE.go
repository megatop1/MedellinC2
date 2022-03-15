/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (

	//"github.com/manifoldco/promptui"
	"log"

	"github.com/megatop1/MedellinC2/data"
	"github.com/spf13/cobra"
	//"os"
	//"errors"
)

// winEXECmd represents the winEXE command
var winEXECmd = &cobra.Command{
	Use:   "winEXE",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("winEXE called")
		createNewLauncher()
	},
}

func init() {
	launcherCmd.AddCommand(winEXECmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// winEXECmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// winEXECmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createNewLauncher() { //function to construct our launcher
	namePromptContent := promptContent{ //prompt user to enter a name for the listener
		"Please provide a launcher name",
		"What would you like to name your launcher?: ",
	}
	name := promptGetInput(namePromptContent) //capture the name as an input from the user

	listenerPromptContent := promptContent{ //prompt user to enter a name for the listener
		"Please enter the listener for the launcher",
		"What listener would you like your launcher to listen for connections over?: ",
	}
	listener := promptGetInput(listenerPromptContent) //capture the name as an input from the user

	data.InsertLauncher(name, listener)

	log.Println("Inserted launcher successfully")
}
