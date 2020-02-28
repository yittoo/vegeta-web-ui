package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// DefaultCommands is a type that defines commands that should be run depending on prefered javascript bundler
type DefaultCommands struct {
	Build     string
	Start     string
	Serve     string
	Install   string
	Directory string
}

func setupBuildCommands(isYarn bool, pwd string) DefaultCommands {
	var buildCommands DefaultCommands
	if isYarn {
		buildCommands = DefaultCommands{
			Build:     "build",
			Start:     "start",
			Serve:     "serve",
			Install:   "",
			Directory: pwd,
		}
	} else {
		buildCommands = DefaultCommands{
			Build:     "run build",
			Start:     "start",
			Serve:     "run serve",
			Install:   "install",
			Directory: "",
		}
	}
	return buildCommands
}

func startReactApp() {
	os.Chdir("client")
	defer os.Chdir("..") // after this function call is over return the currentWorkingDirectory back to root
	isYarn := true

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error geting current working directory")
	}

	// absolute path to yarn executible, if it exists allow process to continue
	p, err := exec.LookPath("yarn")
	if err != nil {
		// absolute path to npm executible, if it exists allow process to continue
		p, err = exec.LookPath("npm")
		if err != nil {
			log.Fatal("You need to have either yarn or npm installed. If both are installed yarn will be preferred over. Make sure they are available under PATH environment variable")
		}
		isYarn = false
	}
	fmt.Printf("\n[+] Found Javascript package bundler at: %v\n", p)

	exeName := "yarn"
	if !isYarn {
		exeName = "npm"
	}
	buildCommands := setupBuildCommands(isYarn, pwd)
	// fmt.Printf("\n%v\n", buildCommands.Directory)

	cwdFlag := fmt.Sprintf("--cwd=\"%v\"", buildCommands.Directory)

	fmt.Println("[+] Installing necessary Javascript packages in client folder...")
	cmd := exec.Command(exeName, buildCommands.Install, cwdFlag)
	fmt.Printf("    %v\n", cmd)
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error installing and checking Javascript packages:\n%v", err)
	}
	fmt.Println("[+] Installing packages completed.")

	fmt.Println("[+] Building necessary Javascript packages in client folder...")
	cmd = exec.Command(exeName, buildCommands.Build, cwdFlag)
	fmt.Printf("    %v\n", cmd)
	err = cmd.Run()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error building Javascript packages:\n%v\nExiting...", err))
	}
	fmt.Println("[+] Building packages completed.")
}
