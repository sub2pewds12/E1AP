# Code Generation Process

The E1AP library is largely auto-generated from ASN.1 specifications. This document explains how the generation pipeline works.

## Overview

The generation process is driven by Python scripts located in the `generator_scripts/` directory. The main entry point is `main.py`.

```bash
python3 generator_scripts/main.py -i extracted_sections/ -o e1ap_ies/
```

## Pipeline Stages

### 1. Loading Files
The script reads `.txt` files from the input directory (`extracted_sections/`). These files contain the raw ASN.1 text.
*   **Skipping**: Files listed in `config/skip_files.txt` are ignored.
*   **Context**: The loader tracks the filename and line number for every line of text to provide useful error messages.

### 2. Parsing (`asn1_parser.py`)
The `ASN1Parser` class processes the loaded lines. It uses a state machine to handle multi-line definitions.
*   **Simple Types**: Parses `INTEGER`, `ENUMERATED`, `BIT STRING`, `OCTET STRING`.
*   **Complex Types**: Parses `SEQUENCE`, `CHOICE`, `SEQUENCE OF`.
*   **Classes**: Parses `CLASS` definitions (e.g., `E1AP-PROTOCOL-IES`).
*   **Parameterized Types**: Handles `ProtocolIE-Container`, `ProtocolExtensionContainer`, etc., by resolving them against the parsed `CLASS` definitions.

### 2.5. Patching (`asn1_patches.py`)
After parsing, the `ASN1Patcher` applies manual fixes.
*   **Hardcoded Definitions**: Some definitions are too complex or irregular for the parser. These are defined manually in `config/patches.yaml` (or similar) and injected into the definitions list.
*   **Fixes**: Specific fields or types can be modified post-parsing to correct known issues in the spec or parser output.

### 3. Code Generation (`go_generator.py`)
The `GoCodeGenerator` class iterates over the definitions and produces Go files.
*   **Type Mapping**:
    *   `INTEGER` -> `aper.Integer` (int64)
    *   `BIT STRING` -> `aper.BitString`
    *   `OCTET STRING` -> `aper.OctetString`
    *   `ENUMERATED` -> `aper.Enumerated` (int)
    *   `SEQUENCE` -> `struct`
    *   `CHOICE` -> `struct` with a `Present` field and a `Choice` interface/value.
*   **File Organization**:
    *   `e1ap_constants.go`: All `INTEGER` and `ID` constants.
    *   `e1ap_common_types.go`: Simple types (aliases to basic types).
    *   `[Type].go`: Complex types (`SEQUENCE`, `CHOICE`, `ENUMERATED`) get their own files.
    *   `[Type]Extensions.go`: Extension containers are generated in separate files for type safety.

### 4. Formatting
Finally, the script runs `goimports` and `gofmt` on the output directory to ensure the generated code adheres to Go standards.

## Templates

The generation logic uses string templates defined in:
*   **`go_templates.py`**: Defines the structure of the Go structs.
*   **`method_templates.py`**: Defines the `Encode` and `Decode` methods. These methods use the `aper` library to perform the actual serialization.

## Adding/Modifying Definitions

If you need to fix a bug in the generated code:
1.  **Check the ASN.1**: Is the input text correct?
2.  **Check the Parser**: Is `asn1_parser.py` misinterpreting the syntax?
3.  **Use a Patch**: If it's a one-off issue, add a patch in `asn1_patches.py` or `config/patches.yaml`.
4.  **Modify Templates**: If the Go code structure is wrong, edit `go_templates.py` or `method_templates.py`.
