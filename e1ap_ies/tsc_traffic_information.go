package e1ap_ies

// TSCTrafficInformation From: 9_4_5_Information_Element_Definitions.txt:2256
type TSCTrafficInformation struct {
	Periodicity      int64   `asn1:"mandatory"`
	BurstArrivalTime *[]byte `asn1:"optional"`
}
