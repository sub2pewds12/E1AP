package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PDCPCount is a generated SEQUENCE type.
type PDCPCount struct {
	PDCPSN aper.Integer `aper:"lb:0,ub:262143,mandatory,ext"`
	HFN    aper.Integer `aper:"lb:0,ub:4294967295,mandatory,ext"`
}
