package resolver

import (
	"net/http"

	appEmailDTO "github.com/d3ta-go/ddd-mod-email/modules/email/application/dto/email"
	"github.com/d3ta-go/ms-email-graphql-api/interface/http-apps/graphql/echo/graphql/resolver/email"
)

// SendEmail send Email
func (r *RootResolver) SendEmail(args email.SendEmailRequest) (*email.SendEmailResResolver, error) {
	var etRxs *email.SendEmailResResolver

	// identity
	i, err := r.SetIdentity(r.context)
	if err != nil {
		return etRxs, &GraphQLError{Code: http.StatusInternalServerError, Message: err.Error()}
	}
	if !i.IsLogin || i.IsAnonymous {
		return etRxs, &GraphQLError{Code: http.StatusForbidden, Message: "Forbidden Access"}
	}
	i.RequestInfo.RequestObject = "graphql.email.send"
	i.RequestInfo.RequestAction = "EXECUTE"

	// request
	req := appEmailDTO.SendEmailReqDTO{
		TemplateCode:   args.Input.TemplateCode,
		From:           &appEmailDTO.MailAddressDTO{Email: args.Input.From.Email, Name: args.Input.From.Name},
		To:             &appEmailDTO.MailAddressDTO{Email: args.Input.To.Email, Name: args.Input.To.Name},
		CC:             args.Input.TranslateCC(),
		BCC:            args.Input.TranslateBCC(),
		TemplateData:   args.Input.TranslateTD(),
		ProcessingType: args.Input.ProcessingType,
	}

	// send to application service
	resp, err := r.appEmail.EmailSvc.Send(&req, i)
	if err != nil {
		return etRxs, &GraphQLError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	// response
	etRxs = &email.SendEmailResResolver{Model: resp}

	return etRxs, nil
}
