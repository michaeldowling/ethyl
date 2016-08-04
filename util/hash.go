package util

import (
    sha "github.com/michaeldowling/ethyl/crypto/sha3"
    "encoding/hex"
)

func EthereumHash(value []byte) (string, error) {

    d := sha.NewKeccak256()
    d.Write(value);
    hashValue := d.Sum(nil);

    return "0x" + hex.EncodeToString(hashValue), nil;

}
