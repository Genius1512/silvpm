//go:build windows
// +build windows

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func downloadApplication(appName string) error {
	resp, err := http.Get(fmt.Sprintf(
		"%sapplications/%s/%s.exe",
		baseURL,
		appName,
		appName,
	))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	os.Mkdir(os.Getenv("appdata")+"\\silvpm", os.ModePerm)
	f, err := os.Create(os.Getenv("appdata") + "\\silvpm\\" + appName + ".exe")
	if err != nil {
		return err
	}
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}
	defer f.Close()

	return nil
}
