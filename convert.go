package docling

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func (c *Client) ProcessFileWithOptions(ctx context.Context, files []File, targetType TargetType, opts ...ConvertOption) (ConvertResponse, error) {
	var options ConvertOptions
	for _, opt := range opts {
		opt(&options)
	}
	return c.ProcessFile(ctx, ProcessFileRequest{
		Files:          files,
		TargetType:     targetType,
		ConvertOptions: options,
	})
}

func writeFormFile(w *multipart.Writer, f File) error {
	ff, err := w.CreateFormFile("files", filepath.Base(f.Name()))
	if err != nil {
		return err
	}
	_, err = io.Copy(ff, f)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) ProcessFile(ctx context.Context, req ProcessFileRequest) (ConvertResponse, error) {
	body, contentType := c.processFileBody(req)
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.apiURL("convert/file"), body)
	if err != nil {
		return ConvertResponse{}, err
	}
	r.Header.Set("Content-Type", contentType)
	var resp ConvertResponse
	err = c.Do(r, &resp)
	if err != nil {
		return ConvertResponse{}, err
	}
	return resp, nil
}

func (c *Client) ProcessFileAsyncWithOptions(ctx context.Context, files []File, targetType TargetType, opts ...ConvertOption) (AsyncResponse, error) {
	var options ConvertOptions
	for _, opt := range opts {
		opt(&options)
	}
	return c.ProcessFileAsync(ctx, ProcessFileRequest{
		Files:          files,
		TargetType:     targetType,
		ConvertOptions: options,
	})
}

func (c *Client) ProcessFileAsync(ctx context.Context, req ProcessFileRequest) (AsyncResponse, error) {
	body, contentType := c.processFileBody(req)
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.apiURL("convert/file/async"), body)
	if err != nil {
		return AsyncResponse{}, err
	}
	r.Header.Set("Content-Type", contentType)
	if len(c.apiKey) > 0 {
		r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	}
	var resp AsyncResponse
	err = c.Do(r, &resp)
	if err != nil {
		return AsyncResponse{}, err
	}
	return resp, nil
}

type ProcessFileRequest struct {
	Files      []File
	TargetType TargetType
	ConvertOptions
}

func (c *Client) ProcessURLWithOptions(ctx context.Context, srcs []Source, target Target, opts ...ConvertOption) (ConvertResponse, error) {
	var options ConvertOptions
	for _, opt := range opts {
		opt(&options)
	}
	return c.ProcessURL(ctx, ProcessURLRequest{
		Options: options,
		Sources: srcs,
		Target:  target,
	})
}

func (c *Client) ProcessURL(ctx context.Context, req ProcessURLRequest) (ConvertResponse, error) {
	r, err := c.NewRequest(ctx, http.MethodPost, "convert/source", req)
	if err != nil {
		return ConvertResponse{}, err
	}
	var resp ConvertResponse
	err = c.Do(r, &resp)
	if err != nil {
		return ConvertResponse{}, err
	}
	return resp, nil
}

func (c *Client) ProcessURLAsyncWithOptions(ctx context.Context, srcs []Source, target Target, opts ...ConvertOption) (AsyncResponse, error) {
	var options ConvertOptions
	for _, opt := range opts {
		opt(&options)
	}
	return c.ProcessURLAsync(ctx, ProcessURLRequest{
		Options: options,
		Sources: srcs,
		Target:  target,
	})
}

func (c *Client) ProcessURLAsync(ctx context.Context, req ProcessURLRequest) (AsyncResponse, error) {
	r, err := c.NewRequest(ctx, http.MethodPost, "convert/source/async", req)
	if err != nil {
		return AsyncResponse{}, err
	}
	var resp AsyncResponse
	err = c.Do(r, &resp)
	if err != nil {
		return AsyncResponse{}, err
	}
	return resp, nil
}

func (c *Client) processFileBody(req ProcessFileRequest) (io.Reader, string) {
	pr, pw := io.Pipe()
	w := multipart.NewWriter(pw)
	go func() {
		var err error
		defer func() {
			pw.CloseWithError(err)
		}()
		err = multipartEncode(w, req.ConvertOptions)
		if err != nil {
			return
		}
		for _, f := range req.Files {
			err = writeFormFile(w, f)
			if err != nil {
				return
			}
		}
		if req.TargetType != "" {
			ff, err := w.CreateFormField("target_type")
			if err != nil {
				return
			}
			_, err = fmt.Fprint(ff, req.TargetType)
			if err != nil {
				return
			}
		}
		err = w.Close()
		if err != nil {
			return
		}
	}()
	return pr, w.FormDataContentType()
}

type ProcessURLRequest struct {
	Options ConvertOptions `json:"options"`
	Sources []Source       `json:"sources"`
	Target  Target         `json:"target"`
}

type ConvertResponse struct {
	Document Document `json:"document"`
	Status   string   `json:"status"` // "pending" "started" "failure" "success" "partial_success" "skipped"
	Errors   []struct {
		ComponentType string `json:"component_type"` // "document_backend" "model" "doc_assembler" "user_input"
		ModuleName    string `json:"module_name"`
		ErrorMessage  string `json:"error_message"`
	} `json:"errors"`
	ProcessingTime float64 `json:"processing_time"`
	Timings        map[string]struct {
		Scope           string    `json:"scope"` // "page" "document"
		Count           int       `json:"count"`
		Times           []float64 `json:"times"`
		StartTimestamps []string  `json:"start_timestamps"`
	} `json:"timings"`
}

type AsyncResponse struct {
	TaskID       string `json:"task_id"`
	TaskType     string `json:"task_type"` // "convert" "chunk"
	TaskStatus   string `json:"task_status"`
	TaskPosition int    `json:"task_position"`
	TaskMeta     struct {
		NumDocs      int `json:"num_docs"`
		NumProcessed int `json:"num_processed"`
		NumSucceeded int `json:"num_succeeded"`
		NumFailed    int `json:"num_failed"`
	} `json:"task_meta"`
}

type Source interface {
	Kind() SourceKind
}

type SourceKind string

const (
	SourceKindFile SourceKind = "file"
	SourceKindHTTP SourceKind = "http"
	SourceKindS3   SourceKind = "s3"
)

type SourceFile struct {
	Base64String string `json:"base64_string"` // required
	Filename     string `json:"filename"`      // required
}

func (s SourceFile) Kind() SourceKind {
	return SourceKindFile
}

func (s SourceFile) MarshalJSON() ([]byte, error) {
	type Alias SourceFile
	return json.Marshal(struct {
		Kind SourceKind `json:"kind"`
		Alias
	}{
		Kind:  SourceKindFile,
		Alias: Alias(s),
	})
}

type SourceHTTP struct {
	URL     string            `json:"url"`               // required
	Headers map[string]string `json:"headers,omitempty"` // default: {}
}

func (s SourceHTTP) Kind() SourceKind {
	return SourceKindHTTP
}

func (s SourceHTTP) MarshalJSON() ([]byte, error) {
	type Alias SourceHTTP
	return json.Marshal(struct {
		Kind SourceKind `json:"kind"`
		Alias
	}{
		Kind:  SourceKindHTTP,
		Alias: Alias(s),
	})
}

type SourceS3 struct {
	Endpoint  string `json:"endpoint"`             // required
	VerifySSL *bool  `json:"verify_ssl,omitempty"` // default: true
	AccessKey string `json:"access_key"`           // required
	SecretKey string `json:"secret_key"`           // required
	Bucket    string `json:"bucket"`               // required
	KeyPrefix string `json:"key_prefix,omitempty"` // default: ""
}

func (s SourceS3) Kind() SourceKind {
	return SourceKindS3
}

func (s SourceS3) MarshalJSON() ([]byte, error) {
	type Alias SourceS3
	return json.Marshal(struct {
		Kind SourceKind `json:"kind"`
		Alias
	}{
		Kind:  SourceKindS3,
		Alias: Alias(s),
	})
}

type Target interface {
	Kind() TargetKind
}

type TargetInBody struct{}

func (t TargetInBody) Kind() TargetKind {
	return TargetKindInBody
}

func (t TargetInBody) MarshalJSON() ([]byte, error) {
	type Alias TargetInBody
	return json.Marshal(struct {
		Kind TargetKind `json:"kind"`
		Alias
	}{
		Kind:  TargetKindInBody,
		Alias: Alias(t),
	})
}

type TargetPut struct {
	URL string `json:"url"` // required
}

func (t TargetPut) Kind() TargetKind {
	return TargetKindPut
}

func (t TargetPut) MarshalJSON() ([]byte, error) {
	type Alias TargetPut
	return json.Marshal(struct {
		Kind TargetKind `json:"kind"`
		Alias
	}{
		Kind:  TargetKindPut,
		Alias: Alias(t),
	})
}

type TargetS3 struct {
	Endpoint  string `json:"endpoint"`             // required
	VerifySSL *bool  `json:"verify_ssl,omitempty"` // default: true
	AccessKey string `json:"access_key"`           // required
	SecretKey string `json:"secret_key"`           // required
	Bucket    string `json:"bucket"`               // required
	KeyPrefix string `json:"key_prefix,omitempty"` // default: ""
}

func (t TargetS3) Kind() TargetKind {
	return TargetKindS3
}

func (t TargetS3) MarshalJSON() ([]byte, error) {
	type Alias TargetS3
	return json.Marshal(struct {
		Kind TargetKind `json:"kind"`
		Alias
	}{
		Kind:  TargetKindS3,
		Alias: Alias(t),
	})
}

type TargetZip struct{}

func (t TargetZip) Kind() TargetKind {
	return TargetKindZip
}

func (t TargetZip) MarshalJSON() ([]byte, error) {
	type Alias TargetZip
	return json.Marshal(struct {
		Kind TargetKind `json:"kind"`
		Alias
	}{
		Kind:  TargetKindZip,
		Alias: Alias(t),
	})
}

type TargetKind string

const (
	TargetKindInBody TargetKind = "inbody"
	TargetKindPut    TargetKind = "put"
	TargetKindS3     TargetKind = "s3"
	TargetKindZip    TargetKind = "zip"
)

type TargetType string

const (
	TargetTypeInBody TargetType = "inbody"
	TargetTypeZip    TargetType = "zip"
)

type Document struct {
	Filename string
	Contents []Content
}

func (d Document) MarkdownContent() string {
	for _, content := range d.Contents {
		if content.Format() == ToMarkdown {
			return content.String()
		}
	}
	return ""
}

func (d Document) HTLMContent() string {
	for _, content := range d.Contents {
		if content.Format() == ToHTML {
			return content.String()
		}
	}
	return ""
}

func (d Document) JSONContent() string {
	for _, content := range d.Contents {
		if content.Format() == ToJSON {
			return content.String()
		}
	}
	return ""
}

func (d Document) DocTagsContent() string {
	for _, content := range d.Contents {
		if content.Format() == ToDocTags {
			return content.String()
		}
	}
	return ""
}

func (d *Document) UnmarshalJSON(data []byte) error {
	aux := &struct {
		Filename       string           `json:"filename"`
		MDContent      string           `json:"md_content"`
		JSONContent    *json.RawMessage `json:"json_content"` // we need a pointer to json.RawMessage since the json.RawMessage is not nil if the parameter is set to null
		HTMLContent    string           `json:"html_content"`
		DocTagsContent string           `json:"doctags_content"`
	}{}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	d.Filename = aux.Filename
	if aux.MDContent != "" {
		d.Contents = append(d.Contents, MarkdownContent(aux.MDContent))
	}
	if aux.JSONContent != nil {
		d.Contents = append(d.Contents, JSONContent(*aux.JSONContent))
	}
	if aux.HTMLContent != "" {
		d.Contents = append(d.Contents, HTMLContent(aux.HTMLContent))
	}
	if aux.DocTagsContent != "" {
		d.Contents = append(d.Contents, DocTagsContent(aux.DocTagsContent))
	}
	return nil
}

type Content interface {
	Format() ToFormat
	fmt.Stringer
}

type MarkdownContent string

func (c MarkdownContent) Format() ToFormat {
	return ToMarkdown
}

func (c MarkdownContent) String() string {
	return string(c)
}

type JSONContent json.RawMessage

func (c JSONContent) Format() ToFormat {
	return ToJSON
}

func (c JSONContent) String() string {
	return string(c)
}

type HTMLContent string

func (c HTMLContent) Format() ToFormat {
	return ToHTML
}

func (c HTMLContent) String() string {
	return string(c)
}

type DocTagsContent string

func (c DocTagsContent) Format() ToFormat {
	return ToDocTags
}

func (c DocTagsContent) String() string {
	return string(c)
}

type File interface {
	Name() string
	io.Reader
}

type FileReader struct {
	Filename string
	io.Reader
}

func (fr FileReader) Name() string {
	return fr.Filename
}

func FileReaderFromFile(filename string) (FileReader, error) {
	f, err := os.Open(filename)
	if err != nil {
		return FileReader{}, err
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		return FileReader{}, err
	}
	return FileReader{
		Filename: filename,
		Reader:   bytes.NewReader(data),
	}, nil
}
