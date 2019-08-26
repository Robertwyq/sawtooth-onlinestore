/**
* Author：Robert WYQ
* Date：2019-7-22 ——  2019-8-31
* Description: To create an onlinestore APP base on the sawtooth1.1.5
* ------------------------------------------------------------------------------
*/

package utils

import (
	"crypto/sha512" // hash algorithm
	"encoding/hex"
	"strings"
)
// return with encryption hex code
func Hex_encryption(str string) string {
	hash := sha512.New() 		// return new hash
	hash.Write([]byte(str))
	hashBytes := hash.Sum(nil)  // define new hashbyte，return examine code
	return strings.ToLower(hex.EncodeToString(hashBytes)) // hex code，into unicode lower form
}
