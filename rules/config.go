package rules

import "regexp"

type Config struct {
	CheckLowercase         bool     `yaml:"check_lowercase" json:"check_lowercase"`
	CheckEnglishOnly       bool     `yaml:"check_english_only" json:"check_english_only"`
	CheckSpecialSymbols    bool     `yaml:"check_special_symbols" json:"check_special_symbols"`
	CheckSensitiveData     bool     `yaml:"check_sensitive_data" json:"check_sensitive_data"`
	CustomSensitiveRegexps []string `yaml:"sensitive_regexps" json:"sensitive_regexps"`
	CustomSensitiveSubstrs []string `yaml:"sensitive_ident_substrings" json:"sensitive_ident_substrings"`
}

var (
	currentConfig = Config{
		CheckLowercase:      true,
		CheckEnglishOnly:    true,
		CheckSpecialSymbols: true,
		CheckSensitiveData:  true,
	}
	customSensitiveRegexps []*regexp.Regexp
)

func GetConfig() Config {
	return currentConfig
}

func SetConfig(c Config) {
	currentConfig = c
	rebuildCustomRegexps()
}

func ApplySettings(conf any) {
	m, ok := conf.(map[string]any)
	if !ok {
		return
	}
	cfg := currentConfig
	if v, ok := m["check_lowercase"].(bool); ok {
		cfg.CheckLowercase = v
	}
	if v, ok := m["check_english_only"].(bool); ok {
		cfg.CheckEnglishOnly = v
	}
	if v, ok := m["check_special_symbols"].(bool); ok {
		cfg.CheckSpecialSymbols = v
	}
	if v, ok := m["check_sensitive_data"].(bool); ok {
		cfg.CheckSensitiveData = v
	}
	if v, ok := m["sensitive_regexps"]; ok {
		cfg.CustomSensitiveRegexps = extractStringSlice(v)
	}
	if v, ok := m["sensitive_ident_substrings"]; ok {
		cfg.CustomSensitiveSubstrs = extractStringSlice(v)
	}
	SetConfig(cfg)
}

func extractStringSlice(v any) []string {
	switch vv := v.(type) {
	case []string:
		return vv
	case []any:
		out := make([]string, 0, len(vv))
		for _, item := range vv {
			if s, ok := item.(string); ok {
				out = append(out, s)
			}
		}
		return out
	default:
		return nil
	}
}

func rebuildCustomRegexps() {
	customSensitiveRegexps = customSensitiveRegexps[:0]
	for _, pattern := range currentConfig.CustomSensitiveRegexps {
		if re, err := regexp.Compile(pattern); err == nil {
			customSensitiveRegexps = append(customSensitiveRegexps, re)
		}
	}
}
