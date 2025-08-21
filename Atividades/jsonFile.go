package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)
func mainJson(){
	 var pass Flys
	 listaVoos, err := GetJsonFile("flights-1m.json", pass)
	 if err != nil {
		fmt.Println("Error reading JSON file:", err)
	 }
	 var data string
	 fmt.Println("Digite a data do voo no formato YYYY-MM-DD")
	 fmt.Scanln(&data)
	 for i, voo := range listaVoos {
		if voo.FL_DATE == data{
			fmt.Println(i,"Dados do Voo: Data do voo",voo.FL_DATE, "| Demora no ar",voo.DEP_DELAY,"| Demora no Pouso", voo.ARR_DELAY,"| Tempo de voo", voo.AIR_TIME,"| Distancia de voo", voo.DISTANCE,"| Tempo de Decolagem", voo.DEP_TIME,"! Tempo de Pouso", voo.ARR_TIME)
		}
	 }
}
type Flys struct{
	FL_DATE string `json:"FL_DATE"`
	DEP_DELAY int `json:"DEP_DELAY"`
	ARR_DELAY int `json:"ARR_DELAY"`
	AIR_TIME int `json:"AIR_TIME"`
	DISTANCE int `json:"DISTANCE"`
	DEP_TIME float32 `json:"DEP_TIME"`
	ARR_TIME float32 `json:"ARR_TIME"`
}
func GetJsonFile(filepath string, jStorw Flys) ([]Flys, error) {
	var arrStru []Flys
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
