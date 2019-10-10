package pipeline

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type Reporter struct {
	results      []Result
	successColor *color.Color
	failureColor *color.Color
}

func NewReporter() *Reporter {
	return &Reporter{
		results:      []Result{},
		successColor: color.New(color.FgGreen),
		failureColor: color.New(color.FgRed),
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

func (r *Reporter) RecordTestSuccess(suiteID string, testID, entryID int, expectedLog, actualLog *TestExpectedLog) {
	r.recordResult(Result{
		suiteID:     suiteID,
		testID:      testID,
		entryID:     entryID,
		success:     true,
		expectedLog: expectedLog,
		actualLog:   actualLog,
	})
}

func (r *Reporter) RecordTestFailure(suiteID string, testID, entryID int, expectedLog, actualLog *TestExpectedLog) {
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

func (r *Reporter) recordResult(result Result) {
	r.results = append(r.results, result)
	r.printResult(result)
}

func (r *Reporter) printResult(result Result) {
	if result.success {
		r.successColor.Println(fmt.Sprintf("PASS (%s)", result.Header()))
	} else if result.customErr != nil {
		r.failureColor.Println(fmt.Sprintf("FAIL (%s)", result.Header()))
		r.failureColor.Println("    " + result.customErr.Error())
	} else {
		r.failureColor.Println(fmt.Sprintf("FAIL (%s)", result.Header()))
		r.failureColor.Println("    Expected:")
		r.failureColor.Println(result.expectedLog.String(8))
		r.failureColor.Println("    Actual:  ")
		r.failureColor.Println(result.actualLog.String(8))
	}
}

type Result struct {
	suiteID     string
	testID      int
	entryID     int
	success     bool
	expectedLog *TestExpectedLog
	actualLog   *TestExpectedLog
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
