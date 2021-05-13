module github.com/fmy1993/blockchain-real-estate

go 1.16

require (
	github.com/Knetic/govaluate v3.0.1-0.20171022003610-9aa49832a739+incompatible // indirect
	github.com/Shopify/sarama v1.28.0 // indirect
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/cloudflare/cfssl v1.5.0 // indirect
	github.com/coreos/bbolt v1.3.5 // indirect
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/fsouza/go-dockerclient v1.7.2 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-critic/go-critic v0.5.5 // indirect
	github.com/go-ini/ini v1.62.0
	github.com/go-logfmt/logfmt v0.5.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/mock v1.5.0 // indirect
	github.com/golang/protobuf v1.5.1 // indirect
	github.com/golangci/golangci-lint v1.38.0 // indirect
	github.com/google/btree v1.0.1 // indirect
	github.com/google/certificate-transparency-go v1.1.0 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/gostaticanalysis/analysisutil v0.7.1 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/hyperledger/fabric v1.4.11
	github.com/hyperledger/fabric-amcl v0.0.0-20200424173818-327c9e2cf77a // indirect
	github.com/hyperledger/fabric-config v0.1.0 // indirect
	github.com/hyperledger/fabric-protos-go v0.0.0-20210318103044-13fdee960194
	github.com/hyperledger/fabric-sdk-go v1.0.0
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/jmoiron/sqlx v1.3.1 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/magiconair/properties v1.8.4 // indirect
	github.com/mattn/go-runewidth v0.0.10 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/nbutton23/zxcvbn-go v0.0.0-20210217022336-fa2cb2858354 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/pelletier/go-toml v1.8.1 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/procfs v0.6.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/robfig/cron/v3 v3.0.1
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/spf13/afero v1.6.0 // indirect
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.7.0
	github.com/sykesm/zap-logfmt v0.0.4 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20201229170055-e5319fda7802 // indirect
	github.com/urfave/cli v1.22.5 // indirect
	github.com/zmap/zcrypto v0.0.0-20210318145614-34856aada1be // indirect
	go.etcd.io/bbolt v1.3.5 // indirect
	go.etcd.io/etcd v3.3.25+incompatible // indirect
	go.opencensus.io v0.22.4 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2 // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/mod v0.4.2 // indirect
	golang.org/x/net v0.0.0-20210316092652-d523dce5a7f4 // indirect
	golang.org/x/sys v0.0.0-20210320140829-1e4c9ba3b0c4 // indirect
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba // indirect
	google.golang.org/genproto v0.0.0-20210322173543-5f0e89347f5a // indirect
	google.golang.org/grpc v1.36.0 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect

//github.com/fmy1993/blockchain-real-estate v0.1.1
)

replace github.com/coreos/bbolt v1.3.5 => go.etcd.io/bbolt v1.3.5

replace go.etcd.io/bbolt v1.3.5 => github.com/coreos/bbolt v1.3.5

replace blockchain-real-estate/github.com/togettoyou/blockchain-real-estate v0.0.0-20210225030431-330c0099e6b4 => blockchain-real-estate/github.com/fmy1993/blockchain-real-estate v0.1.2

replace google.golang.org/grpc v1.36.0 => google.golang.org/grpc v1.26.0
