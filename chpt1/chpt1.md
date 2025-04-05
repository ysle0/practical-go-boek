# Go Greeter Application

A simple command-line greeter application written in Go that demonstrates basic CLI functionality with flag parsing.

## Features

- Interactive command-line greeter that asks for your name
- Customizable number of greetings
- Optional HTML page generation with greeting message
- Simple error handling and input validation

## Installation

```bash
# Clone the repository
git clone <repository-url>

# Navigate to the project directory
cd practical-go-boek

# Build the application
go build -o greeter chpt1/flag-parse/main.go
```

## Usage

The application requires at least one flag:

```bash
./greeter -n <number-of-greetings> [-o <output-page-name>]
```

### Required Flags

- `-n`: The number of times to display the greeting (must be greater than 0)

### Optional Flags

- `-o`: Generate an HTML page with the given name (e.g., `-o greeting` will create `greeting.gohtml`)

## Examples

Basic greeting (5 times):

```bash
./greeter -n 5
```

Generate greeting with HTML page:

```bash
./greeter -n 3 -o welcome
```

## How It Works

1. The application parses command-line arguments using the `flag` package
2. It validates that the number of greetings is greater than 0
3. It prompts the user to enter their name
4. If the `-o` flag is provided, it generates an HTML page with the greeting
5. It prints the greeting message the specified number of times

## Error Handling

The application handles various error conditions:

- Invalid flags
- Missing required flags
- Empty user input for name
- File access or creation issues
