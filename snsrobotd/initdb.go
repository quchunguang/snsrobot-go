package main

import (
	"database/sql"
	"strings"
)

const pgsql = `
DROP TABLE IF EXISTS users;
CREATE TABLE users(
    _id serial NOT NULL,
    _username character varying(100),
    _password character varying(30),
    _group character varying(100),
    _date timestamp,
    _ipv4 character varying(16),
    _rank double precision,
    CONSTRAINT users_pkey PRIMARY KEY (_id)
);

DROP INDEX IF EXISTS user_rank;
CREATE INDEX user_rank ON users (_rank);

DROP TABLE IF EXISTS edges;
CREATE TABLE edges(
    _id serial NOT NULL,
    _source integer NOT NULL,
    _target integer NOT NULL,
    _weight double precision,
    CONSTRAINT edges_pkey PRIMARY KEY (_id)
);
`

func InitDB(db *sql.DB) {
	tx, err := db.Begin()
	checkErr(err)

	defer func() {
		_ = tx.Rollback()
	}()

	sqlStmtSlice := strings.Split(pgsql, ";\n")
	for _, q := range sqlStmtSlice {
		_, err := tx.Exec(q)
		checkErr(err)
	}

	err = tx.Commit()
}
