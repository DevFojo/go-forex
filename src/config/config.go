package config

var HistoricalRateURL string
var DatabasePath string

func init() {
	DatabasePath = "../../dev.db"
	HistoricalRateURL = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml"
}
