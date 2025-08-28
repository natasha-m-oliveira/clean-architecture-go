# Clean Architecture Go Template

## Purpose

This project aims to provide a template application using **Clean Architecture** in Go (Golang). It serves as a starting point for new projects, promoting separation of concerns, testability, and maintainability.

## Project Structure

The structure follows clean architecture principles, separating domain, use cases, adapters (interfaces), and infrastructure layers.

```
cmd/
internal/
  adapter/
  core/
  infrastructure/
prisma/
test/
```

## How to Run Unit Tests

To run all unit tests in the project, use:

```sh
go test ./... -coverprofile=coverage.out && ./exclude-from-code-coverage.sh && go tool cover -html=coverage.out
```

## How to Run the Project

1. **Configure environment variables**  
   Copy the `.env.example` file to `.env` and adjust as needed.

2. **Install dependencies**  
   Run:
   ```sh
   go mod tidy
   ```
3. **Synchronize your schema with the database**  
  To create the database (if it doesn't exist) and apply your schema, run:
  ```sh
  go run github.com/steebchen/prisma-client-go db push
  ```

4. **(Optional) Re-generate the Prisma client**  
  If you only need to re-generate the client, run:
  ```sh
  go run github.com/steebchen/prisma-client-go generate
  ```

5. **Start the application**  
  ```sh
  go run cmd/main.go
  ```

## Technologies Used

- Go 1.23+
- Prisma Client Go(Auto-generated query builder)
- Gin (HTTP framework)
- Viper (Configuration)
- Testify (Testing)

---

Feel free to customize this template to fit your project's