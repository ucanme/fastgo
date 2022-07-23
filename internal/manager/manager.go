package manager

import (
	"fmt"
	"github.com/ucanme/fastgo/library/db"
	"github.com/ucanme/fastgo/models"
	"sort"
)

type manager struct {
	ProductLineMap map[int]models.ProductionLine
	ProductionLineStationMap map[int]map[string]models.Station
	ProductionLineStationSort map[int][]models.Station
	ProductionLineStationIndex map[int]map[string]int

}

var Manager manager
func Init()  {
	productionLineList := []models.ProductionLine{}
	err := db.DB().Find(&productionLineList).Error
	if err != nil{
		panic(err)
	}
	fmt.Println("productionLineList------",productionLineList)
	Manager.ProductLineMap = map[int]models.ProductionLine{}
	for _,v := range productionLineList{
		Manager.ProductLineMap[v.ProductionLineId] = v
	}

	stationList := []models.Station{}
	err = db.DB().Find(&stationList).Error
	if err != nil{
		panic(err)
	}

	fmt.Println("stationList",stationList)
	Manager.ProductionLineStationMap= map[int]map[string]models.Station{}
	for _,v := range stationList{
		if _,ok := Manager.ProductionLineStationMap[v.ProductionLineId];!ok{
			Manager.ProductionLineStationMap[v.ProductionLineId] = map[string]models.Station{}
		}
		Manager.ProductionLineStationMap[v.ProductionLineId][v.StationCode] = v
	}



	Manager.ProductionLineStationSort =  map[int][]models.Station{}

	Manager.ProductionLineStationIndex = map[int]map[string]int{}



	for _,v := range stationList{
		if _ , ok := Manager.ProductionLineStationSort[v.ProductionLineId];!ok{
			Manager.ProductionLineStationSort[v.ProductionLineId] = []models.Station{}
		}
		Manager.ProductionLineStationSort[v.ProductionLineId] = append(Manager.ProductionLineStationSort[v.ProductionLineId],v)
	}

	for _,stations := range 	Manager.ProductionLineStationSort{
		sort.Slice(stations, func(i, j int) bool {
			return stations[i].StationID < stations[j].StationID
		})
	}

	for _,stations := range Manager.ProductionLineStationSort{
		for index,station := range stations{
			if _,ok := Manager.ProductionLineStationIndex[station.ProductionLineId];!ok{
				Manager.ProductionLineStationIndex[station.ProductionLineId] = map[string]int{}
			}
			Manager.ProductionLineStationIndex[station.ProductionLineId][station.StationCode] = index
		}
	}

}
