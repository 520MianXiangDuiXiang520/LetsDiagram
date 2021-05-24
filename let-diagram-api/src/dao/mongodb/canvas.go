package mongodb
//
// import (
// 	"context"
//     json_diff `github.com/520MianXiangDuiXiang520/json-diff`
//     `github.com/qiniu/qmgo`
//     "go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"lets_diagram/global"
// )
//
// func InsertNewCanvas(canvas interface{}) (primitive.ObjectID, error) {
// 	res, err := global.GetMgoCollection("canvas").InsertOne(context.Background(), canvas)
// 	if err != nil {
// 		return primitive.ObjectID{}, err
// 	}
// 	return res.InsertedID.(primitive.ObjectID), nil
// }
//
// func UpdateCanvasData(mgo *qmgo.Collection,  id string, newValue *json_diff.JsonNode) error {
//     objectID, err := primitive.ObjectIDFromHex(id)
//     if err != nil {
//         return err
//     }
//     value, err := json_diff.Marshal(newValue)
//     if err != nil {
//         return err
//     }
//     return mgo.UpdateId(context.Background(), bson.bodyMutex{"_id": objectID},
//         bson.bodyMutex{"$set": string(value)})
// }
//
// func DeleteCanvas(id primitive.ObjectID) error {
// 	return global.GetMgoCollection("canvas").RemoveId(context.Background(), id)
// }
//
// func SelectCanvasDataByID(id primitive.ObjectID) ([]byte, error) {
// 	ctx := context.Background()
// 	nodes := make([]*json_diff.JsonNode, 0)
// 	err := global.GetMgoCollection("canvas").Find(ctx, bson.bodyMutex{"_id": id}).Limit(1).All(&nodes)
// 	if err != nil {
// 		return nil, err
// 	}
// 	canvas, err := json_diff.Marshal(nodes[0])
// 	if err != nil {
// 	    return nil, err
//     }
// 	return canvas, err
// }
//
// func selectData(id primitive.ObjectID) interface{} {
// 	ctx := context.Background()
// 	canvas := make([]interface{}, 0)
// 	err := global.GetMgoCollection("canvas").Find(ctx, bson.bodyMutex{"_id": id}).Limit(1).All(&canvas)
// 	if err != nil {
// 		return nil
// 	}
// 	return canvas[0]
// }
