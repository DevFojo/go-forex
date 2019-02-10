package rates

import (
	"fmt"
	"github.com/mrfojo/go-forex/src/database"
	"github.com/mrfojo/go-forex/src/utils"
	"time"
)

type LatestRate struct {
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}

const latestRateQuery = "SELECT currency, rate FROM rates WHERE date = (SELECT date FROM rates ORDER BY  date DESC LIMIT 1)  ORDER BY currency"

func GetLatest() LatestRate {

	rows, err := database.Db.Query(latestRateQuery)
	utils.ProcessError(err)
	defer rows.Close()

	rates := make(map[string]float64, 1)

	for rows.Next() {
		var (
			currency     string
			currencyRate float64
		)
		err := rows.Scan(&currency, &currencyRate)
		utils.ProcessError(err)

		if _, exists := rates[currency]; !exists {
			rates[currency] = currencyRate
		}
	}
	return LatestRate{
		Base:  "EUR",
		Rates: rates,
	}
}

const getRateByDateQuery = "SELECT currency, rate FROM rates WHERE DATE (date) = '%v' ORDER BY currency"

func GetRatesByDate(date time.Time) LatestRate {

	rows, err := database.Db.Query(fmt.Sprintf(getRateByDateQuery, date.Format(utils.TimeLayout)))
	utils.ProcessError(err)
	defer rows.Close()

	rates := make(map[string]float64, 1)

	for rows.Next() {
		var (
			currency     string
			currencyRate float64
		)
		err := rows.Scan(&currency, &currencyRate)
		utils.ProcessError(err)

		if _, exists := rates[currency]; !exists {
			rates[currency] = currencyRate
		}
	}
	return LatestRate{
		Base:  "EUR",
		Rates: rates,
	}
}
