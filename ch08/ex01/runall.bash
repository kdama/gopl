#!/usr/bin/env bash

TZ=US/Eastern    go run server/main.go -port 8010 &
TZ=Asia/Tokyo    go run server/main.go -port 8020 &
TZ=Europe/London go run server/main.go -port 8030 &

go run client/main.go NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030
