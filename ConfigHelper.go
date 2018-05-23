package ConfigHelper

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// GetConfigWithDefault sets configInstances to the config value from the specified json config file or sets to point to defaultConfig
func GetConfigWithDefault(filename string, defaultConfig interface{}, configInstance interface{}) error {

	configJSON, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Printf("Error reading config file (%s), creating default config file...\n", err)

		defaultConfigJSON, _ := json.MarshalIndent(defaultConfig, "", "  ")
		err = ioutil.WriteFile(filename, defaultConfigJSON, 0644)

		if err != nil {
			log.Fatalf("Error creating default config: %s\n", err)
		}

		configJSON = defaultConfigJSON
	}

	err = json.Unmarshal(configJSON, configInstance)
	if err != nil {
		log.Fatalf("Error unmarshaling config json: %s\n", err)
		return err
	}

	return nil
}
