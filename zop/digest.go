package zop

import (
	"crypto/md5"
	"encoding/base64"
)

func digest(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
