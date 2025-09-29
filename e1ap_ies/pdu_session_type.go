package e1ap_ies

// PDUSessionType represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1899
type PDUSessionType int32

const (
	PDUSessionType_Ipv4         PDUSessionType = 0
	PDUSessionType_Ipv6         PDUSessionType = 1
	PDUSessionType_Ipv4v6       PDUSessionType = 2
	PDUSessionType_Ethernet     PDUSessionType = 3
	PDUSessionType_Unstructured PDUSessionType = 4
)
