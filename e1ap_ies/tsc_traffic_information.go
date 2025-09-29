package e1ap_ies

// TSCTrafficInformation represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2256
type TSCTrafficInformation struct {
	Periodicity      int64   `asn1:"lb:1,ub:640000,mandatory"`
	BurstArrivalTime *[]byte `asn1:"optional"`
}
