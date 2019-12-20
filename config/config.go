package config

import "database/sql"

func GetDB() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "doadmin"
	dbPass := "flgkgp2q2mmxuumj"
	dbName := "compose"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(202.91.14.3:3313)/"+dbName)
	return
}
