package scalar

import (
	"errors"
	"strconv"
)

// Int64 scalar type
type Int64 uint64

// ImplementsGraphQLType implement GraphQLType
func (Int64) ImplementsGraphQLType(name string) bool {
	return name == "Int64"
}

// UnmarshalGraphQL unmarshal GraphQL
func (i *Int64) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case int, int32, int64:
		if input.(int64) < 0 {
			val := 0
			*i = Int64(val)
		} else {
			*i = Int64(input.(int64))
		}
	case float32, float64:
		if input.(float64) < 0 {
			val := 0
			*i = Int64(val)
		} else {
			*i = Int64(input.(float64))
		}
	case uint, uint32, uint64:
		*i = Int64(input.(int64))
	case string:
		var value int64
		value, err = strconv.ParseInt(input, 10, 64)
		if err == nil {
			*i = Int64(value)
		}
	default:
		err = errors.New("wrong type")
	}
	return err
}

// MarshalJSON marshal JSON
func (i Int64) MarshalJSON() ([]byte, error) {
	b := []byte(strconv.FormatInt(int64(i), 10))
	return b, nil
}
