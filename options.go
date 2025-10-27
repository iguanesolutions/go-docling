package docling

type ConvertOptions struct {
	FromFormats                     []FromFormat             `json:"from_formats,omitempty"`                       // default: all formats
	ToFormats                       []ToFormat               `json:"to_formats,omitempty"`                         // default: ["md"]
	ImageExportMode                 ImageExportMode          `json:"image_export_mode,omitempty"`                  // default: "embedded"
	DoOCR                           *bool                    `json:"do_ocr,omitempty"`                             // default: true
	ForceOCR                        bool                     `json:"force_ocr,omitempty"`                          // default: false
	OCREngine                       OCREngine                `json:"ocr_engine,omitempty"`                         // default: "easyocr"
	OCRLang                         []string                 `json:"ocr_lang,omitempty"`                           // default: empty
	PDFBackend                      PDFBackend               `json:"pdf_backend,omitempty"`                        // default: "dlparse_v4"
	TableMode                       TableMode                `json:"table_mode,omitempty"`                         // default: "accurate"
	TableCellMatching               *bool                    `json:"table_cell_matching,omitempty"`                // default: true
	Pipeline                        Pipeline                 `json:"pipeline,omitempty"`                           // default: "standard"
	PageRange                       []int                    `json:"page_range,omitempty"`                         // default: [1,9223372036854776000]
	DocumentTimeout                 *int                     `json:"document_timeout,omitempty"`                   // default: 604800
	AbortOnError                    bool                     `json:"abort_on_error,omitempty"`                     // default: false
	DoTableStructure                *bool                    `json:"do_table_structure,omitempty"`                 // default: true
	IncludeImages                   *bool                    `json:"include_images,omitempty"`                     // default: true
	ImagesScale                     *float64                 `json:"images_scale,omitempty"`                       // default: 2.0
	MDPageBreakPlaceholder          string                   `json:"md_page_break_placeholder,omitempty"`          // default: ""
	DoCodeEnrichment                bool                     `json:"do_code_enrichment,omitempty"`                 // default: false
	DoFormulaEnrichment             bool                     `json:"do_formula_enrichment,omitempty"`              // default: false
	DoPictureClassification         bool                     `json:"do_picture_classification,omitempty"`          // default: false
	DoPictureDescription            bool                     `json:"do_picture_description,omitempty"`             // default: false
	PictureDescriptionAreaThreshold *float64                 `json:"picture_description_area_threshold,omitempty"` // default: 0.05
	PictureDescriptionLocal         *PictureDescriptionLocal `json:"picture_description_local,omitempty"`          // default: nil
	PictureDescriptionAPI           *PictureDescriptionAPI   `json:"picture_description_api,omitempty"`            // default: nil
	VLMPipelineModel                *VLMPipelineModel        `json:"vlm_pipeline_model,omitempty"`                 // default: nil
	VLMPipelineModelLocal           *VLMPipelineModelLocal   `json:"vlm_pipeline_model_local,omitempty"`           // default: nil
	VLMPipelineModelAPI             *VLMPipelineModelAPI     `json:"vlm_pipeline_model_api,omitempty"`             // default: nil
}

type ConvertOption func(*ConvertOptions)

func WithFromFormats(formats ...FromFormat) ConvertOption {
	return func(o *ConvertOptions) {
		o.FromFormats = append(o.FromFormats, formats...)
	}
}

func WithToFormats(formats ...ToFormat) ConvertOption {
	return func(o *ConvertOptions) {
		o.ToFormats = append(o.ToFormats, formats...)
	}
}

func WithImageExportMode(mode ImageExportMode) ConvertOption {
	return func(o *ConvertOptions) {
		o.ImageExportMode = mode
	}
}

func WithDoOCR(enable bool) ConvertOption {
	return func(o *ConvertOptions) {
		o.DoOCR = &enable
	}
}

func WithForceOCR(enable bool) ConvertOption {
	return func(o *ConvertOptions) {
		o.ForceOCR = enable
	}
}

func WithOCREngine(engine OCREngine) ConvertOption {
	return func(o *ConvertOptions) {
		o.OCREngine = engine
	}
}

func WithOCRLang(langs ...string) ConvertOption {
	return func(o *ConvertOptions) {
		o.OCRLang = append(o.OCRLang, langs...)
	}
}

func WithPDFBackend(backend PDFBackend) ConvertOption {
	return func(o *ConvertOptions) {
		o.PDFBackend = backend
	}
}

func WithTableMode(mode TableMode) ConvertOption {
	return func(o *ConvertOptions) {
		o.TableMode = mode
	}
}

func WithTableCellMatching(enable bool) ConvertOption {
	return func(o *ConvertOptions) {
		o.TableCellMatching = &enable
	}
}

func WithPipeline(pipeline Pipeline) ConvertOption {
	return func(o *ConvertOptions) {
		o.Pipeline = pipeline
	}
}

func WithPageRange(pages ...int) ConvertOption {
	return func(o *ConvertOptions) {
		o.PageRange = append(o.PageRange, pages...)
	}
}

func WithDocumentTimeout(timeout int) ConvertOption {
	return func(o *ConvertOptions) {
		o.DocumentTimeout = &timeout
	}
}

func WithAbortOnError(enable bool) ConvertOption {
	return func(o *ConvertOptions) {
		o.AbortOnError = enable
	}
}

func WithDoTableStructure(enable bool) ConvertOption {
	return func(o *ConvertOptions) {
		o.DoTableStructure = &enable
	}
}

func WithIncludeImages(enable bool) ConvertOption {
	return func(o *ConvertOptions) {
		o.IncludeImages = &enable
	}
}

func WithImagesScale(scale float64) ConvertOption {
	return func(o *ConvertOptions) {
		o.ImagesScale = &scale
	}
}

func WithMDPageBreakPlaceholder(placeholder string) ConvertOption {
	return func(o *ConvertOptions) {
		o.MDPageBreakPlaceholder = placeholder
	}
}

func WithDoCodeEnrichment(enable bool) ConvertOption {
	return func(o *ConvertOptions) {
		o.DoCodeEnrichment = enable
	}
}

func WithDoFormulaEnrichment(enable bool) ConvertOption {
	return func(o *ConvertOptions) {
		o.DoFormulaEnrichment = enable
	}
}

func WithDoPictureClassification(enable bool) ConvertOption {
	return func(o *ConvertOptions) {
		o.DoPictureClassification = enable
	}
}

func WithDoPictureDescription(enable bool) ConvertOption {
	return func(o *ConvertOptions) {
		o.DoPictureDescription = enable
	}
}

func WithPictureDescriptionAreaThreshold(threshold float64) ConvertOption {
	return func(o *ConvertOptions) {
		o.PictureDescriptionAreaThreshold = &threshold
	}
}

func WithPictureDescriptionLocal(local *PictureDescriptionLocal) ConvertOption {
	return func(o *ConvertOptions) {
		o.PictureDescriptionLocal = local
	}
}

func WithPictureDescriptionAPI(api *PictureDescriptionAPI) ConvertOption {
	return func(o *ConvertOptions) {
		o.PictureDescriptionAPI = api
	}
}

func WithVLMPipelineModel(model *VLMPipelineModel) ConvertOption {
	return func(o *ConvertOptions) {
		o.VLMPipelineModel = model
	}
}

func WithVLMPipelineModelLocal(model *VLMPipelineModelLocal) ConvertOption {
	return func(o *ConvertOptions) {
		o.VLMPipelineModelLocal = model
	}
}

func WithVLMPipelineModelAPI(model *VLMPipelineModelAPI) ConvertOption {
	return func(o *ConvertOptions) {
		o.VLMPipelineModelAPI = model
	}
}

type FromFormat string

const (
	FromDOCX        FromFormat = "docx"
	FromPPTX        FromFormat = "pptx"
	FromHTML        FromFormat = "html"
	FromImage       FromFormat = "image"
	FromPDF         FromFormat = "pdf"
	FromASCIIDoc    FromFormat = "asciidoc"
	FromMarkdown    FromFormat = "md"
	FromCSV         FromFormat = "csv"
	FromXLSX        FromFormat = "xlsx"
	FromXMLUspto    FromFormat = "xml_uspto"
	FromXMLJats     FromFormat = "xml_jats"
	FromMetsGbs     FromFormat = "mets_gbs"
	FromJSONDocling FromFormat = "json_docling"
	FromAudio       FromFormat = "audio"
)

type ToFormat string

const (
	ToMarkdown      ToFormat = "md"
	ToJSON          ToFormat = "json"
	ToHTML          ToFormat = "html"
	ToHTMLSplitPage ToFormat = "html_split_page"
	ToText          ToFormat = "text"
	ToDocTags       ToFormat = "doctags"
)

type ImageExportMode string

const (
	ImageExportModePlaceholder ImageExportMode = "placeholder"
	ImageExportModeEmbedded    ImageExportMode = "embedded"
	ImageExportModeReferenced  ImageExportMode = "referenced"
)

type OCREngine string

const (
	OCREngineEasyOCR   OCREngine = "easyocr"
	OCREngineOCRMac    OCREngine = "ocrmac"
	OCREngineRapidOCR  OCREngine = "rapidocr"
	OCREngineTesserOCR OCREngine = "tesserocr"
	OCREngineTesseract OCREngine = "tesseract"
)

type PDFBackend string

const (
	PDFBackendPyPDFium2 PDFBackend = "pypdfium2"
	PDFBackendDLParseV1 PDFBackend = "dlparse_v1"
	PDFBackendDLParseV2 PDFBackend = "dlparse_v2"
	PDFBackendDLParseV4 PDFBackend = "dlparse_v4"
)

type TableMode string

const (
	TableModeFast     TableMode = "fast"
	TableModeAccurate TableMode = "accurate"
)

type Pipeline string

const (
	PipelineStandard Pipeline = "standard"
	PipelineVLM      Pipeline = "vlm"
	PipelineASR      Pipeline = "asr"
)

type PictureDescriptionLocal struct {
	GenerationConfig GenerationConfig `json:"generation_config,omitzero"` // default: {"max_new_tokens":200,"do_sample":false}
	Prompt           string           `json:"prompt,omitempty"`           // default: "Describe this image in a few sentences."
	RepoID           string           `json:"repo_id"`                    // required
}

type GenerationConfig struct {
	DoSample     bool `json:"do_sample"`
	MaxNewTokens int  `json:"max_new_tokens"`
}

type PictureDescriptionAPI struct {
	Concurrency int               `json:"concurrency,omitempty"` // default: 1
	Headers     map[string]string `json:"headers,omitempty"`     // default: {}
	Params      map[string]any    `json:"params,omitempty"`      // default: {}
	Prompt      string            `json:"prompt,omitempty"`      // default: "Describe this image in a few sentences."
	Timeout     int               `json:"timeout,omitempty"`     // default: 20
	URL         string            `json:"url"`                   // required
}

type VLMPipelineModel string

const (
	VLMPipelineModelSmolDocling         VLMPipelineModel = "smoldocling"
	VLMPipelineModelSmolDoclingVLLM     VLMPipelineModel = "smoldocling_vllm"
	VLMPipelineModelGraniteVision       VLMPipelineModel = "granite_vision"
	VLMPipelineModelGraniteVisionVLLM   VLMPipelineModel = "granite_vision_vllm"
	VLMPipelineModelGraniteVisionOllama VLMPipelineModel = "granite_vision_ollama"
	VLMPipelineModelGotOCR2             VLMPipelineModel = "got_ocr_2"
)

type VLMPipelineModelLocal struct {
	ExtraGenerationConfig map[string]any        `json:"extra_generation_config,omitempty"` // default: {"max_new_tokens":800,"do_sample":false} see: https://huggingface.co/docs/transformers/en/main_classes/text_generation#transformers.GenerationConfig
	InferenceFramework    InferenceFramework    `json:"inference_framework"`               // required
	Prompt                string                `json:"prompt,omitempty"`                  // default: "Convert this page to docling."
	RepoID                string                `json:"repo_id"`                           // required
	ResponseFormat        ResponseFormat        `json:"response_format"`                   // required
	Scale                 float64               `json:"scale,omitempty"`                   // default: 2.0
	TransformersModelType TransformersModelType `json:"transformers_model_type,omitempty"` // default: "automodel"
}

type ResponseFormat string

const (
	ResponseFormatDocTags   ResponseFormat = "doctags"
	ResponseFormatMarkdown  ResponseFormat = "markdown"
	ResponseFormatHTML      ResponseFormat = "html"
	ResponseFormatOTSL      ResponseFormat = "otsl"
	ResponseFormatPlainText ResponseFormat = "plaintext"
)

type InferenceFramework string

const (
	InferenceFrameworkMLX          InferenceFramework = "mlx"
	InferenceFrameworkTransformers InferenceFramework = "transformers"
	InferenceFrameworkVLLM         InferenceFramework = "vllm"
)

type TransformersModelType string

const (
	TransformersModelTypeAutoModel                TransformersModelType = "automodel"
	TransformersModelTypeAutoModelVision2Seq      TransformersModelType = "automodel-vision2seq"
	TransformersModelTypeAutoModelCausalLM        TransformersModelType = "automodel-causallm"
	TransformersModelTypeAutoModelImageTextToText TransformersModelType = "automodel-imagetexttotext"
)

type VLMPipelineModelAPI struct {
	Concurrency    int               `json:"concurrency,omitempty"` // default: 1
	Headers        map[string]string `json:"headers,omitempty"`     // default: {}
	Params         map[string]any    `json:"params,omitempty"`      // default: {}
	Prompt         string            `json:"prompt,omitempty"`      // default: "Convert this page to docling."
	ResponseFormat ResponseFormat    `json:"response_format"`       // required
	Scale          float64           `json:"scale,omitempty"`       // default: 2.0
	Timeout        int               `json:"timeout,omitempty"`     // default: 60
	URL            string            `json:"url"`                   // required
}

func Ptr[T any](v T) *T {
	return &v
}
