package logutil

import (
	"bufio"
	"bytes"

	"github.com/go-logfmt/logfmt"
)

func LogfmtString(m map[string][]string) (string, error) {
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	e := logfmt.NewEncoder(w)
	for k, vs := range m {
		for _, v := range vs {
			if err := e.EncodeKeyval(k, v); err != nil {
				return "", err
			}
		}
	}
	if err := e.EndRecord(); err != nil {
		return "", err
	}
	return buf.String(), nil
}
