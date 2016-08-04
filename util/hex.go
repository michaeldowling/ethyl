/*
 * (C) Copyright 2016 Michael Dowling (http://github.com/michaeldowling) and others.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package util

import (
    "strings"
    "math"
)

var HEXTABLE = map[string]int64{
    "0":0,
    "1":1,
    "2":2,
    "3":3,
    "4":4,
    "5":5,
    "6":6,
    "7":7,
    "8":8,
    "9":9,
    "a":10,
    "b":11,
    "c":12,
    "d":13,
    "e":14,
    "f":15,
};

func ToInt64(hex string) (int64, error) {

    // Do we have a "0x" in the beginning?
    hex = strings.ToLower(strings.Replace(hex, "0x", "", 1));

    // # of chars
    var total int64 = 0;
    numOfChars := len(hex);
    for i := 0; i < numOfChars; i++ {

        exponent := int64(math.Pow(16, float64((numOfChars - i - 1))));
        num := HEXTABLE[string(hex[i])] * exponent;
        total += num;

    }

    return total, nil;

}
