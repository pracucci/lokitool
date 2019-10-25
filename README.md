# `lokitool`: tooling for Grafana Loki


## Promtail pipeline unit testing

**How it works**:

```
usage: lokitool test pipeline --config.file=CONFIG.FILE [<flags>] <test-file>...

Run unit tests for Promtail pipeline.

Flags:
  -h, --help                     Show context-sensitive help (also try --help-long and --help-man).
      --verbose                  Enable verbose output.
      --config.file=CONFIG.FILE  Promtail config file.

Args:
  <test-file>  The unit test file.
```

**Example**:

```
$ ./lokitool test pipeline --config.file=./examples/test-pipeline/promtail-config.yaml ./examples/test-pipeline/test.yaml
..

========== SUMMARY ===========

All 2 tests have passed
```


## Release

Run `make release VERSION=0.1.2` and follow the instructions.


## License

Apache License 2.0, see [LICENSE](./LICENSE).
