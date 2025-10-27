# Go Docling

Go bindings for [docling-serve](https://github.com/docling-project/docling-serve).

## Installation

You can install the Docling Go Library using `go get`:

```sh
go get github.com/iguanesolutions/go-docling@latest
```

## Usage

Basic example:

```go
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
			FromFormats: []docling.FromFormat{docling.FromPDF},
			ToFormats:   []docling.ToFormat{docling.ToMarkdown},
		},
	})
	if err != nil {
		log.Fatal("failed to process file", err)
	}
	fmt.Println(resp.Document.MarkdownContent())
}
```

## Endpoints implementation

Not all endpoints are implemented since we only needed to convert documents to markdown format.

### Health

- [x] Health

### Convert

- [x] Process Url
- [x] Process File
- [x] Process Url Async
- [x] Process File Async

### Chunk

- [ ] Chunk Sources With Hybridchunker As Async Task
- [ ] Chunk Files With Hybridchunker As Async Task
- [ ] Chunk Sources With Hybridchunker
- [ ] Chunk Files With Hybridchunker
- [ ] Chunk Sources With Hierarchicalchunker As Async Task
- [ ] Chunk Files With Hierarchicalchunker As Async Task
- [ ] Chunk Sources With Hierarchicalchunker
- [ ] Chunk Files With Hierarchicalchunker

### Tasks

- [x] Task Status Poll
- [x] Task Result

### Clear

- [ ] Clear Converters
- [ ] Clear Results

