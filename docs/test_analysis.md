# E1AP Test Suite Analysis

This document analyzes the three core test files in the E1AP project. These tests verify the correctness of the APER encoding/decoding logic, ensuring that Go structs are correctly serialized to binary (and printed as Hex) and deserialized back to identical Go structs.

## Testing Standards
All tests follow the **Round-Trip Pattern** mandated by the project requirements:
1.  **Construct**: Create a complex Go struct with nested IEs.
2.  **Encode**: Serialize it to APER binary.
3.  **Output**: Log the **Encoded Hex String** (Mentor Requirement).
4.  **Decode**: Parse the binary back into a new Go struct.
5.  **Verify**: Log the **Decoded Struct** and assert deep equality with the original.

---

## 1. `gnbcuucpe1setuprequest_test.go`
**Purpose**: The "Prototype" Test.
This file was created to explicitly replicate the mentor's testing style for a single, complex message: `GNBCUCPE1SetupRequest`.

### Key Features
*   **Complex PDU Construction**: Manually builds a deeply nested PDU including optional fields (`GNBCUCPName`) and lists (`TransportLayerAddressInfo`).
*   **Strict Type Assertion**: casts the generic `interface{}` message body to `*e1ap.InitiatingMessage` and then to the specific PDU struct.
*   **Deep Quality Check**: Uses `reflect.DeepEqual` to strictly verify that the input and output objects are identical field-for-field.
*   **Null Byte Handling**: Includes a workaround for null-terminated strings if present in the decoder.

**Output Example**:
```text
Encoded PDU (len=46): 0004002A00030039...
Decoded Message Structure: &{TransactionID:{Value:123} ...}
```

---

## 2. `e1ap_pdu_test.go`
**Purpose**: The "Main Coverage" Suite.
This file tests the three Top-Level PDU types defined in E1AP to ensure the dispatcher logic (`E1apEncode`/`E1apDecode`) works for all message categories.

### Test Cases
*   **`TestE1AP_PDU_InitiatingMessage`**: Verifies standard request messages (e.g., `GNBCUCPE1SetupRequest`).
*   **`TestE1AP_PDU_SuccessfulOutcome`**: Verifies robust response handling (e.g., `GNBCUCPE1SetupResponse`), including complex lists like `SupportedPLMNs`.
*   **`TestE1AP_PDU_UnsuccessfulOutcome`**: Verifies failure response handling (e.g., `GNBCUCPE1SetupFailure`), checking `Cause` and `CriticalityDiagnostics`.

### Verification Logic
Instead of `reflect.DeepEqual`, these tests use **Explicit Field Checks** (e.g., `if decodedReq.TransactionID.Value != 123`). This allows for more granular error reporting if a specific field fails to decode.

---

## 3. `E1AP_test.go`
**Purpose**: The "Helper & Manual" Suite.
This file contains unit tests for specific low-level encoding mechanisms and manual helper functions.

### Key Tests
*   **`Test_E1New`**: A quick manual test for creating and encoding a PDU.
*   **`Test_Sequence`**: Specifically validates the generic `Sequence[T]` helper logic used for LIST types (fixing the slice covariance issue).
*   **`Test_Bitstring`**: Verification of manual `BIT STRING` constraint handling.
*   **`Test_BITSTRING_Struct`**: Verifies the wrapper generation for Bit Strings.

**Refactoring Note**: This file was recently standardized to use `t.Logf` and `%X` hex formatting to match the other two files, replacing previous `fmt.Println` calls.

---

## Summary
| Test File | Focus | Method | Output Style |
| :--- | :--- | :--- | :--- |
| `gnbcuucpe1setuprequest_test.go` | Single Complex PDU | Deep Equality | Verbose Hex + Struct |
| `e1ap_pdu_test.go` | PDU Dispatcher Types | Field Assertion | Verbose Hex + Struct |
| `E1AP_test.go` | Low-level Helpers | Unit Checks | Verbose Hex + Struct |

All three files now strictly adhere to the requirement of outputting the **Encoded Hex String** for verification.

---

## How to Verify Output
To verify the Hex string manually (e.g., `0004002A00030039...`), you can read the octets against the E1AP specification (`e1ap_constants.go`):

| Hex Bytes | Meaning | Value | Notes |
| :--- | :--- | :--- | :--- |
| **00** | PDU Present | `0` | InitiatingMessage |
| **04** | ProcedureCode | `4` | `id-E1Setup` |
| **00** | Criticality | `0` | `Reject` |
| **2A** | Structure Length | `42` | Length of bytes following this header |
| **00 03**| IE Count | `3` | Number of IEs in this PDU |
| **00 39**| Protocol IE ID | `57` | `id-TransactionID` |
| **00 09**| Protocol IE ID | `9` | `id-gNB-CU-CP-Name` |
| **00 56**| Protocol IE ID | `86` | `id-TransportLayerAddressInfo` |

If the hex string starts with `00 04 ...` and contains `00 39`, you confirm it is a valid **E1 Setup Request** with a **Transaction ID**.
