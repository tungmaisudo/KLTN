package Data

import (
	"time"
)

type Data struct {
	Activated  time.Time `json:"activated"`
	Bbl        string    `json:"bbl"`
	Bin        string    `json:"bin"`
	Boro       string    `json:"boro"`
	Borocd     string    `json:"borocd"`
	Borocode   string    `json:"borocode"`
	Boroct2010 string    `json:"boroct2010"`
	Boroname   string    `json:"boroname"`
	City       string    `json:"city"`
	Coundist   string    `json:"coundist"`
	Ct2010     string    `json:"ct2010"`
	DoittID    string    `json:"doitt_id"`
	Lat        string    `json:"lat"`
	Location   string    `json:"location"`
	LocationT  string    `json:"location_t"`
	Lon        string    `json:"lon"`
	Name       string    `json:"name"`
	Ntacode    string    `json:"ntacode"`
	Ntaname    string    `json:"ntaname"`
	Objectid   string    `json:"objectid"`
	Postcode   string    `json:"postcode"`
	Provider   string    `json:"provider"`
	Remarks    string    `json:"remarks"`
	Sourceid   string    `json:"sourceid"`
	Ssid       string    `json:"ssid"`
	TheGeom    TheGeom   `json:"the_geom"`
	Type       string    `json:"type"`
	X          string    `json:"x"`
	Y          string    `json:"y"`
}

type TheGeom struct {
	Type        string    `json:"type"`
	Coordinates []float32 `json:"coordinates"`
}
