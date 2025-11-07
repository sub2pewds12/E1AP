package e1ap

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/lvdund/ngap/aper"
)

func TestGNBCUCPE1SetupRequest(t *testing.T) {

	originalMsg := GNBCUCPE1SetupRequest{
		TransactionID: TransactionID{Value: aper.Integer(123)},
		GNBCUCPName: &GNBCUCPName{
			Value: aper.OctetString("my-gNB-CU-CP-Name-1"),
		},
		TransportLayerAddressInfo: &TransportLayerAddressInfo{
			Value: aper.BitString{
				Bytes:   []byte{0x0A, 0x0B, 0x0C, 0x0D},
				NumBits: 32,
			},
		},
		ExtendedGNBCUCPName: nil,
	}

	var buf bytes.Buffer

	if err := originalMsg.Encode(&buf); err != nil {

		t.Fatalf("Encoding failed: %v", err)
	}
	encodedBytes := buf.Bytes()

	t.Logf("Encoded PDU (len=%d): %X", len(encodedBytes), encodedBytes)

	decodedMsg := GNBCUCPE1SetupRequest{}
	err, diags := decodedMsg.Decode(encodedBytes)
	if err != nil {
		t.Fatalf("Decoding failed: %v", err)
	}
	if len(diags) > 0 {

		t.Logf("Received non-critical diagnostics: %+v", diags)
	}

	if !reflect.DeepEqual(originalMsg, decodedMsg) {

		t.Errorf("Mismatch after round-trip.\nOriginal: %+v\nDecoded:  %+v", originalMsg, decodedMsg)
	} else {
		t.Logf("Successfully performed round-trip for GNBCUCPE1SetupRequest.")
	}
}
