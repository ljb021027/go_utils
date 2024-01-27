module github.com/ljb021027/unit

go 1.13

require (
	github.com/Knetic/govaluate v3.0.0+incompatible
	github.com/coreos/bbolt v0.0.0-00010101000000-000000000000 // indirect
	github.com/coreos/etcd v3.3.27+incompatible
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/fatih/structs v1.1.0
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/mock v1.4.4
	github.com/google/btree v1.0.1 // indirect
	github.com/google/go-querystring v1.0.0
	github.com/google/uuid v1.2.0
	github.com/google/wire v0.4.0
	github.com/gorilla/schema v1.2.0
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/json-iterator/go v1.1.11
	github.com/juju/ratelimit v1.0.1
	github.com/mitchellh/mapstructure v1.4.1
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/prometheus/client_golang v1.11.0 // indirect
	github.com/shopspring/decimal v1.2.0
	github.com/soheilhy/cmux v0.1.5 // indirect
	github.com/stretchr/testify v1.8.1
	github.com/tmc/grpc-websocket-proxy v0.0.0-20201229170055-e5319fda7802 // indirect
	github.com/traefik/yaegi v0.10.0
	github.com/vmihailenco/msgpack/v5 v5.0.0-beta.1
	github.com/xiang90/probing v0.0.0-20190116061207-43a291ad63a2 // indirect
	go.etcd.io/bbolt v1.3.8 // indirect
	go.uber.org/zap v1.19.1 // indirect
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac // indirect
	google.golang.org/genproto v0.0.0-20211021150943-2b146023228c // indirect
	google.golang.org/grpc v1.41.0 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

//replace github.com/coreos/bbolt v1.3.4 => github.com/coreos/bbolt v1.3.3
replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.4

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
