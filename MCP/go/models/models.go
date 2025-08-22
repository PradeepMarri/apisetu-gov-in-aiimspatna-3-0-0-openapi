package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// AcademicCertificateSchema represents the AcademicCertificateSchema schema from the OpenAPI specification
type AcademicCertificateSchema struct {
	Name string `json:"name"`
	Issuedate string `json:"issueDate"`
	Validfromdate string `json:"validFromDate"`
	Issuedto map[string]interface{} `json:"IssuedTo"`
	Language string `json:"language"`
	Status string `json:"status"`
	TypeField string `json:"type"`
	Issuedby map[string]interface{} `json:"IssuedBy"`
	Issuedat string `json:"issuedAt"`
	Number int `json:"number"`
	Certificatedata map[string]interface{} `json:"CertificateData"`
}

// ConsentArtifactSchema represents the ConsentArtifactSchema schema from the OpenAPI specification
type ConsentArtifactSchema struct {
	Signature map[string]interface{} `json:"signature"`
	Consent map[string]interface{} `json:"consent"`
}
