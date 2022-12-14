package helpers

import (
	"fmt"
	"strings"
)

/*
 Copyright (C) 2022 Eunimart Omnichannel Pvt Ltd. (www.eunimart.com)
 All rights reserved.
 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.
 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.
 You should have received a copy of the GNU General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
func ApplySearch(query string, fields []string) (string, []interface{}) {

	var field_array []string
	var values []interface{}

	for _, field := range fields {

		temp1 := fmt.Sprintf("%v ILIKE ?", field)

		temp2 := "%" + query + "%"

		field_array = append(field_array, temp1)
		values = append(values, temp2)
	}

	field_string := strings.Join(field_array, " OR ")

	return field_string, values
}
