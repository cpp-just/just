package commands

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Install(basePath string, packages []string) error {
	if len(packages) == 0 {
		return errors.New("no packages were specified")
	}

	for _, pkg := range packages {
		gitPath := fmt.Sprintf("%s%s", basePath, pkg)
		packagePath := fmt.Sprintf("packages/%s", pkg)

		_, err := os.Stat(packagePath)
		if err == nil {
			fmt.Printf("%s is already installed. Update it with just update %s. Skipping.\n", pkg, pkg)
			continue
		}

		const tempZipFilename = ".justtemp.zip"

		err = downloadReleaseZip(gitPath, tempZipFilename)
		if err != nil {
			return fmt.Errorf("failed to download the package (%w)", err)
		}

		// don't care about the error
		defer os.Remove(tempZipFilename)

		err = unzipPackage(packagePath, tempZipFilename)
		if err != nil {
			return fmt.Errorf("failed to unzip the package (%w)", err)
		}

		justRepoFilePath := fmt.Sprintf("%s/.justrepo", packagePath)
		justRepoFile, err := os.Create(justRepoFilePath)
		if err != nil {
			os.RemoveAll(packagePath)
			return errors.New("failed to create the .justrepo file")
		}
		defer justRepoFile.Close()

		_, err = justRepoFile.WriteString(gitPath)
		if err != nil {
			os.RemoveAll(packagePath)
			return errors.New("failed to write to the .justrepo file")
		}
	}

	return nil
}

func downloadReleaseZip(gitPath string, tempZipFilename string) error {
	releasePath := gitPath + "/releases/latest/download/just.zip"

	resp, err := http.Get(releasePath)
	if err != nil {
		return fmt.Errorf("failed to download just.zip from %s because of an http error", releasePath)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("failed to download just.zip from %s. Status code: %d", releasePath, resp.StatusCode)
	}

	zipFile, err := os.Create(tempZipFilename)
	if err != nil {
		return fmt.Errorf("failed to create the temporary %s file", tempZipFilename)
	}
	defer zipFile.Close()

	_, err = io.Copy(zipFile, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write the content to the zip file")
	}

	return nil
}

func unzipPackage(packagePath string, tempZipFilename string) error {
	archive, err := zip.OpenReader(tempZipFilename)
	if err != nil {
		return fmt.Errorf("failed to read %s", tempZipFilename)
	}
	defer archive.Close()

	err = os.MkdirAll(packagePath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create the package path %s", packagePath)
	}

	for _, file := range archive.File {
		readCloser, err := file.Open()
		if err != nil {
			return fmt.Errorf("failed to unzip file %s", file.Name)
		}

		fullPath := filepath.Join(packagePath, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(fullPath, file.Mode())
			continue
		}

		parentDirectory := ""
		if lastSeperator := strings.LastIndex(fullPath, string(os.PathSeparator)); lastSeperator > -1 {
			parentDirectory = fullPath[:lastSeperator]
		}

		err = os.MkdirAll(parentDirectory, file.Mode())
		if err != nil {
			return fmt.Errorf("failed to create parent directory %s for file %s", parentDirectory, file.Name)
		}

		outFile, err := os.OpenFile(fullPath, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, file.Mode())
		if err != nil {
			return fmt.Errorf("failed to create out file %s", fullPath)
		}
		defer outFile.Close()

		_, err = io.Copy(outFile, readCloser)
		if err != nil {
			return fmt.Errorf("failed to copy contents for out file %s", fullPath)
		}
	}

	return nil
}