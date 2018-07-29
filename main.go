package main

import (
	"log"

	"gopkg.in/AlecAivazis/survey.v1"
)

func main() {

	xMap := map[string][]string{
		"linux":     []string{"386", "amd64", "ppc64", "ppc64le", "mips64", "mips64le", "ARM64", "ARMv7", "ARMv6", "ARMv5"},
		"darwin":    []string{"amd64", "ARM64", "ARMv7", "ARMv6", "ARMv5"},
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
		survey.AskOne(a, pickedArqs[i], nil)
	}
	log.Print(pickedPlatforms)
	log.Print(pickedArqs)
}
