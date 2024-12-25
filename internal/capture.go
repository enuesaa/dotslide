package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/enuesaa/dotslide/internal/slides"
	pdfapi "github.com/pdfcpu/pdfcpu/pkg/api"
	pdfmodel "github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/playwright-community/playwright-go"
	"gopkg.in/yaml.v3"
)

func NewCaptureCtl(config Config) CaptureCtl {
	return CaptureCtl{
		config: config,
	}
}

type CaptureCtl struct {
	config Config

	pw *playwright.Playwright
	browser playwright.Browser
}

func (c *CaptureCtl) Capture() error {
	if err := os.Mkdir(".capture", os.ModePerm); err != nil {
		return err
	}
	defer os.RemoveAll(".capture")

	if err := c.setup(); err != nil {
		return err
	}
	if err := c.launch(); err != nil {
		return err
	}
	defer c.close()

	slideLength, err := c.countSlideLength()
	if err != nil {
		return err
	}

	for slideNumbr := range slideLength {
		if err := c.captureSlide(slideNumbr); err != nil {
			return err
		}
	}

	if err := c.mergeCaptures(slideLength); err != nil {
		return err
	}
	return nil
}

func (c *CaptureCtl) setup() error {
	return playwright.Install()
}

func (c *CaptureCtl) close() error {
	if c.browser != nil {
		if err := c.browser.Close(); err != nil {
			return err
		}
		c.browser = nil
	}
	if c.pw != nil {
		if err := c.pw.Stop(); err != nil {
			return err
		}
		c.pw = nil
	}
	return nil
}

func (c *CaptureCtl) launch() error {
	pw, err := playwright.Run()
	if err != nil {
		return err
	}
	c.pw = pw

	browser, err := c.pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	if err != nil {
		return err
	}
	c.browser = browser

	return nil
}

func (c *CaptureCtl) captureSlide(slideNumber int) error {
	if c.browser == nil {
		return fmt.Errorf("browser not launched")
	}
	page, err := c.browser.NewPage(playwright.BrowserNewPageOptions{
		Viewport: &playwright.Size{
			Width: 1500,
			Height: 720,
		},
	})
	if err != nil {
		return err
	}
	defer page.Close()

	url := fmt.Sprintf("http://localhost:%d/?slide=%d", c.config.Port, slideNumber)
	if _, err := page.Goto(url); err != nil {
		return err
	}
	time.Sleep(5 * time.Second)

	pdfops := playwright.PagePdfOptions{
		Path: playwright.String(c.fmtCaptureOutput(slideNumber)),
		Width: playwright.String("1500"),
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

func (c *CaptureCtl) fmtCaptureOutput(slideNumber int) string {
	return fmt.Sprintf(".capture/slide%d.pdf", slideNumber)
}

// TODO: mv this to another struct
func (c *CaptureCtl) countSlideLength() (int, error) {
	path := filepath.Join(c.config.Workdir, ".slide.yml")

	data, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}

	var dotslide slides.DotSlide
	if err := yaml.Unmarshal(data, &dotslide); err != nil {
		return 0, err
	}

	return len(dotslide.Slides), nil
}

func (c *CaptureCtl) mergeCaptures(slideLengrh int) error {
	files := []string{}
	for slideNumber := range slideLengrh {
		files = append(files, c.fmtCaptureOutput(slideNumber))
	}
	conf := pdfmodel.NewDefaultConfiguration()

	return pdfapi.MergeCreateFile(files, "capture.pdf", false, conf)
}
