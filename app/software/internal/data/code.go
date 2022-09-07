package data

import (
	"context"
	"stb-library/app/software/internal/biz"
	"stb-library/app/software/internal/model"

	"gopkg.in/mgo.v2/bson"
)

var _ biz.CodeRepo = (*codeRepo)(nil)

type codeRepo struct {
	data *Data
}

func NewCodeRepo(d *Data) biz.CodeRepo {
	return &codeRepo{
		data: d,
	}
}

func (u *codeRepo) Delete(ctx context.Context, token string) error {
	// rediser.DelCode(u.data.rds, token)
	return nil
}

// db.code.aggregate( [
//
//	{ $match : { key : "bbb" ,"_id":{$nin:[ObjectId("630f6a86c466a68e1a651315")]}}
//	},
//	{ $sample :{ size : 3 }
//	}
//
// ] );
func (u *codeRepo) GetCodes(num int, lan string, keys []string, filters []string) ([]bson.M, error) {
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
				"language": lan,
				"key": bson.M{
					"$in": keys,
				},
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
	cur, err := u.data.mongoClient.CollectionDB.Collection("code").Aggregate(u.data.mongoClient.Ctx, where)
	if err != nil {
		return nil, err
	}
	var result []bson.M
	defer cur.Close(u.data.mongoClient.Ctx)
	for cur.Next(u.data.mongoClient.Ctx) {
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

func (u *codeRepo) Create(arg []model.Code) error {
	list := []interface{}{}
	for _, cod := range arg {
		m := bson.M{
			"language":    cod.Language,
			"key":         cod.Key,
			"content":     cod.Content,
			"code_length": len(cod.Content),
		}
		list = append(list, m)
	}

	if _, err := u.data.mongoClient.InsertMany("code", list); err != nil {
		return err
	}
	return nil
}

func (u *codeRepo) GetHeaderCode(num int, lan string, filters []string) ([]bson.M, error) {
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
				"language": lan,
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

func (u *codeRepo) CreateHeaders(arg []model.Code) error {
	list := []interface{}{}
	for _, cod := range arg {
		m := bson.M{
			"language":    cod.Language,
			"content":     cod.Content,
			"code_length": len(cod.Content),
		}
		list = append(list, m)
	}

	if _, err := u.data.mongoClient.InsertMany("code_header", list); err != nil {
		return err
	}
	return nil
}
