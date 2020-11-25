package resolver

import "fmt"

// GraphQLError represent GraphQLError
type GraphQLError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error function
func (e GraphQLError) Error() string {
	return fmt.Sprintf("error [%d]: %s", e.Code, e.Message)
}

// Extensions error Extensions
func (e GraphQLError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
	}
}
