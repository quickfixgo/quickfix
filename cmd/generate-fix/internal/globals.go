package internal

import (
	"fmt"
	"sort"

	"github.com/quickfixgo/quickfix/datadictionary"
)

type fieldTypeMap map[string]*datadictionary.FieldType

var (
	globalFieldTypesLookup fieldTypeMap
	GlobalFieldTypes       []*datadictionary.FieldType
)

// Sort fieldtypes by name.
type byFieldName []*datadictionary.FieldType

func (n byFieldName) Len() int           { return len(n) }
func (n byFieldName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n byFieldName) Less(i, j int) bool { return n[i].Name() < n[j].Name() }

func getGlobalFieldType(f *datadictionary.FieldDef) (t *datadictionary.FieldType, err error) {
	var ok bool
	t, ok = globalFieldTypesLookup[f.Name()]
	if !ok {
		err = fmt.Errorf("Unknown global type for %v", f.Name())
	}

	return
}

func BuildGlobalFieldTypes(specs []*datadictionary.DataDictionary) {
	globalFieldTypesLookup = make(fieldTypeMap)
	for _, spec := range specs {
		for _, field := range spec.FieldTypeByTag {
			if oldField, ok := globalFieldTypesLookup[field.Name()]; ok {
				// Merge old enums with new.
				if len(oldField.Enums) > 0 && field.Enums == nil {
					field.Enums = make(map[string]datadictionary.Enum)
				}

				for enumVal, enum := range oldField.Enums {
					if _, ok := field.Enums[enumVal]; !ok {
						// Verify an existing enum doesn't have the same description. Keep newer enum.
						okToKeepEnum := true
						for _, newEnum := range field.Enums {
							if newEnum.Description == enum.Description {
								okToKeepEnum = false
								break
							}
						}

						if okToKeepEnum {
							field.Enums[enumVal] = enum
						}
					}
				}
			}

			globalFieldTypesLookup[field.Name()] = field
		}
	}

	GlobalFieldTypes = make([]*datadictionary.FieldType, len(globalFieldTypesLookup))
	i := 0
	for _, fieldType := range globalFieldTypesLookup {
		GlobalFieldTypes[i] = fieldType
		i++
	}

	sort.Sort(byFieldName(GlobalFieldTypes))
}
