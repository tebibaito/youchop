package main

import (
	"context"
	"fmt"
	"youchop/ytchat"
	"youchop/ytchop"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) domready(ctx context.Context) {
	runtime.WindowSetMinSize(ctx, 1024, 768)
	runtime.WindowSetMaxSize(ctx, 1024, 768)
}

func (a *App) Download(url string, chopDataList []ytchop.ChopData) string {
	runtime.LogPrint(a.ctx, "Download Start")
	var option runtime.OpenDialogOptions
	option.CanCreateDirectories = true
	option.Title = "保存先を選択"
	dir, _ := runtime.OpenDirectoryDialog(a.ctx, option)
	ytchop.Download(url, dir, chopDataList)
	return "Download Done!"
}

func (a *App) GetChatCount(url string) []ytchat.CountData {
	result, err := ytchat.GetChatCount(url)
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
	}
	return result
}
