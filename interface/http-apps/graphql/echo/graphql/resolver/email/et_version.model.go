package email

// EmailTemplateVersion represent EmailTemplateVersion
type EmailTemplateVersion struct {
	ID              uint64 `json:"ID"`
	Version         string `json:"version"`
	SubjectTpl      string `json:"subjectTpl"`
	BodyTpl         string `json:"bodyTpl"`
	EmailTemplateID uint64 `json:"emailTemplateID"`
}
