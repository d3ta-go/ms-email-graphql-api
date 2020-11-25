package email

import (
	appEmailDTOET "github.com/d3ta-go/ddd-mod-email/modules/email/application/dto/email_template"
	"github.com/d3ta-go/ms-email-graphql-api/interface/http-apps/graphql/echo/graphql/scalar"
)

// ETFindByCodeResolver represent ETFindByCodeResolver
type ETFindByCodeResolver struct {
	Model *appEmailDTOET.ETFindByCodeResDTO
}

// ID field resolver
func (r *ETFindByCodeResolver) ID() *scalar.Uint64 {
	id := scalar.Uint64(r.Model.Data.ID)
	return &id
}

// UUID field resolver
func (r *ETFindByCodeResolver) UUID() *string {
	return &r.Model.Data.UUID
}

// Code field resolver
func (r *ETFindByCodeResolver) Code() *string {
	return &r.Model.Data.Code
}

// Name field resolver
func (r *ETFindByCodeResolver) Name() *string {
	return &r.Model.Data.Name
}

// IsActive field resolver
func (r *ETFindByCodeResolver) IsActive() *bool {
	return &r.Model.Data.IsActive
}

// EmailFormat field resolver
func (r *ETFindByCodeResolver) EmailFormat() *string {
	return &r.Model.Data.EmailFormat
}

// DefaultVersionID field resolver
func (r *ETFindByCodeResolver) DefaultVersionID() *scalar.Uint64 {
	id := scalar.Uint64(r.Model.Data.DefaultVersionID)
	return &id
}

// DefaultTemplateVersion field resolver
func (r *ETFindByCodeResolver) DefaultTemplateVersion() *ETVersionResolver {
	etRxs := &ETVersionResolver{Model: &EmailTemplateVersion{
		ID:              r.Model.Data.DefaultTemplateVersion.ID,
		Version:         r.Model.Data.DefaultTemplateVersion.Version,
		SubjectTpl:      r.Model.Data.DefaultTemplateVersion.SubjectTpl,
		BodyTpl:         r.Model.Data.DefaultTemplateVersion.BodyTpl,
		EmailTemplateID: r.Model.Data.DefaultTemplateVersion.EmailTemplateID,
	}}
	return etRxs
}
