# Multi-Binary Application Example

This example demonstrates creating a container with multiple binaries that work together.

## Build and Run

1. Build both applications:
   ```bash
   go build -o server server.go
   go build -o worker worker.go
   ```

2. Create the container with multiple binaries:
   ```bash
   gocontainer build \
     --binary ./server:/usr/local/bin/server \
     --binary ./worker:/usr/local/bin/worker \
     --label "app=multi-binary-example" \
     --label "version=1.0.0" \
     --outdir ./container-output
   ```

3. Import into Docker and run:
   ```bash
   # Import the container
   docker import ./container-output multi-app:latest
   
   # Run the server (which calls the worker)
   docker run -p 8080:8080 multi-app:latest /usr/local/bin/server
   ```

4. Test the application:
   ```bash
   curl http://localhost:8080
   ```

## Features Demonstrated

- Multiple binary inclusion
- Custom binary paths in container
- Inter-binary communication
- Organized application structure