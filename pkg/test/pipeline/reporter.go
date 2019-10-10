package pipeline

import (
	"fmt"
	"strings"
)

type Reporter struct {
	results []Result
}

func NewReporter() *Reporter {
	return &Reporter{
		results: []Result{},
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
		fmt.Println(fmt.Sprintf("PASS (%s)", result.Header()))
	} else if result.customErr != nil {
		fmt.Println(fmt.Sprintf("FAIL (%s)", result.Header()))
		fmt.Println("    " + result.customErr.Error())
	} else {
		fmt.Println(fmt.Sprintf("FAIL (%s)", result.Header()))
		fmt.Println("    Expected:")
		fmt.Println(result.expectedLog.String(8))
		fmt.Println("    Actual:  ")
		fmt.Println(result.actualLog.String(8))
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
