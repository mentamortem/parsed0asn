package repo

import (
	"database/sql"
	rpsl "github.com/frederic-arr/rpsl-go"
	
)

func ConnectToDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
func InsertRPSLObjects(db *sql.DB, data []rpsl.Object) error {

}
