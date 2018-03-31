package query

import (
	"github.com/graphql-go/graphql"
)

var QueryMap = map[string]*graphql.Field{}

func init() {
	QueryMap["course"] = getCoursesQuery()
}
