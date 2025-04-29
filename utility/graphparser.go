package utility

import (
	"os"

	"regexp"

	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

type FieldsMap map[string]string

type AllEntitlement map[string]string

func ParseSchema() (FieldsMap, AllEntitlement, error) {
	schemaFilePath := "../schema.graphql"
	body, err := os.ReadFile(schemaFilePath)
	if err != nil {
		return nil, nil, err
	}
	doc, err := gqlparser.LoadSchema(&ast.Source{Input: string(body)})
	if err != nil {
		return nil, nil, err
	}
	mapping, err := ExtractEntitlementIdentifiers(string(body))
	if err != nil {
		return nil, nil, err
	}
	allFieldMap := make(FieldsMap)
	for typeName, def := range doc.Types {
		if validateString(typeName) && len(def.Fields) > 0 {
			fieldMap := make(map[string]string)
			for _, field := range def.Fields {
				if validateString(field.Name) {
					fieldMap[field.Name] = field.Type.String()
					allFieldMap[field.Name] = field.Type.String()
				}
			}
		}
	}
	return allFieldMap, mapping, nil
}

func ExtractEntitlementIdentifiers(sdl string) (AllEntitlement, error) {
	result := make(AllEntitlement)

	regex := regexp.MustCompile(`(?s)key:\s*"(.*?)".*?entitlementIdentifier:\s*"(.*?)"`)

	matches := regex.FindAllStringSubmatch(sdl, -1)
	for _, match := range matches {
		key := match[1]
		entitlementIdentifier := match[2]
		result[key] = entitlementIdentifier
	}

	return result, nil
}
