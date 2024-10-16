// Code generated by scorify, DO NOT EDIT.
// Import generated from "github.com/scorify/smb@v0.0.1"

package checks

import (
	"github.com/scorify/schema"

	smb "github.com/scorify/smb"
)

func init() {
	schema, err := schema.Describe(smb.Schema{})
	if err != nil {
		panic(err)
	}

	Checks["smb"] = Check{
		Func:     smb.Run,
		Validate: smb.Validate,
		Schema:   schema,
	}
}