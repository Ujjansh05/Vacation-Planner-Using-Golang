# VacationPlanner

A small Go API that generates vacation itineraries through `langchaingo` using OpenAI-compatible providers (OpenAI and OpenRouter).

The app exposes two endpoints:
- `POST /Vacation/create` to start generating a vacation idea.
- `GET /Vacation/:id` to fetch generation status and result.

Generation runs asynchronously in a goroutine, so you create first and then poll by ID.

## Tech Stack

- Go
- Gin (`github.com/gin-gonic/gin`) for HTTP routing
- LangChainGo (`github.com/tmc/langchaingo`) with OpenAI provider
- UUIDs via `github.com/google/uuid`

## Project Structure

```text
VacationPlanner/
|-- main.go
|-- go.mod
|-- go.sum
|-- .gitignore
|-- chains/
|   |-- structs.go
|   `-- vacations.go
`-- routes/
    |-- structs.go
    `-- vacations.go
```

## How It Works

1. Client sends vacation preferences to `POST /Vacation/create`.
2. API returns a generated UUID immediately with `completed: false`.
3. Background goroutine calls an OpenAI-compatible provider using `langchaingo`.
4. Generated itinerary is saved in memory and marked `completed: true`.
5. Client polls `GET /Vacation/:id` until completed.

## Prerequisites

- Go installed (project `go.mod` declares `go 1.25.0`)
- A valid API key with available quota:
  - OpenAI key, or
  - OpenRouter key

## Run Locally

From project directory:

```bat
cd C:\Users\ujjan\Music\Go\VacationPlanner
set "OPENAI_API_KEY=YOUR_OPENAI_KEY"
go run main.go
```

OpenRouter test example:

```bat
cd C:\Users\ujjan\Music\Go\VacationPlanner
set "OPENROUTER_API_KEY=YOUR_OPENROUTER_KEY"
set "OPENROUTER_MODEL=openai/gpt-4o-mini"
go run main.go
```

Server starts at:

```text
http://localhost:8080
```

## API Usage

### 1. Create Vacation Idea

Request:

```http
POST /Vacation/create
Content-Type: application/json
```

Body:

```json
{
  "favourite_season": "summer",
  "hobbies": ["hiking", "photography"],
  "budget": 1500
}
```

`curl` (cmd.exe):

```bat
curl -X POST http://localhost:8080/Vacation/create -H "Content-Type: application/json" -d "{\"favourite_season\":\"summer\",\"hobbies\":[\"hiking\",\"photography\"],\"budget\":1500}"
```

Example response:

```json
{
  "id": "c1e0f8d0-48a3-4b2e-a4ee-25dcf58c69d0",
  "completed": false
}
```

### 2. Poll Result by ID

Request:

```http
GET /Vacation/:id
```

`curl`:

```bat
curl http://localhost:8080/Vacation/c1e0f8d0-48a3-4b2e-a4ee-25dcf58c69d0
```

Possible responses:

- While processing:

```json
{
  "id": "c1e0f8d0-48a3-4b2e-a4ee-25dcf58c69d0",
  "completed": false,
  "idea": ""
}
```

- After completion:

```json
{
  "id": "c1e0f8d0-48a3-4b2e-a4ee-25dcf58c69d0",
  "completed": true,
  "idea": "..."
}
```

## Environment Variables

- `OPENROUTER_API_KEY` (optional, preferred for OpenRouter): if set, the app routes requests through OpenRouter.
- `OPENROUTER_MODEL` (optional): defaults to `openai/gpt-4o-mini`.
- `OPENROUTER_BASE_URL` (optional): defaults to `https://openrouter.ai/api/v1`.
- `OPENAI_API_KEY` (optional fallback): used when `OPENROUTER_API_KEY` is not set.
- `OPENAI_MODEL` (optional): defaults to `gpt-4o-mini`.
- `PORT` (optional): Gin defaults to `8080` when not set.

## Security Notes

- Never commit API keys.
- Keep keys in environment variables or untracked local files like `.env`.
- If a key is exposed, rotate/revoke it immediately.
- `.gitignore` is configured to exclude common secret and local artifact files.

## Troubleshooting

- `OPENROUTER_API_KEY or OPENAI_API_KEY is not set`
  - Set one of those variables in the same terminal session before `go run`.
- `OpenAI generation error: ... 429 ... exceeded your current quota`
  - Billing/quota issue in OpenAI account or project.
  - Add credits/payment method, confirm project budget, then use a valid key.
- `completed` stays `false`
  - Check server logs right after `POST /Vacation/create` for the real generation error.

## Current Limitations

- Data is stored in memory (`Vacations` slice) and is lost on restart.
- No persistence/database layer.
- No synchronization for concurrent writes to shared in-memory slice.
- No auth/rate limiting for public endpoints.

## Suggested Next Improvements

1. Add persistent storage (PostgreSQL or SQLite).
2. Add mutex protection around shared in-memory data if kept.
3. Add structured error field in API response for failed generation jobs.
4. Add tests for routes and chain logic.
5. Add Dockerfile and `.env.example` for easier setup.
