package e1ap_ies

// CauseProtocol represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:162
type CauseProtocol int32

const (
	CauseProtocol_TransferSyntaxError                          CauseProtocol = 0
	CauseProtocol_AbstractSyntaxErrorReject                    CauseProtocol = 1
	CauseProtocol_AbstractSyntaxErrorIgnoreAndNotify           CauseProtocol = 2
	CauseProtocol_MessageNotCompatibleWithReceiverState        CauseProtocol = 3
	CauseProtocol_SemanticError                                CauseProtocol = 4
	CauseProtocol_AbstractSyntaxErrorFalselyConstructedMessage CauseProtocol = 5
	CauseProtocol_Unspecified                                  CauseProtocol = 6
)
