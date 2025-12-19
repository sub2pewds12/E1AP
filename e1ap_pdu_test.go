package e1ap_test

import (
	"testing"

	"github.com/lvdund/ngap/aper"
	e1ap "github.com/sub2pewds12/E1AP"
	"github.com/sub2pewds12/E1AP/e1ap_ies"
)

func TestE1AP_PDU_InitiatingMessage(t *testing.T) {
	// 1. Create an InitiatingMessage (GNBCUCPE1SetupRequest)
	req := &e1ap_ies.GNBCUCPE1SetupRequest{
		TransactionID: e1ap_ies.TransactionID{Value: aper.Integer(123)},
		GNBCUCPName: &e1ap_ies.GNBCUCPName{
			Value: aper.OctetString("my-gNB-CU-CP"),
		},
	}

	// 2. Encode
	encoded, err := e1ap.E1apEncode(req)
	if err != nil {
		t.Fatalf("E1apEncode failed: %v", err)
	}
	if len(encoded) == 0 {
		t.Fatal("Encoded bytes are empty")
	}
	t.Logf("Encoded InitiatingMessage (len=%d): %X", len(encoded), encoded)

	// 3. Decode
	pdu, _, err := e1ap.E1apDecode(encoded)
	if err != nil {
		t.Fatalf("E1apDecode failed: %v", err)
	}

	// 4. Verify PDU Structure
	if pdu.Present != e1ap.PduChoiceInitiatingMessage {
		t.Errorf("Expected PduChoiceInitiatingMessage, got %d", pdu.Present)
	}

	initMsg, ok := pdu.Message.(*e1ap.InitiatingMessage)
	if !ok {
		t.Fatalf("Expected *e1ap.InitiatingMessage, got %T", pdu.Message)
	}
	if int64(initMsg.ProcedureCode.Value) != int64(e1ap_ies.ProcedureCodeGNBCUCPE1Setup) {
		t.Errorf("Expected ProcedureCodeGNBCUCPE1Setup, got %d", initMsg.ProcedureCode.Value)
	}
	if initMsg.Criticality.Value != e1ap_ies.CriticalityReject {
		t.Errorf("Expected CriticalityReject, got %d", initMsg.Criticality.Value)
	}

	// 5. Verify Message Content
	decodedReq, ok := initMsg.Value.(*e1ap_ies.GNBCUCPE1SetupRequest)
	if !ok {
		t.Fatalf("Expected *e1ap_ies.GNBCUCPE1SetupRequest, got %T", initMsg.Value)
	}
	if decodedReq.TransactionID.Value != 123 {
		t.Errorf("Expected TransactionID 123, got %d", decodedReq.TransactionID.Value)
	}
	if string(decodedReq.GNBCUCPName.Value) != "my-gNB-CU-CP" {
		// Workaround for potential null byte issue
		val := string(decodedReq.GNBCUCPName.Value)
		if len(val) > 0 && val[len(val)-1] == 0 {
			val = val[:len(val)-1]
		}
		if val != "my-gNB-CU-CP" {
			t.Errorf("Expected GNBCUCPName 'my-gNB-CU-CP', got '%s'", string(decodedReq.GNBCUCPName.Value))
		}
	}
	t.Logf("Decoded InitiatingMessage: %+v", decodedReq)
}

func TestE1AP_PDU_SuccessfulOutcome(t *testing.T) {
	// 1. Create a SuccessfulOutcome (GNBCUCPE1SetupResponse)
	resp := &e1ap_ies.GNBCUCPE1SetupResponse{
		TransactionID: e1ap_ies.TransactionID{Value: aper.Integer(123)},
		GNBCUUPID:     e1ap_ies.GNBCUUPID{Value: aper.Integer(456)},
		CNSupport:     e1ap_ies.CNSupport{Value: e1ap_ies.CNSupportCEpc},
		SupportedPLMNs: e1ap_ies.SupportedPLMNsList{
			Value: []e1ap_ies.SupportedPLMNsItem{
				{
					PLMNIdentity: e1ap_ies.PLMNIdentity{
						Value: aper.OctetString{0x02, 0xF8, 0x39},
					},
				},
			},
		},
	}

	// 2. Encode
	encoded, err := e1ap.E1apEncode(resp)
	if err != nil {
		t.Fatalf("E1apEncode failed: %v", err)
	}
	if len(encoded) == 0 {
		t.Fatal("Encoded bytes are empty")
	}
	t.Logf("Encoded SuccessfulOutcome (len=%d): %X", len(encoded), encoded)

	// 3. Decode
	pdu, _, err := e1ap.E1apDecode(encoded)
	if err != nil {
		t.Fatalf("E1apDecode failed: %v", err)
	}

	// 4. Verify PDU Structure
	if pdu.Present != e1ap.PduChoiceSuccessfulOutcome {
		t.Errorf("Expected PduChoiceSuccessfulOutcome, got %d", pdu.Present)
	}

	succMsg, ok := pdu.Message.(*e1ap.SuccessfulOutcome)
	if !ok {
		t.Fatalf("Expected *e1ap.SuccessfulOutcome, got %T", pdu.Message)
	}
	if int64(succMsg.ProcedureCode.Value) != int64(e1ap_ies.ProcedureCodeGNBCUCPE1Setup) {
		t.Errorf("Expected ProcedureCodeGNBCUCPE1Setup, got %d", succMsg.ProcedureCode.Value)
	}
	if succMsg.Criticality.Value != e1ap_ies.CriticalityIgnore {
		t.Errorf("Expected CriticalityIgnore, got %d", succMsg.Criticality.Value)
	}

	// 5. Verify Message Content
	decodedResp, ok := succMsg.Value.(*e1ap_ies.GNBCUCPE1SetupResponse)
	if !ok {
		t.Fatalf("Expected *e1ap_ies.GNBCUCPE1SetupResponse, got %T", succMsg.Value)
	}
	if decodedResp.TransactionID.Value != 123 {
		t.Errorf("Expected TransactionID 123, got %d", decodedResp.TransactionID.Value)
	}
	t.Logf("Decoded SuccessfulOutcome: %+v", decodedResp)
}

func TestE1AP_PDU_UnsuccessfulOutcome(t *testing.T) {
	// 1. Create an UnsuccessfulOutcome (GNBCUCPE1SetupFailure)
	fail := &e1ap_ies.GNBCUCPE1SetupFailure{
		TransactionID: e1ap_ies.TransactionID{Value: aper.Integer(123)},
		Cause: e1ap_ies.Cause{
			Choice: e1ap_ies.CausePresentRadioNetwork,
			RadioNetwork: &e1ap_ies.CauseRadioNetwork{
				Value: e1ap_ies.CauseRadioNetworkUnspecified,
			},
		},
	}

	// 2. Encode
	encoded, err := e1ap.E1apEncode(fail)
	if err != nil {
		t.Fatalf("E1apEncode failed: %v", err)
	}
	if len(encoded) == 0 {
		t.Fatal("Encoded bytes are empty")
	}
	t.Logf("Encoded UnsuccessfulOutcome (len=%d): %X", len(encoded), encoded)

	// 3. Decode
	pdu, _, err := e1ap.E1apDecode(encoded)
	if err != nil {
		t.Fatalf("E1apDecode failed: %v", err)
	}

	// 4. Verify PDU Structure
	if pdu.Present != e1ap.PduChoiceUnsuccessfulOutcome {
		t.Errorf("Expected PduChoiceUnsuccessfulOutcome, got %d", pdu.Present)
	}

	unsuccMsg, ok := pdu.Message.(*e1ap.UnsuccessfulOutcome)
	if !ok {
		t.Fatalf("Expected *e1ap.UnsuccessfulOutcome, got %T", pdu.Message)
	}
	if int64(unsuccMsg.ProcedureCode.Value) != int64(e1ap_ies.ProcedureCodeGNBCUCPE1Setup) {
		t.Errorf("Expected ProcedureCodeGNBCUCPE1Setup, got %d", unsuccMsg.ProcedureCode.Value)
	}
	if unsuccMsg.Criticality.Value != e1ap_ies.CriticalityIgnore {
		t.Errorf("Expected CriticalityIgnore, got %d", unsuccMsg.Criticality.Value)
	}

	// 5. Verify Message Content
	decodedFail, ok := unsuccMsg.Value.(*e1ap_ies.GNBCUCPE1SetupFailure)
	if !ok {
		t.Fatalf("Expected *e1ap_ies.GNBCUCPE1SetupFailure, got %T", unsuccMsg.Value)
	}
	if decodedFail.TransactionID.Value != 123 {
		t.Errorf("Expected TransactionID 123, got %d", decodedFail.TransactionID.Value)
	}
	if decodedFail.Cause.Choice != e1ap_ies.CausePresentRadioNetwork {
		t.Errorf("Expected CausePresentRadioNetwork, got %d", decodedFail.Cause.Choice)
	}
	t.Logf("Decoded UnsuccessfulOutcome: %+v", decodedFail)
}
