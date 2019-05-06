package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/MrFojo/go-forex/src/utils"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/MrFojo/go-forex/src/models"
)

func shiftPath(p string) string { 
 
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 { 
 
		return "/"
	}
	return p[i:]
}

const InvalidMethodError = "The requested resource does not support http method '%v'."
const InvalidResourceError = "The requested resource '%v' is not available."

func HandleRateRequests(w http.ResponseWriter, r *http.Request) { 
 
	switch urlPath := shiftPath(r.URL.Path); urlPath { 
 
	case "/latest":
		{ 
 
			getLatestRate(&w, r)
			return
		}
	case "/analyze":
		{ 
 
			getAnalyzedRate(&w, r)
			return
		}
	default:
		{ 
 
			{ 
 
				date, err := utils.ExtractDate(urlPath)
				if err != nil { 
 
					http.Error(w, fmt.Sprintf(InvalidResourceError, urlPath), 400)

				} else { 
 
					getRatesByDate(&w, r, date)
				}
			}
		}
	}
}

func getAnalyzedRate(w *http.ResponseWriter, r *http.Request) { 
 
	if r.Method == http.MethodGet { 
 
		analyzedRate := rates.GetAnalyzeRate()
		responseJSON, err := json.Marshal(analyzedRate)
		utils.ProcessError(err)
		(*w).WriteHeader(200)
		(*w).Header().Set("Content-Type", "application/json")
		(*w).Write(responseJSON)
	} else { 
 
		http.Error(*w, fmt.Sprintf(InvalidMethodError, r.Method), 405)
	}
}

func getLatestRate(w *http.ResponseWriter, r *http.Request) { 
 
	if r.Method == http.MethodGet { 
 
		latestRate := rates.GetLatest()
		responseJSON, err := json.Marshal(latestRate)
		utils.ProcessError(err)
		(*w).WriteHeader(200)
		(*w).Header().Set("Content-Type", "application/json")
		(*w).Write(responseJSON)
	} else { 
 
		http.Error(*w, fmt.Sprintf(InvalidMethodError, r.Method), 405)
	}
}

func getRatesByDate(w *http.ResponseWriter, r *http.Request, date time.Time) { 
 
	if r.Method == http.MethodGet { 
 
		rates := rates.GetRatesByDate(date)
		responseJSON, err := json.Marshal(rates)
		utils.ProcessError(err)
		(*w).WriteHeader(200)
		(*w).Header().Set("Content-Type", "application/json")
		(*w).Write(responseJSON)
	} else { 
 
		http.Error(*w, fmt.Sprintf(InvalidMethodError, r.Method), 405)
	}
}
