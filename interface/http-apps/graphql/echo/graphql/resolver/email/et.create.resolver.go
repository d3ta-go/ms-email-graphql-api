package email

import appEmailDTOET "github.com/d3ta-go/ddd-mod-email/modules/email/application/dto/email_template"

// ETCreateResResolver represent ETCreateResponse Resolver
type ETCreateResResolver struct {
	Model *appEmailDTOET.ETCreateResDTO
}

// Code field resolver
func (r *ETCreateResResolver) Code() *string {
	return &r.Model.Code
}

// Version field resolver
func (r *ETCreateResResolver) Version() *string {
	return &r.Model.Version
}
