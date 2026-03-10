# Portfolio

Personal portfolio website built with Go and [templ](https://templ.guide).

## Prerequisites

- [Go](https://go.dev/dl/) 1.22+
- [templ](https://templ.guide) CLI

### Installing templ

```bash
go install github.com/a-h/templ/cmd/templ@latest
```

Make sure `$GOPATH/bin` is in your `PATH`.

## Getting Started

```bash
# Initialize module
go mod init github.com/alexbrunn/portfolio
go mod tidy

# Generate templ files
templ generate

# Run the application
go run ./cmd/web
```

## Project Structure

```
├── cmd/web/          Entry point
├── internal/
│   ├── app/          Application logic
│   ├── config/       Configuration
│   └── web/          Routes and server
└── ui/
    ├── html/         templ templates
    └── static/       CSS, JS, images
```
