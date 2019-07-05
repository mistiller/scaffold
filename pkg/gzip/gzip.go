package gzip

import(
	"fmt"
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

func Zip(payload []byte) []byte {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	w.Write([]byte(payload))
	w.Close()

	return b.Bytes()
}

func Unzip(payload []byte) (by []byte, err error) {
	b := bytes.NewBuffer(payload)
	r, err := gzip.NewReader(b)
	if err != nil {
		return by, fmt.Errorf("Decompress payload: %v", err)
	}
	defer r.Close()

	by, err = ioutil.ReadAll(r)
	if err != nil {
		return by, fmt.Errorf("Extract payload: %v", err)
	}
	
	return by, nil
}