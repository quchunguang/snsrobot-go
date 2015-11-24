package main

import (
	"database/sql"
	"log"
	"strings"
)

const pgsql = `
DROP TABLE IF EXISTS users;
CREATE TABLE users(
    u_id serial NOT NULL,
    u_name character varying(100) NOT NULL,
    u_pass character varying(30) NOT NULL,
    u_group character varying(100) NOT NULL,
    u_rank double precision NOT NULL,
    u_date timestamp,
    u_ipv4 character varying(16),
    CONSTRAINT users_pkey PRIMARY KEY (u_id)
);

-- DROP INDEX IF EXISTS user_rank;
-- CREATE INDEX user_rank ON users (u_rank);

DROP TABLE IF EXISTS edges;
CREATE TABLE edges(
    e_id serial NOT NULL,
    e_source integer NOT NULL,
    e_target integer NOT NULL,
    e_weight double precision NOT NULL,
    CONSTRAINT edges_pkey PRIMARY KEY (e_id)
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
	checkErr(err)
	log.Println("Database initialized.")
}
