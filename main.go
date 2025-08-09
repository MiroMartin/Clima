package main

import (
	"clima/cmd/httpRec"
	jsonsave "clima/cmd/jsonSave"
	"fmt"
)

func main() {

	var url string = "https://api.open-meteo.com/v1/forecast?latitude=-36.3569&longitude=-56.7235&current=apparent_temperature,temperature_2m&forecast_days=1"
	json := httpRec.HttpReq(url)
	var datos jsonsave.Datos = jsonsave.JsonSaver(json)
	termperatura := datos.Current.Temperature_2m
	sensacion_termica := datos.Current.Apparent_temperature
	fmt.Printf("La temperatura actual es %.1f y la sensacion termica es %.1f\n", termperatura, sensacion_termica)
}
