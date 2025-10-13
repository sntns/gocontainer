module github.com/sntns/gocontainer

go 1.22

require (
	github.com/spf13/cobra v1.10.1
	github.com/stretchr/testify v1.11.1
)

require github.com/opencontainers/go-digest v1.0.0

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/opencontainers/image-spec v1.1.1
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/pflag v1.0.10 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// Exclude examples from main module
exclude github.com/sntns/gocontainer/examples/webserver v0.0.0

exclude github.com/sntns/gocontainer/examples/multi-binary v0.0.0

exclude github.com/sntns/gocontainer/examples/with-config v0.0.0
