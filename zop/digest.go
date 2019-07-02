package zop

import (
	"crypto/md5"
	"encoding/base64"
)

// Digest generates the digest for the given string by MD5 hash and base64 encoding.
func Digest(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
