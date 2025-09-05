# HeliTodo
**A minimal Todo service with PostgreSQL, migrations, HTTP API, and unit tests.**

## Configuration
Edit `configs/config.json` if needed. Defaults should work with the provided `docker-compose.yml`.

## Run 
```bash
make run
```
This runs `docker-compose up --build`, starting PostgreSQL and the app.


## Database migrations
Migration files live in `migrations/` and are applied automatically on container startup (if your entrypoint handles it). If you prefer manual application, you can exec into the app container and run your migration tool there.

Files included:
- `migrations/0001_create_todo_items.up.sql`
- `migrations/0001_create_todo_items.down.sql`


## Tests
Run all unit tests:
```bash
make test
```

## Project layout
- `internal/todo/adapters/postgres/`: PostgreSQL repository
- `internal/todo/adapters/http/`: HTTP handlers and routes
- `internal/todo/usecase/`: Core application logic and tests
- `migrations/`: SQL migration files
- `pkg/`: shared packages (config, server)

## Architecture: Hexagonal (Ports & Adapters)
This project follows Hexagonal Architecture to keep domain logic independent from frameworks and external systems.

- Domain (`internal/todo/domain/`):
  - Core business entities and rules. No framework or I/O dependencies.

- Ports (`internal/todo/ports/`):
  - Interfaces that define what the domain needs from the outside world (e.g., repositories, validators). These are implemented by adapters.

- Use Cases (`internal/todo/usecase/`):
  - Application services orchestrating domain operations via ports. Contains business workflows and unit tests.

- Adapters (`internal/todo/adapters/`):
  - Implementations of ports and entrypoints:
    - `postgres/`: data persistence adapter implementing repository ports.
    - `http/`: transport adapter exposing HTTP endpoints to the outside.
    - `validator/`: input validation adapter.

## Notes
- Unit tests mock or fake external services so tests do not depend on a live database.
- See `internal/todo/adapters/postgres/repository_test.go` and `internal/todo/usecase/service_test.go` for examples.


## API endpoints
- Base URL: `http://localhost:8080/api/v1`

- Create Todo
  - Method: `POST`
  - Path: `/todos`
  - Request JSON:
    ```json
    {
      "description": "Write tests",
      "due_date": "2025-12-25T10:00:00Z"
    }
    ```
  - Response 201 JSON:
    ```json
    {
      "id": "<uuid-or-id>",
      "description": "Write tests",
      "due_date": "2025-12-25T10:00:00Z"
    }
    ```

Example curl
```bash
curl -sS -X POST http://localhost:8080/api/v1/todos \
  -H 'Content-Type: application/json' \
  -d '{"description":"Write tests","due_date":"2030-01-01T10:00:00Z"}' | jq
```