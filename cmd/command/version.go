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

package command

import (
	"fmt"
	"log"

	"github.com/geekros/ubuntu-server/pkg/version"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

// Version creates and returns a new Cobra command for displaying version information.
// This command prints the application's name, version number, and site.
func Version() *cobra.Command {
	// Define the version command
	command := &cobra.Command{
		Use:     "version",            // Command name to be used on the CLI
		Short:   "Get version number", // Short description for help command
		Long:    "Get version number", // Long description for help command
		Example: "geekros version",    // Example usage of the command
		Run:     versionRun,           // Function to execute when command is called
	}

	return command
}

// versionRun is the execution function for the version command.
// It prints the application version information to the log.
func versionRun(cmd *cobra.Command, args []string) {
	log.Println(color.Gray.Text(fmt.Sprintf("[version] %s v%s (%s)", version.Name, version.Number, version.Site)))
}
