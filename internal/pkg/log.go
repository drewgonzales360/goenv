// ///////////////////////////////////////////////////////////////////////
// Copyright 2023 Drew Gonzales
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// ///////////////////////////////////////////////////////////////////////
package pkg

import (
	"os"

	"github.com/fatih/color"
)

func Success(mesg string) {
	color.Green("😎 %s", mesg)
}

func Error(mesg string) {
	color.Red("😭 %s", mesg)
}

func Info(mesg string) {
	color.White("😃 %s", mesg)
}

func Warn(mesg string) {
	color.Yellow("😥 %s", mesg)
}

func Debug(mesg string) {
	logLevel := os.Getenv("GOENV_LOG")
	if logLevel == "DEBUG" {
		color.Blue("🤔 %s", mesg)
	}
}
