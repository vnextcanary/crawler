// Copyright 2019 vnextcanary Core Team.  All rights reserved.
// LICENSE: Use of this source code is governed by AGPL-3.0.
// license that can be found in the LICENSE file.
package configs

import (
	"database/sql"

	"github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB // global variable to share it between main and the HTTP handler
var err error

func Initdb() {

	Db, err = sql.Open("mysql", DbConnStr)
	if err != nil {
		LogErr(err)
	}
	Db.SetMaxOpenConns(10000)
	Db.SetMaxIdleConns(1000)
	var version string
	Db.QueryRow("SELECT VERSION()").Scan(&version)
	logrus.Info("Connected to Mysql ok, version: " + version)
	err = Db.Ping() // This DOES open a connection if necessary. This makes sure the database is accessible
	if err != nil {
		LogErr(err)
	}

	//defer Db.Close()
}
