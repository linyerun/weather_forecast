package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"weather_forecast/db"
	"weather_forecast/entity"
)

func main() {
	db.Init()                 // 初始化数据库
	respMsg, err := GetData() // 获取远程Json数据
	if err != nil {
		panic(err)
	}
	saveRespMsg(respMsg) //保存数据
}

func saveRespMsg(respMsg *RespMsg) {
	// 保存城市信息
	cityInfo := respMsg.CityInfo
	data := respMsg.Data
	cityMsg := entity.NewCityMsg(cityInfo.City, cityInfo.Citykey, cityInfo.Parent, cityInfo.UpdateTime, data.Shidu, data.Quality, data.Wendu, data.Ganmao, data.Pm25, data.Pm10)
	weathers := make([]*entity.ForecastWeather, 0, len(respMsg.Data.Forecast)+1)

	// 保存天气信息
	item := respMsg.Data.Yesterday
	weather := entity.NewForecastWeather(item.Ymd, item.Date, item.High, item.Low, item.Week, item.Sunrise, item.Sunset, item.Aqi, item.Fx, item.Fl, item.Type, item.Notice)
	weathers = append(weathers, weather)
	for _, item := range respMsg.Data.Forecast {
		weather := entity.NewForecastWeather(item.Ymd, item.Date, item.High, item.Low, item.Week, item.Sunrise, item.Sunset, item.Aqi, item.Fx, item.Fl, item.Type, item.Notice)
		weathers = append(weathers, weather)
	}

	// 异步打印信息
	ch := make(chan struct{}, 1)
	go func() {
		for _, item := range respMsg.Data.Forecast {
			fmt.Printf("date: %s, low: %s, high: %s\n", item.Ymd, item.Low, item.High)
		}
		ch <- struct{}{}
	}()

	// 保存到数据库
	db.InsertCityMsgAndForecastWeathers(cityMsg, weathers)

	// 等待异步信息打印完
	<-ch
}

func GetData() (respMsg *RespMsg, err error) {
	// 获取数据
	resp, err := http.Get("http://t.weather.sojson.com/api/weather/city/101030100")
	if err != nil {
		return nil, err
	}
	defer func() {
		err = resp.Body.Close()
	}()

	// 读取数据
	respMsg = new(RespMsg)
	if err = json.NewDecoder(resp.Body).Decode(respMsg); err != nil {
		return nil, err
	}

	return
}

type RespMsg struct {
	Message  string `json:"message"`
	Status   int    `json:"status"`
	Date     string `json:"date"`
	Time     string `json:"time"`
	CityInfo struct {
		City       string `json:"city"`
		Citykey    string `json:"citykey"`
		Parent     string `json:"parent"`
		UpdateTime string `json:"updateTime"`
	} `json:"cityInfo"`
	Data struct {
		Shidu    string  `json:"shidu"`
		Pm25     float64 `json:"pm25"`
		Pm10     float64 `json:"pm10"`
		Quality  string  `json:"quality"`
		Wendu    string  `json:"wendu"`
		Ganmao   string  `json:"ganmao"`
		Forecast []struct {
			Ymd     string `json:"ymd"`
			Date    string `json:"date"`
			High    string `json:"high"`
			Low     string `json:"low"`
			Week    string `json:"week"`
			Sunrise string `json:"sunrise"`
			Sunset  string `json:"sunset"`
			Aqi     int    `json:"aqi"`
			Fx      string `json:"fx"`
			Fl      string `json:"fl"`
			Type    string `json:"type"`
			Notice  string `json:"notice"`
		} `json:"forecast"`
		Yesterday struct {
			Date    string `json:"date"`
			High    string `json:"high"`
			Low     string `json:"low"`
			Ymd     string `json:"ymd"`
			Week    string `json:"week"`
			Sunrise string `json:"sunrise"`
			Sunset  string `json:"sunset"`
			Aqi     int    `json:"aqi"`
			Fx      string `json:"fx"`
			Fl      string `json:"fl"`
			Type    string `json:"type"`
			Notice  string `json:"notice"`
		} `json:"yesterday"`
	} `json:"data"`
}
