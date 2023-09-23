package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type AppVersionStruct struct {
	Version string `json:"version"`
}

func Version(w http.ResponseWriter, r *http.Request) {
	v := os.Getenv("version")
	fmt.Println(v)
	vs := AppVersionStruct{Version: v}
	vjs, err := json.Marshal(vs)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(vjs)
}
