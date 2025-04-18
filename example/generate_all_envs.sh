#!/bin/bash

# Script to generate .env files for all clients
# Usage: ./generate_all_envs.sh

# Path to the template file
TEMPLATE_FILE="template.properties"

# Directory where this script is located
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Change to the script directory
cd "$SCRIPT_DIR"

# Get a list of all CSV files in the current directory
CSV_FILES=$(ls *.csv)

# Loop through each CSV file
for CSV_FILE in $CSV_FILES; do
    # Extract client name from CSV filename (remove .csv extension)
    CLIENT_NAME="${CSV_FILE%.csv}"

    # Output file name
    OUTPUT_FILE="${CLIENT_NAME}.env"

    echo "Generating .env file for client: $CLIENT_NAME"

    # Run the .env generator
    ../csv2env generate -t "$TEMPLATE_FILE" -c "$CSV_FILE" -o "$OUTPUT_FILE"

    # Check if the command was successful
    if [ $? -eq 0 ]; then
        echo "Successfully generated $OUTPUT_FILE"
    else
        echo "Failed to generate $OUTPUT_FILE"
    fi

    echo "-----------------------------------"
done

echo "All .env files have been generated."
