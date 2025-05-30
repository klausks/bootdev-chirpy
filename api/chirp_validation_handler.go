package api

import (
	"encoding/json"
	"net/http"

	"github.com/klausks/bootdev-chirpy/internal/webutils"
)

func HandleValidateChirpReq(rw http.ResponseWriter, req *http.Request) {
	type body struct {
		Body string `json:"body"`
	}

	type validResponse struct {
		Valid bool `json:"valid"`
	}

	reqBody := body{}
	if err := json.NewDecoder(req.Body).Decode(&reqBody); err != nil {
		webutils.RespondWithError(rw, 500, "Something went wrong")
	}

	if len(reqBody.Body) > 140 {
		webutils.RespondWithError(rw, 400, "Chirp is too long")
	} else {
		webutils.RespondWithJSON(rw, 200, validResponse{Valid: true})
	}
}
