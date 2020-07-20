package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"k8s.io/apimachinery/pkg/util/yaml"
)

func printUsage() {
	fmt.Fprintf(os.Stderr,
		`Usage: k8syaml2json < $inputYamlFile > $newlineSeparatedOutputJsonFile

k8syaml2json converts the yaml document or documents input via stdin into json,
and outputs them on stdout.
It follows kubernetes semantics for both yaml and json.

The output is simple newline separated json blobs, one per input yaml document.

Multiple yaml documents may be separated by '---'.
`)
}

func main() {
	if len(os.Args) != 1 {
		printUsage()
		os.Exit(1)
	}
	yamlToJSON := yaml.NewYAMLOrJSONDecoder(os.Stdin, 8096)
	for {
		var obj interface{}
		err := yamlToJSON.Decode(&obj)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Deserialization failed: %v", err)
		}
		jsonBytes, err := json.Marshal(obj)
		if err != nil {
			panic(fmt.Sprintf("unable to json marshal a k8s object: this should not happen: %v", err))
		}
		if _, err := os.Stdout.Write(jsonBytes); err != nil {
			log.Fatalf("Error writing to stdout: %v", err)
		}
	}
}
