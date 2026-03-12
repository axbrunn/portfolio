# Portfolio

Personal portfolio website built with Go and [templ](https://templ.guide).

[Ardanlabs Service Starter Kit](https://github.com/ardanlabs/service)

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
