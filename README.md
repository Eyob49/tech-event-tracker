# Tech Event Tracker

A CLI tool to manage and discover tech events in Addis Ababa.

## Usage

List all events:
```bash
go run . list'
```

Add an event:
```bash
go run . add "Event Name" "2024-08-25" "18:00" "Location"
```

Search events:
```bash
go run . search "keyword"
```

## Setup

```bash
go get github.com/mattn/go-sqlite3
sqlite3 tech-events.db < schema.sql
go run . list
```
