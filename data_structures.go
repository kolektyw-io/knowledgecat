package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Namespace struct {
	Mnemonic string
	Name     string
}

type Article struct {
	File     string    `json:"file"`
	Title    string    `json:"title"`
	Children []Article `json:"children"`
}

type NamespaceFile struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	Mnemonic    string `json:"mnemonic"`
	Entries     []struct {
		File     string    `json:"file"`
		Title    string    `json:"title"`
		Children []Article `json:"children"`
	} `json:"entries"`
}

func retrieveAvailableNamespaces() []Namespace {
	namespaces := make([]Namespace, 0)
	err := filepath.Walk(*datadir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {

			namespaceData := NamespaceFile{}
			f, e := os.ReadFile(filepath.Join(path, "namespace.json"))
			if e != nil {
				return nil
			}
			e = json.Unmarshal(f, &namespaceData)
			if e != nil {
				panic(e)
			}
			namespace := Namespace{Mnemonic: namespaceData.Mnemonic, Name: namespaceData.Name}
			namespace.Name = fmt.Sprintf("%s", namespaceData.Name)
			namespaces = append(namespaces, namespace)

		}
		return nil
	})

	if err != nil {
		panic(err)
	}
	return namespaces
}

func getArticles(namespace string) []Article {
	articles := make([]Article, 0)
	namespaceData := NamespaceFile{}
	err := filepath.Walk(filepath.Join(*datadir, namespace), func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			f, e := os.ReadFile(filepath.Join(path, "namespace.json"))
			if e != nil {
				return nil
			}
			e = json.Unmarshal(f, &namespaceData)
			if e != nil {
				panic(e)
			}

		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, elem := range namespaceData.Entries {
		articles = append(articles, Article{File: elem.File, Title: elem.Title})
	}
	return articles
}
