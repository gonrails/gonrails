package rabbits

//
// import (
// 	"github.com/gonrails/gonrails/models"
// )
//
// func orderHandle() {
// 	for {
// 		select {
// 		case msg := <-orderChan:
// 			processMsg(msg)
// 		}
// 	}
// }
//
// func processMsg(msg *OrderMessage) (err error) {
// 	var record models.BetOrder
//
// 	if err := models.DB().Model(&models.BetOrder{}).First(&record, msg.ID).Error; err != nil {
// 		return err
// 	}
//
// 	switch msg.Event {
// 	case "Create":
// 		record.HandleCreate()
// 	case "Check":
// 		record.HandleCheck()
// 	}
//
// 	return nil
// }
