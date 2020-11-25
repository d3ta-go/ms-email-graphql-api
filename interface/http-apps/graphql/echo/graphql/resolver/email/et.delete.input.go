package email

import "encoding/json"

// ETDeleteRequest represent
type ETDeleteRequest struct {
	Code string
}

// ToJSON covert to JSON
func (i *ETDeleteRequest) ToJSON() []byte {
	json, err := json.Marshal(i)
	if err != nil {
		return nil
	}
	return json
}
