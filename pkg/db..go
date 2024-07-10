package pkg

import (
	"fmt"

	c "reservation_service/configs"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sqlx.DB, error) {
	config := c.Load()
	conn := fmt.Sprintf(`host = %s port = %d user = %s dbname = %s password = %s sslmode = disable`,
		config.DB_HOST, config.DB_PORT, config.DB_USER, config.DB_NAME, config.DB_PASSWORD)

	db, err := sqlx.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
