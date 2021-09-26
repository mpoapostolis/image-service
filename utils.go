package main

import (
    "fmt"
    "os"
    "strings"

    "github.com/google/uuid"
    "github.com/h2non/bimg"
)

// A new folder is created at the root of the project.
func createFolder(dirname string) error {
    _, err := os.Stat(dirname)
    if os.IsNotExist(err) {
        errDir := os.MkdirAll(dirname, 0755)
        if errDir != nil {
            return errDir
        }
    }
    return nil
}

// The mime type of the image is changed, it is compressed and then saved in the specified folder.
func imageProcessing(buffer []byte, quality int, dirname string) (string, error) {
    filename := strings.Replace(uuid.New().String(), "-", "", -1) + ".webp"

    converted, err := bimg.NewImage(buffer).Convert(bimg.WEBP)
    if err != nil {
        return filename, err
    }

    processed, err := bimg.NewImage(converted).Process(bimg.Options{Quality: quality})
    if err != nil {
        return filename, err
    }

    writeError := bimg.Write(fmt.Sprintf("./"+dirname+"/%s", filename), processed)
    if writeError != nil {
        return filename, writeError
    }

    return filename, nil
}