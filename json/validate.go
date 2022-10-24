package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type VDCInfo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Aggregation *struct {
		AVZoneIds []string `json:"availableZoneIds"`
		Projects  []*struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"projects"`
	} `json:"aggregation"`
	SubVdcs []*VDCInfo `json:"subVdcs"`
}

type Response struct {
	Data []*VDCInfo `json:"data"`
}

func main() {
	b, _ := os.ReadFile("./data.json")
	r := &Response{}
	json.Unmarshal(b, r)

	show(r.Data)	
}

func show(in []*VDCInfo){
	if in == nil {
		return 
	}
	for _, v := range in {
		if v.Aggregation != nil {
	 	
		fmt.Println("Name:", v.Name, " -- AVZoneIds:", v.Aggregation.AVZoneIds)	
			for _, p := range v.Aggregation.Projects{
				fmt.Println("pid:",p.Id, " -- ", p.Name)
		}
	}
	show(v.SubVdcs)
	}
	
}
