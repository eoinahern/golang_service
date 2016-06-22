package entities




func NewCity(){
	city := new(City)
	return city
}



type City struct {
	name string
	latitude float64
	longitude float64
}






