package oas3

type Oas3 struct {
	Openapi string   `json:"openapi"`
	Info    Info     `json:"info"`
	Servers []Server `json:"servers"`
	Paths   Paths    `json:"paths"`
}

type Info struct {
	Title          string `json:"title"`
	Version        string `json:"version"`
	Description    string `json:"description"`
	TermsOfService string `json:"termsOfService"`
	Contact        struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		URL   string `json:"url"`
	} `json:"contact"`
	License struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"license"`
}

type Server struct {
	URL         string                    `json:"url"`
	Description string                    `json:"description,omitempty"`
	Variables   map[string]ServerVariable `json:"variables,omitempty"`
}

type ServerVariable struct {
	// TODO
}

type Paths map[string]Path

type Path struct {
	Path       string       `json:"-"`
	Get        *Operation   `json:"get,omitempty"`
	Put        *Operation   `json:"put,omitempty"`
	Post       *Operation   `json:"post,omitempty"`
	Delete     *Operation   `json:"delete,omitempty"`
	Options    *Operation   `json:"delete,omitempty"`
	Head       *Operation   `json:"delete,omitempty"`
	Patch      *Operation   `json:"patch,omitempty"`
	Trace      *Operation   `json:"delete,omitempty"`
	Servers    []*Server    `json:"servers,omitempty"`
	// TODO: Parameters can be Reference Object - needs investigation
	Parameters []*Parameter `json:"parameters,omitempty"`
	// TODO: Handle Specification Extensions? https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.0.md#specificationExtensions
}

type Operation struct {
	OperationId string              `json:"operationId,omitempty"`
	Description string              `json:"description,omitempty"`
	Parameters  []*Parameter        `json:"parameters,omitempty"`
	Summary     string              `json:"summary,omitempty"`
	Responses   map[string]Response `json:"responses,omitempty"`
}

type Parameter struct {
	Name        string `json:"name"`
	In          string `json:"in"`
	Description string `json:"description,omitempty"`
	Required    bool   `json:"required"`
	Style       string `json:"style,omitempty"`
	Schema      struct {
		// TODO
	} `json:"schema"`
}

var parameterIn = []string{
	"path",
	"query",
	"header",
	"cookie",
}

type Response struct {
	Description string                     `json:"description"`
	Content     map[string]ResponseContent `json:"content"`
}

type ResponseContent struct {
	Schema struct {
		Ref string `json:"$ref"`
	} `json:"schema"`
}

type Components struct {
	Schemas map[string]Schema `json:"schemas"`
}

type Schema struct {
	Type       string   `json:"type"`
	Required   []string `json:"required"`
	Properties struct {
		Name Property `json:"name"`
		Tag  Property `json:"tag"`
		Code Property `json:"code"`
	} `json:"properties"`
}

type Property struct {
	Type   string `json:"type"`
	Format string `json:"format"`
}
