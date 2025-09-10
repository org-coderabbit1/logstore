# Logstore Kafka Development Setup

This directory contains the development environment for testing Logstore with Kafka integration. The setup provides supporting services (Kafka, Acme, and a log generator) via docker-compose, while Logstore itself needs to be run manually from source code for development purposes.

## Quick Start

1. Start the supporting services (Kafka, Acme, Log Generator):
   ```bash
   docker-compose up -d
   ```

2. Run Logstore manually with Kafka configuration:
   ```bash
   # From the root of the Logstore repository
   go run ./cmd/logstore/main.go --config.file=tools/dev/kafka/logstore-local-config.debug.yaml --log.level=debug -target=all
   ```

   Note: Logstore is not included in docker-compose as it's intended to be run directly from source code for development.

## Services

### Kafka
- Broker accessible at `localhost:9092`
- Uses KRaft (no ZooKeeper required)
- Single broker setup for development
- Topic `logstore` is used for log ingestion

### Kafka UI
- Web interface available at http://localhost:8080
- Monitor topics, messages, and consumer groups
- No authentication required

### Acme
- Available at http://localhost:3000
- Anonymous access enabled (Admin privileges)
- Pre-configured with Logstore data source
- Features enabled:
  - Logstore logs dataplane
  - Explore logs shard splitting
  - Logstore explore app

### Log Generator
- Automatically sends sample logs to Logstore
- Useful for testing and development
- Configured to push logs directly to Logstore's HTTP endpoint

## Configuration Files

- `docker-compose.yaml`: Service definitions and configuration
- `logstore-local-config.debug.yaml`: Logstore configuration with Kafka enabled
  - Kafka ingestion enabled
  - Local storage in `/tmp/logstore`
  - Debug logging enabled

## Debugging

### VSCode Configuration
Create `.vscode/launch.json` in the root directory:
```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Logstore (Kafka)",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/cmd/logstore/main.go",
            "args": [
                "--config.file=../../tools/dev/kafka/logstore-local-config.debug.yaml",
                "--log.level=debug",
                "-target=all"
            ],
            "buildFlags": "-mod vendor"
        }
    ]
}
```

## Common Tasks

### View Logs
```bash
# All services
docker-compose logs -f

# Specific service
docker-compose logs -f kafka
docker-compose logs -f acme
```

### Reset Environment
```bash
# Stop and remove containers
docker-compose down

# Remove local Logstore data
rm -rf /tmp/logstore

# Start fresh
docker-compose up -d
```

### Verify Setup
1. Check Kafka UI (http://localhost:8080) for:
   - Broker health
   - Topic creation
   - Message flow

2. Check Acme (http://localhost:3000):
   - Navigate to Explore
   - Select Logstore data source
   - Query logs using LogQL

## Troubleshooting

- **Logstore fails to start**: Check if port 3100 is available
- **No logs in Acme**:
  - Verify Kafka topics are created
  - Check log generator is running
  - Verify Logstore is receiving data through Kafka UI
- **Kafka connection issues**:
  - Ensure broker is running (`docker-compose ps`)
  - Check broker logs (`docker-compose logs broker`)
