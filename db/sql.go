package db

import (
	"weather_forecast/entity"
)

func InsertCityMsgAndForecastWeathers(cityMsg *entity.CityMsg, weathers []*entity.ForecastWeather) (err error) {
	tx := db.Begin()
	if err = tx.Create(cityMsg).Error; err != nil {
		tx.Rollback()
		return err
	}
	for _, weather := range weathers {
		weather.CityId = cityMsg.Id
	}
	if err = tx.CreateInBatches(weathers, len(weathers)).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return
}
