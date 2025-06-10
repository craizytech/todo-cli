# 📝 Todo CLI App in Go

> A beginner-to-advanced CLI Todo app built in Go to simulate real-world development, from writing clean code to testing, CI/CD, and persistence.

---

## 🚀 Project Overview

This project is a **Command Line Interface (CLI) Todo Application** written in **Go (Golang)**. It was created as part of my learning journey to:

- Master the Go programming language
- Understand modular and layered application design
- Practice writing idiomatic, maintainable, and testable code
- Set up CI/CD automation for tests
- Learn data persistence using file and databases

---

## 📌 Goals

| Stage | Goal |
| --- | --- |
| ✅ Phase 1 | In-memory CLI app with clean separation of concerns |
| ✅ Phase 2 | Write unit tests for domain and business logic |
| ⏳ Phase 3 | Setup GitHub Actions for CI/CD pipeline |
| ⏳ Phase 4 | Persist tasks to a local JSON file |
| ⏳ Phase 5 | Use a relational database (e.g., PostgreSQL) for data storage |
| ⏳ Phase 6 | Introduce features like due dates, task categories, etc. |
| ⏳ Phase 7 | Improve UX (CLI experience), error handling, and observability |

---

## 🗂️ Project Structure

```
todo-cli/
├── cmd/              # CLI entry point (main.go)
│   └── main.go
├── store/            # In-memory or file/database-based task storage
│   └── memory.go
├── task/             # Domain and business logic
│   ├── model.go      # Task struct and methods
│   └── manager.go    # Business logic (Add/Delete/Complete/List)
├── ui/               # CLI interaction logic (printing, input)
│   └── cli.go
├── go.mod            # Module definition
├── go.sum            # Module checksums
└── README.md
```

---

## 🧠 Philosophy

This project adheres to:

- **Clean Architecture**: Core logic is independent of frameworks and UI
- **Separation of Concerns**: Each package has a single responsibility
- **Testability First**: Code is written to be testable with clear interface boundaries
- **Discipline in Craftsmanship**: Focus on clean, readable, and purposeful code

---

## 🛠️ Features

- Add a new task
- List all current tasks
- Mark a task as complete
- Delete a task
- In-memory storage with plans to expand to file and database
- Designed with interfaces to allow extensibility

---

## ✅ Prerequisites

- Go 1.21+ installed on your system
- Terminal access (macOS, Linux, or Windows)
- Basic understanding of terminal navigation and `go` commands

---

## 🚀 Running the App

### 1. Clone the repository

```bash
git clone https://github.com/your-username/todo-cli.git
cd todo-cli
```

### 2. Run the application

```bash
go run ./cmd/main.go
```

You’ll see a CLI menu to add, complete, delete, or list tasks.

---

## 🧪 Running Tests

Unit tests are written using Go’s built-in testing package.

### Run All Tests

```bash
go test ./... -v
```

### Run with Coverage

```bash
go test ./... -cover
```

### Generate HTML Coverage Report

```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

This opens an HTML report showing how well your code is covered by tests.

---

## 🏗️ Build Executable

You can build a static binary to run your CLI app as a native executable:

```bash
go build -o todo-cli ./cmd/main.go
./todo-cli
```

---

## 📡 Planned CI/CD Integration

In upcoming stages, the app will include:

- ✅ Linting using `go vet` and `golangci-lint`
- ✅ Test execution with `go test`
- 🔄 GitHub Actions workflow:
  - On every push and pull request
  - Runs tests and reports status
  - Ensures `main` branch is always green

---

## 🧱 Architecture Summary

| Layer | Description |
| --- | --- |
| `task` | Core domain (task struct, behavior) |
| `manager` | Business rules (add, mark done, delete) |
| `store` | Interface + in-memory storage |
| `ui/cli` | User interface layer (input/output) |
| `cmd` | CLI entry point (main.go) |

---

## 💡 Engineering Habits Practiced

This project is used to reinforce:

- Interface-driven design
- Layered architecture
- Test-driven mindset
- Writing maintainable Go code
- Building CI/CD pipelines
- Using meaningful error handling and logging
- Commenting and documenting responsibly

---

## 🧾 Commands Cheatsheet

| Command | Purpose |
| --- | --- |
| `go run ./cmd/main.go` | Run the application |
| `go test ./...` | Run unit tests |
| `go test ./... -cover` | Check code coverage |
| `go build -o todo-cli ./cmd/main.go` | Compile binary |
| `go vet ./...` | Check for code issues |
| `golangci-lint run` | Lint code (if configured) |

---

## 🙋 About Me

- 👨‍💻 **Name**: Eammon
- 🎓 **Background**: Computer Science graduate from Kenya
- 💡 **Interests**: Software engineering, clean code, cloud, and security
- 🎯 **Goal**: Become a world-class backend/cloud engineer

---

## 🤝 Contributions

This project is currently personal, but for practice:

1. Fork the repo
2. Create your branch: `git checkout -b feature/foo`
3. Commit changes: `git commit -m "Add foo feature"`
4. Push: `git push origin feature/foo`
5. Open a Pull Request

CI will automatically test your changes.

---

## ⚖️ License

This project is for educational purposes. You’re welcome to fork, study, or contribute, but no license is applied.

---

## 📎 Useful Links

- Go Documentation
- Effective Go
- Go CLI Tips
- Clean Architecture (Uncle Bob)
- golangci-lint

---

## 🙇‍♂️ Acknowledgements

Learning using ChatGPT
