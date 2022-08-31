package data

// func randSelectRoom(param interface{}, p *core.ElementHandleArgs) error {
// 	pm := param.(*chatList)
// 	where := []bson.M{
// 		{
// 			"$match": bson.M{
// 				"name": bson.M{
// 					"$nin": pm.RoomID},
// 			},
// 		},
// 		{
// 			"$sample": bson.M{
// 				"size": 9,
// 			},
// 		},
// 	}
// 	cur, err := core.Mdb.CollectionDB.Collection("chatroom").Aggregate(core.Mdb.Ctx, where)
// 	if err != nil {
// 		return err
// 	}
// 	var result []bson.M
// 	defer cur.Close(core.Mdb.Ctx)
// 	for cur.Next(core.Mdb.Ctx) {
// 		var res bson.M
// 		if err := cur.Decode(&res); err != nil {
// 			return err
// 		}
// 		result = append(result, res)
// 	}
// 	if err := cur.Err(); err != nil {
// 		return err
// 	}
// 	core.SendJSON(p.Res, http.StatusOK, core.SendMap{"success": true, "data": result})
// 	return nil
// }
