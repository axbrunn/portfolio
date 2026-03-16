# Portfolio

Personal portfolio website built with Go, [Templ](https://templ.guide) and MySQL.

Inspired by the [Ardanlabs Service Starter Kit](https://github.com/ardanlabs/service).

## Tech Stack

- **Go** — backend
- **Templ** — server-side HTML templating
- **HTMX** — dynamic UI without JavaScript frameworks
- **MySQL** — database
- **Docker** — containerized database
- **golang-migrate** — database migrations

## Prerequisites

- [Go](https://go.dev/dl/) 1.22+
- [Docker](https://www.docker.com/) + Docker Compose
- [Templ CLI](https://templ.guide)
- [golang-migrate CLI](https://github.com/golang-migrate/migrate)

### Installing Templ

```bash
go install github.com/a-h/templ/cmd/templ@latest
```

### Installing golang-migrate

```bash
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Make sure `$GOPATH/bin` is in your `PATH`.

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/alexbrunn/portfolio
cd portfolio
```

### 2. Configure environment

```bash
cp .env.example .env
```

Edit `.env` to set your database credentials:

```env
DB_ROOT_PASSWORD=secret_root
DB_DATABASE=portfolio
DB_USER=portfolio_user
DB_PASSWORD=secret_password
DB_PORT=3306
DB_DSN=mysql://portfolio_user:secret_password@tcp(localhost:3306)/portfolio?parseTime=true

APP_PORT=8080
APP_ENV=development
```

### 3. Start the database

```bash
docker compose up -d
```

### 4. Run migrations

```bash
make db/migrations/up
```

### 5. Run the application

```bash
make run/web
```

The application is now running at `http://localhost:8080`.

## Makefile

| Command | Description |
|---|---|
| `make run/web` | Run the application |
| `make db/connect` | Connect to MySQL as app user |
| `make db/connect/root` | Connect to MySQL as root |
| `make db/migrations/new name=<name>` | Create new migration files |
| `make db/migrations/up` | Run all pending migrations |
| `make db/migrations/down` | Roll back migrations |

## Docker

```bash
# Start containers in the background
docker compose up -d

# Follow logs
docker compose logs -f

# Stop containers
docker compose down

# Stop and remove volumes (⚠️ deletes all data)
docker compose down -v
```

## Development

```bash
# Install dependencies
go mod tidy

# Generate templ files
templ generate

# Run the application
make run/web
```# Portfolio

Personal portfolio website built with Go, [Templ](https://templ.guide) and MySQL.

Inspired by the [Ardanlabs Service Starter Kit](https://github.com/ardanlabs/service).

## Tech Stack

- **Go** — backend
- **Templ** — server-side HTML templating
- **HTMX** — dynamic UI without JavaScript frameworks
- **MySQL** — database
- **Docker** — containerized database
- **golang-migrate** — database migrations

## Prerequisites

- [Go](https://go.dev/dl/) 1.22+
- [Docker](https://www.docker.com/) + Docker Compose
- [Templ CLI](https://templ.guide)
- [golang-migrate CLI](https://github.com/golang-migrate/migrate)

### Installing Templ

```bash
go install github.com/a-h/templ/cmd/templ@latest
```

### Installing golang-migrate

```bash
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Make sure `$GOPATH/bin` is in your `PATH`.

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/alexbrunn/portfolio
cd portfolio
```

### 2. Configure environment

```bash
cp .env.example .env
```

Edit `.env` to set your database credentials:

```env
DB_ROOT_PASSWORD=secret_root
DB_DATABASE=portfolio
DB_USER=portfolio_user
DB_PASSWORD=secret_password
DB_PORT=3306
DB_DSN=mysql://portfolio_user:secret_password@tcp(localhost:3306)/portfolio?parseTime=true

APP_PORT=8080
APP_ENV=development
```

### 3. Start the database

```bash
docker compose up -d
```

### 4. Run migrations

```bash
make db/migrations/up
```

### 5. Run the application

```bash
make run/web
```

The application is now running at `http://localhost:8080`.

## Makefile

| Command | Description |
|---|---|
| `make run/web` | Run the application |
| `make db/connect` | Connect to MySQL as app user |
| `make db/connect/root` | Connect to MySQL as root |
| `make db/migrations/new name=<name>` | Create new migration files |
| `make db/migrations/up` | Run all pending migrations |
| `make db/migrations/down` | Roll back migrations |

## Docker

```bash
# Start containers in the background
docker compose up -d

# Follow logs
docker compose logs -f

# Stop containers
docker compose down

# Stop and remove volumes (⚠️ deletes all data)
docker compose down -v

```

## Development

```bash
# Install dependencies
go mod tidy

# Generate templ files
templ generate

# Run the application
make run/web
```