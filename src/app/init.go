package app

import (
	"github.com/mrfojo/go-forex/src/database"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/mrfojo/go-forex/src/config"
	"github.com/mrfojo/go-forex/src/utils"
)

type RateXML struct {
	XMLName xml.Name `xml:"Envelope"`
	Cube    struct {
		Cube []struct {
			Time string `xml:"time,attr"`
			Cube []struct {
				Currency string `xml:"currency,attr"`
				Rate     string `xml:"rate,attr"`
			} `xml:"Cube"`
		} `xml:"Cube"`
	} `xml:"Cube"`
}

func InitializeData() {

	data, err := http.Get(config.HistorialRateURL)
	utils.ProcessError(err)

	defer data.Body.Close()
	body, err := ioutil.ReadAll(data.Body)
	utils.ProcessError(err)

	var rateXML RateXML
	xml.Unmarshal(body, &rateXML)

	statement, err := database.Db.Prepare("CREATE TABLE IF NOT EXISTS rates (id INTEGER PRIMARY KEY, date TIMESTAMP,  currency TEXT, rate FLOAT)")
	utils.ProcessError(err)

	_, err = statement.Exec()
	utils.ProcessError(err)

	insertRateStatement, err := database.Db.Prepare("INSERT INTO rates (date, currency, rate) VALUES (?, ?, ?)")

	for _, r := range rateXML.Cube.Cube {
		var date, _ = time.Parse("2006-01-02", r.Time)
		for _, c := range r.Cube {
			_, err = insertRateStatement.Exec(date.Format("2006-01-02T15:04:05.999999999"), c.Currency, c.Rate)
			utils.ProcessError(err)
		}
	}
}
