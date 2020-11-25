package email

import appEmailDTOET "github.com/d3ta-go/ddd-mod-email/modules/email/application/dto/email_template"

// ETSetActiveResResolver represent ETSetActiveRespnse Resolver
type ETSetActiveResResolver struct {
	Model *appEmailDTOET.ETSetActiveResDTO
}

// Code field resolver
func (r *ETSetActiveResResolver) Code() *string {
	return &r.Model.Code
}

// IsActive field resolver
func (r *ETSetActiveResResolver) IsActive() *bool {
	return &r.Model.IsActive
}
