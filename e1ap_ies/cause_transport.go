package e1ap_ies

// CauseTransport represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:207
type CauseTransport int32

const (
	CauseTransport_Unspecified                  CauseTransport = 0
	CauseTransport_TransportResourceUnavailable CauseTransport = 1
	CauseTransport_UnknownTNLAddressForIAB      CauseTransport = 2
)
