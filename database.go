package main

import "database/sql"

func initDB(path string) (*sql.DB, error) {
    db, err := sql.Open("sqlite3", path)
    if err != nil {
       return nil, err
    }

	// Test the connection
	if err = db.Ping(); err != nil{
		return nil, err
	}

	return db, nil
}

func listEvents(db *sql.DB) ([]Event, error) {
	rows, err := db.Query("SELECT id, name, date, time, location FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next(){
		var e Event
		err := rows.Scan(&e.ID, &e.Name, &e.Date, &e.Time, &e.Location)
		if err != nil{
			return nil, err
		}
		events = append(events, e)
	}

	if err := rows.Err(); err != nil{
		return nil, err
	}

	return events, nil
}

func addEvent(db *sql.DB, event Event) error {
    query := `INSERT INTO events (name, date, time, location) VALUES (?, ?, ?, ?) RETURNING id`

	if err := db.QueryRow(query, event.Name, event.Date, event.Time, event.Location).Scan(&event.ID); err != nil{
		return err
	}

	return nil
}

func searchEvents(db *sql.DB, query string) ([]Event, error) {
    rows, err := db.Query("SELECT id, name, date, time, location FROM events WHERE name LIKE ?", "%"+query+"%")
	if err != nil {
		return nil, err
	}
    defer rows.Close()

	var events []Event
	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Name, &e.Date, &e.Time, &e.Location)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}