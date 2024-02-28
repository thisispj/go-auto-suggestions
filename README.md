
# Autocomplete Feature in Go (Golang)

This repository contains the source code for implementing a modern autocomplete feature in Go (Golang) using Gin middleware and Redis cache. The autocomplete feature provides real-time suggestions to users as they type into input fields, enhancing the user experience of web applications.


![App Screenshot](https://github.com/thisispj/go-auto-suggestions/blob/main/Screenshot.png)


## Getting Started

Follow these instructions to get the project up and running on your local machine.


#### Clone the project

```bash
  git clone https://github.com/thisispj/go-auto-suggestions.git
```

#### Go to the project directory

```bash
  cd go-auto-suggestions
```

#### Initialise the Go Module

```bash
go mod init autocomplete
```

#### Install dependencies

```bash
  go mod tidy
```

#### Start the server

To start the server, run the following command in the project directory:

```bash
  go run main.go
```

The server will start running on port 8080 by default.
## Testing the Autocomplete Feature

Once the server is running, you can test the autocomplete feature by opening a web browser and navigating to the following URL:

```bash
  http://localhost:8080/
```
You should see a simple HTML page with an input field for searching banks. Start typing in the input field, and you'll notice autocomplete suggestions appearing in real-time based on the input.

## Contributing

Contributions are welcome! If you'd like to contribute to this project, please follow these guidelines:

Fork the repository.
Create a new branch for your feature or bug fix.
Make your changes and ensure that the code is well-tested.
Commit your changes with descriptive commit messages.
Push your changes to your fork.
Submit a pull request to the main branch of the original repository.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
