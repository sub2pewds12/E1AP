from typing import Dict, List, Optional, Any

class ASN1Definition:
    def __init__(self, name: str, def_type: str, source_file: str = "", source_line: int = -1, full_text: str = ""):
        self.name: str = name
        self.type: str = def_type
        self.source_file: str = source_file
        self.source_line: int = source_line
        self.full_text: str = full_text
        self.min_val: Optional[str] = None
        self.max_val: Optional[str] = None
        self.enum_values: List[str] = []
        self.ies: List[Dict[str, Any]] = []
        self.is_constant: bool = False
        self.is_extensible: bool = False
        self.alias_of: Optional[str] = None
        self.of_type: Optional[str] = None
        self.is_builtin: bool = False
        self.is_synthetic: bool = False