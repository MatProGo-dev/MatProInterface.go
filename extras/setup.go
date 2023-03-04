/*
setup.go
Description:
	This script is designed to help setup the gurobi.go dependency which is needed to run goop.
*/

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	versionTool "github.com/hashicorp/go-version"
)

// Type Definitions
type SetupInfo struct {
	InitialDirectory  string
	GurobiGoDirectory string
}

// Functions

/*
GetDefaultSetupInfo
Description:

	Defines the default setup flags for the setup script.
*/
func GetDefaultSetupInfo() (SetupInfo, error) {
	// Create Default Struct
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return SetupInfo{}, fmt.Errorf("There was an error collecting the user's home directory: %v", err)
	}

	sf0 := SetupInfo{
		InitialDirectory:  fmt.Sprintf("%v/go/pkg/mod/github.com/MatProGo-dev/", homeDir),
		GurobiGoDirectory: fmt.Sprintf("%v/go/pkg/mod/github.com/MatProGo-dev/", homeDir),
	}

	// Get current working directory
	tempPwd, err := os.Getwd()
	if err != nil {
		return sf0, err
	}
	sf0.InitialDirectory = tempPwd
	log.Printf("The current directory is \"%v\"", tempPwd)

	return sf0, nil
	// gurobiVersionList, err := StringsToGurobiGoVersionInfoList(gurobiDirectories)
	// if err != nil {
	// 	return mlf, err
	// }

	// highestVersion, err := FindHighestVersion(gurobiVersionList)
	// if err != nil {
	// 	return mlf, err
	// }

	// // Write the highest version's directory into the GurobiHome variable
	// mlf.GurobiHome = fmt.Sprintf("/Library/gurobi%v%v%v/mac64", highestVersion.MajorVersion, highestVersion.MinorVersion, highestVersion.TertiaryVersion)

	// return mlf, nil

}

/*
SafeGetGurobiGo
Description:

	Run the safe go get which is necessary for this library which uses cgo.
*/
func SafeGetGurobiGo() error {
	err := exec.Command("go", "get", "-d", "github.com/MatProGo-dev/Gurobi.go/gurobi").Run()
	if err != nil {
		return fmt.Errorf("There was an issue running go get: %v", err)
	}

	// Return nil
	return nil

}

/*
SetupMostRecentGurobigo
Description:

	Runs the gen.go file from the directory where the most recent gurobi.go version has been installed.
*/
func SetupMostRecentGurobigo(siIn *SetupInfo) error {

	// Search through the pkg Library for all instances of gurobi.go
	libraryContents, err := os.ReadDir(siIn.GurobiGoDirectory)
	if err != nil {
		return err
	}
	gurobiGoDirectories := []string{}
	for _, content := range libraryContents {
		if content.IsDir() && strings.Contains(content.Name(), "gurobi.go") {
			fmt.Println(content.Name())
			gurobiGoDirectories = append(gurobiGoDirectories, content.Name())
		}
	}

	if len(gurobiGoDirectories) == 0 {
		return fmt.Errorf("No gurobi.go directories were found at the specified directory: %v", siIn.GurobiGoDirectory)
	}

	log.Printf("Identified %v versions of gurobi.go.\n", len(gurobiGoDirectories))

	// Convert Directories into gurobi.go Version Info
	gurobiGoVersionList := []*versionTool.Version{}
	for _, directoryName := range gurobiGoDirectories {
		tempVersionPointer, err := versionTool.NewVersion(directoryName[len("gurobi.go@v"):])
		if err != nil {
			if err != nil {
				return err
			}
		}
		gurobiGoVersionList = append(gurobiGoVersionList, tempVersionPointer)
	}

	// fmt.Println(gurobiGoVersionList)
	// fmt.Println(gurobiGoVersionList[0].GreaterThan(gurobiGoVersionList[3]))

	// Get highest version
	highestVersionIndex := 0
	for versionIndex, tempVersion := range gurobiGoVersionList {
		if tempVersion.GreaterThan(gurobiGoVersionList[highestVersionIndex]) {
			highestVersionIndex = versionIndex
		}
	}

	log.Printf("Highest Version: %v\n", gurobiGoVersionList[highestVersionIndex])

	// Change directory to the one with the highest version
	targetDirectory := fmt.Sprintf("%v%v", siIn.GurobiGoDirectory, gurobiGoDirectories[highestVersionIndex])
	err = os.Chdir(targetDirectory)
	if err != nil {
		return fmt.Errorf("There was an issue changing the directory: %v", err)
	}
	defer os.Chdir(siIn.InitialDirectory) // Change directory back to the original one

	// Run gen.go
	err = exec.Command("go", "generate").Run()
	if err != nil {
		return fmt.Errorf("There was an issue running \"go generate\" in the new directory (was sudo used?): %v", err)
	}

	log.Printf("Completed the running of go generate from %v.", targetDirectory)

	// If all Searches have passed, then return no errors.
	return nil
}

/*
SetUpLog
Description:

	Creates a log file in the new directory "logs"
	1. Checks to see if log file already exists. If it exists, then the function deletes the old log file.
	2. Creates a new log file.
	3. Sets the log module to point to new log file.
	4. Adds initial message to log file.
*/
func SetUpLog() error {
	// Constants
	logFileName := "extras/setup_log.txt"

	// Check to see if logFile exists
	_, err := os.Stat(logFileName)
	if os.IsNotExist(err) {
		//Do Nothing. The later lines will create the file.
	} else {
		//Delete the old file.
		err = os.Remove(logFileName)
		if err != nil {
			return err
		}
	}

	// Create Logging file
	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		// log.Fatal(err)
		return err
	}
	log.SetOutput(io.MultiWriter(file, os.Stdout))

	// Create initial file
	log.Println("Log file created.")

	return nil

}

func main() {

	// Setup Log
	err := SetUpLog()
	if err != nil {
		fmt.Printf("There was an issue setting up the log file: %v", err)
	}

	SetupInfo, err := GetDefaultSetupInfo()
	if err != nil {
		log.Printf("There was an error collecting default setup info: %v", err)
	}

	// Call go get for gurobi.go
	err = SafeGetGurobiGo()
	if err != nil {
		log.Printf("There was an issue running SafeGetGurobiGo(): %v", err)
		os.Exit(1)
	}

	// Setup the most recently installed gurobi.go
	err = SetupMostRecentGurobigo(&SetupInfo)
	if err != nil {
		log.Printf("There was an issue running SetupMostREcentGurobigo: %v", err)
	}

	fmt.Println(SetupInfo)

	// Next, parse the arguments to make_lib and assign values to the mlf appropriately
	//sf, err = ParseMakeLibArguments(sf)
}
