package util

import (
    "strings"
    "log"

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
    log.Println("Hex:  " + hex);
    hex = strings.ToLower(strings.Replace(hex, "0x", "", 1));
    log.Println("Hex:  " + hex);

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
