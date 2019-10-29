package main

import (
	"database/sql"
	"encoding/json"

	_ "github.com/lib/pq"
)

var DB_ *sql.DB

func init_db_connect() {
	ConnectString := "user=" + Config_.Postgres.Username + " password=" + Config_.Postgres.Password + " host=" + Config_.Postgres.Server + " port=" + Config_.Postgres.Port + " dbname=" + Config_.Postgres.Database + " sslmode=disable"
	DB_, _ = sql.Open("postgres", ConnectString)
	DB_.SetConnMaxLifetime(3600000)
	DB_.SetMaxIdleConns(100)
	DB_.SetMaxOpenConns(100)
}

func sql_router(request webRequest) (resp webResponse, status int) {
	var val string
	if DB_.Ping() != nil {
		resp.Body = DB_.Ping().Error()
		return resp, 502
	}

	param, err := json.Marshal(request)
	if err != nil {
		resp.Body = err.Error()
		return resp, 501
	}
	err = DB_.QueryRow("SELECT " + Config_.Postgres.Router + "('" + string(param) + "')").Scan(&val)
	if err != nil {
		resp.Body = err.Error()
		return resp, 502
	}
	err = json.Unmarshal([]byte(val), &resp)
	if err != nil {
		resp.Body = err.Error()
		return resp, 501
	}
	return resp, 200
}
