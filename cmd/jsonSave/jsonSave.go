package jsonsave

import (
	"encoding/json"
	"log"
)

type Datos struct {
	Current Current `json:"current"`
}

type Current struct {
	Time                 string  `json:"time"`
	Interval             int     `json:"interval"`
	Apparent_temperature float64 `json:"apparent_temperature"`
	Temperature_2m       float64 `json:"temperature_2m"`
}

func JsonSaver(jsonResponse string) Datos {

	var datos Datos
	err := json.Unmarshal([]byte(jsonResponse), &datos)
	if err != nil {
		log.Fatal(err)
	}
	return datos
}
