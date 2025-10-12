# GoContainer

[![Go Version](https://img.shields.io/github/go-mod/go-version/sntns/gocontainer)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/sntns/gocontainer)](https://goreportcard.com/report/github.com/sntns/gocontainer)

A tool to generate Docker/OCI containers for Go applications without requiring Docker or complex build setups.

## Features

- 🚀 Generate OCI-compliant containers directly from Go binaries
- 🔧 Cross-platform support (Linux, macOS, Windows)
- 📦 Minimal container images with only necessary files
- 🏥 Health check support
- 🏷️ Custom labels and metadata
- 📁 File and directory copying capabilities
- 🎯 No Docker daemon required

## Installation

### From Source

```bash
go install github.com/sntns/gocontainer/cmd/gocontainer@latest
```

### From Releases

Download the latest binary from the [releases page](https://github.com/sntns/gocontainer/releases).

## Quick Start

1. Build your Go application:
   ```bash
   go build -o myapp main.go
   ```

2. Generate an OCI container:
   ```bash
   gocontainer build --binary myapp --outdir ./container-output
   ```

3. The generated container can be used with any OCI-compatible runtime (Docker, Podman, etc.):
   ```bash
   # Import into Docker
   docker import ./container-output myapp:latest
   
   # Or use with Podman
   podman import ./container-output myapp:latest
   ```

### Try the Demo

If you have the repository cloned, you can quickly try GoContainer with our examples:

```bash
# Clone and run demo
git clone https://github.com/sntns/gocontainer.git
cd gocontainer
make demo

# The demo creates a container from the webserver example
# Check the output in examples/webserver/container-output/
```

## Usage

### Basic Container Creation

Create a container with a single binary:

```bash
gocontainer build \
  --binary ./myapp \
  --outdir ./output
```

### Advanced Usage

Create a container with multiple binaries, files, labels, and health checks:

```bash
gocontainer build \
  --binary ./myapp:/usr/local/bin/myapp \
  --binary ./helper:/usr/local/bin/helper \
  --copy ./config:/etc/myapp \
  --copy ./static:/var/www \
  --label "version=1.0.0" \
  --label "maintainer=me@example.com" \
  --healthcheck "--interval=30s --timeout=3s --retries=3 CMD [\"myapp\", \"health\"]" \
  --outdir ./container-output
```

### Command Options

- `--binary <file>[:<name>]`: Binary to include in container (can be used multiple times)
- `--copy <source>[:<destination>]`: Copy files or directories to container (can be used multiple times)  
- `--label <key=value>`: Add metadata labels (can be used multiple times)
- `--healthcheck <config>`: Add health check instruction
- `--outdir <directory>`: Output directory for the OCI container (required)

### Health Check Format

The health check flag accepts Docker-style health check syntax:

```bash
--healthcheck "--interval=30s --timeout=3s --retries=3 CMD [\"myapp\", \"health\"]"
```

## Examples

See the [examples](examples/) directory for complete usage examples:

- [Simple web server](examples/webserver/)
- [Multi-binary application](examples/multi-binary/)
- [Application with configuration files](examples/with-config/)

## Library Usage

GoContainer can also be used as a library in your Go applications:

```go
package main

import (
    "github.com/sntns/gocontainer/pkg/container"
)

func main() {
    // Create a new container
    cont, err := container.New()
    if err != nil {
        panic(err)
    }

    // Add binaries
    err = cont.AddBinary([]string{"./myapp"})
    if err != nil {
        panic(err)
    }

    // Set labels
    err = cont.SetLabels("version=1.0.0", "app=myapp")
    if err != nil {
        panic(err)
    }

    // Save the container
    err = cont.Save("./output")
    if err != nil {
        panic(err)
    }
}
```

## How It Works

GoContainer generates OCI-compliant container images by:

1. **Binary Analysis**: Detecting the target platform (OS/architecture) of Go binaries
2. **Layer Creation**: Building filesystem layers with your binaries and files
3. **Metadata Generation**: Creating proper OCI image configuration with labels, health checks, and entry points
4. **OCI Layout**: Outputting a complete OCI image layout that can be imported into any container runtime

The generated containers are minimal and contain only the files you specify, resulting in smaller image sizes and improved security.

## Requirements

- Go 1.19 or later
- Target binaries must be statically linked (default for Go)

## Contributing

We welcome contributions! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details on how to contribute to this project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Related Projects

- [OCI Image Specification](https://github.com/opencontainers/image-spec)
- [Go Containers](https://github.com/google/go-containerregistry)
- [BuildKit](https://github.com/moby/buildkit)