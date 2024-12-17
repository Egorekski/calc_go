# Calculator

**Calculator** is a simple calculator written in Go, implemented as both an HTTP server and a CLI application. It performs basic arithmetic operations (addition, subtraction, multiplication, division) and returns results in JSON format for HTTP or directly in the console for CLI.

---

### Translations

- en [English](./README.md)
- ru [Русский](./README_ru.md)

---

## Features

- Built with **Go**, ensuring high performance and simplicity.
- Supports two modes of operation:
    - **HTTP Server** mode for API-based integration.
    - **CLI** mode for quick calculations via the terminal.
- Supports basic operations:
    - **Addition**
    - **Subtraction**
    - **Multiplication**
    - **Division**
- Handles numbers in `float` format.
- Simple and intuitive interfaces for both HTTP and CLI.
- Error handling for edge cases like division by zero.

---

## Getting Started

### Prerequisites

- **Go** (minimum version: `1.18`)  
  [Download and install Go](https://golang.org/dl/).

---

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/http-calculator.git
   ```

---

### Usage
#### HTTP Server Mode

1. Start the server:

   Uncomment `app.RunServer()` and comment `app.Run()` in `main.go` file.

   The server will start on http://localhost:8080, or you can write your own port by writing it as env variable.
   Example:
    ```bash
    export PORT='your-port' && go run ./cmd/main.go 
   ```

2. Use an HTTP client (e.g., curl or Postman) to interact with the API.

#### Request Format

Send a POST request with JSON data containing the string with expression as follows:
```json
{
  "expression": "2+2"    
}
```

#### Example Request
```bash
curl -X POST http://localhost:8080 \
-H "Content-Type: application/json" \
-d '{"expression": "2+2"}'
```

#### Response:
```json
{
    "result": 4
}
```

---

### CLI Mode

1. Run the application in CLI mode:

   Uncomment `app.Run()` and comment `app.RunServer()` in `main.go` file.
   Then run the application:
    ```bash
    go run ./cmd/main.go
    ```

2. Follow the interactive prompts to perform operations. For example:
```bash

2024/12/07 11:48:33 Welcome to the CLI Calculator!
2024/12/07 11:48:33 input expression
1+1
2024/12/07 11:48:37 1+1 = 2
2024/12/07 11:48:37 input expression
exit
2024/12/07 11:48:43 aplication was successfully closed
```

3. Write `exit` to exit the program

---

### Project Structure
```graphql
calc_go/
├── cmd/ 
     └── main.go                          # Entry point of the application
├── internal/                             # Contains HTTP server implementation
        └── application/
                ├── application.go        # Application
                └── application_test.go   # Tests for application
├── pkg/           
     └── calculation/
              ├── calculation.go          # Utility functions for calculations
              ├── calculation_test.go     # Tests for calculations
              └── errors.go               # Separate file for returned errors
└── README.md                             # Project documentation
```

---

### Future Plans

- Add support for advanced mathematical operations (e.g., exponentiation, square root).
- Implement logging for requests and results.
- Add Docker support for containerized deployment.
- Extend CLI mode with additional interactive features.

---

### Contributing

Contributions are welcome! If you’d like to contribute:

1. Fork the repository.
2. Create a feature branch: `git checkout -b feature-name`.
3. Commit your changes: `git commit -m "Add new feature"`.
4. Push to the branch: `git push origin feature-name`.
5. Open a pull request.

---

### License

This project is licensed under the MIT License. See the LICENSE file for details.
