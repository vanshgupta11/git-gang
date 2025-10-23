package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/vanshgupta11/git-gang/leetcode-auto/internal/config"
	"github.com/vanshgupta11/git-gang/leetcode-auto/internal/submitter"
)

func main() {
	cfg := &config.Config{}

	// Parse command-line flags
	flag.StringVar(&cfg.LeetCodeSession, "session", "", "LeetCode session token")
	flag.StringVar(&cfg.CSRFToken, "csrf", "", "LeetCode CSRF token")
	flag.IntVar(&cfg.MaxQuestions, "max-questions", 3691, "Maximum question number to consider")
	flag.IntVar(&cfg.TargetSuccess, "target", 1, "Target number of successful submissions")
	flag.Parse()

	// Validate required flags
	if cfg.LeetCodeSession == "" || cfg.CSRFToken == "" {
		fmt.Println("Error: Both session and csrf tokens are required")
		flag.Usage()
		os.Exit(1)
	}

	// Create submitter and run
	s := submitter.New(cfg)
	if err := s.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}