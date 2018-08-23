package main

import (
	"github.com/fatih/color"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/briandowns/spinner"

	"gopkg.in/AlecAivazis/survey.v1"
)

func main() {

	xMap := map[string][]string{
		"linux":     []string{"386", "amd64", "ppc64", "ppc64le", "mips64", "mips64le", "arm64", "ARMv7", "ARMv6", "ARMv5"},
		"darwin":    []string{"amd64", "arm64", "ARMv7", "ARMv6", "ARMv5"},
		"windows":   []string{"386", "amd64"},
		"freebsd":   []string{"amd64", "ARMv7", "ARMv6"},
		"netbsd":    []string{"amd64", "ARMv7", "ARMv6"},
		"dragonfly": []string{"amd64", "ARMv7", "ARMv6"},
		"plan9":     []string{"amd64", "ARMv7", "ARMv6"},
		"solaris":   []string{"amd64", "ARMv7", "ARMv6"},
	}

	platforms := []string{}
	for k := range xMap {
		platforms = append(platforms, k)
	}
	log.Print(platforms)

	q1 := &survey.MultiSelect{
		Message:  "Platforms do you want to compile for:",
		Options:  platforms,
		PageSize: 8,
	}
	pickedPlatforms := []string{}
	survey.AskOne(q1, &pickedPlatforms, nil)

	pickedArqs := [8][]string{}
	for i := range pickedPlatforms {
		a := &survey.MultiSelect{
			Message: "Select architectures for " + pickedPlatforms[i] + ":",
			Options: xMap[pickedPlatforms[i]],
		}
		survey.AskOne(a, &pickedArqs[i], nil)
	}
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond) // Build our new spinner
	s.Color("red", "bold") // Set the spinner color to a bold red
	s.Suffix = "  building" // Append text after the spinner


	s.Start()
	for i := range pickedPlatforms {
		for _, ark := range pickedArqs[i] {
			compile(pickedPlatforms[i], ark)
		}
	}
	s.Stop()
	color.HiGreen("Builds complete!")
}
func compile(goos string, goark string) {
	cmd := exec.Command("go", "build", "-o", fmt.Sprintf("./bin/%s-%s", goos, goark), "./main.go")
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, fmt.Sprintf("GOOS=%s", goos))
	cmd.Env = append(cmd.Env, fmt.Sprintf("GOARCH=%s", goark))

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
