package main

import (
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	factor := 1
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if f, err := strconv.Atoi(r.FormValue("factor")); nil == err {
			factor = f
		}
		if latency, err := strconv.Atoi(r.FormValue("latency")); nil == err {
			time.Sleep(time.Duration(factor*latency) * time.Millisecond)
		}
		payload, err := strconv.Atoi(r.FormValue("payload"))
		if nil != err {
			payload = 1
		}
		io.WriteString(w, strings.Repeat("*", payload))
	})
	http.ListenAndServe(":8080", nil)
}