#!/bin/bash

# Seed script to populate the database with test data
echo "Seeding the database with test data..."

# Set environment variables
source ../.env

# Execute the seeding queries
psql $DATABASE_URL <<EOF
INSERT INTO users (username, password, email, created_at)
VALUES ('testuser', 'password123', 'testuser@example.com', NOW());

INSERT INTO urls (long_url, short_url, created_at, user_id, clicks)
VALUES
('https://example.com', 'exmpl1', NOW(), 1, 0),
('https://anotherexample.com', 'anoth1', NOW(), 1, 0);
EOF

echo "Database seeding completed."
