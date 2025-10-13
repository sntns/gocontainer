# Simple Web Server Example

This example demonstrates creating a container for a simple HTTP web server.

## Build and Run

1. Build the Go application:
   ```bash
   go build -o webserver main.go
   ```

2. Create the container:
   ```bash
   gocontainer build \
     --binary ./webserver \
     --label "app=webserver" \
     --label "version=1.0.0" \
     --healthcheck "--interval=30s --timeout=3s --retries=3 CMD [ \"./webserver\" ]" \
     --outdir ./container-output
   ```

3. Import into Docker and run:
   ```bash
   # Import the container
   docker import ./container-output webserver:latest
   
   # Run the container
   docker run -p 8080:8080 -e PORT=8080 webserver:latest /webserver
   ```

4. Test the server:
   ```bash
   curl http://localhost:8080
   curl http://localhost:8080/health
   ```

## Features Demonstrated

- Basic binary inclusion
- Custom labels for metadata
- Health check configuration
- Environment variable usage
- HTTP server in a container