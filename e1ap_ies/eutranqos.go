package e1ap_ies

// EUTRANQOS represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1053
type EUTRANQOS struct {
	QCI                                  int64                                `asn1:"lb:0,ub:255,mandatory,ext"`
	EUTRANallocationAndRetentionPriority EUTRANAllocationAndRetentionPriority `asn1:"mandatory,ext"`
	GbrQosInformation                    *GBRQosInformation                   `asn1:"optional,ext"`
}
