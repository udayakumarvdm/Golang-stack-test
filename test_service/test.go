package test

import (
	"encoding/json"
	entity "stack/entity"
	"time"
)

type Response struct {
	Arr []Dt `json:"arr"`
}

type Dt struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

func Test() (*entity.TestEntity, error) {

	input := []byte(`{
		"arr":[
			{
				"start_date": "2019-05-29",
				"end_date": "2019-08-30"
	
			},
			{
				"start_date": "2019-05-29", 
				"end_date": null
			}
		]
	}`)

	var jsonMap Response

	err := json.Unmarshal(input, &jsonMap)
	if err != nil {
		return nil, err
	}

	var res entity.TestEntity

	for _, v := range jsonMap.Arr {
		//Convert string date to datetime
		layout := "2006-01-02"
		if v.StartDate != "" && v.EndDate != "" {

			strcnvTotimeFrom, err := time.Parse(layout, v.StartDate)
			if err != nil {
				return nil, err
			}

			strcnvTotimeTo, err := time.Parse(layout, v.EndDate)
			if err != nil {
				return nil, err
			}

			res.StartDate = strcnvTotimeFrom
			res.EndDate = strcnvTotimeTo
		}
	}
	return &res, nil
}
