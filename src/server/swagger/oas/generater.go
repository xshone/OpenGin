package oas

import (
	"fmt"
	"time"
)

func GenerateOpenApiSpec(outputPath string, debug, withoutPkgName bool) error {
	startTime := time.Now()
	p, err := newParser(".", "", "", debug, withoutPkgName)
	if err != nil {
		return err
	}
	createErr := p.CreateOASFile(outputPath)
	if createErr == nil {
		processTime := time.Since(startTime)
		fmt.Println("[OAS3] Doc file 'openapi.json' has been created successfully. Time elapsed:", processTime.Seconds(), "s")
	} else {
		fmt.Println("[OAS3 Error]:", createErr)
	}

	return createErr
}
