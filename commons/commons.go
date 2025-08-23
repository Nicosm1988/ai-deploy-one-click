package commons

import "strings"

type AppConfigInterface interface {
	SetAppName(string)
	SetAppBriefDescription(string)
	SetAppLongDescription(string)
	SetAppSystemPrompt(string)
	SetAppUrl(string)
	SetAppGitHubSource(string)
	SetAppAiModel(string, int)
	SetAppApiKey(string)
}

type AppConfig struct {
	AppName             string `json:"name"`
	AppBriefDescription string `json:"shortDescription"`
	AppLongDescription  string `json:"extendedDescription"`
	AppUrl              string `json:"appUrl"`
	AppGitHubSource     string `json:"gitHubSource"`
	AppAiModel          string `json:"aiModel"`
	AppSystemPrompt     string `json:"systemPrompt"`
	AppApiKey           string `json:"apiKey"`
}

func (a *AppConfig) SetAppName(text string) {
	a.AppName = text
}

func (a *AppConfig) SetAppUrl(text string) {
	a.AppUrl = text
}

func (a *AppConfig) SetAppLongDescription(text string) {
	a.AppLongDescription = text
}

func (a *AppConfig) SetAppBriefDescription(text string) {
	a.AppBriefDescription = text
}

func (a *AppConfig) SetAppAiModel(option string, optionIndex int) {
	a.AppAiModel = option
}

func (a *AppConfig) SetAppGitHubSource(text string) {
	a.AppGitHubSource = text
}

func (a *AppConfig) SetAppApiKey(text string) {
	switch {
	case strings.HasPrefix(a.AppAiModel, "Gemini"):
		a.AppApiKey = "GOOGLE_API_KEY=" + text
	case strings.HasPrefix(a.AppAiModel, "Claude"):
		a.AppApiKey = "ANTHROPIC_API_KEY=" + text
	default:
		a.AppApiKey = "OPENAI_API_KEY=" + text
	}
}

func (a *AppConfig) SetAppSystemPrompt(text string) {
	a.AppSystemPrompt = text
}
