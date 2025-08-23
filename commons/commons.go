package commons

import (
	"os"
	"strings"
)

const TemplateString string = `import type { AIModelKey } from './ai-models';

export const appConfig: {
    name: string;
    shortDescription: string;
    extendedDescription: string;
    aiModel: AIModelKey;
    systemPrompt: string;
    appUrl: string;
    gitHubSource: string;
} =`

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

type FileInterface interface {
	WriteContent() error
	ReadContent() ([]byte, error)
	CopyContent(string) error
}

type File struct {
	Path    string
	Content string
}

func (f *File) WriteContent() error {
	return os.WriteFile(f.Path, []byte(f.Content), 0777)
}

func (f *File) ReadContent() ([]byte, error) {
	content, err := os.ReadFile(f.Path)
	if err != nil {
		return nil, err
	} else {
		return content, nil
	}
}

func (f *File) CopyContent(filePath string) error {
	content, err := f.ReadContent()
	if err != nil {
		return err
	} else {
		fl := NewFile(filePath, string(content))
		errFl := fl.WriteContent()
		if errFl != nil {
			return errFl
		}
		return nil
	}
}

func NewFile(path, content string) *File {
	return &File{
		Path:    path,
		Content: content,
	}
}
