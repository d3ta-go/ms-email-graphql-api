package resolver

import (
	"encoding/json"
	"testing"

	"github.com/d3ta-go/ms-email-graphql-api/interface/http-apps/graphql/echo/graphql/resolver/email"
	"github.com/stretchr/testify/assert"
)

func TestRootResolver_SendEmail(t *testing.T) {
	rsx, h, err := newRootResolver(t)
	if err != nil {
		t.Errorf("Error.newRootResolver: %s", err.Error())
	}

	if rsx != nil {

		viper, err := h.GetViper("test-data")
		if err != nil {
			t.Errorf("GetViper: %s", err.Error())
		}
		testData := viper.GetStringMapString("test-data.graphql-api.email.email.interface-layer.features.send.request")
		testDataET := viper.GetStringMapString("test-data.graphql-api.email.email.interface-layer.features.send.request.email-template-data")

		// client request
		jsonReq := `{
    "input": {
		"templateCode": "` + testData["email-template-code"] + `",
		"from": { "email": "` + testData["from-email"] + `", "name": "` + testData["from-name"] + `" },
		"to": { "email": "` + testData["to-email"] + `", "name": "` + testData["to-name"] + `" },
		"cc": [
			{ "email": "` + testData["cc-email-01"] + `", "name": "` + testData["cc-name-01"] + `" },
			{ "email": "` + testData["cc-email-02"] + `", "name": "` + testData["cc-name-02"] + `" }
		],
		"bcc": [
			{ "email": "` + testData["bcc-email-01"] + `", "name": "` + testData["bcc-name-01"] + `" },
			{ "email": "` + testData["bcc-email-02"] + `", "name": "` + testData["bcc-name-02"] + `" }
		],
		"templateData": [
			{ "fieldName": "Header.Name", "fieldValue" : "` + testDataET["header-name"] + `"},
			{ "fieldName": "Body.UserAccount", "fieldValue": "` + testDataET["body-user-account"] + `"},
			{ "fieldName": "Body.ActivationURL", "fieldValue": "` + testDataET["body-activation-url"] + `"},
			{ "fieldName": "Footer.Name", "fieldValue": "` + testDataET["footer-name"] + `"}
		],
		"processingType": "` + testData["processing-type"] + `"
	}
}`

		var req1 email.SendEmailRequest
		if err := json.Unmarshal([]byte(jsonReq), &req1); err != nil {
			t.Errorf("jsonReq.Unmarshal: %s", err.Error())
			return
		}

		type args struct {
			args email.SendEmailRequest
		}
		tests := []struct {
			name    string
			args    args
			want    *email.SendEmailResResolver
			wantErr bool
		}{
			// TODO: Add test cases.
			{
				name:    "OK",
				args:    args{args: req1},
				wantErr: false,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {

				got, err := rsx.SendEmail(tt.args.args)
				if (err != nil) != tt.wantErr {
					t.Errorf("RootResolver.SendEmail() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				/*
					if !reflect.DeepEqual(got.Model.TemplateCode, tt.want.Model.TemplateCode) {
						t.Errorf("RootResolver.SendEmail() = %v, want %v", got.Model.TemplateCode, tt.want.Model.TemplateCode)
					}
				*/
				if assert.NoError(t, err) {
					if assert.NotNil(t, got) {
						viper.Set("test-data.graphql-api.email.email.interface-layer.features.send.response.json", string(got.Model.ToJSON()))
						if err := viper.WriteConfig(); err != nil {
							t.Errorf("Error: viper.WriteConfig(), %s", err.Error())
						}
						t.Logf("RESPONSE.graphql.SendEmail: %s", got.Model.ToJSON())
					}
				}
			})
		}
	}
}
