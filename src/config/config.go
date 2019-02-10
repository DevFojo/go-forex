package config

var HistorialRateURL string
var DatabasePath string

func init() {
	DatabasePath = "../../dev.db"
	HistorialRateURL = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml"
}
