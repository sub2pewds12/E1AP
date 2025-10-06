package e1ap_ies

import "github.com/lvdund/ngap/aper"

// ProtocolIEField is the manually defined type for the generic IE field.
// The Value is an OctetString, which holds raw bytes for later decoding.
type ProtocolIEField struct {
	Id          ProtocolIEID
	Criticality Criticality
	Value       aper.OctetString `aper:"openType"` // CHANGED from OpenType
}

// ProtocolExtensionContainer is a list of extension fields.
type ProtocolExtensionContainer []ProtocolExtensionField

// --- Protocol IE Single Container ---
// This is used for choice extensions. It's fundamentally a single IE field.
type ProtocolIESingleContainer struct {
	Value ProtocolIEField
}

// PrivateIEContainer is a list of private IE fields.
type PrivateIEContainer []PrivateIEField
