package entity

import "time"

type CityMsg struct {
	Id         int64   `gorm:"column:id"`
	City       string  `gorm:"column:city;type:varchar(50)"`
	CityKey    string  `gorm:"column:citykey;type:varchar(50)"`
	Parent     string  `gorm:"column:parent;type:varchar(50)"`
	UpdateTime string  `gorm:"column:update_time;type:varchar(10)"`
	Shidu      string  `gorm:"column:shidu;type:varchar(10)"`
	Pm25       float64 `gorm:"column:pm25;type:varchar(10)"`
	Pm10       float64 `gorm:"column:pm10;type:varchar(10)"`
	Quality    string  `gorm:"column:quality;type:varchar(10)"`
	Wendu      string  `gorm:"column:wendu;type:varchar(10)"`
	Ganmao     string  `gorm:"column:ganmao;type:varchar(50)"`
}

func NewCityMsg(city, citykey, parent, updateTime, shidu, quality, wendu, ganmao string, pm25, pm10 float64) *CityMsg {
	return &CityMsg{
		City:       city,
		CityKey:    citykey,
		Parent:     parent,
		UpdateTime: updateTime,
		Shidu:      shidu,
		Pm25:       pm25,
		Pm10:       pm10,
		Quality:    quality,
		Wendu:      wendu,
		Ganmao:     ganmao,
	}
}

func (CityMsg) TableName() string {
	return "city_msg"
}

type ForecastWeather struct {
	Id      int64     `gorm:"column:id"`
	CityId  int64     `gorm:"column:city_id"` //外键
	Ymd     time.Time `gorm:"column:ymd"`
	Date    string    `gorm:"column:date;type:varchar(20)"`
	High    string    `gorm:"column:high;type:varchar(10)"`
	Low     string    `gorm:"column:low;type:varchar(10)"`
	Week    string    `gorm:"column:week;type:varchar(10)"`
	Sunrise string    `gorm:"column:sunrise;type:varchar(20)"`
	Sunset  string    `gorm:"column:sunset;type:varchar(20)"`
	Aqi     int       `gorm:"column:aqi"`
	Fx      string    `gorm:"column:fx;type:varchar(10)"`
	Fl      string    `gorm:"column:fl;type:varchar(10)"`
	Type    string    `gorm:"column:type;type:varchar(10)"`
	Notice  string    `gorm:"column:notice;type:varchar(50)"`
}

func NewForecastWeather(ymd string, date, high, low, week, sunrise, sunset string, aqi int, fx, fl, weatherType, notice string) *ForecastWeather {
	t, err := time.Parse("2006-01-02", ymd)
	if err != nil {
		panic(err)
	}
	return &ForecastWeather{
		Ymd:     t,
		Date:    date,
		High:    high,
		Low:     low,
		Week:    week,
		Sunrise: sunrise,
		Sunset:  sunset,
		Aqi:     aqi,
		Fx:      fx,
		Fl:      fl,
		Type:    weatherType,
		Notice:  notice,
	}
}

func (ForecastWeather) TableName() string {
	return "forecast_weather"
}
