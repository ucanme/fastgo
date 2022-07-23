package manager

import (
	"fmt"
	"github.com/ucanme/fastgo/library/db"
	"github.com/ucanme/fastgo/models"
)

type manager struct {
	ProductLineMap map[int]models.ProductionLine
	ProductionLineStationMap map[int]map[string]models.Station
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

}
