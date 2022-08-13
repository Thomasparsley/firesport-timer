from .ms import Ms
from .status import Status, DEFAULT_ID, UNDEFINED_ID, STOP_ID, RUN_ID


class Line:
    __slots__ = ["time", "status"]

    def __init__(
        self,
        time: Ms = Ms(),
        status: Status = Status.from_id(DEFAULT_ID),
    ):
        self.time = time
        self.status = status

    def __str__(self):
        return str(self.time)

    def __eq__(self, o: object) -> bool:
        if not isinstance(o, Line):
            return False

        return self.time == o.time and self.status == o.status

    @classmethod
    def from_raw_data(cls, raw_time: str, raw_id: str):
        return cls(
            Ms.from_str(raw_time),
            Status.from_str_id(raw_id),
        )

    @classmethod
    def from_raw_countdown_data(cls, raw_time: str):
        line = cls(
            Ms.from_str(raw_time),
            Status.from_id(UNDEFINED_ID),
        )

        if line.is_zero():
            line.status = Status.from_id(STOP_ID)
        else:
            line.status = Status.from_id(RUN_ID)

        return line

    def is_zero(self) -> bool:
        return self.time.is_zero()
