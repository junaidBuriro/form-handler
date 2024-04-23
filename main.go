package main

import (
	"fmt"

	"github.com/junaidBuriro/form_handler/userform"
)

func main() {
	// Initialize form service with XML parser
	xmlParser := &userform.XMLFormParser{}
	formService := userform.NewFormService(xmlParser, &userform.PDFGeneratorImpl{})

	// Example submitted data
	submittedData := map[string]string{
		"program_language": "A",
		"other":            "Some other programming experience",
		"code_repos":       "golang.zip",
		"first_job":        "Careem",
		// Add more submitted data as needed
	}

	// Parse the form schema file and generate PDF
	if err := formService.ParseAndGeneratePDF("form_schema.xml", submittedData); err != nil {
		fmt.Println("Error generating PDF:", err)
		return
	}

	fmt.Println("PDF generated successfully.")
}
