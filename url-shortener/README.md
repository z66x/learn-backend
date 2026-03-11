# URL Shortener

A fast, minimal URL shortener built with Go and Svelte.

🔗 **Live:** [zwix-url.up.railway.app](https://zwix-url.up.railway.app)

## Demo
![demo](./screenshots/demo.png)

## Stack
- **Backend:** Go + Chi
- **Frontend:** Svelte + TypeScript
- **Database:** PostgreSQL
- **DB Layer:** sqlx + golang-migrate

## Project Structure
```
url-shortener/
├── backend/
│   ├── cmd/api/         # entry point
│   ├── internal/
│   │   ├── db/          # database connection + repositories
│   │   ├── handlers/    # http handlers
│   │   └── models/      # data models
│   └── migrations/      # sql migrations
└── frontend/
    └── src/routes/      # svelte pages
```

## Getting Started

### Prerequisites
- Go 1.25+
- Node.js 20+
- Docker

### 1. Start the database
```bash
docker compose up -d
```

### 2. Run migrations
```bash
migrate -path backend/migrations \
  -database "postgres://postgres:postgres@localhost:5432/urlshortener?sslmode=disable" up
```

### 3. Start the backend
```bash
cd backend
cp .env.example .env
go run cmd/api/main.go
```

### 4. Start the frontend
```bash
cd frontend
npm install
npm run dev
```

Open `http://localhost:5173`

## API

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/shorten` | Shorten a URL |
| GET | `/:code` | Redirect to original URL |

### POST /shorten
```json
// request
{ "url": "https://example.com/very/long/url" }

// response
{
  "short_url": "https://zwix-url-api.up.railway.app/abc123",
  "original_url": "https://example.com/very/long/url"
}
```

## Deployment

Deployed on Railway:
- **Frontend:** [zwix-url.up.railway.app](https://zwix-url.up.railway.app)
- **Backend:** [zwix-url-api.up.railway.app](https://zwix-url-api.up.railway.app)
- **Database:** Railway managed PostgreSQL

## Roadmap
- [ ] Redis caching
- [ ] Click analytics
- [ ] Custom slugs
- [ ] Link expiry
- [ ] Auth + dashboard