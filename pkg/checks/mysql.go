// Code generated by scorify, DO NOT EDIT.
// Import generated from "github.com/scorify/mysql@v0.0.0"

package checks

import (
	mysql "github.com/scorify/mysql"
)

func init() {
	Checks["mysql"] = Check{
		Func:   mysql.Run,
		Schema: ConvertSchema(mysql.Schema{}),
	}
}