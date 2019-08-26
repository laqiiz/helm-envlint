package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/bronze1man/yaml2json/y2jLib"
	"github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {

	prev, _ := filepath.Abs(".")

	dir := flag.String("d", ".", "helm directory")
	leftValues := flag.String("l", "", "helm values file path")
	rightValues := flag.String("r", "", "helm values file path")
	flag.StringVar(dir, "dir", ".", "helm directory")
	flag.StringVar(leftValues, "left", "", "helm values file path")
	flag.StringVar(rightValues, "right", "", "helm values file path")
	flag.Parse()

	if *leftValues == "" {
		log.Fatal("could not get left values path")
	}
	if *rightValues == "" {
		log.Fatal("could not get right values path")
	}

	leftYaml, err := exec.Command("helm", "template", prev+"/"+*dir, "-f", prev+"/"+*leftValues).CombinedOutput()
	if err != nil {
		log.Println(string(leftYaml))
		log.Fatal(err)
	}

	rightYaml, err := exec.Command("helm", "template", prev+"/"+*dir, "-f", prev+"/"+*rightValues).CombinedOutput()
	if err != nil {
		log.Println(string(leftYaml))
		log.Fatal(err)
	}

	tripLefts := strings.Split(string(leftYaml), "---")
	tripRights := strings.Split(string(rightYaml), "---")

	var result []map[string]interface{}

	for i := range tripLefts {

		leftJSON := tripLefts[i]
		rightJSON := tripRights[i]

		if leftJSON == "" {
			continue
		}

		var left bytes.Buffer
		if err := y2jLib.TranslateStream(bytes.NewReader([]byte(leftJSON)), &left); err != nil {
			log.Println("YAML parse error")
			log.Fatal(err)
		}

		var right bytes.Buffer
		if err := y2jLib.TranslateStream(bytes.NewReader([]byte(rightJSON)), &right); err != nil {
			log.Println("YAML parse error")
			log.Fatal(err)
		}

		differ := gojsondiff.New()
		jsonDiff, err := differ.Compare(left.Bytes(), right.Bytes())
		if err != nil {
			log.Println("Diff calc error")
			log.Fatal(err)
		}

		if len(jsonDiff.Deltas()) == 0 {
			continue
		}

		deltaDiff, err :=  formatter.NewDeltaFormatter().FormatAsJson(jsonDiff)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, deltaDiff)
	}

	out, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Println("JSON Marshal error")
		log.Fatal(err)
	}

	fmt.Println(string(out))
}
