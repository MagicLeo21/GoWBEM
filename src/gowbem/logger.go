//    Copyright (C) 2017  Chen Yan
//
//    This program is free software: you can redistribute it and/or modify
//    it under the terms of the GNU General Public License as published by
//    the Free Software Foundation, either version 3 of the License, or
//    (at your option) any later version.
//
//    This program is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//    GNU General Public License for more details.
//
//    You should have received a copy of the GNU General Public License
//    along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
//    Author: Chen Yan <leochenlinux@gmail.com>

package gowbem

import (
	"fmt"
	"log"
)

var loggerEnabled bool = false

func SetLoggerEnabled(enabled bool) {
	loggerEnabled = enabled
}

func IsLoggerEnabled() bool {
	return loggerEnabled
}

func loggerPrint(format string, args ...interface{}) {
	if true == loggerEnabled {
		log.Print(fmt.Sprintf(format, args...))
	}
}
