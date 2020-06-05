// +build mage

package main

import (
	"fmt"
	"github.com/magefile/mage/sh"
	"os"
	"os/exec"
	"path"
	"log"
	"time"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
)

const (
	BUILD_DIR string = "./bin"
	BUILD_BINARY string = "name-of-binary.exe"
	DEPLOY_DIR string = "some/dir"
)

var (
	buildVersion string
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

// A build step that requires additional params, or platform specific steps for example
func Build() error {
	mg.Deps(InstallDeps)
	fmt.Println("Building...")

	versionCmd := exec.Command("git", "rev-parse", "--short", "HEAD")
	versionCmdOutput, err := versionCmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	buildVersion = string(versionCmdOutput)
	buildTime := time.Now().UTC().Format("2006-01-02_15:04:05")
	fmt.Println(buildTime)
	//buildTime = "Hallo"

	cmd := exec.Command("go", "build", "-ldflags", "-X main.version=" + buildVersion + " -X main.buildTime=" + buildTime, "-o", path.Join(BUILD_DIR, BUILD_BINARY), ".")
	return cmd.Run()
}

// Copies binary, template and config file to Virtualbox Programme folder
func Deploy() error {
	mg.Deps(Build)
	fmt.Println("Deploying...")
	err := sh.Copy(path.Join(DEPLOY_DIR, BUILD_BINARY), path.Join(BUILD_DIR, BUILD_BINARY))
	err  = sh.Copy(path.Join(DEPLOY_DIR, "template.html"), "template.html")
	err  = sh.Copy(path.Join(DEPLOY_DIR, "config.json"), "config.json")
	return err
}

// Manage your deps, or running package managers.
func InstallDeps() error {
	fmt.Println("Installing Deps...")
	cmd := exec.Command("go", "get", "github.com/stretchr/piglatin")
	return cmd.Run()
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll("bin")
}
