package e1ap

import (
	"bytes"
)

func E1apEncode(msg MessageEncoder) (wire []byte, err error) {
	var buf bytes.Buffer

	if err = msg.Encode(&buf); err == nil {
		wire = buf.Bytes()
	}
	return
}
