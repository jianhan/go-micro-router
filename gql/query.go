package gql

import (
	"errors"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/graphql-go/graphql"
	"github.com/jianhan/go-micro-courses/proto/course"
)

var QueryMap = map[string]*graphql.Field{}

func init() {
	QueryMap["courses"] = getCoursesQuery(course.NewCoursesClient("", nil))
}

func GetProtoTimestamp(params graphql.ResolveParams, field, format string) (*timestamp.Timestamp, bool, error) {
	if _, found := params.Args[field]; !found {
		return nil, false, nil
	}
	if startStr, ok := params.Args[field].(string); ok {
		t, err := time.Parse(format, startStr)
		if err != nil {
			return nil, true, errors.New("start time is not a valid RFC3339 format")
		}
		pt, err := ptypes.TimestampProto(t)
		if err != nil {
			return nil, true, errors.New("can not parse start time")
		}
		return pt, true, nil
	}
	return nil, false, nil
}
