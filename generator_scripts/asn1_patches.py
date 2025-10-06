from common_types import (
    IntegerDefinition,
    EnumDefinition,
    SequenceDefinition,
    ListDefinition,
    InformationElement,
)
from typing import List


class ASN1Patcher:
    def get_hardcoded_definitions(self) -> List:
        definitions = []

        diag_item = SequenceDefinition(
            name="CriticalityDiagnostics-IE-Item",
            source_file="manual_patch",
            full_text="MANUALLY CREATED DEFINITION FOR CriticalityDiagnostics-IE-Item",
        )

        diag_item.ies.extend(
            [
                InformationElement(
                    ie="iECriticality",
                    type="Criticality",
                    presence="mandatory",
                    criticality="ignore",
                ),
                InformationElement(
                    ie="iE-ID",
                    type="ProtocolIE-ID",
                    presence="mandatory",
                    criticality="ignore",
                ),
                InformationElement(
                    ie="typeOfError",
                    type="TypeOfError",
                    presence="mandatory",
                    criticality="ignore",
                ),
            ]
        )
        definitions.append(diag_item)

        diag_list = ListDefinition(
            name="CriticalityDiagnostics-IE-List",
            source_file="manual_patch",
            full_text="MANUALLY CREATED DEFINITION FOR CriticalityDiagnostics-IE-List",
        )
        diag_list.of_type = "CriticalityDiagnostics-IE-Item"
        definitions.append(diag_list)

        return definitions
