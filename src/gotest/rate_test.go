package gotest

import (
	"github.com/mrfojo/go-forex/src/models"
	"testing"
	"time"
)

func TestGetLatestRate(t *testing.T) {
	latestRate := rates.GetLatest()
	if latestRate == nil {
		t.Error("Latest rate is nil")
	}
	if len(latestRate.Rates) <= 0 {
		t.Error("Rates is empty")
	}
}
func TestGetRateByDate(t *testing.T) {
	date, _ := time.Parse("2006-01-02", "2019-02-01")
	dayRate := rates.GetRatesByDate(date)
	if dayRate == nil {
		t.Error("Day rate is nil")
	}
	if len(dayRate.Rates) <= 0 {
		t.Error("Rates is empty")
	}
}
func TestGetAnalyzedRate(t *testing.T) {
	analyzedRate := rates.GetAnalyzeRate()
	if analyzedRate == nil {
		t.Error("Analyzed rate is nil")
	}
	if len(analyzedRate.RatesAnalyses) <= 0 {
		t.Error("Analyzed Rates is empty")
	}
}
