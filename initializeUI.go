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
	err := os.Chdir("client")
	if err != nil {
		fmt.Printf("\n[+] Client folder not found, skipping javascript build process. If you are running through Docker, this is perfectly fine as client will build in the next container.\n")
		return
	}
	defer os.Chdir("..") // after this function call is over return the currentWorkingDirectory back to root

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error geting current working directory")
	}

	exeName, isYarn, err := findJavascriptBundler()
	must(err)
	buildCommands := setupBuildCommands(isYarn, pwd)

	err = installJavascriptPackages(exeName, buildCommands.Install)
	must(err)
	err = buildJavascriptPackages(exeName, buildCommands.Build)
	must(err)
}

func findJavascriptBundler() (string, bool, error) {
	isYarn := true

	// absolute path to yarn executible, if it doesn't exist check for npm
	p, err := exec.LookPath("yarn")
	if err != nil {
		// absolute path to npx executible, if it exists allow process to continue
		p, err = exec.LookPath("npx")
		if err != nil {
			return "", isYarn, fmt.Errorf("You need to have either yarn or npx(comes with npm) installed. If both are installed yarn will be preferred over. Make sure they are available under PATH environment variable")
		}
		isYarn = false
	}
	fmt.Printf("\n[+] Found Javascript package bundler at: %v\n", p)

	exeName := "yarn"
	if !isYarn {
		exeName = "npm"
	}
	return exeName, isYarn, nil
}

func installJavascriptPackages(exeName string, flags []string) error {
	fmt.Println("[+] Installing necessary Javascript packages in client folder...")
	cmd := exec.Command(exeName, flags...)
	fmt.Printf("    %v\n", cmd)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Error installing and checking Javascript packages:\n%v", err)
	}
	fmt.Println("[+] Installing packages completed.")
	return nil
}

func buildJavascriptPackages(exeName string, flags []string) error {
	fmt.Println("[+] Building necessary Javascript packages in client folder...")
	cmd := exec.Command(exeName, flags...)
	fmt.Printf("    %v\n", cmd)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error building Javascript packages:\n%v\nExiting", err)
	}
	fmt.Println("[+] Building packages completed.")
	return nil
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
