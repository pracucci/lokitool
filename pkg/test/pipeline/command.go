package pipeline

import (
	"fmt"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/grafana/loki/pkg/cfg"
	"github.com/grafana/loki/pkg/logentry/stages"
	"github.com/grafana/loki/pkg/promtail/config"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	exitCodeTestsSuccess  = 0
	exitCodeTestsFailure  = 1
	exitCodeConfigFailure = 2
)

func RunUnitTestsCommand(configFile string, testFiles ...string) int {
	var config config.Config

	// Load Promtail config
	if err := cfg.Unmarshal(&config, cfg.YAML(&configFile)); err != nil {
		fmt.Fprintln(os.Stderr, "Unable to load config file at", configFile, ":", err.Error())
		return exitCodeConfigFailure
	}

	// Ensure there's at least 1 scrape config
	if len(config.ScrapeConfig) == 0 {
		fmt.Fprintln(os.Stderr, "No scrape config found at", configFile)
		return exitCodeConfigFailure
	}

	// Load tests
	testSuites, err := LoadTestSuites(testFiles)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	// Build pipelines
	pipelines := map[string]*stages.Pipeline{}
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))

	for _, scrapeConfig := range config.ScrapeConfig {
		jobName := scrapeConfig.JobName
		pipelineStages := scrapeConfig.PipelineStages

		pipeline, err := stages.NewPipeline(logger, pipelineStages, &jobName, prometheus.DefaultRegisterer)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Unable to initialize pipeline for job", jobName, ":", err.Error())
			return exitCodeConfigFailure
		}

		pipelines[jobName] = pipeline
	}

	// Run tests
	reporter := NewReporter()

	for _, testSuite := range testSuites {
		testSuite.Run(reporter, pipelines)
	}

	if reporter.Failed() {
		return exitCodeTestsFailure
	} else {
		return exitCodeTestsSuccess
	}
}
