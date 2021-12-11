package utils

import (
	"encoding/json"
	"fmt"
)


func PrintMaptoJSON(data interface{}) {
	jsonData, err := json.Marshal(data)
	if err!=nil {
		panic("error dalam membuat json")
	}
	fmt.Println(string(jsonData))
}