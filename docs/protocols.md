# E1AP Protocol Implementation

This library implements the E1 Application Protocol (E1AP) as defined in 3GPP TS 38.463. E1AP is the control plane interface between the gNB-CU-CP (Control Plane) and gNB-CU-UP (User Plane).

## Message Structure

E1AP messages are encapsulated in a top-level PDU (`E1AP_PDU`). This PDU is a `CHOICE` of three types:

1.  **`InitiatingMessage`**: Starts a procedure (e.g., `BearerContextSetupRequest`).
2.  **`SuccessfulOutcome`**: A positive response to an initiating message (e.g., `BearerContextSetupResponse`).
3.  **`UnsuccessfulOutcome`**: A negative response to an initiating message (e.g., `BearerContextSetupFailure`).

Each of these wrappers contains:
*   `ProcedureCode`: Identifies the procedure (e.g., `id-BearerContextSetup`).
*   `Criticality`: `reject`, `ignore`, or `notify`.
*   `Value`: The actual message content (an `OPEN TYPE` container).

## Encoding/Decoding

The library uses **Aligned Packet Encoding Rules (APER)** for serialization, provided by the `github.com/lvdund/ngap/aper` package.

### Encoding Flow
1.  Create a specific message struct (e.g., `BearerContextSetupRequest`).
2.  Wrap it in an `InitiatingMessage` (or other wrapper).
3.  Wrap that in an `E1AP_PDU`.
4.  Call `E1apEncode(pdu)`.
    *   This calls `pdu.Encode()`.
    *   The PDU encoder writes the `ProcedureCode` and `Criticality`.
    *   It then encodes the inner message into a byte buffer (the `OPEN TYPE` container).
    *   Finally, it writes the length and bytes of the container.

### Decoding Flow
1.  Call `E1apDecode(bytes)`.
2.  The function reads the `Present` field (Initiating/Successful/Unsuccessful).
3.  It reads the `ProcedureCode` and `Criticality`.
4.  It reads the `OPEN TYPE` container (the message body) as a byte slice.
5.  It calls `createMessage()` (in `E1APMessageIE.go`) to instantiate the correct struct based on the `ProcedureCode`.
6.  It calls `Decode()` on the instantiated struct, passing the container bytes.

## Information Elements (IEs)

IEs are generated as Go structs in the `e1ap_ies` package.

*   **SEQUENCE**: Mapped to a struct. Optional fields are pointers or have a `Presence` flag (depending on the implementation details, usually handled by the APER library's `Optional` support or pointer nil checks).
*   **CHOICE**: Mapped to a struct with a `Present` integer and a `Choice` interface or value holder.
*   **ENUMERATED**: Mapped to a struct wrapping an `int` (via `aper.Enumerated`).
*   **SEQUENCE OF**: Mapped to a struct wrapping a slice.

## Extensions

E1AP supports protocol extensions. These are handled via `ProtocolExtensionContainer`. The generator creates specific extension structs (e.g., `BearerContextSetupRequestExtensions`) to provide type-safe access to known extensions.
