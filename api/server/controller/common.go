package controller

import (
	"net/url"

	"go.mongodb.org/mongo-driver/bson"
)

func BuildFilterFromGinQuery(query url.Values, queriableParams []string) bson.M {
	filter := bson.M{}
	for _, param := range queriableParams {
		value := query.Get(param)
		if value != "" {
			filter[param] = value
		}
	}
	return filter
}
