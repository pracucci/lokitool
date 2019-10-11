package pipeline

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/grafana/loki/pkg/logentry/stages"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type TestSuiteConfig struct {
	Filepath  string     `yaml:"-"`
	JobName   string     `yaml:"job_name"`
	TestCases []TestCase `yaml:"tests"`
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

type TestSuite struct {
	config *TestSuiteConfig
}

func NewTestSuite(testFile string) (*TestSuite, error) {
	config, err := LoadTestSuiteConfig(testFile)
	if err != nil {
		return nil, err
	}

	return &TestSuite{
		config: config,
	}, nil
}

func NewTestSuites(testFiles []string) ([]*TestSuite, error) {
	testSuites := make([]*TestSuite, 0, len(testFiles))

	for _, testFile := range testFiles {
		testSuite, err := NewTestSuite(testFile)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("Unable to load test file %s", testFile))
		}

		testSuites = append(testSuites, testSuite)
	}

	return testSuites, nil
}

func (s *TestSuite) Run(reporter *Reporter, pipelines map[string]*stages.Pipeline) {
	suiteID := filepath.Base(s.config.Filepath)

	// Get the pipeline for the actual job
	pipeline, ok := pipelines[s.config.JobName]
	if !ok {
		reporter.RecordSuiteError(suiteID, fmt.Errorf("No pipeline found for the job '%s'", s.config.JobName))
		return
	}

	// Run test cases
	for i, testCase := range s.config.TestCases {
		testID := i + 1

		testCase.Run(suiteID, testID, pipeline, reporter)
	}
}
