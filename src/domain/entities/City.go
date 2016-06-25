package entities




func NewCity() (*City) {
	city := new(City)
	return city
}



type City struct {
	name string 	`json:"name"`
	latitude float64 	`json :"latitude"`
	longitude float64	`json: "longitude"`
}






