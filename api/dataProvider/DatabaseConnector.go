package dataprovider

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

type DatabaseConnector struct {
	Db *sql.DB
}

func (connector *DatabaseConnector) Connect() error {
	var err error
	connector.Db, err = sql.Open("sqlite", "data.db")
	if err != nil {
		return err
	}

	_, err = connector.Db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            mail TEXT PRIMARY KEY,
			secret BLOB,
			mailbox_username TEXT,
            mailbox_password BLOB,
            mailbox_server TEXT,
            mailbox_port INTEGER
        );
	`)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (connector *DatabaseConnector) Close() {

	if connector.Db != nil {
		connector.Db.Close()
	}
}

func (connector *DatabaseConnector) Get() *sql.DB {
	return connector.Db
}
