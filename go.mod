module blockchain-real-estate

go 1.16

require (
	// github.com/fmy1993/blockchain-real-estate v0.1.0

	github.com/Shopify/sarama v1.28.0 // indirect
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/cloudflare/cfssl v0.0.0-20190409034051-768cd563887f // indirect
	github.com/fsouza/go-dockerclient v1.7.2 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.62.0
	github.com/golang/mock v1.4.3 // indirect
	github.com/golang/protobuf v1.5.1 // indirect
	github.com/google/certificate-transparency-go v1.0.21 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0 // indirect
	github.com/hashicorp/go-version v1.2.1 // indirect
	github.com/hyperledger/fabric v1.4.11
	github.com/hyperledger/fabric-amcl v0.0.0-20200424173818-327c9e2cf77a // indirect
	github.com/hyperledger/fabric-protos-go v0.0.0-20200707132912-fee30f3ccd23 // indirect
	github.com/hyperledger/fabric-sdk-go v1.0.0-beta2
	github.com/miekg/pkcs11 v1.0.3 // indirect
	github.com/mitchellh/mapstructure v1.3.2 // indirect
	github.com/onsi/gomega v1.9.0 // indirect
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7 // indirect
	github.com/pelletier/go-toml v1.8.0 // indirect
	github.com/prometheus/client_golang v1.1.0 // indirect
	github.com/robfig/cron v1.2.0
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/spf13/afero v1.3.1 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.7.1 // indirect
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.7.0
	github.com/sykesm/zap-logfmt v0.0.4 // indirect
	github.com/togettoyou/blockchain-real-estate v0.0.0-20210225030431-330c0099e6b4
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/crypto v0.0.0-20210317152858-513c2a44f670 // indirect
	google.golang.org/grpc v1.36.0 // indirect

)

replace blockchain-real-estate/github.com/togettoyou/blockchain-real-estate v0.0.0-20210225030431-330c0099e6b4 => blockchain-real-estate/github.com/fmy1993/blockchain-real-estate v0.1.0

replace blockchain-real-estate/github.com/robfig/cron v1.2.0 => github.com/robfig/cron/v3 v3.0.1
