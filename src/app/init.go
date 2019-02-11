package app

import (
	"encoding/xml"
	"github.com/mrfojo/go-forex/src/database"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/mrfojo/go-forex/src/config"
	"github.com/mrfojo/go-forex/src/utils"
)

type dailyRate struct {
	Time string `xml:"time,attr"`
	Cube []struct {
		Currency string `xml:"currency,attr"`
		Rate     string `xml:"rate,attr"`
	} `xml:"Cube"`
}

type rateXML struct {
	XMLName xml.Name `xml:"Envelope"`
	Cube    struct {
		Cube []dailyRate `xml:"Cube"`
	} `xml:"Cube"`
}

func getInitialRates() *rateXML {
	data, err := http.Get(config.HistoricalRateURL)
	utils.ProcessError(err)

	defer data.Body.Close()
	body, err := ioutil.ReadAll(data.Body)
	utils.ProcessError(err)

	var rateXML rateXML
	xml.Unmarshal(body, &rateXML)
	return &rateXML
}

func EnsureInitializeData() {

	ensureCreateRatesTable()
	if !checkIfRatesHasRecords() {
		rates := getInitialRates()
		if len(rates.Cube.Cube) > 0 {
			saveRates(&rates.Cube.Cube)
		}
	}
}

func saveRates(dailyRates *[]dailyRate) {
	var params []interface{}
	saveRatesCommand := "INSERT INTO rates (date, currency, rate) VALUES "

	for _, r := range *dailyRates {
		var date, _ = time.Parse(utils.TimeLayout, r.Time)
		for _, c := range r.Cube {
			saveRatesCommand += " (?, ?, ?)"
			params = append(params, date.Format("2006-01-02T15:04:05.999999999"), c.Currency, c.Rate)
		}
	}

	statement, err := database.Db.Prepare(saveRatesCommand)
	utils.ProcessError(err)

	_, err = statement.Exec(params)
	utils.ProcessError(err)
}

func checkIfRatesHasRecords() bool {
	const getFirstRateRecord = "SELECT * FROM rates"

	statement, err := database.Db.Prepare(getFirstRateRecord)
	utils.ProcessError(err)

	rows, err := statement.Query()
	utils.ProcessError(err)

	if rows.Next() {
		return true
	}
	return false

}

func ensureCreateRatesTable() {
	const createDbCommand = "CREATE TABLE IF NOT EXISTS rates (id INTEGER PRIMARY KEY, date TIMESTAMP,  currency TEXT, rate FLOAT)"

	statement, err := database.Db.Prepare(createDbCommand)
	utils.ProcessError(err)

	_, err = statement.Exec()
	utils.ProcessError(err)
}
