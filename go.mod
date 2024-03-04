module collection.com/main

go 1.21.4

replace collection.com/config => /home/andreamboni/collection-ws/config

replace collection.com/router => /home/andreamboni/collection-ws/router

replace collection.com/language => /home/andreamboni/collection-ws/handler/language

replace collection.com/item => /home/andreamboni/collection-ws/handler/item

replace collection.com/author => /home/andreamboni/collection-ws/handler/author

replace collection.com/models => /home/andreamboni/collection-ws/models

replace collection.com/handler => /home/andreamboni/collection-ws/handler

replace collection.com/collection => /home/andreamboni/collection-ws/handler/collection

replace collection.com/country => /home/andreamboni/collection-ws/handler/country

replace collection.com/publisher => /home/andreamboni/collection-ws/handler/publisher

require (
	collection.com/config v1.0.0
	collection.com/router v1.0.0
)

require (
	collection.com/author v1.0.0 // indirect
	collection.com/collection v1.0.0 // indirect
	collection.com/country v1.0.0 // indirect
	collection.com/handler v1.0.0 // indirect
	collection.com/item v1.0.0 // indirect
	collection.com/language v1.0.0 // indirect
	collection.com/models v1.0.0 // indirect
	collection.com/publisher v1.0.0 // indirect
	github.com/bytedance/sonic v1.10.1 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20230717121745-296ad89f973d // indirect
	github.com/chenzhuoyu/iasm v0.9.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/gin-contrib/cors v1.5.0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.9.1 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.15.5 // indirect
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.5 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.1.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	golang.org/x/arch v0.5.0 // indirect
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/net v0.16.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
