package internal

import (
	"github.com/playwright-community/playwright-go"
)

func Capture() error {
	installops := playwright.RunOptions{
		Browsers: []string{"chromium"},
	}
	if err := playwright.Install(&installops); err != nil {
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

	if _, err := page.Goto("https://example.com"); err != nil {
		return err
	}

	pdfops := playwright.PagePdfOptions{
		Path:   playwright.String("output.pdf"),
		Format: playwright.String("A4"),
	}

	if _, err := page.PDF(pdfops); err != nil {
		return err
	}

	return nil
}
