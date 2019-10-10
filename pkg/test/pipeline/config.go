package pipeline

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/prometheus/common/model"
	yaml "gopkg.in/yaml.v2"
)

type TestSuiteConfig struct {
	Filepath string       `yaml:"-"`
	JobName  string       `yaml:"job_name"`
	Tests    []TestConfig `yaml:"tests"`
}

type TestConfig struct {
	InputLogs    []string          `yaml:"input_logs"`
	ExpectedLogs []TestExpectedLog `yaml:"expected_logs"`
}

type TestExpectedLog struct {
	Timestamp time.Time      `yaml:"timestamp"`
	Entry     string         `yaml:"entry"`
	Labels    model.LabelSet `yaml:"labels"`
}

func LoadTestSuiteConfig(file string) (*TestSuiteConfig, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	config := TestSuiteConfig{
		Filepath: file,
	}

	if err = yaml.Unmarshal(content, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (l *TestExpectedLog) Equal(other *TestExpectedLog) bool {
	if other == nil {
		return false
	}

	if !l.Timestamp.Equal(other.Timestamp) {
		return false
	}

	if l.Entry != other.Entry {
		return false
	}

	if !l.Labels.Equal(other.Labels) {
		return false
	}

	return true
}

func (l *TestExpectedLog) String(indent int) string {
	prefix := strings.Repeat(" ", indent)
	lines := []string{
		prefix + fmt.Sprintf("Timestamp: %s", l.Timestamp.Format(time.RFC3339Nano)),
		prefix + fmt.Sprintf("Entry:     %s", l.Entry),
		prefix + fmt.Sprintf("Labels:    %s", l.Labels.String()),
	}

	return strings.Join(lines, "\n")
}
