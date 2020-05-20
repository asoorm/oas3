package oas3

import (
	"encoding/json"
	"testing"
)

func TestOas3(t *testing.T) {
	docs := Oas3{
		Info: Info{
			Title:          "Swagger Petstore",
			Version:        "1.0.0",
			Description:    "A sample API that uses a petstore as an example to demonstrate features of OAS3",
			TermsOfService: "http://swagger.io/terms/",
			Contact: struct {
				Name  string `json:"name"`
				Email string `json:"email"`
				URL   string `json:"url"`
			}{
				Name:  "Swagger API Team",
				Email: "apiteam@swagger.io",
				URL:   "http://swagger.io",
			},
			License: struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			}{
				Name: "Apache 2.0",
				URL:  "https://www.apache.org/licenses/LICENSE-2.0.html",
			},
		},
		Servers: []Server{
			{
				URL: "http://petstore.swagger.io/api",
			},
		},
	}

	data, _ := json.MarshalIndent(docs, "", "  ")
	t.Log(string(data))
}

oas3OperationId = "oas3.operationId"
oas3Description = "oas3.description"
oas3Parameter   = "oas3.parameter"
oas3Summary     = "oas3.summary"
oas3Response   = "oas3.response"
// oas3.operationId handlerTest
// oas3.description example description
// oas3.parameter
func handlerTest() {

}
