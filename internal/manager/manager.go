package manager

import (
	"fmt"
	"github.com/ucanme/fastgo/library/db"
	"github.com/ucanme/fastgo/models"
)

type manager struct {
	ProductLineMap map[int]models.ProductionLine
	ProductionLineStationMap map[int]models.Station
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

	Manager.ProductionLineStationMap= map[int]models.Station{}
	for _,v := range stationList{
		Manager.ProductionLineStationMap[v.ProductionLineId] = v
	}

}
