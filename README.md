# Real-Time Task Kanban Board

A lightweight Kanban-style task management application built with **Go, SQLite, Templ, and HTMX**. This project focuses on learning backend web development through Go's standard HTTP library, server-rendered interfaces, database integration, and modular application architecture.

## Features

- **Kanban Workflow:** A task board organized around three stages — To Do, In Progress, and Done.
- **User Registration:** A signup flow with password hashing using bcrypt.
- **Database Integration:** SQLite integration for persistent data storage and user management.
- **Server-Side Rendering:** HTML interfaces rendered using Templ components.
- **HTMX Integration:** A foundation for interactive web experiences without requiring a separate frontend framework.
- **Modular Architecture:** Separated server initialization, HTTP handlers, database logic, and UI components.
- **Static Asset Serving:** Dedicated static files for CSS and other frontend assets.

> **Development status:** The project is under active development. User registration and initial database integration are implemented. Login, cookie-based session management, and complete task CRUD and drag-and-drop functionality are planned next.

## Tech Stack

| Technology | Purpose |
|---|---|
| [Go](https://go.dev/) | Backend development and HTTP server |
| `net/http` | HTTP routing and request handling |
| [SQLite](https://www.sqlite.org/) | Lightweight relational database |
| [Templ](https://templ.guide/) | Type-safe, component-based HTML rendering |
| [HTMX](https://htmx.org/) | Dynamic interactions using HTML attributes |
| HTML5 & CSS3 | Structure and styling |
| [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) | Password hashing |

## Project Structure

```text
Daily-Tasks-Go-Server/
├── Main.go              # Application entry point
├── server.go            # Server initialization and HTTP handlers
├── components/          # Templ UI components and pages
├── static/
│   └── css/
│       └── style.css    # Application styles
├── go.mod               # Go module and dependencies
├── go.sum               # Dependency checksums
└── kanban.db            # Local SQLite database (generated at runtime)
```

*The structure above reflects the current project layout; additional packages and files may be introduced as development progresses.*

## Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) installed
- [Git](https://git-scm.com/) installed
- [Templ CLI](https://templ.guide/) for generating Go code from `.templ` files when required

### 1. Clone the repository

```bash
git clone https://github.com/stanlyvarghese/Daily-Tasks-Go-Server.git
cd Daily-Tasks-Go-Server
```

### 2. Download dependencies

```bash
go mod download
```

### 3. Generate Templ code

If you have modified `.templ` files or the generated Go files are missing, install the Templ CLI and generate the required code:

```bash
go install github.com/a-h/templ/cmd/templ@latest
templ generate
```

### 4. Run the application

```bash
go run .
```

Open the local address configured by the application in your browser. Check the server's startup output for the port and URL.

## Application Architecture

The application uses a straightforward server-rendered architecture:

1. **HTTP server:** Go's `net/http` handles incoming requests and routes them to the appropriate handlers.
2. **Page rendering:** Templ components generate HTML for the application interface.
3. **Database layer:** SQLite stores application data, with database initialization handled by the server.
4. **Frontend interactions:** HTMX provides a path toward updating interface elements without full-page reloads.
5. **Static assets:** CSS files are served separately to keep presentation and server logic organized.

## Development Roadmap

- [x] Initialize the Go HTTP server.
- [x] Set up Templ components and application pages.
- [x] Configure static CSS serving.
- [x] Integrate SQLite and database initialization.
- [x] Implement user registration and password hashing.
- [ ] Implement login and authentication.
- [ ] Add secure, cookie-based session management.
- [ ] Implement task creation, editing, deletion, and database persistence.
- [ ] Support moving tasks between Kanban stages.
- [ ] Add drag-and-drop task interactions.
- [ ] Improve input validation, error handling, and automated testing.

## Learning Objectives

This project serves as a practical exploration of:

- Building web servers using Go's standard library.
- Structuring a maintainable Go application.
- Rendering reusable server-side UI components.
- Connecting Go applications to a relational database.
- Handling user registration and password security.
- Implementing session-based authentication.
- Building interactive web interfaces with HTMX.

## Future Improvements

Planned improvements include complete authentication, authorization, persistent task management, responsive layouts, and automated tests for HTTP handlers and database operations.



---

**Author:** [Stanly Varghese](https://github.com/stanlyvarghese)

**Repository:** [Daily-Tasks-Go-Server](https://github.com/stanlyvarghese/Daily-Tasks-Go-Server)
