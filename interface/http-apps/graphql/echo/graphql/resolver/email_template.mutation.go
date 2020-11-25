package resolver

import (
	"encoding/json"
	"net/http"

	appEmailDTOET "github.com/d3ta-go/ddd-mod-email/modules/email/application/dto/email_template"
	"github.com/d3ta-go/ms-email-graphql-api/interface/http-apps/graphql/echo/graphql/resolver/email"
)

// ETCreate represent ETCreate
func (r *RootResolver) ETCreate(args email.ETCreateRequest) (*email.ETCreateResResolver, error) {
	var etRxs *email.ETCreateResResolver

	// identity
	i, err := r.SetIdentity(r.context)
	if err != nil {
		return etRxs, &GraphQLError{Code: http.StatusInternalServerError, Message: err.Error()}
	}
	if !i.IsLogin || i.IsAnonymous {
		return etRxs, &GraphQLError{Code: http.StatusForbidden, Message: "Forbidden Access"}
	}
	i.RequestInfo.RequestObject = "graphql.email.template.create"
	i.RequestInfo.RequestAction = "CREATE"

	// request
	jsonInput := args.Input.ToJSON()
	var req appEmailDTOET.ETCreateReqDTO
	if err := json.Unmarshal(jsonInput, &req); err != nil {
		return etRxs, &GraphQLError{Code: http.StatusBadRequest, Message: err.Error()}
	}
	if err := req.Validate(); err != nil {
		return etRxs, &GraphQLError{Code: http.StatusBadRequest, Message: err.Error()}
	}

	// send to application service
	resp, err := r.appEmail.EmailTemplateSvc.Create(&req, i)
	if err != nil {
		return etRxs, &GraphQLError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	// response
	etRxs = &email.ETCreateResResolver{Model: resp}

	return etRxs, nil
}

// ETUpdate update existing EmailTemplate
func (r *RootResolver) ETUpdate(args email.ETUpdateRequest) (*email.ETUpdateResResolver, error) {
	var etRxs *email.ETUpdateResResolver

	// identity
	i, err := r.SetIdentity(r.context)
	if err != nil {
		return etRxs, &GraphQLError{Code: http.StatusInternalServerError, Message: err.Error()}
	}
	if !i.IsLogin || i.IsAnonymous {
		return etRxs, &GraphQLError{Code: http.StatusForbidden, Message: "Forbidden Access"}
	}
	i.RequestInfo.RequestObject = "graphql.email.template.update"
	i.RequestInfo.RequestAction = "UPDATE"

	// request
	jsonInput := args.ToJSON()
	var req appEmailDTOET.ETUpdateReqDTO
	if err := json.Unmarshal(jsonInput, &req); err != nil {
		return etRxs, &GraphQLError{Code: http.StatusBadRequest, Message: err.Error()}
	}

	// send to application service
	resp, err := r.appEmail.EmailTemplateSvc.Update(&req, i)
	if err != nil {
		return etRxs, &GraphQLError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	// response
	etRxs = &email.ETUpdateResResolver{Model: resp}

	return etRxs, nil
}

// ETSetActive set existing EmailTemplate active status
func (r *RootResolver) ETSetActive(args email.ETSetActiveRequest) (*email.ETSetActiveResResolver, error) {
	var etRxs *email.ETSetActiveResResolver

	// identity
	i, err := r.SetIdentity(r.context)
	if err != nil {
		return etRxs, &GraphQLError{Code: http.StatusInternalServerError, Message: err.Error()}
	}
	if !i.IsLogin || i.IsAnonymous {
		return etRxs, &GraphQLError{Code: http.StatusForbidden, Message: "Forbidden Access"}
	}
	i.RequestInfo.RequestObject = "graphql.email.template.set-active"
	i.RequestInfo.RequestAction = "UPDATE"

	// request
	jsonInput := args.ToJSON()
	var req appEmailDTOET.ETSetActiveReqDTO
	if err := json.Unmarshal(jsonInput, &req); err != nil {
		return etRxs, &GraphQLError{Code: http.StatusBadRequest, Message: err.Error()}
	}

	// send to application service
	resp, err := r.appEmail.EmailTemplateSvc.SetActive(&req, i)
	if err != nil {
		return etRxs, &GraphQLError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	// response
	etRxs = &email.ETSetActiveResResolver{Model: resp}

	return etRxs, nil
}

// ETDelete delete existing EmailTemplate
func (r *RootResolver) ETDelete(args email.ETDeleteRequest) (*email.ETDeleteResResolver, error) {
	var etRxs *email.ETDeleteResResolver

	// identity
	i, err := r.SetIdentity(r.context)
	if err != nil {
		return etRxs, &GraphQLError{Code: http.StatusInternalServerError, Message: err.Error()}
	}
	if !i.IsLogin || i.IsAnonymous {
		return etRxs, &GraphQLError{Code: http.StatusForbidden, Message: "Forbidden Access"}
	}
	i.RequestInfo.RequestObject = "graphql.email.template.delete"
	i.RequestInfo.RequestAction = "DELETE"

	// request
	jsonInput := args.ToJSON()
	var req appEmailDTOET.ETDeleteReqDTO
	if err := json.Unmarshal(jsonInput, &req); err != nil {
		return etRxs, &GraphQLError{Code: http.StatusBadRequest, Message: err.Error()}
	}
	if err := req.Validate(); err != nil {
		return etRxs, &GraphQLError{Code: http.StatusBadRequest, Message: err.Error()}
	}

	// send to application service
	resp, err := r.appEmail.EmailTemplateSvc.Delete(&req, i)
	if err != nil {
		return etRxs, &GraphQLError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	// response
	etRxs = &email.ETDeleteResResolver{Model: resp}

	return etRxs, nil
}
