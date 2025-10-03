package e1ap_ies

import "github.com/lvdund/ngap/aper"

// TSCTrafficInformation From: 9_4_5_Information_Element_Definitions.txt:2256
// ASN.1 Data Type: SEQUENCE
type TSCTrafficInformation struct {
	Periodicity      aper.Integer      `aper:"mandatory"`
	BurstArrivalTime *aper.OctetString `aper:"optional"`
}
