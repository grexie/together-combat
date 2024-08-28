package auth

import (
	"regexp"

	html "html/template"
	text "text/template"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CampaignAction = string

var togetherCampaignActionPrefix = "together:"
var togetherCampaignActionRegexp = regexp.MustCompilePOSIX("^together:")
var togetherCampaignActionNamingConventionRegex = regexp.MustCompilePOSIX("^[a-z0-9\\-_]+(:[a-z0-9\\-_]+)?$")

const (
	CampaignActionChallenge CampaignAction = "together:challenge"
)

type Campaign interface {
	ID() primitive.ObjectID
	SetID(id primitive.ObjectID)
	WithID(id primitive.ObjectID) Campaign

	App() primitive.ObjectID
	SetApp(app primitive.ObjectID)
	WithApp(app primitive.ObjectID) Campaign

	Action() *CampaignAction
	SetAction(action CampaignAction)
	UnsetAction()
	WithAction(action CampaignAction) Campaign
	WithoutAction() Campaign

	Name() string
	SetName(name string)
	WithName(name string) Campaign

	SubjectTemplate() string
	SubjectTemplateParsed() *text.Template
	SetSubjectTemplate(subjectTemplate string)
	WithSubjectTemplate(subjectTemplate string) Campaign

	HTMLTemplate() string
	HTMLTemplateParsed() *html.Template
	SetHTMLTemplate(htmlTemplate string)
	WithHTMLTemplate(htmlTemplate string) Campaign

	TextTemplate() string
	TextTemplateParsed() *text.Template
	SetTextTemplate(textTemplate string)
	WithTextTemplate(textTemplate string) Campaign
}

type CreateCampaignOptions interface {
	App() primitive.ObjectID
	SetApp(app primitive.ObjectID)
	WithApp(app primitive.ObjectID) CreateCampaignOptions

	Action() *CampaignAction
	SetAction(action CampaignAction)
	UnsetAction()
	WithAction(action CampaignAction) CreateCampaignOptions
	WithoutAction() CreateCampaignOptions

	Name() string
	SetName(name string)
	WithName(name string) CreateCampaignOptions

	SubjectTemplate() string
	SetSubjectTemplate(subjectTemplate string)
	WithSubjectTemplate(subjectTemplate string) CreateCampaignOptions

	HTMLTemplate() string
	SetHTMLTemplate(htmlTemplate string)
	WithHTMLTemplate(htmlTemplate string) CreateCampaignOptions

	TextTemplate() string
	SetTextTemplate(textTemplate string)
	WithTextTemplate(textTemplate string) CreateCampaignOptions
}

type UpdateCampaignOptions interface {
	Name() *string
	SetName(name string)
	UnsetName()
	WithName(name string) UpdateCampaignOptions
	WithoutName() UpdateCampaignOptions

	SubjectTemplate() *string
	SetSubjectTemplate(subjectTemplate string)
	UnsetSubjectTemplate()
	WithSubjectTemplate(subjectTemplate string) UpdateCampaignOptions
	WithoutSubjectTemplate() UpdateCampaignOptions

	HTMLTemplate() *string
	SetHTMLTemplate(htmlTemplate string)
	UnsetHTMLTemplate()
	WithHTMLTemplate(htmlTemplate string) UpdateCampaignOptions
	WithoutHTMLTemplate() UpdateCampaignOptions

	TextTemplate() *string
	SetTextTemplate(textTemplate string)
	UnsetTextTemplate()
	WithTextTemplate(textTemplate string) UpdateCampaignOptions
	WithoutTextTemplate() UpdateCampaignOptions
}