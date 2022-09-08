package main

import "encoding/json"

func map2ToJson(param map[string][]string /*interface{}*/) string {
	dataType, _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}