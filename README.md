# Project: Central Finance API

This README provides step-by-step instructions to get the Central Finance API up and running using Docker Compose.

---

## 🔍 Overview

`Central Finance API` is a Go-based backend application that connects to a PostgreSQL database. We use Docker Compose to simplify running both the API service and its database.

---

## 🛠 Prerequisites

Before you begin, ensure you have the following installed:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- (Optional) [Make](https://www.gnu.org/software/make/) if you prefer `make` shortcuts

---

## ⚙️ Environment Variables

Create a `.env` file in the project root with the following variables:

```dotenv
# .env
DATABASE_URL=postgres://<DB_USER>:<DB_PASSWORD>@db:5432/<DB_NAME>?sslmode=disable
GO_PORT=8080
```

- Replace `<DB_USER>`, `<DB_PASSWORD>`, and `<DB_NAME>` with your credentials.
- The hostname `db` matches the service name in `docker-compose.yml`.

---

## 📂 Project Structure

```
Central-Finance/
├── cmd/
│   └── server/
│       └── main.go
├── db/
│   └── postgres.go
├── migrations/
│   └── 0001_init.up.sql
├── docker-compose.yml
├── .env.example
└── go.mod
```

---

## 🚀 Running with Docker Compose

1. **Copy the example env file**

   ```bash
   cp .env.example .env
   # then edit .env with your real credentials
   ```

2. **Build and start the services**

   ```bash
   docker-compose up --build -d
   ```

   - `--build` forces a rebuild of the Go service image.
   - `-d` runs containers in detached mode.

3. **Verify services are running**

   ```bash
   docker-compose ps
   ```

   You should see two services: `central-finance-db` (Postgres) and `central-finance-api` (Go server).

4. **Run database migrations**

   If you have migration scripts in `migrations/`, run:

   ```bash
   docker-compose exec central-finance-api migrate -path /app/migrations -database "$DATABASE_URL" up
   ```

   > Adjust the `migrate` command if you use a different CLI or path.

5. **Check logs (optional)**

   ```bash
   docker-compose logs -f central-finance-api
   ```

   This will tail the API logs to your terminal.

6. **Test the API**

   By default, the API listens on port `8080`. You can hit the health endpoint:

   ```bash
   curl http://localhost:8080/health
   ```

   You should receive a JSON response confirming service health.

---

## 🛑 Stopping & Cleaning Up

- To stop services:

  ```bash
  docker-compose down
  ```

- To remove volumes (e.g., to reset the database):

  ```bash
  docker-compose down -v
  ```

---

## 📚 Further Reading

- [Docker Compose Reference](https://docs.docker.com/compose/)
- [Go Database SQL Package](https://pkg.go.dev/database/sql)
- [Migrate CLI Docs](https://github.com/golang-migrate/migrate)

---

