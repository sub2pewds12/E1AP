package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PDUSessionType From: 9_4_5_Information_Element_Definitions.txt:1899
const (
	PDUSessionTypeIpv4         aper.Enumerated = 0
	PDUSessionTypeIpv6         aper.Enumerated = 1
	PDUSessionTypeIpv4v6       aper.Enumerated = 2
	PDUSessionTypeEthernet     aper.Enumerated = 3
	PDUSessionTypeUnstructured aper.Enumerated = 4
)

type PDUSessionType struct {
	Value aper.Enumerated
}
