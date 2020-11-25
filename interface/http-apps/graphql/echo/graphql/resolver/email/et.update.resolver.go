package email

import appEmailDTOET "github.com/d3ta-go/ddd-mod-email/modules/email/application/dto/email_template"

// ETUpdateResResolver represent ETUpdateResponse Resolver
type ETUpdateResResolver struct {
	Model *appEmailDTOET.ETUpdateResDTO
}

// Code field resolver
func (r *ETUpdateResResolver) Code() *string {
	return &r.Model.Code
}

// Version field resolver
func (r *ETUpdateResResolver) Version() *string {
	return &r.Model.Version
}
