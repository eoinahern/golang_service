package cron

import (
	"domain/model"
	"fmt"
	"strings"
	"net/http"
	"strconv"
)



//1.call external service for each city in my db!!
//2. data for each call returns json obj

func LoadServiceDataPerCity(){
	
	var apikey string = "mykey"
	dbconn := model.NewDatabase("eoin", "pass","weather_app");
	citydao := model.NewCityDAO(dbconn)
	dailyweatherdao := model.NewDailyWeatherDAO(dbconn);
	cities := citydao.GetAllCities()
	
	
	for _, cityval := range cities {
		
		//call external service forecast.io
		lat := strconv.FormatFloat(cityval.Latitude,E,8,64)
		longit := strconv.FormatFloat(cityval.Longitude,E,8,64)
		
		cal := fmt.Sprintf("https://api.forecast.io/forecast/?/?,?" ,key, lat, longit);
		resp, err := http.Get(cal)
		}
	
}



//4. add each dailyweather element to the dailweather database

func getCityData(dailyweather []DailyWeather){
	
	
	keyStrings := make([]string,  17)
	values := make([]interface{}, 0)
	
	for key, weatherval := range cities {
		
		
		 keyStrings := append(keyStrings, `(?, ?, ?,
		 ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
		
		}
	
	stmt := fmt.Sprintf(`INSERT into dailyweather  
	() VALUES %s`, strings.Join(keyStrings, ","))
	
}





