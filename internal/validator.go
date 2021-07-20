package internal

import (
	"github.com/xeipuuv/gojsonschema"
	"net/url"
	"path/filepath"
)

func Validate(schema string, document string) (*gojsonschema.Result, error) {
	if res, err := AbsolutizeUrl(schema); err != nil {
		return nil, err
	} else {
		schema = res
	}
	if res, err := AbsolutizeUrl(document); err != nil {
		return nil, err
	} else {
		document = res
	}
	schemaRefLoader := gojsonschema.NewReferenceLoader(schema)
	documentRefLoader := gojsonschema.NewReferenceLoader(document)
	return gojsonschema.Validate(schemaRefLoader, documentRefLoader)
}

func AbsolutizeUrl(input string) (string, error) {
	u, err := url.Parse(input)
	if err != nil {
		return "", err
	}
	if u.Path != "" {
		if p, err := filepath.Abs(u.Path); err != nil {
			return "", err
		} else {
			u.Path = p
		}
	}
	if u.Scheme == "" && u.Host == "" {
		u.Scheme = "file"
	} else if u.Scheme == "" {
		u.Scheme = "https"
	}
	return u.String(), nil
}
