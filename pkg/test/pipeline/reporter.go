package pipeline

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

var (
	successColor = color.New(color.FgGreen)
	failureColor = color.New(color.FgRed)
)

type Reporter struct {
	results []Result
	verbose bool
}

func NewReporter(verbose bool) *Reporter {
	return &Reporter{
		results: []Result{},
		verbose: verbose,
	}
}

func (r *Reporter) Failed() bool {
	for _, result := range r.results {
		if !result.success {
			return true
		}
	}

	return false
}

func (r *Reporter) RecordTestSuccess(suiteID string, testID, entryID int, expectedLog, actualLog *TestLog) {
	r.recordResult(Result{
		suiteID:     suiteID,
		testID:      testID,
		entryID:     entryID,
		success:     true,
		expectedLog: expectedLog,
		actualLog:   actualLog,
	})
}

func (r *Reporter) RecordTestFailure(suiteID string, testID, entryID int, expectedLog, actualLog *TestLog) {
	r.recordResult(Result{
		suiteID:     suiteID,
		testID:      testID,
		entryID:     entryID,
		success:     false,
		expectedLog: expectedLog,
		actualLog:   actualLog,
	})
}

func (r *Reporter) RecordTestError(suiteID string, testID int, err error) {
	r.recordResult(Result{
		suiteID:   suiteID,
		testID:    testID,
		success:   false,
		customErr: err,
	})
}

func (r *Reporter) RecordSuiteError(suiteID string, err error) {
	r.recordResult(Result{
		suiteID:   suiteID,
		success:   false,
		customErr: err,
	})
}

func (r *Reporter) Summary() string {
	summary := []string{
		"",
		"========== SUMMARY ===========",
		"",
	}

	// Print and count failures
	failures := 0

	for _, result := range r.results {
		if result.success {
			continue
		}

		// In non-verbose mode we need to print failures
		if !r.verbose {
			summary = append(summary, result.Report(true))
		}

		failures++
	}

	// Display the number of failures (if any)
	if failures > 0 {
		summary = append(summary, failureColor.Sprintf("Failed %d tests out of %d", failures, len(r.results)))
	} else {
		summary = append(summary, successColor.Sprintf("All %d tests have passed", len(r.results)))
	}

	return strings.Join(summary, "\n")
}

func (r *Reporter) recordResult(result Result) {
	r.results = append(r.results, result)

	fmt.Print(result.Report(r.verbose))
}

type Result struct {
	suiteID     string
	testID      int
	entryID     int
	success     bool
	expectedLog *TestLog
	actualLog   *TestLog
	customErr   error
}

func (r *Result) Header() string {
	header := make([]string, 0, 3)

	if r.suiteID != "" {
		header = append(header, fmt.Sprintf("suite: %s", r.suiteID))
	}

	if r.testID != 0 {
		header = append(header, fmt.Sprintf("test: #%d", r.testID))
	}

	if r.entryID != 0 {
		header = append(header, fmt.Sprintf("entry: #%d", r.entryID))
	}

	return strings.Join(header, " ")
}

func (r *Result) Report(verbose bool) string {
	if verbose {
		return r.reportVerbose()
	} else {
		return r.reportMinimal()
	}
}

func (r *Result) reportMinimal() string {
	if r.success {
		return successColor.Sprint(".")
	} else {
		return failureColor.Sprint("F")
	}
}

func (r *Result) reportVerbose() string {
	var report []string

	if r.success {
		report = []string{
			successColor.Sprintf("PASS (%s)", r.Header()),
		}
	} else if r.customErr != nil {
		report = []string{
			failureColor.Sprintf("FAIL (%s)", r.Header()),
			failureColor.Sprintf("    %s", r.customErr.Error()),
		}
	} else {
		report = []string{
			failureColor.Sprintf("FAIL (%s)", r.Header()),
			failureColor.Sprint("    Expected:"),
			failureColor.Sprint(r.expectedLog.String(8)),
			failureColor.Sprint("    Actual:  "),
			failureColor.Sprint(r.actualLog.String(8)),
		}
	}

	return strings.Join(report, "\n") + "\n"
}
