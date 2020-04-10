package main

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

var closureRegex = regexp.MustCompile("\\[(.+?)\\]")

func main() {
	readConfig()
	tidy(viper.GetString("origin_path"), viper.GetString("destination_path"))
}

func readConfig(){
	viper.SetConfigFile("./config.yaml")
	if err := viper.ReadInConfig();
		err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}

func tidy(originPath string, destinationPath string) {
	_ = filepath.Walk(originPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			isFound, tag := FindTag(info.Name())
			if isFound {
				newPathName := createPath(destinationPath, tag)
				_ = os.Rename(path, generateDestPath(newPathName, info))
				log.Println("tag:", tag, ", filename:", info.Name())
			}
		}
		return nil
	})
}

func FindTag(input string) (bool, string) {
	results := closureRegex.FindStringSubmatch(input)
	if len(results) <= 0 {
		return false, ""
	}
	return true, results[1]
}

func generateDestPath(destPath string, info os.FileInfo) string {
	return destPath + string(os.PathSeparator) + info.Name()
}

func createPath(path string, tag string) string {
	newPathName := path + string(os.PathSeparator) + tag
	_ = os.MkdirAll(newPathName, os.ModePerm)
	return newPathName
}
