module github.com/pracucci/lokitool

go 1.13

require (
	github.com/alecthomas/units v0.0.0-20190924025748-f65c72e2690d // indirect
	github.com/blang/semver v3.5.1+incompatible // indirect
	github.com/bmatcuk/doublestar v1.1.1 // indirect
	github.com/coreos/go-systemd v0.0.0-20190321100706-95778dfbb74e // indirect
	github.com/cortexproject/cortex v0.2.1-0.20191003165238-857bb8476e59 // indirect
	github.com/fatih/color v1.7.0
	github.com/go-kit/kit v0.9.0
	github.com/gocql/gocql v0.0.0-20181124151448-70385f88b28b // indirect
	github.com/gogo/protobuf v1.3.0 // indirect
	github.com/grafana/loki v6.7.8+incompatible
	github.com/grpc-ecosystem/grpc-gateway v1.9.6 // indirect
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v1.1.0
	github.com/prometheus/common v0.7.0
	go.etcd.io/etcd v0.0.0-20190815204525-8f85f0dc2607 // indirect
	go.opencensus.io v0.22.1 // indirect
	go.uber.org/atomic v1.4.0 // indirect
	golang.org/x/net v0.0.0-20190923162816-aa69164e4478 // indirect
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e // indirect
	golang.org/x/time v0.0.0-20190921001708-c4c64cad1fd0 // indirect
	golang.org/x/tools v0.0.0-20190925134113-a044388aa56f // indirect
	google.golang.org/api v0.10.0 // indirect
	google.golang.org/appengine v1.6.3 // indirect
	google.golang.org/genproto v0.0.0-20190916214212-f660b8655731 // indirect
	google.golang.org/grpc v1.23.1 // indirect
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	gopkg.in/yaml.v2 v2.2.4
)

replace github.com/weaveworks/common => github.com/sandlis/weaveworks-common v0.0.0-20190822064708-8fa0a1ca9d89

replace github.com/hpcloud/tail => github.com/grafana/tail v0.0.0-20191024143944-0b54ddf21fe7

// Override reference that causes an error from Go proxy - see https://github.com/golang/go/issues/33558
replace k8s.io/client-go => k8s.io/client-go v0.0.0-20190620085101-78d2af792bab

// Override reference causing proxy error.  Otherwise it attempts to download https://proxy.golang.org/golang.org/x/net/@v/v0.0.0-20190813000000-74dc4d7220e7.info
replace golang.org/x/net => golang.org/x/net v0.0.0-20190923162816-aa69164e4478
