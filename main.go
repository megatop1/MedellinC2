/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/megatop1/MedellinC2/cmd"
	"github.com/megatop1/MedellinC2/data" //custom data package
)

func main() {
	//Open the database connection
	data.OpenDatabase()
	cmd.Execute()
	//Keep the server consistently running

}
