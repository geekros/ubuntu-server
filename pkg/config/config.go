// Copyright 2025 GEEKROS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gookit/color"
	"gopkg.in/yaml.v3"
)

var Get = &Config{}

// Config struct holds the configuration details for the application.
type Config struct {
	Path      string  `yaml:"-" json:"path"`
	Workspace string  `yaml:"-" json:"workspace"`
	Runtime   string  `yaml:"-" json:"runtime"`
	Server    service `yaml:"server" json:"server"`
}

// service struct defines the server-related configuration.
type service struct {
	Mode         string        `yaml:"mode" json:"mode"`
	Port         int           `yaml:"port" json:"port"`
	ReadTimeout  time.Duration `yaml:"read_timeout" json:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout" json:"write_timeout"`
}

// New initializes a new Config instance with default paths.
// It sets the workspace, configuration file path, and runtime directory.
func New() *Config {
	// Define the default workspace directory.
	workspace := "/opt/geekros"

	// Define the default configuration file path.
	configPath := filepath.Join(workspace, "/release/config.yaml")

	// Return a new Config instance with the default values.
	return &Config{
		Path:      configPath,
		Workspace: workspace,
		Runtime:   filepath.Join(workspace, "/runtime"),
	}
}

// LoadConfig loads the configuration from the file.
// If the file does not exist, it initializes default values and creates the file.
func (c *Config) LoadConfig() *Config {
	// Check if the configuration file exists.
	if _, err := os.Stat(c.Path); os.IsNotExist(err) {
		// Set default server configuration.
		c.Server.Mode = "debug"
		c.Server.Port = 10800
		c.Server.ReadTimeout = 60 * time.Second
		c.Server.WriteTimeout = 60 * time.Second

		// Set default multimodal platform configuration.
		c.UpdateConfig()
	}

	// Read the configuration file.
	file, err := os.ReadFile(c.Path)
	if err != nil {
		// Return the default configuration if reading fails.
		return c
	}

	// Parse the YAML configuration file into the Config struct.
	err = yaml.Unmarshal(file, c)
	if err != nil {
		// Return the default configuration if parsing fails.
		return c
	}

	return c
}

// UpdateConfig saves the current configuration to the file.
func (c *Config) UpdateConfig() error {
	// Convert the Config struct to YAML format.
	data, err := yaml.Marshal(c)
	if err != nil {
		// Return the error if marshalling fails.
		return err
	}

	// Write the YAML data to the configuration file.
	err = os.WriteFile(c.Path, data, 0644)
	if err != nil {
		log.Println(color.Yellow.Text(fmt.Sprintf("[config] %s", "failed to update config file")))
		// Return the error if writing fails.
		return err
	}

	return nil
}
