package e1ap_ies

// PDUSessionResourceToSetupItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1836
type PDUSessionResourceToSetupItem struct {
	PDUSessionID                               int64                             `asn1:"lb:0,ub:255,mandatory,ext"`
	PDUSessionType                             PDUSessionType                    `asn1:"mandatory,ext"`
	SNSSAI                                     SNSSAI                            `asn1:"mandatory,ext"`
	SecurityIndication                         SecurityIndication                `asn1:"mandatory,ext"`
	PDUSessionResourceDLAMBR                   *int64                            `asn1:"lb:0,ub:4000000000000,optional,ext"`
	NGULUPTNLInformation                       UPTNLInformation                  `asn1:"mandatory,ext"`
	PDUSessionDataForwardingInformationRequest *DataForwardingInformationRequest `asn1:"optional,ext"`
	PDUSessionInactivityTimer                  *int64                            `asn1:"lb:1,ub:7200,optional,ext"`
	ExistingAllocatedNGDLUPTNLInfo             *UPTNLInformation                 `asn1:"optional,ext"`
	NetworkInstance                            *int64                            `asn1:"lb:1,ub:256,optional,ext"`
	DRBToSetupListNGRAN                        []DRBToSetupItemNGRAN             `asn1:"lb:1,ub:Item,mandatory,ext"`
}
