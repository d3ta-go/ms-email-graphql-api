package email

// EmailTemplate represent EmailTemplate
type EmailTemplate struct {
	ID                     uint64                `json:"ID"`
	UUID                   string                `json:"uuid"`
	Code                   string                `json:"code"`
	Name                   string                `json:"name"`
	IsActive               bool                  `json:"isActive"`
	EmailFormat            string                `json:"emailFormat"`
	DefaultVersionID       uint64                `json:"defaultVersionID"`
	DefaultTemplateVersion *EmailTemplateVersion `json:"defaultTemplate"`
}
