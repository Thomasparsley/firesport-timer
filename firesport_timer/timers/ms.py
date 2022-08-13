from typing import Final, Any

MS: Final = 1
MS_LIMIT: Final = MS * 1000
SEC: Final = MS_LIMIT
SEC_LIMIT: Final = 60 * SEC
MIN: Final = SEC_LIMIT

class Ms:
    __slots__ = ["value"]

    def __init__(self, ms: int = 0):
        self.value = ms

    def __eq__(self, other: Any) -> bool:
        if not isinstance(other, Ms):
            return False
        
        return self.value == other.value

    def __lt__(self, other: Any) -> bool:
        if not isinstance(other, Ms):
            return False
        
        return self.value < other.value

    def __gt__(self, other: Any) -> bool:
        if not isinstance(other, Ms):
            return False
        
        return self.value > other.value

    def __add__(self, other: Any):
        if isinstance(other, Ms):
            return Ms(self.value + other.value)
        elif isinstance(other, int):
            return Ms(self.value + other)
        else:
            raise NotImplementedError

    def __sub__(self, other: Any):
        if isinstance(other, Ms):
            return Ms(self.value - other.value)
        elif isinstance(other, int):
            return Ms(self.value - other)
        else:
            raise NotImplementedError

    def __mul__(self, other: Any):
        if isinstance(other, Ms):
            return Ms(self.value * other.value)
        elif isinstance(other, int):
            return Ms(self.value * other)
        else:
            raise NotImplementedError

    def __str__(self) -> str:
        min = (self.value // MIN) % 60
        sec = (self.value // SEC) % 60
        ms = self.value % MS_LIMIT

        return f"{min}:{sec:02}.{ms:03}"

    @classmethod
    def new_from_sec(cls, value: int) -> "Ms":
        return cls(value * SEC)

    @classmethod
    def new_from_min(cls, value: int) -> "Ms":
        return cls(value * MIN)

    @classmethod
    def from_str(cls, input: str) -> "Ms":
        return Ms(int(input))

    def is_zero(self) -> bool:
        return self.value == 0