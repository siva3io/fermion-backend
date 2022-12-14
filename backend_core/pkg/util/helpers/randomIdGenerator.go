package helpers

import (
	"math/rand"
	"strconv"
	"time"
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
func GenerateRandID(min int, max int, prefix string) string {
	rand.Seed(time.Now().UnixNano())
	if min == 0 {
		min = 1000
	}
	if max == 0 {
		max = 9999
	}
	randomNum := rand.Intn(max-min+1) + min
	randomId := strconv.Itoa(randomNum)
	return prefix + randomId
}
