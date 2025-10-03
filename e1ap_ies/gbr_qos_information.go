package e1ap_ies

import "github.com/lvdund/ngap/aper"

// GBRQosInformation From: 9_4_5_Information_Element_Definitions.txt:1165
// ASN.1 Data Type: SEQUENCE
type GBRQosInformation struct {
	ERABMaximumBitrateDL    aper.Integer `aper:"mandatory,reject,ext"`
	ERABMaximumBitrateUL    aper.Integer `aper:"mandatory,reject,ext"`
	ERABGuaranteedBitrateDL aper.Integer `aper:"mandatory,reject,ext"`
	ERABGuaranteedBitrateUL aper.Integer `aper:"mandatory,reject,ext"`
}
