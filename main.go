package main

import (
	"encoding/json"
	"fmt"

	"github.com/AstraBert/ai-deploy-one-click/commons"
	"github.com/rivo/tview"
)

var config commons.AppConfig

func main() {
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
			app.Stop()
			if jsonData, err := json.MarshalIndent(config, "", "\t"); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println(string(jsonData))
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
