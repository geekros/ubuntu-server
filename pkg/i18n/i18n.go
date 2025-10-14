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

package i18n

import (
	"github.com/geekros/ubuntu-server/pkg/i18n/drives"
	"github.com/geekros/ubuntu-server/pkg/i18n/language"
)

// Get is a global instance of I18n used for accessing language settings.
var Get = &I18n{}

// I18n struct holds internationalization settings for the application.
// Language: The current language code (e.g., "en" for English).
type I18n struct {
	// Language represents the current language setting (e.g., "en" for English, "zh" for Chinese).
	Language string
}

// New creates and returns a new I18n instance with default language set to English.
func New() *I18n {
	return &I18n{
		Language: "en", // Default language is English
	}
}

func (l *I18n) GetSystemLanguage() string {
	return drives.GetSystemLanguage()
}

// SetLanguage sets the language for the I18n instance based on the provided acceptLanguage string.
func (l *I18n) SetLanguage(acceptLanguage string) *I18n {
	// Set the language based on the first two characters of the acceptLanguage string.
	if len(acceptLanguage) >= 2 {
		l.Language = acceptLanguage[:2]
	}

	return l
}

// Lang retrieves the localized string for a given code based on the current language setting.
func (l *I18n) Lang(code int) string {
	// Retrieve the localized string based on the current language setting.
	string := language.GetEnLanguage(code)

	// Retrieve the localized string for the current language setting.
	if l.Language == "zh" {
		// Get the Chinese language string.
		string = language.GetZhLanguage(code)
	}

	return string
}
