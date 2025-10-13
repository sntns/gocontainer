# Application with Configuration Files

This example demonstrates creating a container that includes configuration files alongside the binary.

## Build and Run

1. Build the application:
   ```bash
   go build -o configapp main.go
   ```

2. Create the container with binary and configuration:
   ```bash
   gocontainer build \
     --binary ./configapp \
     --copy ./config.json:/etc/myapp/config.json \
     --label "app=configurable-app" \
     --label "version=1.0.0" \
     --healthcheck "--interval=30s --timeout=3s --retries=3 CMD [ \"./configapp\" ]" \
     --outdir ./container-output
   ```

3. Import into Docker and run:
   ```bash
   # Import the container
   docker import ./container-output configapp:latest
   
   # Run the application
   docker run -p 8080:8080 configapp:latest /configapp
   ```

4. Test the application:
   ```bash
   curl http://localhost:8080
   curl http://localhost:8080/health
   ```

## Features Demonstrated

- Binary and configuration file inclusion
- File copying with custom paths
- Configuration loading from container filesystem
- Graceful fallback when config is missing
- Health check endpoint