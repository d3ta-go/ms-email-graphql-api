package email

import (
	"encoding/json"

	appEmailDTOET "github.com/d3ta-go/ddd-mod-email/modules/email/application/dto/email_template"
)

// ETSetActiveRequest represent ETSetActiveRequest
type ETSetActiveRequest struct {
	appEmailDTOET.ETSetActiveReqDTO
}

// ToJSON covert to JSON
func (i *ETSetActiveRequest) ToJSON() []byte {
	json, err := json.Marshal(i)
	if err != nil {
		return nil
	}
	return json
}
