package tl

import (
	"errors"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/tidwall/buntdb"
)

type Database struct {
	DB *buntdb.DB
}

func InitDB() (*Database, error) {

	dbPath, success := os.LookupEnv("TIMELOG_DB")
	if !success || dbPath == "" {
		return nil, errors.New("Export TIMELOG_DB to the database file path: `export TIMELOG_DB`")
	}

	db, err := buntdb.Open(dbPath)
	if err != nil {
		return nil, err
	}

	db.CreateIndex("task", "*", buntdb.IndexJSON("task"))
	db.CreateIndex("project", "*", buntdb.IndexJSON("project"))

	tlDB := Database{db}
	return &tlDB, nil
}

func (database *Database) NewID() string {
	id, err := uuid.NewRandom()
	if err != nil {
		log.Fatalf("Could not generate UUID %+v\n", err)
	}

	return id.String()
}
