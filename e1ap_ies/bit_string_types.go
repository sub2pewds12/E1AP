package e1ap_ies

// AdditionalRRMPriorityIndex represents the ASN.1 BIT STRING type.
type AdditionalRRMPriorityIndex []byte

const (
	AdditionalRRMPriorityIndexMinSize int = 32
	AdditionalRRMPriorityIndexMaxSize int = 32
)


// Dscp represents the ASN.1 BIT STRING type.
type Dscp []byte

const (
	DscpMinSize int = 6
	DscpMaxSize int = 6
)


// FlowLabel represents the ASN.1 BIT STRING type.
type FlowLabel []byte

const (
	FlowLabelMinSize int = 20
	FlowLabelMaxSize int = 20
)


// InterfacesToTrace represents the ASN.1 BIT STRING type.
type InterfacesToTrace []byte

const (
	InterfacesToTraceMinSize int = 8
	InterfacesToTraceMaxSize int = 8
)


// MeasurementsToActivate represents the ASN.1 BIT STRING type.
type MeasurementsToActivate []byte

const (
	MeasurementsToActivateMinSize int = 8
	MeasurementsToActivateMaxSize int = 8
)


// Nid represents the ASN.1 BIT STRING type.
type Nid []byte

const (
	NidMinSize int = 44
	NidMaxSize int = 44
)


// NrCellIdentity represents the ASN.1 BIT STRING type.
type NrCellIdentity []byte

const (
	NrCellIdentityMinSize int = 36
	NrCellIdentityMaxSize int = 36
)


// PortNumber represents the ASN.1 BIT STRING type.
type PortNumber []byte

const (
	PortNumberMinSize int = 16
	PortNumberMaxSize int = 16
)


// ReceiveStatusofPDCPSDU represents the ASN.1 BIT STRING type.
type ReceiveStatusofPDCPSDU []byte

const (
	ReceiveStatusofPDCPSDUMinSize int = 1
	ReceiveStatusofPDCPSDUMaxSize int = 131072
)


// ReportCharacteristics represents the ASN.1 BIT STRING type.
type ReportCharacteristics []byte

const (
	ReportCharacteristicsMinSize int = 36
	ReportCharacteristicsMaxSize int = 36
)


// TransportLayerAddress represents the ASN.1 BIT STRING type.
type TransportLayerAddress []byte

const (
	TransportLayerAddressMinSize int = 1
	TransportLayerAddressMaxSize int = 160
)


