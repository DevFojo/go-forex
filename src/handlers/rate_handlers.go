package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/mrfojo/go-forex/src/utils"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/mrfojo/go-forex/src/models"
)

func ShiftPath(p string) string {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return "/"
	}
	return p[i:]
}

func HandleRateRequests(w http.ResponseWriter, r *http.Request) {
	switch urlPath := ShiftPath(r.URL.Path); urlPath {
	case "/latest":
		{
			getLatestRate(w, r)
			return
		}
	default:
		{
			{
				date, err := utils.ExtractDate(urlPath)
				if err != nil {
					http.Error(w, fmt.Sprintf("The requested resource '%v' is not available.", urlPath), 400)

				} else {
					getRatesByDate(w, r, date)
				}
			}
		}
	}
}

func getLatestRate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		latestRate := rates.GetLatest()
		responseJSON, err := json.Marshal(latestRate)
		utils.ProcessError(err)
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseJSON)
	} else {
		http.Error(w, fmt.Sprintf("The requested resource does not support http method '%v'.", r.Method), 405)
	}
}

//GetByDate : Get latest rate
func getRatesByDate(w http.ResponseWriter, r *http.Request, date time.Time) {
	if r.Method == http.MethodGet {
		rates := rates.GetRatesByDate(date)
		responseJSON, err := json.Marshal(rates)
		utils.ProcessError(err)
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseJSON)
	} else {
		http.Error(w, fmt.Sprintf("The requested resource does not support http method '%v'.", r.Method), 405)
	}
}
