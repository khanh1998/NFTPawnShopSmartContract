package controller

import (
	"net/url"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

func BuildFilterFromGinQuery(query url.Values, queriableParams map[string]string) (bson.M, error) {
	filter := bson.M{}
	for param := range queriableParams {
		value := query.Get(param)
		dataType := queriableParams[param]
		if value != "" {
			switch dataType {
			case "string":
				filter[param] = value
			case "int":
				num, err := strconv.Atoi(value)
				if err != nil {
					return nil, err
				}
				filter[param] = num
			}
		}
	}
	return filter, nil
}
