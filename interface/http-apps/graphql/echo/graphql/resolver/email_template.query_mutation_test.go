package resolver

import (
	"encoding/json"
	"fmt"
	"testing"

	appEmailDTOET "github.com/d3ta-go/ddd-mod-email/modules/email/application/dto/email_template"
	"github.com/d3ta-go/ms-email-graphql-api/interface/http-apps/graphql/echo/graphql/resolver/email"
	"github.com/d3ta-go/system/system/utils"
	"github.com/stretchr/testify/assert"
)

func TestRootResolver_ETListAll(t *testing.T) {
	rsx, _, err := newRootResolver(t)
	if err != nil {
		t.Errorf("Error.newRootResolver: %s", err.Error())
	}

	if rsx != nil {
		tests := []struct {
			name    string
			want    *email.ETListAllResolver
			wantErr bool
		}{
			// TODO: Add test cases.
			{
				name:    "OK",
				want:    &email.ETListAllResolver{Model: &appEmailDTOET.ETListAllResDTO{}},
				wantErr: false,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {

				got, err := rsx.ETListAll()
				if (err != nil) != tt.wantErr {
					t.Errorf("RootResolver.ETListAll() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				/*
					if !reflect.DeepEqual(got, tt.want) {
						t.Errorf("RootResolver.ETListAll() = %v, want %v", got, tt.want)
					}
				*/
				// Assertions
				if assert.NoError(t, err) {
					if assert.NotNil(t, got) {
						t.Logf("RESPONSE.graphql.ETListAll: %s", got.Model.ToJSON())
					}
				}
			})
		}
	}
}

func TestRootResolver_ETCreate(t *testing.T) {
	rsx, h, err := newRootResolver(t)
	if err != nil {
		t.Errorf("Error.newRootResolver: %s", err.Error())
	}

	if rsx != nil {

		viper, err := h.GetViper("test-data")
		if err != nil {
			t.Errorf("GetViper: %s", err.Error())
		}
		testData := viper.GetStringMapString("test-data.graphql-api.email.email-template.interface-layer.features.create.request")

		unique := utils.GenerateUUID()
		etCode := fmt.Sprintf(testData["et-code"], unique)
		etName := fmt.Sprintf(testData["et-name"], unique)
		// client request
		jsonReq := `{
	"input": {
		"code": "` + etCode + `",
		"name": "` + etName + `",
		"isActive": ` + testData["et-is-active"] + `,
		"emailFormat": "` + testData["et-email-format"] + `",
		"template": {
			"subjectTpl": "` + testData["et-tpl-subject"] + `",
			"bodyTpl": "` + testData["et-tpl-body"] + `"
		}
	}
}`

		var req1 email.ETCreateRequest
		if err := json.Unmarshal([]byte(jsonReq), &req1); err != nil {
			t.Errorf("jsonReq.Unmarshal: %s", err.Error())
			return
		}

		type args struct {
			args email.ETCreateRequest
		}
		tests := []struct {
			name    string
			args    args
			want    *email.ETCreateResResolver
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

				got, err := rsx.ETCreate(tt.args.args)
				if (err != nil) != tt.wantErr {
					t.Errorf("RootResolver.ETCreate() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if assert.NoError(t, err) {
					if assert.NotNil(t, got) {

						viper.Set("test-data.graphql-api.email.email-template.interface-layer.features.create.response.json", string(got.Model.ToJSON()))
						viper.Set("test-data.graphql-api.email.email-template.interface-layer.features.find-by-code.request.et-code", etCode)

						viper.Set("test-data.graphql-api.email.email-template.interface-layer.features.update.request.et-code", etCode)
						viper.Set("test-data.graphql-api.email.email-template.interface-layer.features.update.request.et-name", etName)

						viper.Set("test-data.graphql-api.email.email-template.interface-layer.features.delete.request.et-code", etCode)
						viper.Set("test-data.graphql-api.email.email-template.interface-layer.features.set-active.request.et-code", etCode)
						if err := viper.WriteConfig(); err != nil {
							t.Errorf("Error: viper.WriteConfig(), %s", err.Error())
						}

						t.Logf("RESPONSE.graphql.ETCreate: %s", got.Model.ToJSON())
					}
				}
			})
		}
	}
}

func TestRootResolver_ETFindByCode(t *testing.T) {
	rsx, h, err := newRootResolver(t)
	if err != nil {
		t.Errorf("Error.newRootResolver: %s", err.Error())
	}

	if rsx != nil {
		viper, err := h.GetViper("test-data")
		if err != nil {
			t.Errorf("GetViper: %s", err.Error())
		}
		testData := viper.GetStringMapString("test-data.graphql-api.email.email-template.interface-layer.features.find-by-code.request")

		type args struct {
			args email.ETFindByCodeRequest
		}

		args1 := email.ETFindByCodeRequest{}
		args1.Code = testData["et-code"]

		tests := []struct {
			name    string
			args    args
			want    *email.ETFindByCodeResolver
			wantErr bool
		}{
			// TODO: Add test cases.
			{
				name:    "OK",
				args:    args{args: args1},
				want:    &email.ETFindByCodeResolver{Model: &appEmailDTOET.ETFindByCodeResDTO{}},
				wantErr: false,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {

				got, err := rsx.ETFindByCode(tt.args.args)
				if (err != nil) != tt.wantErr {
					t.Errorf("RootResolver.ETFindByCode() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if assert.NoError(t, err) {
					if assert.NotNil(t, got) {

						viper.Set("test-data.graphql-api.email.email-template.interface-layer.features.find-by-code.response.json", string(got.Model.ToJSON()))
						if err := viper.WriteConfig(); err != nil {
							t.Errorf("Error: viper.WriteConfig(), %s", err.Error())
						}

						t.Logf("RESPONSE.graphql.ETFindByCode: %s", got.Model.ToJSON())
					}
				}

			})
		}
	}
}

func TestRootResolver_ETUpdate(t *testing.T) {
	rsx, h, err := newRootResolver(t)
	if err != nil {
		t.Errorf("Error.newRootResolver: %s", err.Error())
	}

	if rsx != nil {

		viper, err := h.GetViper("test-data")
		if err != nil {
			t.Errorf("GetViper: %s", err.Error())
		}
		testData := viper.GetStringMapString("test-data.graphql-api.email.email-template.interface-layer.features.update.request")

		// client request
		jsonReq := `{
	"keys": { "code": "` + testData["et-code"] + `" },
	"data": {
		"name": "` + testData["et-name"] + ` Updated",
		"isActive": ` + testData["et-is-active"] + `,
		"emailFormat": "` + testData["et-email-format"] + `",
		"template": {
			"subjectTpl": "` + testData["et-tpl-subject"] + `",
			"bodyTpl": "` + testData["et-tpl-body"] + `"
		}
	}
}`

		var req1 email.ETUpdateRequest
		if err := json.Unmarshal([]byte(jsonReq), &req1); err != nil {
			t.Errorf("jsonReq.Unmarshal: %s", err.Error())
			return
		}

		type args struct {
			args email.ETUpdateRequest
		}
		tests := []struct {
			name    string
			args    args
			want    *email.ETUpdateResResolver
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

				got, err := rsx.ETUpdate(tt.args.args)
				if (err != nil) != tt.wantErr {
					t.Errorf("RootResolver.ETUpdate() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if assert.NoError(t, err) {
					if assert.NotNil(t, got) {

						viper.Set("test-data.graphql-api.email.email-template.interface-layer.features.update.response.json", string(got.Model.ToJSON()))
						if err := viper.WriteConfig(); err != nil {
							t.Errorf("Error: viper.WriteConfig(), %s", err.Error())
						}

						t.Logf("RESPONSE.graphql.ETUpdate: %s", got.Model.ToJSON())
					}
				}
			})
		}
	}
}

func TestRootResolver_ETSetActive(t *testing.T) {
	rsx, h, err := newRootResolver(t)
	if err != nil {
		t.Errorf("Error.newRootResolver: %s", err.Error())
	}

	if rsx != nil {

		viper, err := h.GetViper("test-data")
		if err != nil {
			t.Errorf("GetViper: %s", err.Error())
		}
		testData := viper.GetStringMapString("test-data.graphql-api.email.email-template.interface-layer.features.set-active.request")

		// client request
		jsonReq := `{
	"keys": { "code": "` + testData["et-code"] + `" },
	"data": {
		"isActive": ` + testData["et-is-active"] + `
	}
}`

		var req1 email.ETSetActiveRequest
		if err := json.Unmarshal([]byte(jsonReq), &req1); err != nil {
			t.Errorf("jsonReq.Unmarshal: %s", err.Error())
			return
		}

		type args struct {
			args email.ETSetActiveRequest
		}
		tests := []struct {
			name    string
			args    args
			want    *email.ETSetActiveResResolver
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

				got, err := rsx.ETSetActive(tt.args.args)
				if (err != nil) != tt.wantErr {
					t.Errorf("RootResolver.ETSetActive() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if assert.NoError(t, err) {
					if assert.NotNil(t, got) {

						viper.Set("test-data.graphql-api.email.email-template.interface-layer.features.set-active.response.json", string(got.Model.ToJSON()))
						if err := viper.WriteConfig(); err != nil {
							t.Errorf("Error: viper.WriteConfig(), %s", err.Error())
						}

						t.Logf("RESPONSE.graphql.ETSetActive: %s", got.Model.ToJSON())
					}
				}
			})
		}
	}
}

func TestRootResolver_ETDelete(t *testing.T) {
	rsx, h, err := newRootResolver(t)
	if err != nil {
		t.Errorf("Error.newRootResolver: %s", err.Error())
	}

	if rsx != nil {

		viper, err := h.GetViper("test-data")
		if err != nil {
			t.Errorf("GetViper: %s", err.Error())
		}
		testData := viper.GetStringMapString("test-data.graphql-api.email.email-template.interface-layer.features.delete.request")

		// client request
		jsonReq := `{
	"code": "` + testData["et-code"] + `"
}`

		var req1 email.ETDeleteRequest
		if err := json.Unmarshal([]byte(jsonReq), &req1); err != nil {
			t.Errorf("jsonReq.Unmarshal: %s", err.Error())
			return
		}

		type args struct {
			args email.ETDeleteRequest
		}
		tests := []struct {
			name    string
			args    args
			want    *email.ETDeleteResResolver
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

				got, err := rsx.ETDelete(tt.args.args)
				if (err != nil) != tt.wantErr {
					t.Errorf("RootResolver.ETDelete() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if assert.NoError(t, err) {
					if assert.NotNil(t, got) {

						viper.Set("test-data.graphql-api.email.email-template.interface-layer.features.delete.response.json", string(got.Model.ToJSON()))
						if err := viper.WriteConfig(); err != nil {
							t.Errorf("Error: viper.WriteConfig(), %s", err.Error())
						}

						t.Logf("RESPONSE.graphql.ETDelete: %s", got.Model.ToJSON())
					}
				}
			})
		}
	}
}
