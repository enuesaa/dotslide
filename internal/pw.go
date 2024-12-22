package internal

import (
	"time"

	"github.com/playwright-community/playwright-go"
)

func Capture() error {
	if err := playwright.Install(); err != nil {
		return err
	}

	pw, err := playwright.Run()
	if err != nil {
		return err
	}
	defer pw.Stop()

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	if err != nil {
		return err
	}
	defer browser.Close()

	page, err := browser.NewPage()
	if err != nil {
		return err
	}

	if _, err := page.Goto("http://localhost:3000/"); err != nil {
		return err
	}

	time.Sleep(5 * time.Second)

	pdfops := playwright.PagePdfOptions{
		Path: playwright.String("output.pdf"),
	}
	if _, err := page.PDF(pdfops); err != nil {
		return err
	}

	return nil
}
