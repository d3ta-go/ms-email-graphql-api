package email

import (
	"github.com/d3ta-go/ms-email-graphql-api/interface/http-apps/graphql/echo/graphql/scalar"
)

/**
type EmailTemplate {
  ID: Uint64
  UUID: String
  Code: String
  Name: String
  IsActive: Boolean
  EmailFormat: String
  DefaultVersionID: Uint64
  DefaultTemplateVersion: EmailTemplateVersion!
}
**/

// ETResolver represent ETResolver
type ETResolver struct {
	Model *EmailTemplate
}

// ID field resolver
func (r *ETResolver) ID() *scalar.Uint64 {
	id := scalar.Uint64(r.Model.ID)
	return &id
}

// UUID field resolver
func (r *ETResolver) UUID() *string {
	return &r.Model.UUID
}

// Code field resolver
func (r *ETResolver) Code() *string {
	return &r.Model.Code
}

// Name field resolver
func (r *ETResolver) Name() *string {
	return &r.Model.Name
}

// IsActive field resolver
func (r *ETResolver) IsActive() *bool {
	return &r.Model.IsActive
}

// EmailFormat field resolver
func (r *ETResolver) EmailFormat() *string {
	return &r.Model.EmailFormat
}

// DefaultVersionID field resolver
func (r *ETResolver) DefaultVersionID() *scalar.Uint64 {
	id := scalar.Uint64(r.Model.DefaultVersionID)
	return &id
}

// DefaultTemplateVersion field resolver
func (r *ETResolver) DefaultTemplateVersion() *ETVersionResolver {
	if r.Model.DefaultTemplateVersion == nil {
		return &ETVersionResolver{Model: &EmailTemplateVersion{}}
	}
	etRxs := &ETVersionResolver{r.Model.DefaultTemplateVersion}
	return etRxs
}
