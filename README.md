# Tech Event Tracker

A CLI tool to manage and discover tech events in Addis Ababa.

## Usage

List all events:

go run . list


Add an event:

go run . add "Event Name" "2024-08-25" "18:00" "Location"


Search events:

go run . search "keyword"


## Setup

```bash
go get github.com/mattn/go-sqlite3
sqlite3 tech-events.db < schema.sql
go run . list
```