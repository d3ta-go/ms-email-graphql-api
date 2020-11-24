package schema

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

var schemaStr = `
# Schema
schema {
  query: Query
  mutation: Mutation
}

# Schema.Module
%s
# End of schema`

// LoadGraphQLSchema load GraphQL Schema
func LoadGraphQLSchema(schemaDir string) (string, error) {
	schemaModuleDir := schemaDir
	if schemaDir == "" {
		schemaModuleDir = "./interface/http-apps/graphql/echo/graphql/schema/schema.graphql"
	}
	files, err := ioutil.ReadDir(schemaModuleDir)
	if err != nil {
		return "", err
	}

	var schemaModuleStr string
	for _, f := range files {
		if filepath.Ext(f.Name()) == ".graphql" {
			fileSchema := fmt.Sprintf("%s/%s", schemaModuleDir, f.Name())

			bstr, err := ioutil.ReadFile(fileSchema)
			if err != nil {
				panic(err)
			}
			schemaModuleStr = schemaModuleStr + string(bstr)
			// fmt.Println(f.Name())
		}
	}

	schemaStr := fmt.Sprintf(schemaStr, schemaModuleStr)
	return schemaStr, nil
}
