package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)
func main(){
	 var pass Passager
	 listaPassageiros, err := GetJsonFile("titanic.json", pass)
	 if err != nil {
		fmt.Println("Error reading JSON file:", err)
	 }
	 for _, viajante := range listaPassageiros {
		fmt.Println(viajante.Name)
	 }
}
type Passager struct{
	PassagerId float32 `json:"PassagerId"`
	Name string `json:"Name"`
}
func GetJsonFile(filepath string, jStorw Passager) ([]Passager, error) {
	var arrStru []Passager
	data, err := os.Open(filepath)
	if err != nil {
		return arrStru,err
	}
	bytes,err:= io.ReadAll(data)
	if err != nil {
		return arrStru,err
	}
	err = json.Unmarshal(bytes, &arrStru)
	if err != nil {
		return arrStru,err
	}
	return arrStru, nil
}
