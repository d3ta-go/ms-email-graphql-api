package email

import appEmailDTO "github.com/d3ta-go/ddd-mod-email/modules/email/application/dto/email"

// SendEmailResResolver represent SendEmailResResolver Resolver
type SendEmailResResolver struct {
	Model *appEmailDTO.SendEmailResDTO
}

// TemplateCode field resolver
func (r *SendEmailResResolver) TemplateCode() *string {
	return &r.Model.TemplateCode
}

// Status field resolver
func (r *SendEmailResResolver) Status() *string {
	return &r.Model.Status
}
