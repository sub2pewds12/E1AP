package e1ap

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/lvdund/ngap/aper"
	"github.com/sub2pewds12/E1AP/e1ap_ies"
)

func TestE1SetupRequest(t *testing.T) {

	originalMsg := &e1ap_ies.GNBCUCPE1SetupRequest{
		TransactionID: e1ap_ies.TransactionID{Value: aper.Integer(123)},
		GNBCUCPName: &e1ap_ies.GNBCUCPName{
			Value: aper.OctetString("my-gNB-CU-CP"),
		},
		TransportLayerAddressInfo: &e1ap_ies.TransportLayerAddressInfo{
			TransportUPLayerAddressesInfoToAddList: &e1ap_ies.TransportUPLayerAddressesInfoToAddList{
				Value: []e1ap_ies.TransportUPLayerAddressesInfoToAddItem{
					{
						IPSecTransportLayerAddress: e1ap_ies.TransportLayerAddress{
							Value: aper.BitString{
								Bytes:   []byte{0x0A, 0x0B, 0x0C, 0x0E},
								NumBits: 32,
							},
						},
						GTPTransportLayerAddressesToAdd: &e1ap_ies.GTPTLAs{
							Value: []e1ap_ies.GTPTLAItem{
								{
									GTPTransportLayerAddresses: e1ap_ies.TransportLayerAddress{
										Value: aper.BitString{
											Bytes:   []byte{0x0A, 0x0B, 0x0C, 0x0D},
											NumBits: 32,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	pduToEncode := &InitiatingMessage{
		ProcedureCode: e1ap_ies.ProcedureCode{Value: e1ap_ies.ProcedureCodeGNBCUPE1Setup},
		Criticality:   e1ap_ies.Criticality{Value: e1ap_ies.CriticalityReject},
		Value:         originalMsg,
	}

	var buf bytes.Buffer
	if err := pduToEncode.Encode(&buf); err != nil {
		t.Fatalf("Encode failed: %v", err)
	}
	encodedBytes := buf.Bytes()
	t.Logf("Encoded PDU (len=%d): %X", len(encodedBytes), encodedBytes)

	decodedPDU, diags, err := E1apDecode(encodedBytes)
	if err != nil {
		t.Fatalf("E1apDecode failed: %v", err)
	}
	if len(diags) > 0 {
		t.Logf("Received non-critical diagnostics: %+v", diags)
	}

	im, ok := decodedPDU.Message.(*InitiatingMessage)
	if !ok {
		t.Fatalf("Decoded PDU.Message is not of expected type *InitiatingMessage")
	}
	decodedMsg, ok := im.Value.(*e1ap_ies.GNBCUCPE1SetupRequest)
	if !ok {
		t.Fatalf("Decoded message is not of expected type *e1ap_ies.GNBCUCPE1SetupRequest")
	}

	if !reflect.DeepEqual(originalMsg, decodedMsg) {
		t.Errorf("Mismatch after round-trip.\nOriginal: %+v\nDecoded:  %+v", originalMsg, decodedMsg)
	} else {
		t.Logf("Successfully performed round-trip for GNBCUCPE1SetupRequest.")
	}
}
