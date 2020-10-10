package mpesa

import (
	"bytes"
	"fmt"
	errors "github.com/paymentsds/mpesa-go-sdk/pkg/mpesa/errors"
	"net/url"
	"reflect"
	"regexp"
)

type Operation struct {
	port              int16
	path              string
	method            HttpMethod
	required          []string
	optional          []string
	mapping           map[string]string
	rules             map[string]string
	missingProperties []string
}

func (o *Operation) detectMissingValues(intent Intent) ([]string, error) {
	e := reflect.ValueOf(o).Elem()
	missingProperties := make([]string)

	for i := 0; i < e.NumField(); i++ {
		for _, field := range o.required {
			if name = e.Type().Field(i).Name; field == name {
				if kind = e.Type().Field(i).Kind; kind == reflect.String() {
					if value := e.Type().Field(i).Interface().(string); value == "" {
						append(missingProperties, field)
					}
				}

				if kind == reflect.Float64() {
					if value := e.Type().Field(i).Interface().(float64); value == 0.0 {
						append(missingProperties, field)
					}
				}
			}
		}
	}

	if len(missingProperties) > 0 {
		return missingProperties, NewMissingPropertiesError(missingProperties)
	}

	return missingProperties, nil
}

func (o *Operation) validateValues(intent Intent) ([]string, error) {
	e := reflect.ValueOf(o).Elem()
	badlyFormatedProperties := make([]string)

	for i := 0; i < e.NumField(); i++ {
		for _, field := range o.required {
			if name = e.Type().Field(i).Name; field == name {
				kind = e.Type().Field(i).Kind

				if kind == reflect.String() {
					r = regexp.MustCompile(o.rules[field])

					if value := e.Type().Field(i).Interface().(string); !r.MatchString(value) {
						append(badlyFormatedProperties, field)
					}
				}

				if kind == reflect.Float64() {
					if value := e.Type().Field(i).Interface().(float64); value < 0 {
						append(badlyFormatedProperties, field)
					}
				}
			}
		}
	}

	if len(badlyFormatedProperties) > 0 {
		return badlyFormatedProperties, NewValidationError(badlyFormatedProperties)
	}

	return badlyFormatedProperties, nil
}

func (o *Operation) URL(c *Configuration) *url.URL {
	// TODO
}

func (o *Operation) Method() string {
	// TODO
}
