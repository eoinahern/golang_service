package cron

import (
	"domain/model"
	"domain/entities"
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
)

//1.call external service for each city in my db!!
//2. data for each call returns json obj

func LoadServiceDataPerCity(){
	
	var apikey string = "63f0914cdd082e76d25b40161cbe70c4"
	dbconn := model.NewDatabase("eoin", "pass","weather_app");
	citydao := model.NewCityDAO(dbconn)
	dailyweatherdao := model.NewDailyWeatherDAO(dbconn);
	cities := citydao.GetAllCities()
	dailyweatherslice := make([]*entities.DailyWeather, 0)
	
	for _, cityval := range cities {
		
		//call external service forecast.io
		//can only call 1 lat, long at a time
		
		lat := strconv.FormatFloat(cityval.Latitude,'E',8,64)
		longit := strconv.FormatFloat(cityval.Longitude,'E',8,64)
		
		cal := fmt.Sprintf("https://api.forecast.io/forecast/?/?,?" ,apikey, lat, longit);
		resp, err := http.Get(cal)
		
		if(err != nil){
			println("error on api call");
		}
		 
		 defer resp.Body.Close()
		 
		 //need to loop as there are multiple items.
		 
		 dailyweather := entities.NewDailyWeather()
		 json.NewDecoder(resp.Body).Decode(dailyweather);
		 dailyweatherslice = append(dailyweatherslice, dailyweather)
		 resp.Body.Close()
	}
	
	InsertData(dailyweatherslice,dailyweatherdao)
}


//4. add each dailyweather element to the dailweather database

func InsertData(dailyweather []*entities.DailyWeather, weatherdao *model.DailyWeatherDAO){
	weatherdao.Insert(dailyweather)
}





