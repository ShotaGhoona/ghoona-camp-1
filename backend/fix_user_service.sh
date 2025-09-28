#!/bin/bash

# Fix all ExecuteInTxWithResult issues by replacing with simpler approach
echo "Fixing user_service.go compilation errors..."

# Get the file
FILE="/Users/yamashitashota/Doc/ghoona/starup/ghoona-camp/ghoona-camp-1/backend/internal/application/usecase/user_service.go"

echo "File exists: $(test -f "$FILE" && echo "yes" || echo "no")"
echo "Making backup..."
cp "$FILE" "${FILE}.backup"

echo "Done with backup"