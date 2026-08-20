CREATE TABLE IF NOT EXISTS events (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  date TEXT NOT NULL,
  time TEXT NOT NULL,
  location TEXT NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);


INSERT INTO events (name, date, time, location) VALUES
('Go Study Group', '2024-08-25', '18:00', 'Abrehot'),
('Web Dev Meetup', '2024-08-27', '19:00', 'Addis Hub'),
('Backend Workshop', '2024-08-29', '17:30', 'Impact Hub');