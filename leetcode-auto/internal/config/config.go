package config

// Config holds the application configuration
type Config struct {
	LeetCodeSession string
	CSRFToken       string
	MaxQuestions    int
	TargetSuccess   int
	
	// File paths
	ProgressFile string
	SkippedFile  string
	
	// Constants
	MaxRandomAttempts     int
	DelayBetweenAttempts int
	OverallAttemptLimit  int
}

// Default returns a new Config with default values
func Default() *Config {
	return &Config{
		MaxQuestions:         3691,
		TargetSuccess:       1,
		MaxRandomAttempts:   4,
		DelayBetweenAttempts: 1500,
		OverallAttemptLimit:  2000,
		ProgressFile:        "progress.json",
		SkippedFile:         "skipped.log",
	}
}