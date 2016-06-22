package entities

func NewDailyWeather() *DailyWeather {
    m := new(DailyWeather)
    return m
}

type DailyWeather struct {

  time int
  summary string
  icon string
  sunriseTime int
  sunsetTime int
  precipProbability float32
  temperatureMin float32
  temperatureMinTime int
  temperatureMax float32
  temperatureMaxTime int
  apparentTemperatureMaxTime int
  dewPoint float32
  windSpeed float32
  humidity float32
  pressure float32
  cloudCover float32
}

func (w *DailyWeather) getTime() int {
    return w.time
}


