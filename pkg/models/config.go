/*
Copyright © 2019 Doppler <support@doppler.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package models

// Config options
type Config struct {
	Token     string `json:"token"`
	Project   string `json:"project"`
	Config    string `json:"config"`
	APIHost   string `json:"api-host"`
	VerifyTLS string `json:"verify-tls"`
}

// ScopedConfig options with their scope
type ScopedConfig struct {
	Token     Pair `json:"token"`
	Project   Pair `json:"project"`
	Config    Pair `json:"config"`
	APIHost   Pair `json:"api-host"`
	VerifyTLS Pair `json:"verify-tls"`
}

// Pair value and its scope
type Pair struct {
	Value  string `json:"value"`
	Scope  string `json:"scope"`
	Source string `json:"source"`
}

// Source where the value came from
type Source int

const (
	FlagSource Source = iota
	ConfigFileSource
	EnvironmentSource
	DefaultValueSource
)

func (s Source) String() string {
	return [...]string{"Flag", "Config File", "Environment", "Default Value"}[s]
}

// Pairs get the pairs for the given config
func Pairs(conf Config) map[string]string {
	return map[string]string{
		"token":      conf.Token,
		"project":    conf.Project,
		"config":     conf.Config,
		"api-host":   conf.APIHost,
		"verify-tls": conf.VerifyTLS,
	}
}
