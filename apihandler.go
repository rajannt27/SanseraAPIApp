package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
	elog.Info(1, "In APIHandler")

	err := json.NewEncoder(w).Encode("Hello test123!!")
	if err != nil {
		elog.Error(1, fmt.Sprintf("Failed to write response %#v", err))
	}
}
