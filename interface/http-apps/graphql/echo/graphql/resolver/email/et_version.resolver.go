package email

import "github.com/d3ta-go/ms-email-graphql-api/interface/http-apps/graphql/echo/graphql/scalar"

/*
type EmailTemplateVersion {
  ID: Uint64
  Version: String
  SubjectTpl: String
  BodyTpl: String
  EmailTemplateID: Uint64
}
*/

// ETVersionResolver represent ETVersionResolver
type ETVersionResolver struct {
	Model *EmailTemplateVersion
}

// ID field resolver
func (r *ETVersionResolver) ID() *scalar.Uint64 {
	id := scalar.Uint64(r.Model.ID)
	return &id
}

// Version field resolver
func (r *ETVersionResolver) Version() *string {
	return &r.Model.Version
}

// SubjectTpl field resolver
func (r *ETVersionResolver) SubjectTpl() *string {
	return &r.Model.SubjectTpl
}

// BodyTpl field resolver
func (r *ETVersionResolver) BodyTpl() *string {
	return &r.Model.BodyTpl
}

// EmailTemplateID field resolver
func (r *ETVersionResolver) EmailTemplateID() *scalar.Uint64 {
	id := scalar.Uint64(r.Model.EmailTemplateID)
	return &id
}
