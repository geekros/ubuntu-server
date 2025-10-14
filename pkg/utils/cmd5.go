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

package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// EncryptMD5 generates an MD5 hash from the input string.
// It performs the following steps:
//  1. Creates a new MD5 hasher
//  2. Writes the input bytes to the hasher
//  3. Calculates the final hash
//  4. Converts the hash to a hexadecimal string
//
// Parameters:
//   - input: The string to be hashed
//
// Returns:
//   - A 32-character hexadecimal string representing the MD5 hash
func EncryptMD5(input string) string {
	hash := md5.New()                    // Initialize MD5 hasher
	hash.Write([]byte(input))            // Write input bytes to hasher
	hashBytes := hash.Sum(nil)           // Generate final hash
	return hex.EncodeToString(hashBytes) // Convert to hex string
}
