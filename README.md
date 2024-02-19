# Book Order API

This repository contains the source code for a Book Order API implemented in Golang using the Mux router and Postgres database, Dockerized for easy deployment. This API allows users to manage book orders, including creating new orders, updating existing orders, deleting orders, and retrieving order information.

## Technologies Used
- Golang: A statically typed, compiled programming language designed for building simple, reliable, and efficient software.
- Mux: A powerful HTTP router for Go, enabling elegant and straightforward routing of HTTP requests.
- Postgres: An advanced open-source relational database known for its reliability, robustness, and performance.
- Docker: A containerization platform that simplifies the process of deploying applications by packaging them with their dependencies into containers.

## Prerequisites
Before running the API locally, ensure you have the following installed:
- Golang: [Installation Guide](https://golang.org/doc/install)
- Docker: [Installation Guide](https://docs.docker.com/get-docker/)
- Postgres: [Installation Guide](https://www.postgresql.org/download/)

## Setup
1. Clone this repository to your local machine.
   ```bash
   git clone https://github.com/yourusername/book-order-api.git
   ```
2. Navigate to the project directory.
   ```bash
   cd book-order-api
   ```
3. Start the Postgres database using Docker.
   ```bash
   docker-compose up -d
   ```
4. Build and run the application.
   ```bash
   go build
   ./book-order-api
   ```
5. The API server should now be running locally on `http://localhost:8080`.


## Contribution
Contributions are welcome! If you find any bugs or have suggestions for improvement, feel free to open an issue or submit a pull request.

## License
This project is licensed under the [MIT License](LICENSE).

---

Feel free to customize this README according to your specific project details and requirements.
