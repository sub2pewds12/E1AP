package e1ap_ies

// MaxIPrate represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1307
type MaxIPrate int32

const (
	MaxIPrate_Bitrate64kbs MaxIPrate = 0
	MaxIPrate_MaxUErate    MaxIPrate = 1
)
