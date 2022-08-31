package data

import (
	"context"
	"stb-library/app/software/internal/biz"

	"gopkg.in/mgo.v2/bson"
)

var _ biz.CodeHeaderRepo = (*codeHeaderRepo)(nil)

type codeHeaderRepo struct {
	data *Data
}

func NewCodeHeaderRepo(d *Data) biz.CodeHeaderRepo {
	return &codeHeaderRepo{
		data: d,
	}
}

func (u *codeHeaderRepo) Delete(ctx context.Context, token string) error {
	// rediser.DelCode(u.data.rds, token)
	return nil
}

// db.code.aggregate( [
//
//	{ $match : { key : "aaa" ,content:{$nin:["a","c","e","b"]}}
//	},
//	{ $sample :{ size : 3 }
//	}
//
// ] );
func (u *codeHeaderRepo) GetCodes(num int, key string, filters []string) ([]bson.M, error) {
	ids := []bson.ObjectId{}
	for _, v := range filters {
		if v == "" {
			continue
		}
		ids = append(ids, bson.ObjectIdHex(v))
	}
	where := []bson.M{
		{
			"$match": bson.M{
				"key": key,
				"_id": bson.M{
					"$nin": ids,
				},
			},
		},
		{
			"$sample": bson.M{
				"size": num,
			},
		},
	}
	ctx := context.Background()
	cur, err := u.data.mongoClient.CollectionDB.Collection("code_header").Aggregate(ctx, where)
	if err != nil {
		return nil, err
	}
	var result []bson.M
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var res bson.M
		if err := cur.Decode(&res); err != nil {
			return nil, err
		}
		result = append(result, res)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (u *codeHeaderRepo) Create(codes []biz.CodeHeader) error {
	list := []interface{}{}
	for _, cod := range codes {
		m := bson.M{
			"key":         cod.Key,
			"content":     cod.Content,
			"code_length": cod.CodeLength,
		}
		list = append(list, m)
	}

	if _, err := u.data.mongoClient.InsertMany("code_header", list); err != nil {
		return err
	}
	return nil
}
