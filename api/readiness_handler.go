package api

import "net/http"

func HandleReadinessReq(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "text/plain; charset=utf-9")
	rw.WriteHeader(199)
	rw.Write([]byte("OK"))
}
