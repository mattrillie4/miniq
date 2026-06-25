# miniq

`miniq` is a small command-line job queue written in Go.

The goal of the project is to build a simple, understandable queue from the ground up using SQLite for storage. It is intentionally small, but it touches on useful backend concepts such as command-line tooling, persistence, job states, retries, and worker-style processing.

This project is currently in early development and is being built step by step as a learning project.

## Why I Built This

I wanted to understand how background job queues work beneath the surface, instead of only using a finished library or framework.

`miniq` is my attempt to break the idea down into a smaller system:

- a CLI for interacting with the queue
- a SQLite database for storing jobs
- a jobs table with queue-related fields
- commands for pushing and inspecting jobs
- future worker commands for claiming, running, retrying, and completing jobs

## Current Features

- Initializes a local SQLite database
- Creates a `jobs` table using a migration
- Pushes jobs onto the queue with a type and payload
- Uses Go's standard `flag` package for command options
- Keeps database setup code separate from job logic

## Tech Stack

- Go
- SQLite
- `database/sql`
- `modernc.org/sqlite`

## Project Structure

```text
miniq/
├── main.go
├── go.mod
├── go.sum
├── internal/
│   ├── db/
│   │   ├── db.go
│   │   └── migrations.go
│   └── jobs/
│       └── jobs.go
└── README.md
```

## Usage

Initialize the database:

```powershell
go run . init
```

Push a job onto the queue:

```powershell
go run . push -type "email" -payload "test@email"
```

Show help:

```powershell
go run . help
```

## Example Job Data

Jobs are stored in SQLite with fields such as:

- `type`
- `payload`
- `status`
- `attempts`
- `max_attempts`
- `locked_by`
- `locked_at`
- `run_after`
- `created_at`
- `updated_at`

This gives the project a foundation for queue features like delayed jobs, retries, and worker locking.

## Roadmap

The next planned features are:

- List jobs from the queue
- Claim the next available queued job
- Mark jobs as completed
- Mark jobs as failed
- Retry failed jobs up to a maximum attempt count
- Add a worker loop
- Improve command help output
- Add tests around job creation and queue behavior

## What I Am Learning

This project is helping me practise:

- structuring a small Go application
- using packages to separate responsibilities
- working with SQLite from Go
- writing command-line interfaces
- modelling job queue state
- building software incrementally

## Status

`miniq` is a work in progress. The current focus is building the core queue operations clearly before adding more advanced features.
