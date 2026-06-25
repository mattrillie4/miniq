package jobs

import "database/sql"

type Job struct {
	ID			int 
	Type 		string
	Payload		string
	Status 		string
	Attempts	int
	CreatedAt 	string
}

func Create(database *sql.DB, jobType string, payload string) error {
	query := `
	INSERT INTO jobs (
	type,
	payload,
	status
	) VALUES (
	?,
	?,
	? 
	);`

	_, err := database.Exec(query, jobType, payload, "queued")
	return err
}
