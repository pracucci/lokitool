package main

import (
	"os"
	"path/filepath"

	"github.com/pracucci/lokitool/pkg/test/pipeline"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	app := kingpin.New(filepath.Base(os.Args[0]), "Tooling for Loki.")
	app.HelpFlag.Short('h')

	testCmd := app.Command("test", "Run unit tests.")
	testPipelineCmd := testCmd.Command("pipeline", "Run unit tests for Promtail pipeline.")
	testPipelineVerbose := testPipelineCmd.Flag("verbose", "Enable verbose output.").Bool()
	testPipelineConfigFile := testPipelineCmd.Flag("config.file", "Promtail config file.").Required().String()
	testPipelineFiles := testPipelineCmd.Arg(
		"test-file",
		"The unit test file.",
	).Required().ExistingFiles()

	parsedCmd := kingpin.MustParse(app.Parse(os.Args[1:]))

	switch parsedCmd {
	case testPipelineCmd.FullCommand():
		os.Exit(pipeline.RunUnitTestsCommand(*testPipelineConfigFile, *testPipelineFiles, *testPipelineVerbose))
	}
}
