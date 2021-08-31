package controller

import (
	"errors"
	"log"
	"net/url"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func BuildFilterFromGinQuery(query url.Values, queriableParams map[string]string) (bson.M, error) {
	filter := bson.M{}
	queryMap := map[string][]string(query)
	for param, values := range queryMap {
		parts := strings.Split(param, ".")
		// I decided to let the param name can contain operator also,
		// param name and operator is separated by a dot
		// for example, id.in=1,2,3,4
		// 'id' is the param name, 'in' is the operator
		// the above param will be convert to filter in MongoDB: { id: $in: [1, 2, 3, 4] }
		var name, operator string
		if len(parts) > 2 {
			return nil, errors.New("param name is not valid")
		}
		if len(parts) == 2 {
			name, operator = parts[0], parts[1]
		} else {
			name = parts[0]
		}
		dataType := queriableParams[name]
		log.Println(param, values, len(values))
		log.Println(name, operator, dataType)
		if len(values) > 0 {
			switch dataType {
			case "string":
				filter[name] = values[0]
			case "strings":
				filter[name] = bson.M{
					"$" + operator: strings.Split(values[0], ","),
				}
			case "int":
				num, err := strconv.Atoi(values[0])
				if err != nil {
					return nil, err
				}
				filter[name] = num
			}
		}

	}
	log.Println(filter)
	return filter, nil
}
