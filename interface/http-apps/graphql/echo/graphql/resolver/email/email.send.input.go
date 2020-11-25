package email

import (
	"encoding/json"

	appEmailDTO "github.com/d3ta-go/ddd-mod-email/modules/email/application/dto/email"
)

// SendEmailRequest represent SendEmailRequest
type SendEmailRequest struct {
	Input SendEmailInput
}

// SendEmailInput represent SendEmailInput
type SendEmailInput struct {
	TemplateCode   string
	From           appEmailDTO.MailAddressDTO
	To             appEmailDTO.MailAddressDTO
	CC             *ccBcc
	BCC            *ccBcc
	TemplateData   []TemplateDataInput
	ProcessingType string
}

type ccBcc []*appEmailDTO.MailAddressDTO

// TemplateDataInput represent
type TemplateDataInput struct {
	FieldName  string
	FieldValue string
}

// TranslateCC translate CC
func (i *SendEmailInput) TranslateCC() []*appEmailDTO.MailAddressDTO {
	if i.CC != nil {
		return i._translate(*i.CC)
	}
	return nil

}

// TranslateBCC translate BCC
func (i *SendEmailInput) TranslateBCC() []*appEmailDTO.MailAddressDTO {
	if i.BCC != nil {
		return i._translate(*i.BCC)
	}
	return nil
}

func (i *SendEmailInput) _translate(from ccBcc) []*appEmailDTO.MailAddressDTO {
	var r []*appEmailDTO.MailAddressDTO
	if from != nil {
		for _, v := range from {
			if v != nil {
				r = append(r, v)
			}
		}
	}
	return r
}

// TranslateTD translate Template Data
func (i *SendEmailInput) TranslateTD() map[string]interface{} {
	r := make(map[string]interface{})
	for _, v := range i.TemplateData {
		r[v.FieldName] = v.FieldValue
	}
	return r
}

// ToJSON covert to JSON
func (i *SendEmailInput) ToJSON() []byte {
	json, err := json.Marshal(i)
	if err != nil {
		return nil
	}
	return json
}
