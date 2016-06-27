package model

import (
	"domain/entities"
	"database/sql"
	"log"
	
)

func NewDailyWeatherDAO(dbconnin *Database) *DailyWeatherDAO {
	d_weather := new(DailyWeatherDAO)
	//dbconn := dbconnin
	return d_weather
	
}

type DailyWeatherDAO struct {
	dbconn Database
	
}



func (dw *DailyWeatherDAO) Insert(weatheritems []entities.DailyWeather){
	//insert objs to ddb
	
	/*for _, val := range weatheritems {
		
		
	}*/
}


func (dw *DailyWeatherDAO) Delete(){
	
}


func (dw *DailyWeatherDAO) Get(city string) ([]entities.DailyWeather){
	
	
	 rows, err := dw.dbconn.mydbconn.Query("SELECT * FROM dailyweather WHERE name = ?;", city)
	 
	 if err != nil {
		 	println( "error calling query")
			log.Fatal(err)
	 }
	 
	 defer rows.Close()
	 newrows := createJsonWeather(rows)
	 rows.Close()
	 return newrows
}

//wasnt sure how to make this more generic just yet.
//params differ etc need to be scanned into slice

func  createJsonWeather(rows *sql.Rows) ([]entities.DailyWeather){
	
	data :=  make([]entities.DailyWeather,0) 
	for rows.Next(){
		
		dailyweather := entities.NewDailyWeather()
		rows.Scan(&dailyweather.Name, &dailyweather.Icon, &dailyweather.Time,
		dailyweather.Summary)			
		data = append(data, *dailyweather)
	}
	
	if(len(data) < 1){
		log.Output(1, "no data returned")
		println("error reading from db")
	}
	  
	return data
}

func (dw *DailyWeatherDAO) Update(){
	
}



