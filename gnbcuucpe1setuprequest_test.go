package e1ap_test

import (
	"reflect"
	"testing"

	"github.com/lvdund/ngap/aper"
	e1ap "github.com/sub2pewds12/E1AP"
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

	encodedBytes, err := e1ap.E1apEncode(originalMsg)
	if err != nil {
		t.Fatalf("E1apEncode failed: %v", err)
	}
	t.Logf("Encoded PDU (len=%d): %X", len(encodedBytes), encodedBytes)

	decodedPDU, diags, err := e1ap.E1apDecode(encodedBytes)
	if err != nil {
		t.Fatalf("E1apDecode failed: %v", err)
	}
	if len(diags) > 0 {
		t.Logf("Received non-critical diagnostics: %+v", diags)
	}

	im, ok := decodedPDU.Message.(*e1ap.InitiatingMessage)
	if !ok {
		t.Fatalf("Decoded message is not *e1ap.InitiatingMessage. Got %T", decodedPDU.Message)
	}
	decodedMsg, ok := im.Value.(*e1ap_ies.GNBCUCPE1SetupRequest)
	if !ok {
		t.Fatalf("Decoded message value is not *e1ap_ies.GNBCUCPE1SetupRequest. Got %T", im.Value)
	}

	// Workaround for APER decoder appending null byte to strings
	if decodedMsg.GNBCUCPName != nil && len(decodedMsg.GNBCUCPName.Value) > 0 && decodedMsg.GNBCUCPName.Value[len(decodedMsg.GNBCUCPName.Value)-1] == 0 {
		decodedMsg.GNBCUCPName.Value = decodedMsg.GNBCUCPName.Value[:len(decodedMsg.GNBCUCPName.Value)-1]
	}

	if !reflect.DeepEqual(originalMsg, decodedMsg) {
		t.Errorf("Mismatch after round-trip.")
		t.Errorf("Original TransactionID: %v", originalMsg.TransactionID)
		t.Errorf("Decoded  TransactionID: %v", decodedMsg.TransactionID)
		if originalMsg.GNBCUCPName != nil && decodedMsg.GNBCUCPName != nil {
			t.Errorf("Original GNBCUCPName: %v", *originalMsg.GNBCUCPName)
			t.Errorf("Decoded  GNBCUCPName: %v", *decodedMsg.GNBCUCPName)
		}
		if originalMsg.TransportLayerAddressInfo != nil && decodedMsg.TransportLayerAddressInfo != nil {
			t.Errorf("Original TransportLayerAddressInfo: %+v", *originalMsg.TransportLayerAddressInfo)
			t.Errorf("Decoded  TransportLayerAddressInfo: %+v", *decodedMsg.TransportLayerAddressInfo)
		}
	} else {
		t.Logf("Successfully performed round-trip for GNBCUCPE1SetupRequest.")
	}
}
