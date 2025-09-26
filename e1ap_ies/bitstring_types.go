package e1ap_ies

// AdditionalRRMPriorityIndex ::= BIT STRING
type Additionalrrmpriorityindex struct {
	Value []byte `asn1:"value"` // The byte array representing the bit string
	Len   uint   `asn1:"len"`   // The number of bits in the bit string
}
const (
	AdditionalrrmpriorityindexMinBits = 32
	AdditionalrrmpriorityindexMaxBits = 32
)

// DRBBStatusTransfer-receiveStatusofPDCPSDU ::= BIT STRING
type DrbbstatustransferReceivestatusofpdcpsdu struct {
	Value []byte `asn1:"value"` // The byte array representing the bit string
	Len   uint   `asn1:"len"`   // The number of bits in the bit string
}
const (
	DrbbstatustransferReceivestatusofpdcpsduMinBits = 1
	DrbbstatustransferReceivestatusofpdcpsduMaxBits = 131072
)

// InterfacesToTrace ::= BIT STRING
type Interfacestotrace struct {
	Value []byte `asn1:"value"` // The byte array representing the bit string
	Len   uint   `asn1:"len"`   // The number of bits in the bit string
}
const (
	InterfacestotraceMinBits = 8
	InterfacestotraceMaxBits = 8
)

// MeasurementsToActivate ::= BIT STRING
type Measurementstoactivate struct {
	Value []byte `asn1:"value"` // The byte array representing the bit string
	Len   uint   `asn1:"len"`   // The number of bits in the bit string
}
const (
	MeasurementstoactivateMinBits = 8
	MeasurementstoactivateMaxBits = 8
)

// NID ::= BIT STRING
type Nid struct {
	Value []byte `asn1:"value"` // The byte array representing the bit string
	Len   uint   `asn1:"len"`   // The number of bits in the bit string
}
const (
	NidMinBits = 44
	NidMaxBits = 44
)

// NR-Cell-Identity ::= BIT STRING
type NrCellIdentity struct {
	Value []byte `asn1:"value"` // The byte array representing the bit string
	Len   uint   `asn1:"len"`   // The number of bits in the bit string
}
const (
	NrCellIdentityMinBits = 36
	NrCellIdentityMaxBits = 36
)

// PortNumber ::= BIT STRING
type Portnumber struct {
	Value []byte `asn1:"value"` // The byte array representing the bit string
	Len   uint   `asn1:"len"`   // The number of bits in the bit string
}
const (
	PortnumberMinBits = 16
	PortnumberMaxBits = 16
)

// QoS-Mapping-Information-dscp ::= BIT STRING
type QosMappingInformationDscp struct {
	Value []byte `asn1:"value"` // The byte array representing the bit string
	Len   uint   `asn1:"len"`   // The number of bits in the bit string
}
const (
	QosMappingInformationDscpMinBits = 6
	QosMappingInformationDscpMaxBits = 6
)

// QoS-Mapping-Information-flow-label ::= BIT STRING
type QosMappingInformationFlowLabel struct {
	Value []byte `asn1:"value"` // The byte array representing the bit string
	Len   uint   `asn1:"len"`   // The number of bits in the bit string
}
const (
	QosMappingInformationFlowLabelMinBits = 20
	QosMappingInformationFlowLabelMaxBits = 20
)

// ReportCharacteristics ::= BIT STRING
type Reportcharacteristics struct {
	Value []byte `asn1:"value"` // The byte array representing the bit string
	Len   uint   `asn1:"len"`   // The number of bits in the bit string
}
const (
	ReportcharacteristicsMinBits = 36
	ReportcharacteristicsMaxBits = 36
)

// TransportLayerAddress ::= BIT STRING
type Transportlayeraddress struct {
	Value []byte `asn1:"value"` // The byte array representing the bit string
	Len   uint   `asn1:"len"`   // The number of bits in the bit string
}

