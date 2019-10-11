package pipeline

import (
	"fmt"
	"strings"
	"time"

	"github.com/prometheus/common/model"
)

type TestLog struct {
	Timestamp time.Time      `yaml:"timestamp"`
	Entry     string         `yaml:"entry"`
	Labels    model.LabelSet `yaml:"labels"`
}

func (l *TestLog) Equal(other *TestLog) bool {
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

func (l *TestLog) Time() time.Time {
	if l.Timestamp.After(time.Unix(0, 0)) {
		return l.Timestamp
	} else {
		return time.Now()
	}
}

func (l *TestLog) String(indent int) string {
	prefix := strings.Repeat(" ", indent)
	lines := []string{
		prefix + fmt.Sprintf("Timestamp: %s", l.Timestamp.Format(time.RFC3339Nano)),
		prefix + fmt.Sprintf("Entry:     %s", l.Entry),
		prefix + fmt.Sprintf("Labels:    %s", l.Labels.String()),
	}

	return strings.Join(lines, "\n")
}
