#!/usr/bin/env bash

# ---- Fetch ----
# go run main.go -fetch > index/index.json

# ---- Search ----
go run main.go "Sleep" < index/index.json
go run main.go "Can't" "Sleep" < index/index.json # AND Search
