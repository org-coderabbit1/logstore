#!/bin/sh

# Create required directories
mkdir -p /tmp/logstore/chunks /tmp/logstore/rules

# Set proper permissions
chown -R nobody:nobody /tmp/logstore
chmod -R 777 /tmp/logstore

# Start Logstore
exec logstore --config.file=/etc/logstore/local-config.yaml 