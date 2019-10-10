package pipeline

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/grafana/loki/pkg/logentry/stages"
	"github.com/pkg/errors"
	"github.com/prometheus/common/model"
)

type TestSuite struct {
	config *TestSuiteConfig
}

func LoadTestSuite(testFile string) (*TestSuite, error) {
	config, err := LoadTestSuiteConfig(testFile)
	if err != nil {
		return nil, err
	}

	return &TestSuite{
		config: config,
	}, nil
}

func LoadTestSuites(testFiles []string) ([]*TestSuite, error) {
	testSuites := make([]*TestSuite, 0, len(testFiles))

	for _, testFile := range testFiles {
		testSuite, err := LoadTestSuite(testFile)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("Unable to load test file %s", testFile))
		}

		testSuites = append(testSuites, testSuite)
	}

	return testSuites, nil
}

func (s *TestSuite) Run(reporter *Reporter, pipelines map[string]*stages.Pipeline) {
	// Get the pipeline for the actual job
	pipeline, ok := pipelines[s.config.JobName]
	if !ok {
		suiteID := filepath.Base(s.config.Filepath)
		reporter.RecordSuiteError(suiteID, fmt.Errorf("No pipeline found for the job '%s'", s.config.JobName))
		return
	}

	// Run test cases
	for i, testConfig := range s.config.Tests {
		s.runTestCase(reporter, i+1, testConfig, pipeline)
	}
}

func (s *TestSuite) runTestCase(reporter *Reporter, testID int, testConfig TestConfig, pipeline *stages.Pipeline) {
	suiteID := filepath.Base(s.config.Filepath)
	actualLogs := make([]TestExpectedLog, 0, len(testConfig.InputLogs))

	// Process all input logs
	for _, inputLog := range testConfig.InputLogs {
		labels := model.LabelSet{}
		extracted := map[string]interface{}{}
		// TODO: allow to mock it
		timestamp := time.Now()
		entry := inputLog

		pipeline.Process(labels, extracted, &timestamp, &entry)

		actualLogs = append(actualLogs, TestExpectedLog{
			Timestamp: timestamp,
			Labels:    labels,
			Entry:     entry,
		})
	}

	// Ensure the number of logs match
	if len(testConfig.ExpectedLogs) != len(actualLogs) {
		reporter.RecordTestError(suiteID, testID, fmt.Errorf("Expected %d logs while got %d in output", len(testConfig.ExpectedLogs), len(actualLogs)))
		return
	}

	// Compare logs
	for i, actualLog := range actualLogs {
		entryID := i + 1

		if actualLog.Equal(&testConfig.ExpectedLogs[i]) {
			reporter.RecordTestSuccess(suiteID, testID, entryID, &testConfig.ExpectedLogs[i], &actualLog)
		} else {
			reporter.RecordTestFailure(suiteID, testID, entryID, &testConfig.ExpectedLogs[i], &actualLog)
		}
	}
}
