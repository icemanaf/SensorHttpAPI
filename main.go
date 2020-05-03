package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/icemanaf/HttpConcepts/config"
	"github.com/icemanaf/HttpConcepts/protos"
	"github.com/segmentio/kafka-go"
)

//Payload sent from the sensor
type Payload struct {
	DeviceID    int     `json:"device_id"`
	LightLevel  int     `json:"light_level"`
	Temperature float32 `json:"temp"`
	Pressure    int     `json:"pressure"`
}

func main() {

	var config, err = config.GetAppConfiguration()

	if err != nil {
		log.Panic("config not set correctly", err)
	}

	fmt.Println("starting the server...")

	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  config.KafkaBrokers,
		Topic:    config.KafKaTopic,
		Balancer: &kafka.LeastBytes{},
	})

	if kafkaWriter != nil {
		fmt.Println("Connected to kafka..")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/lightlevel", func(w http.ResponseWriter, req *http.Request) {
		lightLevelHandler(kafkaWriter, w, req)
	})

	http.ListenAndServe(":8080", mux)
}

func lightLevelHandler(kafkaWriter *kafka.Writer, res http.ResponseWriter, req *http.Request) {

	var buf = make([]byte, req.ContentLength)

	if req.Method != "POST" || req.ContentLength == 0 {

		res.WriteHeader(403)
		res.Write([]byte("bad request"))

		return
	}

	req.Body.Read(buf)

	var pay Payload

	json.Unmarshal(buf, &pay)

	json := string(buf)

	fmt.Printf("device=%v,ight= %v,temp=%3.1f,pressure=%d \n", pay.DeviceID, pay.LightLevel, pay.Temperature, pay.Pressure)

	v := &protos.KafkaMessage{
		Id:                 uuid.New().String(),
		DatetimeCreatedUtc: time.Now().UTC().String(),
		MsgType:            2,
		Payload:            json,
	}

	data, err := proto.Marshal(v)

	if err != nil {
		log.Panic("marshalling error", err)
	}

	err = kafkaWriter.WriteMessages(context.Background(), kafka.Message{
		Value: data,
	})

	if err != nil {
		log.Panic("Error writing to kafka", err)
	}

	fmt.Println(json)

	res.WriteHeader(200)

}
