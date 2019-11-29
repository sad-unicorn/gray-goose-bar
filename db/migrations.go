package db

import (
	"database/sql"
	"log"
)

const createSchemaVersionTable = `CREATE TABLE IF NOT EXISTS schema_version (
  app_name     VARCHAR(10) NOT NULL,
  version INT NOT NULL,
  PRIMARY KEY (app_name))
ENGINE = InnoDB;`
const initializeSchemaVersion = `INSERT IGNORE INTO schema_version SET app_name='ggbar', version=0`

func InitMigrations() {
	tx, err := globalSession.Begin()
	if err != nil {
		panic("error during migration: " + err.Error())
	}

	mustExec(tx, createSchemaVersionTable)
	mustExec(tx, initializeSchemaVersion)
	mustCommit(tx)

	tx = mustStartTransaction()
	for version := currentVersion(tx); len(versions) > version; version = tryAdvance(tx, version + 1) {
		mustCommit(tx)
		tx = mustStartTransaction()
		log.Print("Doing db migration ")
	}
}

func tryAdvance(tx *sql.Tx, toVersion int) int {
	script := versions[toVersion]
	mustExec(tx, script)
	mustExec(tx, "UPDATE schema_version SET version = version + 1 WHERE app_name='ggbar'")
	return currentVersion(tx)
}

func currentVersion(tx *sql.Tx) int {
	ver := 0
	err := tx.QueryRow("SELECT version from schema_version WHERE app_name = 'ggbar'").Scan(&ver)
	if err == sql.ErrNoRows {
		return 0
	} else if err != nil {
		panic("failed to get current schema version: " + err.Error())
	}

	return ver
}

func mustStartTransaction() (tx *sql.Tx) {
	tx, err := globalSession.Begin()
	if err != nil {
		panic("error during migration: " + err.Error())
	}
	return tx
}

func mustExec(tx *sql.Tx, query string) {
	log.Println("Executing: ", query)
	_, err := tx.Exec(	query)
	if err != nil {
		panic("exec of query '"+ query + "' failed: " + err.Error())
	}
}

func mustCommit(tx *sql.Tx) {
	err := tx.Commit()
	if err != nil {
		panic("error while committing: " + err.Error())
	}
	log.Println("Committed.")
}