package email

import (
	"encoding/json"

	appEmailDTOET "github.com/d3ta-go/ddd-mod-email/modules/email/application/dto/email_template"
)

// ETCreateRequest represent ETCreateRequest
type ETCreateRequest struct {
	Input ETCreateInput
}

// ETCreateInput represent ETCreateInput
type ETCreateInput struct {
	appEmailDTOET.ETCreateReqDTO
}

// ToJSON covert to JSON
func (i *ETCreateInput) ToJSON() []byte {
	json, err := json.Marshal(i)
	if err != nil {
		return nil
	}
	return json
}
