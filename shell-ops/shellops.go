package shellops

import (
	"errors"
	"fmt"
	"path"
	"strings"

	"github.com/AstraBert/ai-deploy-one-click/commons"
	"github.com/AstraBert/ai-deploy-one-click/shell"
)

const RepoTemplateURL string = "https://github.com/AstraBert/ai-deploy-one-click-ui"

func SanityCheck(sh *shell.Shell) (string, error) {
	_, errVer := sh.Execute("vercel --help")
	_, errGh := sh.Execute("gh --help")
	_, errGit := sh.Execute("git --help")
	switch {
	case errVer != nil || errGh != nil || errGit != nil:
		return "You should install and log-in to Vercel CLI (vercel), GitHub CLI (gh) and Git (git) before using ai-deploy-one-click, but seems like you are missing at least one!", errors.New("you should install and log-in to Vercel CLI (vercel), GitHub CLI (gh) and Git (git) before using ai-deploy-one-click, but seems like you are missing at least one")
	default:
		return "System check successfully passed: you have all the necessary software installed", nil
	}
}

func CreateGhRepo(config commons.AppConfig, sh *shell.Shell) (string, error) {
	repoSource := config.AppGitHubSource
	command := fmt.Sprintf("gh repo create %s -p %s --public --clone", repoSource, RepoTemplateURL)
	_, err := sh.Execute(command)
	if err != nil {
		return "Repository creation failed, please check your permission to create this GitHub repository and try again.\nError:" + err.Error(), err
	} else {
		return "Repository created successfully!", nil
	}
}

func CopyConfigFile(configFile *commons.File, destinationFile string) (string, error) {
	err := configFile.CopyContent(destinationFile)
	if err != nil {
		return "Sorry, an error occurred while copying the file to the specified destination\nError: " + err.Error(), err
	} else {
		return fmt.Sprintf("Successfully copied %s to %s", configFile.Path, destinationFile), nil
	}
}

func VercelConnectGit(config commons.AppConfig, envVar string, sh *shell.Shell) (string, error) {
	pathDir := path.Base(config.AppGitHubSource)
	_, err := sh.Execute("cd " + pathDir + " && vercel git connect --yes && vercel env add " + strings.Split(envVar, "=")[0] + " production < .env.value && git add . && git commit -m 'feat: automatic first commit from ai-deploy-one-click' && git push origin main")
	if err != nil {
		return "An error occurred while connecting your repository to Vercel: " + err.Error(), err
	} else {
		return "Vercel has been connected successfully to your GitHub repository! Continue on vercel.com or on github.com/" + config.AppGitHubSource, nil
	}
}
