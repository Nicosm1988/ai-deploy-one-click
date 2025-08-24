package shellops

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/AstraBert/ai-deploy-one-click/commons"
	"github.com/AstraBert/ai-deploy-one-click/shell"
)

func errorRunFactory(command string) (string, error) {
	return "", errors.New("there was an error when running " + command)
}

func succRunFactory(command string) (string, error) {
	return command, nil
}

func TestSanityCheck(t *testing.T) {
	sh := shell.NewShell(errorRunFactory)
	sh1 := shell.NewShell(succRunFactory)
	retVal, err := SanityCheck(sh)
	retVal1, err1 := SanityCheck(sh1)
	if retVal != "You should install and log-in to Vercel CLI (vercel), GitHub CLI (gh) and Git (git) before using ai-deploy-one-click, but seems like you are missing at least one!" && err == nil {
		t.Error("Expecting an error to be reported during SanityCheck, got none")
	}
	if retVal1 != "System check successfully passed: you have all the necessary software installed" && err1 != nil {
		t.Error("Expecting no error during SanityCheck execution, got: " + err1.Error())
	}
}

func TestGhRepoCreation(t *testing.T) {
	config := commons.AppConfig{}
	sh := shell.NewShell(errorRunFactory)
	sh1 := shell.NewShell(succRunFactory)
	retVal, err := CreateGhRepo(config, sh)
	retVal1, err1 := CreateGhRepo(config, sh1)
	if !strings.HasPrefix(retVal, "Repository creation failed") && err == nil {
		t.Error("Expecting an error to be reported during CreateGhRepo, got none")
	}
	if retVal1 != "Repository created successfully!" && err1 != nil {
		t.Error("Expecting no error during SanityCheck execution, got: " + err1.Error())
	}
}

func TestCopyConfigFile(t *testing.T) {
	f := commons.NewFile("test-file.txt", "content")
	f.WriteContent()
	succ, err := CopyConfigFile(f, "test-file-copy.txt")
	if err != nil {
		t.Errorf("Expected no error while testing CopyConfigFile, go %s", err.Error())
	}
	if succ != fmt.Sprintf("Successfully copied %s to %s", f.Path, "test-file-copy.txt") {
		t.Errorf("Expecting success message: 'Successfully copied %s to %s', got %s", f.Path, "test-file-copy.txt", succ)
	}
	os.Remove("test-file.txt")
	os.Remove("test-file-copy.txt")
}

func TestVercelConnectGit(t *testing.T) {
	config := commons.AppConfig{}
	sh := shell.NewShell(errorRunFactory)
	sh1 := shell.NewShell(succRunFactory)
	retVal, err := VercelConnectGit(config, ".env", sh)
	retVal1, err1 := VercelConnectGit(config, ".env", sh1)
	if !strings.HasPrefix(retVal, "An error occurred while connecting your repository to Vercel:") && err == nil {
		t.Error("Expecting an error to be reported during VercelConnectGit, got none")
	}
	if !strings.HasPrefix(retVal1, "Vercel has been connected successfully to your GitHub repository!") && err1 != nil {
		t.Error("Expecting no error during SanityCheck execution, got: " + err1.Error())
	}
}
