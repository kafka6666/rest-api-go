# REST API in Go

This project is a REST API built using the Go programming language. It provides
a set of endpoints to perform CRUD operations.

## Features

- Create, Read, Update, and Delete operations
- JSON-based API
- Simple and easy to use

## Requirements

- Go 1.16 or higher
- A modern web browser or a tool like `curl` for testing

## Installation

1. Clone the repository:

```sh
git clone https://github.com/yourusername/rest-api-go.git
```

2. Navigate to the project directory:

```sh
cd rest-api-go
```

3. Install dependencies:

```sh
go mod tidy
```

## Usage

1. Run the server:

```sh
go run main.go
```

2. The server will start on `http://localhost:8080`.

## Endpoints

- `GET /items` - Retrieve all items
- `GET /items/{id}` - Retrieve a specific item by ID
- `POST /items` - Create a new item
- `PUT /items/{id}` - Update an existing item by ID
- `DELETE /items/{id}` - Delete an item by ID

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file
for details.

## Contact

For any inquiries, please contact me at kafihrvrd100@gmail.com
