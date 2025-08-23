package main

import (
	"encoding/json"
	"fmt"
	"path"

	"github.com/AstraBert/ai-deploy-one-click/commons"
	"github.com/AstraBert/ai-deploy-one-click/shell"
	shellops "github.com/AstraBert/ai-deploy-one-click/shell-ops"
	"github.com/rivo/tview"
	"github.com/rvfet/rich-go"
)

var config commons.AppConfig
var configStr string
var configMap map[string]any
var envVar string

func main() {
	sh := shell.DefaultShell()
	checked, err := shellops.SanityCheck(sh)
	if err != nil {
		rich.Error(err.Error())
		return
	} else {
		rich.Info(checked)
	}
	app := tview.NewApplication()
	form := tview.NewForm().
		AddInputField("Application Name", "", 20, nil, config.SetAppName).
		AddInputField("Application URL", "", 20, nil, config.SetAppUrl).
		AddInputField("GitHub Source (username/reponame)", "", 20, nil, config.SetAppGitHubSource).
		AddTextArea("Short Description", "", 40, 0, 0, config.SetAppBriefDescription).
		AddTextArea("Long Description", "", 80, 0, 0, config.SetAppLongDescription).
		AddDropDown("AI Model", []string{"GPT 4.1", "GPT 4o", "GPT 5", "Claude Sonnet 3.5", "Claude Sonnet 3.7", "Claude Sonnet 4", "Gemini 2 Flash", "Gemini 2.5 Flash", "Gemini 2.5 Pro"}, 0, config.SetAppAiModel).
		AddTextArea("System Prompt", "", 100, 0, 0, config.SetAppSystemPrompt).
		AddPasswordField("API Key", "", 10, '*', config.SetAppApiKey).
		AddButton("Save", func() {
			envVar = config.AppApiKey
			app.Stop()
			if jsonData, err := json.MarshalIndent(config, "", "\t"); err != nil {
				fmt.Println("Error:", err)
			} else {
				configStr = string(jsonData)
				json.Unmarshal([]byte(configStr), &configMap)
				delete(configMap, "apiKey")
				if jsonData1, err1 := json.MarshalIndent(configMap, "", "\t"); err1 != nil {
					fmt.Println("Error:", err1)
				} else {
					content := fmt.Sprintf("%s %s\n", commons.TemplateString, string(jsonData1))
					fl := commons.NewFile(".local.config.ts", content)
					fl.WriteContent()
					dotEnv := commons.NewFile(".env", envVar)
					dotEnv.WriteContent()
					ghSucc, ghErr := shellops.CreateGhRepo(config, sh)
					if ghErr != nil {
						rich.Error(ghSucc)
						return
					}
					rich.Info(ghSucc)
					confCopy, confCopyErr := shellops.CopyConfigFile(fl, path.Base(config.AppGitHubSource)+"/lib/app-config.ts")
					if confCopyErr != nil {
						rich.Error(confCopy)
						return
					}
					rich.Info(confCopy)
					_, envCopyErr := shellops.CopyConfigFile(dotEnv, path.Base(config.AppGitHubSource)+"/.env")
					if envCopyErr != nil {
						rich.Error("There was an error copying your environment variable to a .env file, skipping...")
					}
					vercConn, vercConnErr := shellops.VercelConnectGit(config, sh)
					if vercConnErr != nil {
						rich.Error(vercConn)
						return
					}
					rich.Info(vercConn)
				}
			}
		}).
		AddButton("Quit", func() {
			app.Stop()
		})
	form.SetBorder(true).SetTitle("Deploy An AI App Instantly").SetTitleAlign(tview.AlignLeft)
	if err := app.SetRoot(form, true).EnableMouse(true).EnablePaste(true).Run(); err != nil {
		panic(err)
	}
}
