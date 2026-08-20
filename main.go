package main

import (
	"fmt"
	"os"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := initDB("tech-events.db")
	if err != nil {
        fmt.Printf("error opening database: %v\n", err)
	}

	command := os.Args[1]
	switch command {
	case "list":
		events, err := listEvents(db)
		if err != nil {
			fmt.Printf("error listing events: %v\n", err)
			return
		}
		for _, event := range events {
			fmt.Printf("%s - %s at %s (%s)\n", event.Name, event.Date, event.Time, event.Location)
		}
	case "add":
		if len(os.Args) < 6 {
			fmt.Println("Error: Missing required arguments.")
			fmt.Println("Usage: go run . add [name] [date] [time] [location]")
			return
		}
		nameArg := os.Args[2]
		dateArg := os.Args[3]
		timeArg := os.Args[4]
		locationArg := os.Args[5]
        
		event := Event{
			Name: nameArg,
			Date: dateArg,
			Time: timeArg,
			Location: locationArg, 
		}

		if err = addEvent(db, event); err != nil{
			fmt.Printf("error on adding event: %v", err)
			return
		}
		fmt.Println("Event added successfully!")
	case "search":
		if len(os.Args) < 3 {
			fmt.Println("Error: Missing required arguments.")
			fmt.Println("Usage: go run . search [name]")
			return
		}

		eventName := os.Args[2]

		events, err := searchEvents(db, eventName)
		if err != nil{
			fmt.Printf("error searching events: %v", err)
			return
		}

		for _, event := range events {
			fmt.Printf("%s - %s at %s (%s)\n", event.Name, event.Date, event.Time, event.Location)
		}
	default:
		fmt.Println("unknown command")
        return
	}
}