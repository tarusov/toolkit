package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/caarlos0/env/v8"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

// Load config to dst struct.
func Load(dst any, opts ...loadOption) error {

	var options = getDefaultLoadOptions()
	for _, set := range opts {
		set(options)
	}

	if options.envEnabled {
		if err := loadEnv(dst); err != nil {
			return fmt.Errorf("load cfg from env: %w", err)
		}
	}

	if options.jsonFile != "" {
		if err := loadJson(dst, options.jsonFile); err != nil {
			return fmt.Errorf("load cfg from json file: %w", err)
		}
	}

	if options.yamlFile != "" {
		if err := loadYaml(dst, options.yamlFile); err != nil {
			return fmt.Errorf("load cfg from yaml file: %w", err)
		}
	}

	if options.validateEnabled {
		if err := validate(dst); err != nil {
			return fmt.Errorf("validate cfg: %w", err)
		}
	}

	return nil
}

// loadEnv
func loadEnv(dst any) error {
	return env.Parse(dst)
}

// loadJson
func loadJson(dst any, fn string) error {
	data, err := os.ReadFile(fn)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, dst)
}

// loadYaml
func loadYaml(dst any, fn string) error {
	data, err := os.ReadFile(fn)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, dst)
}

// validate
func validate(dst any) error {
	return validator.New().Struct(dst)
}
