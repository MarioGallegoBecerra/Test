package data

type ItemEntity struct {
	Market           MarketEntity
	ShortDescription string
	LongDescription  string
	Price            float64
	Imagen           string
}

type MarketEntity struct {
	Id        string
	Name      string
	Locations []LocationEntity
}

type LocationEntity struct {
	Latitude  string
	Longitude string
	Zip       string
	Address   string
}
