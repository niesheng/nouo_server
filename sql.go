package main

import (
	"encoding/json"

	_ "github.com/lib/pq"
)

func sql_router(request webRequest) (resp webResponse, status int) {
	var val string
	if db.Ping() != nil {
		resp.Body = db.Ping().Error()
		return resp, 502
	}

	param, err := json.Marshal(request)
	if err != nil {
		resp.Body = err.Error()
		return resp, 501
	}
	err = db.QueryRow("SELECT " + Config_.Postgres.Router + "('" + string(param) + "')").Scan(&val)
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
