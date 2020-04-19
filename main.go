package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Payload sent from the sensor
type Payload struct {
	LightLevel  int     `json:"light_level"`
	Temperature float32 `json:"temp"`
	Pressure    int     `json:"pressure"`
}

func main() {

	fmt.Println("starting the server...")

	mux := http.NewServeMux()

	mux.HandleFunc("/lightlevel", lightLevelHandler)

	http.ListenAndServe(":8080", mux)
}

func lightLevelHandler(res http.ResponseWriter, req *http.Request) {

	var buf = make([]byte, req.ContentLength)

	if req.Method != "POST" || req.ContentLength == 0 {

		res.WriteHeader(403)
		res.Write([]byte("bad request"))

		return
	}

	req.Body.Read(buf)

	var pay Payload

	json.Unmarshal(buf, &pay)

	fmt.Printf("light= %v,temp=%3.1f,pressure=%d \n", pay.LightLevel, pay.Temperature, pay.Pressure)

	res.WriteHeader(200)

}
