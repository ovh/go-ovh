package common

import (
	"bytes"
	"encoding/json"
	"fmt"

	"os"

	"github.com/ghodss/yaml"
)

// FormatOutput autmatically formats json based output based on user choice.
// when selected formatter is "pretty", call prettyFormatter callback.
func FormatOutput(v interface{}, prettyFormatter func([]byte)) {
	data, e := json.Marshal(v)
	Check(e)

	switch Format {
	case "pretty":
		prettyFormatter(data)
	case "json":
		jsonFormatter(data)
	case "yaml":
		yamlFormatter(data)
	default:
		fmt.Fprintf(os.Stderr, "Invalid formater %s. Use one of 'pretty', 'json', 'yaml'\n", Format)
		return
	}
}

// FormatOutputDef autmatically formats json based output based on user choice.
// uses yamlFormatter as pretty formatter.
func FormatOutputDef(v interface{}) {
	FormatOutput(v, yamlFormatter)
}

// FormatOutputError prints the "message" field of an API return or falls back on FormatOutputDef if the field does not exist
func FormatOutputError(data []byte) {
	var errorDesc map[string]interface{}
	if err := json.Unmarshal(data, &errorDesc); err != nil {
		// sometimes, the API returns a string instead of a
		// JSON-object for the error. Let's fallback on that
		s := ""
		Check(json.Unmarshal(data, &s))
		errorDesc = map[string]interface{}{"message": s}
	}

	message := errorDesc["message"]
	if message == nil {
		message = errorDesc["error_details"]
	}

	if message != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", message)
	} else {
		FormatOutputDef(data)
	}
}

func jsonFormatter(data []byte) {
	var out bytes.Buffer
	json.Indent(&out, data, "", "  ")
	fmt.Println(out.String())
}

func yamlFormatter(data []byte) {
	out, err := yaml.JSONToYAML(data)
	Check(err)
	fmt.Print(string(out))
}
