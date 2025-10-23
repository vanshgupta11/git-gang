package submitter

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/vanshgupta11/git-gang/leetcode-auto/internal/api"
	"github.com/vanshgupta11/git-gang/leetcode-auto/internal/config"
	"github.com/vanshgupta11/git-gang/leetcode-auto/internal/models"
)

// Submitter handles the LeetCode submission workflow
type Submitter struct {
	cfg    *config.Config
	client *api.Client
	
	progress        map[string]bool
	doneQuestions   []int
	excludeSet      map[int]bool
	cloudflareCount int
}

// New creates a new Submitter instance
func New(cfg *config.Config) *Submitter {
	return &Submitter{
		cfg:         cfg,
		client:      api.New(cfg.LeetCodeSession, cfg.CSRFToken),
		progress:    make(map[string]bool),
		excludeSet:  make(map[int]bool),
	}
}

// Run executes the main submission workflow
func (s *Submitter) Run() error {
	if err := s.loadProgress(); err != nil {
		return fmt.Errorf("load progress: %w", err)
	}

	solved, err := s.fetchSolvedQuestions()
	if err != nil {
		return fmt.Errorf("fetch solved questions: %w", err)
	}
	s.doneQuestions = solved

	s.buildExcludeSet()

	successCount := 0
	attempts := 0
	for successCount < s.cfg.TargetSuccess && 
		s.cloudflareCount < 4 && 
		attempts < s.cfg.OverallAttemptLimit {
		
		attempts++
		qnum := s.getRandomQuestion()
		if qnum == 0 {
			return fmt.Errorf("no eligible questions found")
		}

		result, err := s.processQuestion(qnum)
		if err != nil {
			fmt.Printf("Process question %d: %v\n", qnum, err)
			continue
		}

		if result == "SUCCESS" {
			successCount++
		}

		time.Sleep(time.Duration(s.cfg.DelayBetweenAttempts) * time.Millisecond)
	}

	s.showSummary(successCount, attempts)
	return nil
}

// Additional methods will be implemented here...