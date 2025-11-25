# E1AP Library Architecture

This project implements the E1AP protocol (3GPP TS 38.463) for Go. It uses a code generation approach to create Go structs and methods from ASN.1 definitions.

## Project Structure

The project is organized into the following key directories:

*   **`generator_scripts/`**: Contains the Python scripts responsible for parsing ASN.1 specifications and generating the Go code.
*   **`e1ap_ies/`**: Contains the generated Go code for E1AP Information Elements (IEs), PDUs, and constants.
*   **`config/`**: Configuration files for the generator, including patches and acronym definitions.
*   **`extracted_sections/`**: (Presumably) Contains the raw ASN.1 text extracted from the 3GPP specification.
*   **Root Directory**: Contains core runtime files (`E1APMessageIE.go`, `e1ap_pdu.go`, `encode.go`, `decode.go`) that provide the glue logic for message creation and serialization.

## High-Level Workflow

1.  **Input**: The process starts with ASN.1 definitions (likely extracted from the 3GPP PDF/text).
2.  **Parsing**: The `generator_scripts/main.py` script uses `asn1_parser.py` to parse these definitions into an internal representation (`ASN1Definition` objects).
3.  **Patching**: `asn1_patches.py` applies manual fixes to handle irregularities or errors in the source ASN.1.
4.  **Generation**: `go_generator.py` iterates over the parsed definitions and generates Go files in `e1ap_ies/`. It maps ASN.1 types to Go types, primarily using the `github.com/lvdund/ngap/aper` library for APER encoding/decoding.
5.  **Runtime**: The generated code is used in conjunction with the core files in the root directory to create, encode, and decode E1AP messages.

## Key Components

### Generator (`generator_scripts/`)
*   **`main.py`**: Entry point. Orchestrates loading, parsing, patching, and generation.
*   **`asn1_parser.py`**: A custom recursive descent parser for ASN.1. Handles complex constructs like `CLASS`, `SEQUENCE`, `CHOICE`, and parameterized types.
*   **`go_generator.py`**: Converts parsed definitions into Go code. It handles naming conventions (snake_case files, PascalCase structs) and type mapping.
*   **`go_templates.py` & `method_templates.py`**: Jinja2-like templates (implemented in Python) for generating struct definitions and `Encode`/`Decode` methods.

### Runtime (Root & `e1ap_ies/`)
*   **`E1APMessageIE.go`**: A factory that instantiates specific message structs based on the Procedure Code and Criticality.
*   **`e1ap_pdu.go`**: Defines the top-level `E1AP_PDU` structure and the `InitiatingMessage`, `SuccessfulOutcome`, and `UnsuccessfulOutcome` wrappers.
*   **`encode.go` / `decode.go`**: High-level functions `E1apEncode` and `E1apDecode` that handle the byte stream and delegate to the specific message encoders/decoders.
*   **`e1ap_ies/`**: The bulk of the library. Contains a file for each complex IE (e.g., `BearerContextSetupRequest.go`) and common types (`e1ap_common_types.go`).

## Dependencies

*   **`github.com/lvdund/ngap/aper`**: Used for Aligned Packet Encoding Rules (APER) serialization. This project relies heavily on this library for the low-level encoding details.
