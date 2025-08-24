package commons

import (
	"os"
	"slices"
	"testing"
)

func TestAppConfig(t *testing.T) {
	config := &AppConfig{}
	_, isType := any(config).(AppConfigInterface)
	if !isType {
		t.Error("Expected AppConfig to implement AppConfigInterface, but it does not.")
	}
	config.SetAppAiModel("Gemini 2.5 Flash", 0)
	config.SetAppApiKey("test-api-key")
	config.SetAppBriefDescription("Brief")
	config.SetAppLongDescription("Long")
	config.SetAppGitHubSource("AstraBert/test-app")
	config.SetAppName("test-app")
	config.SetAppSystemPrompt("System")
	config.SetAppUrl("https://this-is-not-a-url.com")

	if config.AppAiModel != "Gemini 2.5 Flash" {
		t.Errorf("expected AppAiModel to be 'Gemini 2.5 Flash', got '%s'", config.AppAiModel)
	}
	if config.AppApiKey != "GOOGLE_API_KEY=test-api-key" {
		t.Errorf("expected AppApiKey to be 'GOOGLE_API_KEY=test-api-key', got '%s'", config.AppApiKey)
	}
	if config.AppBriefDescription != "Brief" {
		t.Errorf("expected AppBriefDescription to be 'Brief', got '%s'", config.AppBriefDescription)
	}
	if config.AppLongDescription != "Long" {
		t.Errorf("expected AppLongDescription to be 'Long', got '%s'", config.AppLongDescription)
	}
	if config.AppGitHubSource != "AstraBert/test-app" {
		t.Errorf("expected AppGitHubSource to be 'AstraBert/test-app', got '%s'", config.AppGitHubSource)
	}
	if config.AppName != "test-app" {
		t.Errorf("expected AppName to be 'test-app', got '%s'", config.AppName)
	}
	if config.AppSystemPrompt != "System" {
		t.Errorf("expected AppSystemPrompt to be 'System', got '%s'", config.AppSystemPrompt)
	}
	if config.AppUrl != "https://this-is-not-a-url.com" {
		t.Errorf("expected AppUrl to be 'https://this-is-not-a-url.com', got '%s'", config.AppUrl)
	}
}

func TestFile(t *testing.T) {
	f := NewFile("test-file.txt", "content")
	_, isType := any(f).(FileInterface)
	if !isType {
		t.Error("Expected File to implement FileInterface, but it does not.")
	}
	writeErr := f.WriteContent()
	if writeErr != nil {
		t.Errorf("Expected no error while writing test-file.txt, got %s", writeErr.Error())
	}
	bts, readErr := f.ReadContent()
	if readErr != nil {
		t.Errorf("Expected no error while reading test-file.txt, got %s", readErr.Error())
	}
	if string(bts) != "content" {
		t.Errorf("Expected 'content' when reading from test-file.txt, got %s", string(bts))
	}
	copyErr := f.CopyContent("test-file-copy.txt")
	if copyErr != nil {
		t.Errorf("Expected no error while copying test-file.txt to test-file-copy.txt, got %s", copyErr.Error())
	}
	f1 := NewFile("test-file-copy.txt", "")
	bts1, readErr1 := f1.ReadContent()
	if readErr1 != nil {
		t.Errorf("Expected no error while reading test-file-copy.txt, got %s", readErr1.Error())
	}
	if !slices.Equal(bts, bts1) {
		t.Errorf("Expected read content from test-file.txt and test-file-copy.txt to be the same, but it is not")
	}
	os.Remove("test-file.txt")
	os.Remove("test-file-copy.txt")
}
