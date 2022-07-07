package controllers

import (
	"fmt"
	"github.com/foruscommunity/go-rest-api-example/pkg/utils/env"
	"net/http"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	name := env.Get("NAME")
	version := env.Get("VERSION")
	cr := env.Get("COPYRIGHT")
	year := time.Now().Year()
	response := fmt.Sprintf("%s\nVersion %s\n%d %s", name, version, year, cr)
	_, _ = fmt.Fprint(w, response)
}
