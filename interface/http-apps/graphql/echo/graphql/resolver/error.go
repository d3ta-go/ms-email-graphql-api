package resolver

import "fmt"

type GraphQLError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e GraphQLError) Error() string {
	return fmt.Sprintf("error [%d]: %s", e.Code, e.Message)
}

func (e GraphQLError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
	}
}
