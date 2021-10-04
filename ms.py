MS = 1
MS_LIMIT = 1000
SEC = MS_LIMIT
MIN = 60 * SEC


class Ms:
    value: int

    def __init__(self, ms: int) -> None:
        self.value = ms

    def __eq__(self, other) -> bool:
        if isinstance(other, Ms):
            return self.value == other.value
        else:
            return False

    def __lt__(self, other) -> bool:
        if isinstance(other, Ms):
            return self.value < other.value
        else:
            return False

    def __str__(self) -> str:
        return self.format()

    def __add__(self, other):
        if isinstance(other, Ms):
            return Ms(self.value + other.value)
        elif isinstance(other, int):
            return Ms(self.value + other)
        else:
            return NotImplemented

    def __sub__(self, other):
        if isinstance(other, Ms):
            return Ms(self.value - other.value)
        elif isinstance(other, int):
            return Ms(self.value - other)
        else:
            return NotImplemented

    def __mul__(self, other):
        if isinstance(other, Ms):
            return Ms(self.value * other.value)
        elif isinstance(other, int):
            return Ms(self.value * other)
        else:
            return NotImplemented

    def format(self) -> str:
        """Format the milliseconds to a string (e.g. 1:04.090)."""

        if self.value < MS:
            return "0.000"

        ms = self.value % MS_LIMIT
        sec = (self.value // SEC) % 60
        min = (self.value // MIN) % 60

        if min > 0:
            return f"{min}:{sec:02}.{ms:03}"
        else:
            return f"{sec}.{ms:03}"


def str_to_ms(time_str: str) -> Ms:
    """Transform a string (e. g. 1:04.090) to a Ms class."""

    result = 0

    if "." not in time_str:
        return Ms(int(time_str))

    if ":" in time_str:
        min, sec_ms = time_str.split(":")
        result += int(min) * MIN

    try:
        sec, ms = sec_ms.split(".")
    except UnboundLocalError:
        sec, ms = time_str.split(".")

    result += int(sec) * SEC
    result += int(ms) % MS_LIMIT

    return Ms(result)