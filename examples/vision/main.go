package main

import (
	"context"
	"fmt"
	"log"

	"github.com/iguanesolutions/go-docling"
)

func main() {
	docli, err := docling.NewClient(docling.ClientConfig{
		BaseURL: "http://127.0.0.1:5001",
	})
	if err != nil {
		log.Fatal("failed to init client", err)
	}
	f, err := docling.FileReaderFromFile("2501.17887v1.pdf")
	if err != nil {
		log.Fatal("failed to open file", err)
	}
	resp, err := docli.ProcessFile(context.Background(), docling.ProcessFileRequest{
		Files:      []docling.File{f},
		TargetType: docling.TargetTypeInBody,
		ConvertOptions: docling.ConvertOptions{
			FromFormats:                     []docling.FromFormat{docling.FromPDF},
			ToFormats:                       []docling.ToFormat{docling.ToMarkdown},
			ImageExportMode:                 docling.ImageExportModePlaceholder,
			IncludeImages:                   docling.Ptr(false),
			DoPictureDescription:            true,
			PictureDescriptionAreaThreshold: docling.Ptr(0.0),
			// docling will make an http call to this endpoint to describe each image in the document
			PictureDescriptionAPI: &docling.PictureDescriptionAPI{
				URL: "https://example/v1/chat/completions",
				Headers: map[string]string{
					"Authorization": "Bearer 1234",
				},
				Params: map[string]any{
					"model": "Vision Model",
				},
				Prompt: "Describe this image in a few sentences.",
			},
		},
	})
	if err != nil {
		log.Fatal("failed to process file", err)
	}
	fmt.Println(resp.Document.MarkdownContent())
}
