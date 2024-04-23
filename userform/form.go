package userform

type FormData map[string]string

// FormParser defines the interface for parsing form schemas
type FormParser interface {
	Parser(filePath string, submittedData map[string]string) (FormData, error)
}

// PDFGenerator defines the interface for generating PDFs
type PDFGenerator interface {
	GeneratePDF(data FormData) error
}

// FormService contains methods for handling forms
type FormService struct {
	Parser    FormParser
	Generator PDFGenerator
}

// NewFormService initializes a new FormService
func NewFormService(parser FormParser, generator PDFGenerator) *FormService {
	return &FormService{
		Parser:    parser,
		Generator: generator,
	}
}

// ParseAndGeneratePDF parses the form schema and generates a PDF
func (fs *FormService) ParseAndGeneratePDF(filePath string, submittedData map[string]string) error {
	formData, err := fs.Parser.Parser(filePath, submittedData)
	if err != nil {
		return err
	}
	return fs.Generator.GeneratePDF(formData)
}
