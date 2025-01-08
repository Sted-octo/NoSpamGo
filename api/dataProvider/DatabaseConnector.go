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
			mailbox_password_nonce BLOB,
			mailbox_server TEXT,
			mailbox_port INTEGER
		);

		CREATE TABLE IF NOT EXISTS filters (
			mail TEXT,
			filter_name TEXT,
			filter_number_of_spam_detected INTEGER,
			PRIMARY KEY (mail, filter_name)
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
