package common

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type MetaData struct {
	ProjectName string
	Path        string
}

func GetMetaData() (MetaData, error) {
	file, err := os.Open(".just/metadata.json")
	if err != nil {
		return MetaData{}, fmt.Errorf("failed to open metadata file (%v)", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return MetaData{}, fmt.Errorf("failed to read from the metadata file (%v)", err)
	}

	var metadata MetaData
	if err := json.Unmarshal(bytes, &metadata); err != nil {
		return MetaData{}, fmt.Errorf("failed to unmarshal metadata (%v)", err)
	}

	return metadata, nil
}

func WriteMetaData(data MetaData) error {
	file, err := os.Create(".just/metadata.json")
	if err != nil {
		return fmt.Errorf("failed to create metadata file (%v)", err)
	}
	defer file.Close()

	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal metadata (%v)", err)
	}

	if _, err := file.Write(bytes); err != nil {
		return fmt.Errorf("failed to write to the metadata file (%v)", err)
	}

	return nil
}
