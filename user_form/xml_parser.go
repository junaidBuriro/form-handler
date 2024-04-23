package userform

import (
	"encoding/xml"
	"io"
	"os"
)

// Form represents the structure of a form
// Structs to represent the XML schema
type xmlFormFields struct {
	Fields   []Field   `xml:"Field"`
	Sections []Section `xml:"Section"`
}

type Field struct {
	Name      string  `xml:"Name,attr"`
	Type      string  `xml:"Type,attr"`
	Optional  bool    `xml:"Optional,attr"`
	FieldType string  `xml:"FieldType,attr"`
	Caption   string  `xml:"Caption"`
	Labels    []Label `xml:"Labels>Label"`
}

type Label struct {
	Name  string `xml:"Name,attr"`
	Value string `xml:",chardata"`
}

type Section struct {
	Name        string    `xml:"Name,attr"`
	Optional    bool      `xml:"Optional,attr"`
	Title       string    `xml:"Title"`
	Contents    []Field   `xml:"Contents>Field"`
	SubSections []Section `xml:"Section"`
}

// type Content struct {
// 	Name      string `xml:"Name,attr"`
// 	Type      string `xml:"Type,attr"`
// 	Optional  bool   `xml:"Optional,attr"`
// 	FieldType string `xml:"FieldType,attr"`
// 	Caption   string `xml:"Caption"`
// }

// XMLFormParser implements FormParser interface for XML parsing
type XMLFormParser struct {
	// Add any dependencies or configurations needed
}

// Parser parses the XML file into a Form struct
func (p *XMLFormParser) Parser(xmlFilePath string, submittedValues map[string]string) (FormData, error) {
	xmlForm := &xmlFormFields{}
	xmlFile, err := os.Open(xmlFilePath)
	if err != nil {
		return nil, err
	}

	defer xmlFile.Close()

	byteValue, _ := io.ReadAll(xmlFile)
	err = xml.Unmarshal(byteValue, xmlForm)
	if err != nil {
		return nil, err
	}

	formData := make(FormData)

	// Extract field captions and values
	for _, field := range xmlForm.Fields {
		if value, ok := submittedValues[field.Name]; ok {
			formData[field.Caption] = value
		}
	}
	// since there can be many sections, going through each of them
	for _, section := range xmlForm.Sections {
		updateSectionFields(section, submittedValues, formData)
	}

	return formData, nil
}

// Function to fetch fields from a section and its sub-sections
func updateSectionFields(section Section, submittedValues map[string]string, formData FormData) {

	// Extract fields from the current section
	for _, field := range section.Contents {
		if value, ok := submittedValues[field.Name]; ok {
			formData[field.Caption] = value
		}
	}

	// Recursively extract fields from sub-sections
	for _, subSection := range section.SubSections {
		updateSectionFields(subSection, submittedValues, formData)
	}
}
