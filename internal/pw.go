package internal

import (
	"time"

	"github.com/playwright-community/playwright-go"
)

func Capture() error {
	return Merge()
	if err := playwright.Install(); err != nil {
		return err
	}

	pw, err := playwright.Run()
	if err != nil {
		return err
	}
	defer pw.Stop()

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		// Headless: playwright.Bool(true),
	})
	if err != nil {
		return err
	}
	defer browser.Close()

	page, err := browser.NewPage(playwright.BrowserNewPageOptions{
		Viewport: &playwright.Size{
			Width: 1280,
			Height: 720,
		},
	})
	if err != nil {
		return err
	}

	if _, err := page.Goto("http://localhost:3000/?slide=1"); err != nil {
		return err
	}

	time.Sleep(5 * time.Second)

	pdfops := playwright.PagePdfOptions{
		Path: playwright.String("output.pdf"),
		Width: playwright.String("1280"),
		Height: playwright.String("800"),
		Outline: playwright.Bool(true),
		PageRanges: playwright.String("1"),
		Margin: &playwright.Margin{
			Top: playwright.String("5"),
			Right: playwright.String("5"),
			Bottom: playwright.String("5"),
			Left: playwright.String("5"),
		},
		PrintBackground: playwright.Bool(true),
	}
	if _, err := page.PDF(pdfops); err != nil {
		return err
	}

	return nil
}
