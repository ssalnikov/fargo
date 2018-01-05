package util

import (
	"strings"

	"github.com/gigovich/fargo/orm/field"
)

var reservedFieldNames = map[string]string{"id": "ID", "pk": "PK", "url": "URL", "uri": "URI"}

func FormatFieldName(name string) (modified string) {
	for _, part := range strings.Split(name, "_") {
		if substitute, ok := reservedFieldNames[strings.ToLower(part)]; ok {
			modified += substitute
		} else {
			modified += strings.Title(part)
		}
	}
	return modified
}

func FormatRecordName(name string) (modifier string) {
	return strings.TrimSuffix(name, "Model") + "Record"
}

func GetTags(f field.Base) string {
	if f.Tags == "" {
		return "`json:\"" + f.Name + "\"`"
	}
	return f.Tags
}
