// Code generated by scorify, DO NOT EDIT.
// Import generated from "github.com/scorify/dns@v0.0.2"

package checks

import (
	dns "github.com/scorify/dns"
)

func init() {
	Checks["dns"] = Check{
		Func:   dns.Run,
		Schema: ConvertSchema(dns.Schema{}),
	}
}