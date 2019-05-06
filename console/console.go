package console

import (
	"fmt"
	"kalista/models"
)

func Run() {
	var betOrder = &models.BetOrder{}
	err := models.DB().Preload("ScoreDetails").Model(betOrder).First(betOrder).Error
	
	if err  != nil{
		panic(err)
	}
	
	fmt.Printf("%v", betOrder)
	fmt.Println("")
	fmt.Printf("%v", betOrder.ScoreDetails)
}
