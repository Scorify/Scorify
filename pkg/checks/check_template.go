// Code generated by scorify, DO NOT EDIT.
// Import generated from "github.com/scorify/check-template@v1.0.3"

package checks

import (
	"github.com/scorify/schema"

	check_template "github.com/scorify/check-template"
)

func init() {
	schema, err := schema.Describe(check_template.Schema{})
	if err != nil {
		panic(err)
	}

	Checks["check_template"] = Check{
		Func:     check_template.Run,
		Validate: check_template.Validate,
		Schema:   schema,
	}
}