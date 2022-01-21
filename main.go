package main

import (
	"log"
	"os"

	"github.com/actions-go/toolkit/core"
)

func runMain() {
	region := os.Getenv("AWS_REGION")
	branchName := os.Getenv("BRANCH_NAME")

	macBuildBranches := []string{"staging", "production", "labs"}

	IS_CROSS_REGION := "false"
	IS_MAC_BUILD_BRANCH := "false"

	if region == "" || branchName == "" {
		log.Fatal("AWS_REGION and BRANCH_NAME must be set")
	}

	if branchName == "production" && region != "us-east-1" {
		IS_CROSS_REGION = "true"
		IS_MAC_BUILD_BRANCH = "true"
	}

	if region == "us-east-1" {
		for _, branch := range macBuildBranches {
			if branch == branchName {
				IS_MAC_BUILD_BRANCH = "true"
				break
			}
		}
	}

	core.SetOutput("IS_CROSS_REGION", IS_CROSS_REGION)
	core.SetOutput("IS_MAC_BUILD_BRANCH", IS_MAC_BUILD_BRANCH)
}

func main() {
	runMain()
}
