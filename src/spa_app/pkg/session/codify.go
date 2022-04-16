/*
Author: Connor Sanders
Copyright: Connor Sanders 2020
Version: 0.0.1
Released: 12/10/2020

-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-
		Golang Frontend Boilerplate V0.0.1
-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-
*/

package session

import (
	"crypto/sha256"
	"encoding/hex"
)

// Runs a SHA 2 function on the string
func SHA(str string) string {
	bytes := []byte(str)
	// Converts string to sha2
	h := sha256.New()                   // new sha256 object
	h.Write(bytes)                      // data is now converted to hex
	code := h.Sum(nil)                  // code is now the hex sum
	codestr := hex.EncodeToString(code) // converts hex to string
	return codestr
}
