# `lokitool`: tooling for Grafana Loki


## Promtail pipeline unit testing

**How it works**:

```
usage: lokitool test pipeline --config.file=CONFIG.FILE <test-file>...

Run unit tests for Promtail pipeline.

Flags:
  -h, --help                     Show context-sensitive help (also try --help-long and --help-man).
      --config.file=CONFIG.FILE  Promtail config file

Args:
  <test-file>  The unit test file.
```

**Example**:

```
$ ./lokitool test pipeline --config.file=./examples/test-pipeline/promtail-config.yaml ./examples/test-pipeline/test.yaml

PASS (suite: test.yaml test: #1 entry: #1)
PASS (suite: test.yaml test: #1 entry: #2)
```
