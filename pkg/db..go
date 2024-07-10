package pkg

import (
	"database/sql"
	"fmt"

	c "Github.com/Project-2/Reservation-Service/configs"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	config := c.Load()
	conn := fmt.Sprintf(`host = %s port = %d user = %s dbname = %s password = %s sslmode = disable`,
		config.DB_HOST, config.DB_PORT, config.DB_USER, config.DB_NAME, config.DB_PASSWORD)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
