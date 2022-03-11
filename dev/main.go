/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/megatop1/MedellinC2/dev/cmd"
	"github.com/megatop1/MedellinC2/dev/data" //custom data package
)

func main() {
	//Open the database connection
	data.OpenDatabase()

	cmd.Execute()
}
