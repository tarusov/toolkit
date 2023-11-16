package config

type (
	// options is list of settings for logger ctor.
	loadOptions struct {
		envEnabled      bool
		validateEnabled bool
		jsonFile        string
		yamlFile        string
	}

	// option if type of modifing func for ctor.
	loadOption func(*loadOptions)
)

// getDefaultLoadOptions return default options.
func getDefaultLoadOptions() *loadOptions {
	return &loadOptions{
		envEnabled:      true,
		validateEnabled: true,
		jsonFile:        "",
		yamlFile:        "",
	}
}

// FromEnv switch load from env.
// Default is enabled.
func FromEnv(enabled bool) loadOption {
	return func(lo *loadOptions) {
		lo.envEnabled = enabled
	}
}

// WithValidation switch validation for config.
// Defaut is enabled.
func WithValidation(enabled bool) loadOption {
	return func(lo *loadOptions) {
		lo.validateEnabled = enabled
	}
}

// FromJsonFile set json file as cfg source.
func FromJsonFile(fn string) loadOption {
	return func(lo *loadOptions) {
		lo.jsonFile = fn
	}
}

// FromYamlFile set json file as cfg source.
func FromYamlFile(fn string) loadOption {
	return func(lo *loadOptions) {
		lo.yamlFile = fn
	}
}
