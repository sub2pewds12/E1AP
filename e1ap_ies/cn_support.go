package e1ap_ies

import "github.com/lvdund/ngap/aper"

// CNSupport From: 9_4_5_Information_Element_Definitions.txt:242
const (
	CNSupportCEpc aper.Enumerated = 0
	CNSupportC5gc aper.Enumerated = 1
	CNSupportBoth aper.Enumerated = 2
)

type CNSupport struct {
	Value aper.Enumerated
}
