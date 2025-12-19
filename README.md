# E1AP Library for Go

This repository implements the **E1 Application Protocol (E1AP)** as defined in **3GPP TS 38.463**. It provides a Go library for encoding and decoding E1AP messages using **Aligned Packet Encoding Rules (APER)**.

The library is primarily auto-generated from ASN.1 specifications to ensure strict compliance with the standard.

## Features

*   **Full E1AP Support**: Implements all messages and Information Elements (IEs) defined in TS 38.463.
*   **APER Encoding**: Efficient binary serialization using the `github.com/lvdund/ngap/aper` library.
*   **Type-Safe**: Generated Go structs with strong typing for all IEs.
*   **High Quality**: Code passes strict static analysis, uses optimized `Sequence[T]` generics for lists, and eliminates redundant object allocations.
*   **Verified**: Comprehensive test suite with "Round-Trip" verification and standard Hex output logging.
*   **Code Generation**: Includes the Python-based generator used to produce the library, allowing for easy updates to newer 3GPP versions.

## Documentation

Detailed documentation is available in the `docs/` directory:

*   [**Usage Guide**](docs/usage.md): How to create, encode, and decode messages.
*   [**Architecture**](docs/architecture.md): High-level design of the library and generator.
*   [**Protocol Implementation**](docs/protocols.md): Details on PDU structures and encoding rules.
*   [**Code Generator**](docs/generator.md): Explanation of the Python scripts used to generate the Go code.
*   [**Test Analysis**](docs/test_analysis.md): Guide to verifying the tests and interpreting the Hex output.

## Installation

```bash
go get github.com/sub2pewds12/E1AP
```

## Quick Start

### Encoding a Message

```go
package main

import (
    "fmt"
    "log"
    "github.com/sub2pewds12/E1AP"
    "github.com/sub2pewds12/E1AP/e1ap_ies"
    "github.com/lvdund/ngap/aper"
)

func main() {
    // Create a specific PDU struct directly
    // The library handles embedding it into the correct E1AP PDU container
    req := &e1ap_ies.GNBCUCPE1SetupRequest{
        TransactionID: e1ap_ies.TransactionID{Value: aper.Integer(123)},
        GNBCUCPName:   &e1ap_ies.GNBCUCPName{Value: aper.OctetString("my-gNB-CU-CP")},
    }

    // Encode
    // E1apEncode automatically handles the PDU header (ProcedureCode, Criticality, etc.)
    bytes, err := e1ap.E1apEncode(req)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Encoded Hex: %X\n", bytes)
}
```

## Verifying Correctness
To verify that the library complies with the E1AP specification, run the standardized test suite. These tests output the **Encoded Hex String** for manual verification against the standard.

```bash
go test -v ./...
```
See [**Test Analysis**](docs/test_analysis.md) for a guide on how to read the output.

## Project Structure

*   `e1ap_ies/`: Generated Go code for E1AP Information Elements.
*   `generator_scripts/`: Python scripts for ASN.1 parsing and code generation.
*   `config/`: Configuration for the generator (patches, skipped files).
*   `docs/`: Project documentation.

## Contributing

This library is generated. **Do not modify the files in `e1ap_ies/` directly.** Instead, modify the generator scripts or templates and regenerate the code. See [Generator Documentation](docs/generator.md) for details.
