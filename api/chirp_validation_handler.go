package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/klausks/bootdev-chirpy/internal/webutils"
)

var profanities = map[string]bool{
	"kerfuffle": true,
	"sharbert":  true,
	"fornax":    true,
}

func HandleValidateChirpReq(rw http.ResponseWriter, req *http.Request) {
	type body struct {
		Body string `json:"body"`
	}

	type validResponse struct {
		CleanedBody string `json:"cleaned_body"`
	}

	reqBody := body{}
	if err := json.NewDecoder(req.Body).Decode(&reqBody); err != nil {
		webutils.RespondWithError(rw, 500, "Something went wrong")
		return
	}

	if len(reqBody.Body) > 140 {
		webutils.RespondWithError(rw, 400, "Chirp is too long")
		return
	}
	webutils.RespondWithJSON(rw, 200, validResponse{CleanedBody: removeProfanity(reqBody.Body)})
}

func removeProfanity(text string) string {
	strings.Split(text, "")
	words := strings.Split(text, " ")
	for i, word := range words {
		if _, exists := profanities[strings.ToLower(word)]; exists {
			words[i] = "****"
		}
	}
	return strings.Join(words, " ")
}
