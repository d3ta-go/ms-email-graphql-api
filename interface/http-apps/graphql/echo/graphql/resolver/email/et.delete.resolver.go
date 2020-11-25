package email

import (
	appEmailDTOET "github.com/d3ta-go/ddd-mod-email/modules/email/application/dto/email_template"
	"github.com/d3ta-go/ms-email-graphql-api/interface/http-apps/graphql/echo/graphql/scalar"
)

// ETDeleteResResolver represent ETDeleteResResolver Resolver
type ETDeleteResResolver struct {
	Model *appEmailDTOET.ETDeleteResDTO
}

// ID field resolver
func (r *ETDeleteResResolver) ID() *scalar.Uint64 {
	id := scalar.Uint64(r.Model.Data.ID)
	return &id
}

// UUID field resolver
func (r *ETDeleteResResolver) UUID() *string {
	return &r.Model.Data.UUID
}

// Code field resolver
func (r *ETDeleteResResolver) Code() *string {
	return &r.Model.Data.Code
}

// Name field resolver
func (r *ETDeleteResResolver) Name() *string {
	return &r.Model.Data.Name
}

// IsActive field resolver
func (r *ETDeleteResResolver) IsActive() *bool {
	return &r.Model.Data.IsActive
}

// EmailFormat field resolver
func (r *ETDeleteResResolver) EmailFormat() *string {
	return &r.Model.Data.EmailFormat
}

// DefaultVersionID field resolver
func (r *ETDeleteResResolver) DefaultVersionID() *scalar.Uint64 {
	id := scalar.Uint64(r.Model.Data.DefaultVersionID)
	return &id
}

// VersionCount field resolver
func (r *ETDeleteResResolver) VersionCount() *scalar.Int64 {
	count := scalar.Int64(r.Model.Data.VersionCount)
	return &count
}
