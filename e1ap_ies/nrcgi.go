package e1ap_ies

import "github.com/lvdund/ngap/aper"

// NRCGI From: 9_4_5_Information_Element_Definitions.txt:1492
// ASN.1 Data Type: SEQUENCE
type NRCGI struct {
	PLMNIdentity   aper.OctetString `aper:"lb:3,ub:3,mandatory,ignore"`
	NRCellIdentity aper.BitString   `aper:"lb:36,ub:36,mandatory"`
}
