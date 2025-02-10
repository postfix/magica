#!/bin/bash

# Define the old and new module paths
OLD_MODULE="github.com/postifx/magika/magika"
NEW_MODULE="github.com/postfix/magika"

# Find all Go files in the repository and replace occurrences of the old module path
find . -type f -name "*.go" -print0 | while IFS= read -r -d '' file; do
    if grep -q "$OLD_MODULE" "$file"; then
        echo "Updating imports in: $file"
        sed -i "s|$OLD_MODULE|$NEW_MODULE|g" "$file"
    fi
done

echo "Import paths updated successfully!"

