package mongodb
//
// import (
//     "fmt"
//     json_diff `github.com/520MianXiangDuiXiang520/json-diff`
//     "go.mongodb.org/mongo-driver/bson/primitive"
//     `io/ioutil`
//     "lets_diagram/config"
//     "lets_diagram/global"
//     "log"
//     "testing"
// )
//
// func init() {
// 	config.InitSetting("../../../config/setting.json")
// 	err := global.InitMongoDBConn(config.GetSetting().MongoDBConn.Host,
// 		config.GetSetting().MongoDBConn.Port, config.GetSetting().MongoDBConn.Database)
// 	if err != nil {
// 		log.Fatalf("fail to init MySQL connection: %v", err)
// 	}
// }
//
// func TestInsertNewCanvas(t *testing.T) {
//     data, err := ioutil.ReadFile("./canvas_test.json")
//     if err != nil {
//         t.Error(err)
//     }
//     node := json_diff.Parse(data)
//     id, err := InsertNewCanvas(node)
//     if err != nil {
//         t.Error(err)
//     }
// 	if id.Hex() == "" {
// 		t.Error("fail to insert")
// 	}
// }
//
// func TestDeleteCanvas(t *testing.T) {
// 	data, err := ioutil.ReadFile("./canvas_test.json")
// 	if err != nil {
// 	    t.Error(err)
//     }
// 	node := json_diff.Parse(data)
// 	id, err := InsertNewCanvas(node)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if id.Hex() == "" {
// 		t.Error("fail to insert")
// 	}
// 	err = DeleteCanvas(id)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }
//
// func TestSelectCanvasDataByID(t *testing.T) {
// 	id := "60753e7dca459f02ed1dd7dc"
// 	objectID, _ := primitive.ObjectIDFromHex(id)
// 	x := selectData(objectID)
// 	fmt.Println(x)
// }
