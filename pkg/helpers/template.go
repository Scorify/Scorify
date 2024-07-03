package helpers

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Template struct {
	Number int
	Name   string
}

func format(format string, value interface{}) string {
	return fmt.Sprintf(format, value)
}

var funcMap = template.FuncMap{
	"format": format,
}

func ConfigTemplate(config string, data Template) string {
	if strings.Contains(config, "{{") {
		tmpl, err := template.New(uuid.New().String()).
			Funcs(funcMap).
			Parse(config)
		if err != nil {
			logrus.WithError(err).Error("failed to parse template")
			return config
		}

		buf := bytes.NewBuffer([]byte{})
		err = tmpl.Execute(buf, data)
		if err != nil {
			logrus.WithError(err).Error("failed to execute template")
			return config
		}

		return buf.String()
	} else {
		return config
	}
}
