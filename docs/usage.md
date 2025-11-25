# Using the E1AP Library

This guide explains how to use the E1AP library to create, encode, and decode messages.

## Installation

```bash
go get github.com/sub2pewds12/E1AP
```

## Creating and Encoding a Message

To send an E1AP message, you need to:
1.  Import the necessary packages.
2.  Create the specific message struct (e.g., `BearerContextSetupRequest`).
3.  Populate its fields (IEs).
4.  Wrap it in the appropriate PDU type (`InitiatingMessage`, `SuccessfulOutcome`, etc.).
5.  Encode it using `e1ap.E1apEncode`.

```go
package main

import (
	"fmt"
	"log"

	"github.com/lvdund/ngap/aper"
	"github.com/sub2pewds12/E1AP"
	"github.com/sub2pewds12/E1AP/e1ap_ies"
)

func main() {
	// 1. Create the inner message
	setupRequest := e1ap_ies.BearerContextSetupRequest{
		ProtocolIEs: e1ap_ies.ProtocolIEContainer{
			// Add IEs here...
            // Example: gNB-CU-CP UE E1AP ID
            List: []e1ap_ies.ProtocolIEField{
                {
                    Id: e1ap_ies.ProtocolIEIDGNB_CU_CP_UE_E1AP_ID,
                    Criticality: e1ap_ies.Criticality{Value: aper.Enumerated(e1ap_ies.CriticalityReject)},
                    Value: e1ap_ies.GNB_CU_CP_UE_E1AP_ID{
                        Value: aper.Integer(12345),
                    },
                },
            },
		},
	}

	// 2. Wrap in InitiatingMessage
	pdu := e1ap.E1AP_PDU{
		Present: e1ap.PduChoiceInitiatingMessage,
		Message: &e1ap.InitiatingMessage{
			ProcedureCode: e1ap_ies.ProcedureCode{Value: e1ap_ies.ProcedureCodeBearerContextSetup},
			Criticality:   e1ap_ies.Criticality{Value: aper.Enumerated(e1ap_ies.CriticalityReject)},
			Value:         &setupRequest,
		},
	}

	// 3. Encode
	bytes, err := e1ap.E1apEncode(&pdu)
	if err != nil {
		log.Fatalf("Encoding failed: %v", err)
	}

	fmt.Printf("Encoded bytes: %x\n", bytes)
}
```

## Decoding a Message

To receive and process an E1AP message:
1.  Receive the byte slice (e.g., from SCTP).
2.  Call `e1ap.E1apDecode`.
3.  Switch on the `Present` field to determine the PDU type.
4.  Type assert the `Value` to the specific message struct.

```go
package main

import (
	"fmt"
	"log"

	"github.com/sub2pewds12/E1AP"
	"github.com/sub2pewds12/E1AP/e1ap_ies"
)

func handleMessage(data []byte) {
	// 1. Decode
	pdu, diagnostics, err := e1ap.E1apDecode(data)
	if err != nil {
		log.Fatalf("Decoding failed: %v", err)
	}
    
    if len(diagnostics) > 0 {
        log.Printf("Diagnostics: %v", diagnostics)
    }

	// 2. Handle based on PDU type
	switch pdu.Present {
	case e1ap.PduChoiceInitiatingMessage:
		initMsg := pdu.Message.(*e1ap.InitiatingMessage)
		processInitiatingMessage(initMsg)
	case e1ap.PduChoiceSuccessfulOutcome:
		succMsg := pdu.Message.(*e1ap.SuccessfulOutcome)
		// Handle successful outcome...
        fmt.Println("Received SuccessfulOutcome", succMsg.ProcedureCode.Value)
	case e1ap.PduChoiceUnsuccessfulOutcome:
		failMsg := pdu.Message.(*e1ap.UnsuccessfulOutcome)
		// Handle unsuccessful outcome...
        fmt.Println("Received UnsuccessfulOutcome", failMsg.ProcedureCode.Value)
	}
}

func processInitiatingMessage(msg *e1ap.InitiatingMessage) {
	switch msg.ProcedureCode.Value {
	case e1ap_ies.ProcedureCodeBearerContextSetup:
		req := msg.Value.(*e1ap_ies.BearerContextSetupRequest)
		fmt.Printf("Received Bearer Context Setup Request. IEs: %v\n", req.ProtocolIEs.List)
	default:
		fmt.Printf("Unknown procedure: %d\n", msg.ProcedureCode.Value)
	}
}
```

## Working with IEs

Most IEs are generated as structs in the `e1ap_ies` package.

*   **Integers**: Wrapped in `aper.Integer` (int64).
*   **Enumerations**: Wrapped in `aper.Enumerated` (int). Use the generated constants (e.g., `e1ap_ies.CriticalityReject`).
*   **Bit Strings**: Wrapped in `aper.BitString` (`{Bytes []byte, NumBits uint64}`).
*   **Octet Strings**: Wrapped in `aper.OctetString` (`[]byte`).

### Lists (SEQUENCE OF)

Lists are typically represented as slices within a struct. For example, `ProtocolIEContainer` contains a `List` field which is a slice of `ProtocolIEField`.

### Extensions

Extensions are accessed via the `ProtocolExtensionContainer` field (if present) in the IE struct. The generated code provides helper structs to access specific extensions in a type-safe manner.
