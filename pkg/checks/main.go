package checks

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/scorify/schema"
)

type Check struct {
	Func     func(ctx context.Context, config string) error
	Validate func(config string) error
	Schema   []*schema.Field
}

var Checks = make(map[string]Check)

func ConvertSchema(schema interface{}) string {
	schemaMap := make(map[string]string)
	schemaVal := reflect.ValueOf(schema)

	for i := 0; i < schemaVal.NumField(); i++ {
		field := schemaVal.Type().Field(i)

		switch field.Type.Kind() {
		case reflect.String:
			schemaMap[field.Tag.Get("json")] = "string"
		case reflect.Int:
			schemaMap[field.Tag.Get("json")] = "int"
		case reflect.Bool:
			schemaMap[field.Tag.Get("json")] = "bool"
		default:
			panic(fmt.Errorf("invalid type"))
		}
	}

	out, err := json.Marshal(schemaMap)
	if err != nil {
		panic(err)
	}

	return string(out)
}
