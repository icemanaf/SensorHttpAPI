# A Simple GO based HTTP API for low powered sensors

[![Build Status](https://travis-ci.com/icemanaf/SensorHttpAPI.svg?branch=master)](https://travis-ci.com/icemanaf/SensorHttpAPI)

This is a very simple HTTP service which integrates with low powered microcontrollers such as the NodeMCU.
Please note that simple devices such as these have trouble POSTing data via HTTPS.

The data is grouped together into a message and posted into a Kafka Topic for consumption by other services.

The KafkaMessage.pb.go file is created by first installing [protoc](https://developers.google.com/protocol-buffers/docs/gotutorial) and then running the following command on the KafkaMessage proto file.

```
    protoc --go_out=$PWD KafkaMessage.proto  
```

run the docker  image

```
    docker run -it  -e TOPIC='MAIN_EVENT_QUEUE' -e BROKERS='192.168.0.85:9092,192.168.0.86:9092' -p  8080:8080 icemanaf/rpi-sensor-api
``