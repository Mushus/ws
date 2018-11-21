package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	cd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get working directory: %v", err)
	}
	docs := filepath.Join(cd, "docs")
	if _, err = workAndGen(cd, docs); err != nil {
		log.Fatalf("failed to generate docs: %v", err)
	}
}

func workAndGen(cd string, docs string) (hasChild bool, err error) {
	files, err := ioutil.ReadDir(cd)
	if err != nil {
		return false, fmt.Errorf("cannot get file list: %v", err)
	}
	hasGoFile := false
	for _, file := range files {
		name := file.Name()
		if strings.HasPrefix(name, ".") {
			continue
		}
		if file.IsDir() {
			nextCd := filepath.Join(cd, name)
			if nextCd == docs {
				continue
			}
			nextDocs := filepath.Join(docs, name)
			hc, err := workAndGen(nextCd, nextDocs)
			hasChild = hc
			if err != nil {
				return hasChild, err
			}
		} else {
			ext := filepath.Ext(name)
			if ext == ".go" {
				hasGoFile = true
				hasChild = true
			}
		}
	}
	if !(hasGoFile || hasChild) {
		return hasChild, nil
	}
	if err := genDocs(cd, docs, hasChild); err != nil {
		return false, err
	}
	return hasChild, nil
}

func genDocs(cd string, docs string, index bool) error {
	if err := os.MkdirAll(docs, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}
	args := []string{"-html", cd}
	if index {
		args = []string{"-html", "-index", cd}
	}
	out, err := exec.Command("godoc", args...).Output()
	if err != nil {
		return fmt.Errorf("failed exec godoc command: %v", err)
	}
	fp := filepath.Join(docs, "index.html")
	file, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("cannot open html docs: %v", err)
	}
	if _, err := file.Write(out); err != nil {
		return fmt.Errorf("cannot open html docs: %v", err)
	}
	return nil
}
