package models

// Question represents a LeetCode question with its metadata
type Question struct {
	ID          int    `json:"id"`
	LeetCodeURL string `json:"leetcode_url"`
	WalkCCURL   string `json:"walkcc_url"`
}

// Progress tracks submission progress for questions
type Progress struct {
	Questions map[string]bool `json:"questions"`
}