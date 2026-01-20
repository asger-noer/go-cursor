module examples

go 1.24.0

toolchain go1.24.1

// Ensure that the go-cursor package is used from the local directory
replace github.com/asger-noer/go-cursor => ../

require (
	github.com/99designs/gqlgen v0.17.86
	github.com/asger-noer/go-cursor v1.0.3
	github.com/vektah/gqlparser/v2 v2.5.31
)

require (
	github.com/agnivade/levenshtein v1.2.1 // indirect
	github.com/go-viper/mapstructure/v2 v2.4.0 // indirect
	github.com/goccy/go-yaml v1.19.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/sosodev/duration v1.3.1 // indirect
	github.com/urfave/cli/v3 v3.6.1 // indirect
	golang.org/x/mod v0.31.0 // indirect
	golang.org/x/sync v0.19.0 // indirect
	golang.org/x/text v0.33.0 // indirect
	golang.org/x/tools v0.40.0 // indirect
)
