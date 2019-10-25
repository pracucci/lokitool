package pipeline

import (
	"fmt"

	"github.com/grafana/loki/pkg/logentry/stages"
)

type TestCase struct {
	InputLogs    []TestLog `yaml:"input_logs"`
	ExpectedLogs []TestLog `yaml:"expected_logs"`
}

func (c *TestCase) Run(suiteID string, testID int, pipeline *stages.Pipeline, reporter *Reporter) {
	actualLogs := make([]TestLog, 0, len(c.InputLogs))

	// Process all input logs
	for _, inputLog := range c.InputLogs {
		labels := inputLog.Labels.Clone()
		extracted := map[string]interface{}{}
		timestamp := inputLog.Time()
		entry := inputLog.Entry

		pipeline.Process(labels, extracted, &timestamp, &entry)

		actualLogs = append(actualLogs, TestLog{
			Timestamp: timestamp,
			Labels:    labels,
			Entry:     entry,
		})
	}

	// Ensure the number of logs match
	if len(c.ExpectedLogs) != len(actualLogs) {
		reporter.RecordTestError(suiteID, testID, fmt.Errorf("Expected %d logs while got %d in output", len(c.ExpectedLogs), len(actualLogs)))
		return
	}

	// Compare logs
	for i, actualLog := range actualLogs {
		entryID := i + 1

		if actualLog.Equal(&c.ExpectedLogs[i]) {
			reporter.RecordTestSuccess(suiteID, testID, entryID, c.ExpectedLogs[i], actualLog)
		} else {
			reporter.RecordTestFailure(suiteID, testID, entryID, c.ExpectedLogs[i], actualLog)
		}
	}
}
