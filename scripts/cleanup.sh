#!/bin/bash

# Cleanup script to remove expired URLs from the database
echo "Starting cleanup of expired URLs..."

# Set environment variables
source ../.env

# Execute the cleanup query
psql $DATABASE_URL <<EOF
DELETE FROM urls
WHERE expiry IS NOT NULL AND expiry < NOW();
EOF

echo "Cleanup completed successfully."
