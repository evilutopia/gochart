package main

import (
	"github.com/zieckey/goini"
	"strings"
)

type PieChart int

func (c *PieChart) Parse(ini *goini.INI) (map[string]string, error) {
	args := make(map[string]string)
	args["ChartType"], _ = ini.Get("ChartType")
	args["Title"], _ = ini.Get("Title")
	args["SubTitle"], _ = ini.Get("SubTitle")
	args["SeriesName"], _ = ini.Get("SeriesName")
	
	DataArray := "[\n"
	kv, _ := ini.GetKvmap(goini.DefaultSection)
	for k, v := range kv {
		if !strings.HasPrefix(k, DataPrefix) {
			continue
		}
		
		key := k[len(DataPrefix):]
		DataArray = DataArray + "['" + key + "' , " + v + "],\n"
	}

	DataArray = DataArray + "]"
		
	args["DataArray"] = DataArray
	return args, nil
}

func (c *PieChart) Template() string {
	return TemplatePieHtml
}

