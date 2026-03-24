package e1ap_test

import (
	"bytes"
	"testing"

	"github.com/lvdund/ngap/aper"
	e1ap "github.com/sub2pewds12/E1AP"
	"github.com/sub2pewds12/E1AP/e1ap_ies"
)

func Test_E1New(t *testing.T) {
	msg := e1ap_ies.GNBCUCPE1SetupRequest{
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
					},
				},
			},
		},
	}
	var b []byte
	var err error
	if b, err = e1ap.E1apEncode(&msg); err != nil {
		t.Fatalf("E1apEncode err: %v", err)
	}
	t.Logf("Encoded (len=%d): %X", len(b), b)

	var pdu e1ap.E1AP_PDU
	var diags []e1ap_ies.CriticalityDiagnosticsIEItem
	if pdu, diags, err = e1ap.E1apDecode(b); err != nil {
		t.Fatalf("E1apDecode err: %v", err)
	}
	if len(diags) > 0 {
		t.Logf("Diagnostics: %+v", diags)
	}
	t.Logf("Decoded PDU: %+v", pdu)

	// Type assertion to access specific fields
	if initiatingMsg, ok := pdu.Message.(*e1ap.InitiatingMessage); ok {
		if decode_msg, ok := initiatingMsg.Value.(*e1ap_ies.GNBCUCPE1SetupRequest); ok {
			t.Logf("Decoded TransactionID: %v", decode_msg.TransactionID)
			if decode_msg.GNBCUCPName != nil {
				t.Logf("Decoded GNBCUCPName: %s", string(decode_msg.GNBCUCPName.Value))
			}
			if decode_msg.TransportLayerAddressInfo != nil {
				t.Logf("Decoded TransportLayerAddressInfo: %+v", decode_msg.TransportLayerAddressInfo)
			}
		}
	}

	// Manual IE Encode/Decode Test (Simulating ies.NewBITSTRING)
	// Using GNBCUCPName as it's a simple wrapper type
	var buf bytes.Buffer
	w := aper.NewWriter(&buf)
	ie := e1ap_ies.GNBCUCPName{Value: aper.OctetString("test-name")}
	if err = ie.Encode(w); err != nil {
		t.Fatalf("err ie encode: %v", err)
	}
	w.Close()
	t.Logf("encode ie: %X", buf.Bytes())

	r := aper.NewReader(&buf)
	decodedIE := e1ap_ies.GNBCUCPName{}
	if err = decodedIE.Decode(r); err != nil {
		t.Fatalf("err ie decode: %v", err)
	}
	t.Logf("Decode ie: %s", string(decodedIE.Value))
}

func Test_Sequence(t *testing.T) {
	// Simulating Test_Sequence using TransportUPLayerAddressesInfoToAddList
	// Note: We don't have a generic NewSequence, so we use the generated struct directly.

	item := e1ap_ies.TransportUPLayerAddressesInfoToAddItem{
		IPSecTransportLayerAddress: e1ap_ies.TransportLayerAddress{
			Value: aper.BitString{
				Bytes:   []byte{0x0A, 0x0B, 0x0C, 0x0E},
				NumBits: 32,
			},
		},
	}

	list := e1ap_ies.TransportUPLayerAddressesInfoToAddList{
		Value: []e1ap_ies.TransportUPLayerAddressesInfoToAddItem{item, item},
	}

	var buf bytes.Buffer
	w := aper.NewWriter(&buf)
	if err := list.Encode(w); err != nil {
		t.Fatalf("encode sequence err: %v", err)
	}
	w.Close()
	t.Logf("Encode: %X", buf.Bytes())

	r := aper.NewReader(&buf)
	decodedList := e1ap_ies.TransportUPLayerAddressesInfoToAddList{}

	if err := decodedList.Decode(r); err != nil {
		t.Fatalf("decode sequence err: %v", err)
	}

	t.Logf("Decode %d items. First Item: %+v", len(decodedList.Value), decodedList.Value[0].IPSecTransportLayerAddress)
}

func Test_Bitstring(t *testing.T) {
	var buf bytes.Buffer
	var err error

	w := aper.NewWriter(&buf)
	b := []byte{0xa0}

	// Using constraints similar to the mentor's example
	if err = w.WriteBitString(
		b,
		1,
		&aper.Constraint{Lb: 1, Ub: 3},
		false); err != nil {
		t.Fatalf("err ie encode: %v", err)
	}
	w.Close()
	t.Logf("encode ie: %X", buf.Bytes())

	r := aper.NewReader(&buf)
	var content []byte
	var nbits uint
	content, nbits, err = r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 3}, false)
	if err != nil {
		t.Fatalf("err ie decode: %v", err)
	}
	t.Logf("Decode ie: %X (nbits=%d)", content, nbits)
}

func Test_BITSTRING_Struct(t *testing.T) {
	// Simulating Test_BITSTRING using a struct wrapper
	// TransportLayerAddress is a BitString wrapper
	ie := e1ap_ies.TransportLayerAddress{
		Value: aper.BitString{
			Bytes:   []byte{0x10, 0x10, 0x10, 0x10},
			NumBits: 32,
		},
	}

	var buf bytes.Buffer
	var err error

	w := aper.NewWriter(&buf)
	if err = ie.Encode(w); err != nil {
		t.Fatalf("err ie encode: %v", err)
	}
	w.Close()

	r := aper.NewReader(&buf)
	newie := e1ap_ies.TransportLayerAddress{}
	if err = newie.Decode(r); err != nil {
		t.Fatalf("err ie decode: %v", err)
	}
	t.Logf("Decoded Struct: %+v", newie)
}
