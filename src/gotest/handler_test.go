package gotest

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/MrFojo/go-forex/src/app"
)

func TestGetLatestRateHandler(t *testing.T) { 
 
	req, err := http.NewRequest("GET", "/rates/latest", nil)
	if err != nil { 
 
		t.Errorf(err.Error())
	}
	rec := executeRequest(req)
	if !isSuccessCode(rec.Code) { 
 
		t.Errorf("Request returned failure code %v", rec.Code)
	}
	if content := rec.Body.String(); strings.Index(content, "rates") < 0 { 
 
		t.Errorf("Returned unexpeted result")
	}
}

func TestGetAnalyzedRateHandler(t *testing.T) { 
 
	req, err := http.NewRequest("GET", "/rates/analyze", nil)
	if err != nil { 
 
		t.Errorf(err.Error())
	}
	rec := executeRequest(req)
	if !isSuccessCode(rec.Code) { 
 
		t.Errorf("Request returned failure code %v", rec.Code)
	}
	if content := rec.Body.String(); strings.Index(content, "rates_analyze") < 0 { 
 
		t.Errorf("Returned unexpeted result")
	}
}

func TestGetRateByTestHandler(t *testing.T) { 
 
	req, err := http.NewRequest("GET", "/rates/2019-02-08", nil)
	if err != nil { 
 
		t.Errorf(err.Error())
	}
	rec := executeRequest(req)
	if !isSuccessCode(rec.Code) { 
 
		t.Errorf("Request returned failure code %v", rec.Code)
	}
	if content := rec.Body.String(); strings.Index(content, "rates") < 0 { 
 
		t.Errorf("Returned unexpeted result")
	}
}

func TestInvalidRoute(t *testing.T) { 
 
	req, err := http.NewRequest("GET", "/rates/", nil)
	if err != nil { 
 
		t.Errorf(err.Error())
	}
	rec := executeRequest(req)
	if rec.Code != 400 { 
 
		t.Errorf("Response code is %v. Expected 400.", rec.Code)
	}

}

func TestInvalidMethod(t *testing.T) { 
 
	req, err := http.NewRequest("POST", "/rates/2019-02-08", nil)
	if err != nil { 
 
		t.Errorf(err.Error())
	}
	rec := executeRequest(req)
	if rec.Code != 405 { 
 
		t.Errorf("Response code is %v. Expected 405.", rec.Code)
	}
}

func isSuccessCode(code int) bool { 
 
	return code >= 100 && code <= 299
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder { 
 
	rec := httptest.NewRecorder()
	app.Handler.ServeHTTP(rec, req)
	return rec
}
