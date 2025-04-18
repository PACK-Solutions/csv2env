# .env Generator

A command-line tool that generates `.env` files from a `.properties` template and CSV files containing property values.

## Overview

This tool takes a `.properties` template file with placeholders in the format `#PLACEHOLDER#` and replaces them with values from a CSV file. The CSV file should have property keys in the header row and values in the subsequent row.

## Installation

```bash
go install github.com/yourusername/csv2env@latest
```

Or clone the repository and build it:

```bash
git clone https://github.com/yourusername/csv2env.git
cd csv2env
go build
```

## Usage

```bash
# Generate a .env file
./csv2env generate -t <path-to-properties-template> -c <path-to-csv-file> -o <path-to-output-env-file>

# Show help
./csv2env --help

# Show help for the generate command
./csv2env generate --help
```

### Command-line Options

#### Generate Command

- `-t, --template`: Path to the `.properties` template file (required)
- `-c, --csv`: Path to the CSV file containing property values (required)
- `-o, --output`: Path to the output `.env` file (default: `.env`)

## Examples

### Template File (template.properties)

```properties
PROPERTY1_KEY=#PROPERTY1_VALUE#
PROPERTY2_KEY=#PROPERTY2_VALUE#
PROPERTY3_KEY=#PROPERTY3_VALUE#
DATABASE_URL=jdbc:mysql://#DB_HOST#:#DB_PORT#/#DB_NAME#
API_KEY=#API_KEY#
DEBUG_MODE=#DEBUG_MODE#
```

### CSV Files

#### client1.csv

```csv
PROPERTY1_VALUE,PROPERTY2_VALUE,PROPERTY3_VALUE,DB_HOST,DB_PORT,DB_NAME,API_KEY,DEBUG_MODE
value1,value2,value3,localhost,3306,mydb,abc123,true
```

#### client2.csv

```csv
PROPERTY1_VALUE,PROPERTY2_VALUE,PROPERTY3_VALUE,DB_HOST,DB_PORT,DB_NAME,API_KEY,DEBUG_MODE
client2-value1,client2-value2,client2-value3,db.example.com,5432,production,xyz789,false
```

### Generate .env File for a Single Client

```bash
./csv2env -template example/template.properties -csv example/client1.csv -output .env
```

### Batch Processing for Multiple Clients

The repository includes a shell script (`example/generate_all_envs.sh`) that demonstrates how to generate .env files for multiple clients in batch:

```bash
cd example
chmod +x generate_all_envs.sh
./generate_all_envs.sh
```

This script will:
1. Find all CSV files in the example directory
2. Generate a corresponding .env file for each client
3. Name the output files based on the client name (e.g., client1.env, client2.env)

### Output (.env)

```
PROPERTY1_KEY=value1
PROPERTY2_KEY=value2
PROPERTY3_KEY=value3
DATABASE_URL=jdbc:mysql://localhost:3306/mydb
API_KEY=abc123
DEBUG_MODE=true
```

## CSV File Format

The CSV file should have the following format:
- First row: Property keys (without the `#` symbols)
- Second row: Property values

Each client should have its own CSV file with the appropriate values.

## Error Handling

The tool will exit with an error message if:
- Required command-line flags are missing
- The template file or CSV file cannot be read
- The CSV file does not have at least a header row and a data row
- The number of headers and values in the CSV file do not match
- The output file cannot be written

## License

MIT
