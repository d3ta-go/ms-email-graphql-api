package email

import (
	"encoding/json"

	appEmailDTOET "github.com/d3ta-go/ddd-mod-email/modules/email/application/dto/email_template"
)

// ETUpdateRequest represent ETUpdateRequest
type ETUpdateRequest struct {
	appEmailDTOET.ETUpdateReqDTO
}

// ToJSON covert to JSON
func (i *ETUpdateRequest) ToJSON() []byte {
	json, err := json.Marshal(i)
	if err != nil {
		return nil
	}
	return json
}
