package internal

import (
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func Merge() error {
	inputFiles := []string{"output1.pdf", "output2.pdf"}
	conf := model.NewDefaultConfiguration()

	return api.MergeCreateFile(inputFiles, "output.pdf", false, conf)
}
