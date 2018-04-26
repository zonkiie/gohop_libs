package gohop_libs

import (
  "bytes"
	"regexp"
)

func RemoveNamespaceFromXML(xml []byte) (result []byte) {
	result = xml
	result = bytes.Replace(result, []byte("<ns1:"), []byte("<"), -1)
	result = bytes.Replace(result, []byte("</ns1:"), []byte("</"), -1)
	re := regexp.MustCompile(" xmlns:ns1=\".+\"")
	result = re.ReplaceAll(result, []byte(""))
	return result
}

