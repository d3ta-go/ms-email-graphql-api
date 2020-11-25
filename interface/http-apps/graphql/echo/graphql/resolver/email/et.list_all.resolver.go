package email

import (
	appEmailDTOET "github.com/d3ta-go/ddd-mod-email/modules/email/application/dto/email_template"
	"github.com/d3ta-go/ms-email-graphql-api/interface/http-apps/graphql/echo/graphql/scalar"
)

/**
# EmailTemplateListItem
type EmailTemplateListItem {
  ID: Uint64
  UUID: String
  Code: String
  Name: String
  IsActive: Boolean
  EmailFormat: String
  DefaultVersionID: Uint64
}

# ETListAll Response
type ETListAllRes {
  Count: Int64
  Data: [EmailTemplateListItem!]!
}
**/

// ETListAllResolver represent ETListAllResolver
type ETListAllResolver struct {
	// Model *ETListAllResponse
	Model *appEmailDTOET.ETListAllResDTO
}

// ETListAllResponse represent ETListAllResponse
type ETListAllResponse struct {
	Count int64            `json:"count"`
	Data  []*EmailTemplate `json:"data"`
}

// Count field resolver
func (r *ETListAllResolver) Count() *scalar.Int64 {
	count := scalar.Int64(r.Model.Count)
	return &count
}

// Data field resolver
func (r *ETListAllResolver) Data() []*ETResolver {
	var etRxs []*ETResolver
	for _, et := range r.Model.Data {
		//etRxs = append(etRxs, &ETResolver{et})
		etRxs = append(etRxs, &ETResolver{Model: &EmailTemplate{
			ID:               et.ID,
			UUID:             et.UUID,
			Code:             et.Code,
			Name:             et.Name,
			IsActive:         et.IsActive,
			EmailFormat:      et.EmailFormat,
			DefaultVersionID: et.DefaultVersionID,
		}})
	}
	return etRxs
}
