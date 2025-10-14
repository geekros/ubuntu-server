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

package main

import (
	"fmt"
	"os"

	"github.com/geekros/ubuntu-server/pkg/i18n"
	"github.com/geekros/ubuntu-server/pkg/version"
	"github.com/spf13/cobra"
)

// main is the entry point for the application.
func main() {
	// Create the root Cobra command for the CLI application.
	cmd := &cobra.Command{
		Use:   version.Name,                                                                                                                  // Command name (from version package)
		Short: version.Describe,                                                                                                              // Short description
		Long:  fmt.Sprintf("%s - %s %s %s (%s)", version.Name, version.Describe, version.Number, i18n.Get.GetSystemLanguage(), version.Site), // Detailed description
	}

	// Hide the default completion and help commands.
	cmd.CompletionOptions.HiddenDefaultCmd = true
	cmd.SetHelpCommand(&cobra.Command{
		Hidden: true,
	})

	// Initialize i18n
	i18n.Get = i18n.New()

	// Execute the root command, starting CLI parsing and command dispatch.
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
