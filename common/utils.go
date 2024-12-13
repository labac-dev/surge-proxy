package common

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
)

func Response(res *http.Response, body []byte) error {
	buff := new(bytes.Buffer)

	if res.Uncompressed && len(body) > 1024 {
		w, err := gzip.NewWriterLevel(buff, gzip.BestCompression)
		if err != nil {
			return err
		}

		_, err = w.Write(body)
		if err != nil {
			return err
		}

		err = w.Close()
		if err != nil {
			return err
		}

		res.Header.Set("Content-Encoding", "gzip")
	} else {
		buff.Write(body)
	}

	res.Header.Set("Content-Length", fmt.Sprint(buff.Len()))
	res.Body = io.NopCloser(buff)

	return nil
}
