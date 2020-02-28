package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// DefaultCommands is a type that defines commands that should be run depending on prefered javascript bundler
type DefaultCommands struct {
	Build   []string
	Start   []string
	Install []string
}

func setupBuildCommands(isYarn bool, pwd string) DefaultCommands {
	var buildCommands DefaultCommands
	if isYarn {
		cwdWrapped := fmt.Sprintf("--cwd=\"%v\"", pwd)
		buildCommands = DefaultCommands{
			Build:   []string{"build", cwdWrapped},
			Start:   []string{"start", cwdWrapped},
			Install: []string{cwdWrapped},
		}
	} else {
		buildCommands = DefaultCommands{
			Build:   []string{"run", "build"},
			Start:   []string{"start"},
			Install: []string{"install"},
		}
	}
	return buildCommands
}

func buildReactApp() {
	os.Chdir("client")
	defer os.Chdir("..") // after this function call is over return the currentWorkingDirectory back to root

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error geting current working directory")
	}

	isYarn := false
	// absolute path to yarn executible, if it doesn't exist check for npm
	p, err := exec.LookPath("yarn")
	if err != nil {
		// absolute path to npx executible, if it exists allow process to continue
		p, err = exec.LookPath("npx")
		if err != nil {
			log.Fatal("You need to have either yarn or npx(comes with npm) installed. If both are installed yarn will be preferred over. Make sure they are available under PATH environment variable")
		}
		isYarn = false
	}
	fmt.Printf("\n[+] Found Javascript package bundler at: %v\n", p)

	exeName := "yarn"
	if !isYarn {
		exeName = "npm"
	}
	buildCommands := setupBuildCommands(isYarn, pwd)

	var cmd *exec.Cmd

	fmt.Println("[+] Installing necessary Javascript packages in client folder...")
	cmd = exec.Command(exeName, buildCommands.Install...)
	fmt.Printf("    %v\n", cmd)
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error installing and checking Javascript packages:\n%v", err)
	}
	fmt.Println("[+] Installing packages completed.")

	fmt.Println("[+] Building necessary Javascript packages in client folder...")
	cmd = exec.Command(exeName, buildCommands.Build...)
	fmt.Printf("    %v\n", cmd)
	err = cmd.Run()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error building Javascript packages:\n%v\nExiting...", err))
	}
	fmt.Println("[+] Building packages completed.")
}
