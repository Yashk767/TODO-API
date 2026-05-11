# Todo API

A simple Todo REST API built in Go using `net/http` and SQLite.

Note: The project also supports an in-memory storage implementation.

## APIs

| Method | Endpoint | Description |
|---|---|---|
| `POST` | `/todos` | Create a todo |
| `GET` | `/todos` | List todos, excluding completed todos by default |
| `GET` | `/todos?include_completed=true` | List todos including completed todos |
| `GET` | `/todos/{id}` | Get todo by ID |
| `PUT` | `/todos/{id}` | Update todo by ID |
| `DELETE` | `/todos/{id}` | Delete todo by ID |

## Run

```bash
go run cmd/api/main.go
```

### Example

##### Create Todo

```
curl -X POST http://localhost:8082/todos \
  -H "Content-Type: application/json" \
  -d '{"text":"Complete todo application","due_date":"2026-05-13T10:00:00Z"}'
```

##### Validations

- `text` is required while creating a todo.
- `text` cannot be empty or only spaces.
- `due_date` is required while creating a todo.
- `due_date` must be a valid timestamp.
- `due_date` cannot be in the past.
- Todo `Id` must be a valid positive number.
- `include_completed` must be either `true` or `false` if provided.
- During update, only provided fields are updated.


#### Notes

- Todos are stored in SQLite by default.
- Todos are sorted by due date in ascending order.
- Completed todos are excluded from the list API by default.
- Use include_completed=true to include completed todos
