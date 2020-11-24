package email

import "github.com/d3ta-go/ms-email-graphql-api/interface/http-apps/graphql/echo/graphql/scalar"

/**
type ETListAll {
  Count: int
  Data: [EmailTemplate]
}
**/

// ETListAllResolver represent ETListAllResolver
type ETListAllResolver struct {
	Model *ETListAllResponse
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
		etRxs = append(etRxs, &ETResolver{et})
	}
	return etRxs
}
