from typing import Final, Any

UNDEFINED_ID: Final = 0
UNDEFINED_NAME: Final = "undefined"
DEFAULT_ID: Final = 1
DEFAULT_NAME: Final = "default"
RUN_ID: Final = 2
RUN_NAME: Final = "run"
STOP_ID: Final = 8
STOP_NAME: Final = "stop"


class Status:
    def __init__(self, id: int, name: str):
        self.id = id
        self.name = name

    def __str__(self):
        return self.name

    def __eq__(self, other: Any) -> bool:
        if not isinstance(other, Status):
            return False

        return self.id == other.id

    @classmethod
    def from_id(cls, id: int) -> "Status":
        if id == DEFAULT_ID:
            return cls(DEFAULT_ID, DEFAULT_NAME)
        elif id == RUN_ID:
            return cls(RUN_ID, RUN_NAME)
        elif id == STOP_ID:
            return cls(STOP_ID, STOP_NAME)
        else:
            return cls(UNDEFINED_ID, UNDEFINED_NAME)

    @classmethod
    def from_str_id(cls, raw_id: str) -> "Status":
        return cls.from_id(int(raw_id))
