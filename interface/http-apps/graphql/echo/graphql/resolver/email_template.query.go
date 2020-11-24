package resolver

import (
	"net/http"

	appEmailDTOET "github.com/d3ta-go/ddd-mod-email/modules/email/application/dto/email_template"
	"github.com/d3ta-go/ms-email-graphql-api/interface/http-apps/graphql/echo/graphql/resolver/email"
)

// ETListAll list all Email Template
func (r *RootResolver) ETListAll() (*email.ETListAllResolver, error) {
	var etRxs *email.ETListAllResolver

	// identity
	i, err := r.SetIdentity(r.context)
	if err != nil {
		return etRxs, &GraphQLError{Code: http.StatusInternalServerError, Message: err.Error()}
	}
	if !i.IsLogin || i.IsAnonymous {
		return etRxs, &GraphQLError{Code: http.StatusForbidden, Message: "Forbidden Access"}
	}
	i.RequestInfo.RequestObject = "graphql.email.template.list-all"
	i.RequestInfo.RequestAction = "READ"

	// request

	// send to application layer
	resp, err := r.appEmail.EmailTemplateSvc.ListAll(i)
	if err != nil {
		return etRxs, &GraphQLError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	// response
	tmpRes := &email.ETListAllResponse{}
	tmpRes.Count = resp.Count
	for _, r := range resp.Data {
		d := &email.EmailTemplate{
			ID:               r.ID,
			UUID:             r.UUID,
			Code:             r.Code,
			Name:             r.Name,
			IsActive:         r.IsActive,
			EmailFormat:      r.EmailFormat,
			DefaultVersionID: r.DefaultVersionID,
		}
		tmpRes.Data = append(tmpRes.Data, d)
	}
	etRxs = &email.ETListAllResolver{Model: tmpRes}

	return etRxs, nil
}

// ETFindByCode find EmailTemplateByCode
func (r *RootResolver) ETFindByCode(args struct{ Code string }) (*email.ETResolver, error) {
	var etRxs *email.ETResolver

	// identity
	i, err := r.SetIdentity(r.context)
	if err != nil {
		return etRxs, &GraphQLError{Code: http.StatusInternalServerError, Message: err.Error()}
	}
	if !i.IsLogin || i.IsAnonymous {
		return etRxs, &GraphQLError{Code: http.StatusForbidden, Message: "Forbidden Access"}
	}
	i.RequestInfo.RequestObject = "graphql.email.template.find-by-code"
	i.RequestInfo.RequestAction = "READ"

	// request
	req := new(appEmailDTOET.ETFindByCodeReqDTO)
	req.Code = args.Code
	if err := req.Validate(); err != nil {
		return etRxs, &GraphQLError{Code: http.StatusBadRequest, Message: err.Error()}
	}

	// send to application layer
	resp, err := r.appEmail.EmailTemplateSvc.FindByCode(req, i)
	if err != nil {
		return etRxs, &GraphQLError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	// response
	rd := resp.Data
	tmpRes := &email.EmailTemplate{
		ID:               rd.ID,
		UUID:             rd.UUID,
		Code:             rd.Code,
		Name:             rd.Name,
		IsActive:         rd.IsActive,
		EmailFormat:      rd.EmailFormat,
		DefaultVersionID: rd.DefaultVersionID,
		DefaultTemplateVersion: &email.EmailTemplateVersion{
			ID:              rd.DefaultTemplateVersion.ID,
			Version:         rd.DefaultTemplateVersion.Version,
			SubjectTpl:      rd.DefaultTemplateVersion.SubjectTpl,
			BodyTpl:         rd.DefaultTemplateVersion.BodyTpl,
			EmailTemplateID: rd.DefaultTemplateVersion.EmailTemplateID,
		},
	}

	etRxs = &email.ETResolver{Model: tmpRes}

	return etRxs, nil
}
