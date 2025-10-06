from typing import Dict, List, Optional, Any
from dataclasses import dataclass, field


class ASN1Definition:
    def __init__(
        self,
        name: str,
        source_file: str = "",
        source_line: int = -1,
        full_text: str = "",
    ):
        self.name: str = name
        self.source_file: str = source_file
        self.source_line: int = source_line
        self.full_text: str = full_text
        self.is_synthetic: bool = False


@dataclass
class InformationElement:
    ie: str
    type: str
    presence: str
    criticality: Optional[str] = None
    id: Optional[str] = None


class IntegerDefinition(ASN1Definition):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.min_val: Optional[str] = None
        self.max_val: Optional[str] = None


class ConstantDefinition(IntegerDefinition):
    def __init__(self, const_type: str, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.const_type: str = const_type


class EnumDefinition(ASN1Definition):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.enum_values: List[str] = []
        self.is_extensible: bool = False


class StringDefinition(ASN1Definition):
    def __init__(self, string_type: str, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.string_type: str = string_type
        self.min_val: Optional[str] = None
        self.max_val: Optional[str] = None


class SequenceDefinition(ASN1Definition):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.ies: List[InformationElement] = []
        self.is_extensible: bool = False


class ChoiceDefinition(ASN1Definition):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.ies: List[InformationElement] = []
        self.is_extensible: bool = False


class ListDefinition(ASN1Definition):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.of_type: Optional[str] = None
        self.min_val: Optional[str] = None
        self.max_val: Optional[str] = None


class AliasDefinition(ASN1Definition):
    def __init__(self, alias_of: str, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.alias_of: str = alias_of
        self.min_val: Optional[str] = None
        self.max_val: Optional[str] = None


class BuiltinDefinition(ASN1Definition):
    pass


class ASN1Procedure:
    """A simple container for a parsed E1AP Elementary Procedure."""

    def __init__(self, name: str):
        self.name: str = name
        self.initiating_message: Optional[str] = None
        self.successful_outcome: Optional[str] = None
        self.unsuccessful_outcome: Optional[str] = None
        self.procedure_code: Optional[str] = None


@dataclass
class ASN1ClassField:
    """Represents a single field within an ASN.1 CLASS definition, like '&id'."""

    name: str
    type: str
    is_optional: bool = False


class ASN1Class(ASN1Definition):
    """Represents a parsed ASN.1 CLASS definition itself."""

    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.fields: Dict[str, ASN1ClassField] = {}
