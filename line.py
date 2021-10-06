import ms
import status


class Line:
    def __init__(self, time: ms.Ms, status: status.Status):
        self.time = time
        self.status = status

    def __str__(self):
        return str(self.time) + " " + str(self.status)

    def __eq__(self, o: object) -> bool:
        if isinstance(o, Line):
            return self.time == o.time and self.status == o.status

    def set_default(self):
        self.time = ms.new()
        self.status = status.get_status_by_id(status.DEFAULT_ID)

    def is_zero(self) -> bool:
        return self.time.is_zero()

    def format_time(self) -> str:
        """
        Formats a time in milliseconds into a 1:02.030 format
        """
        return self.time.format()


def new() -> Line:
    return Line(ms.new(), status.get_status_by_id(status.DEFAULT_ID))


def parse(raw_time: str, raw_id: str) -> Line:
    line_time = ms.new(int(raw_time))
    line_status = status.parse_raw_status(raw_id)

    return Line(line_time, line_status)


def parse_countdown(raw_time: str) -> Line:
    line_time = ms.new(int(raw_time))
    result = Line(line_time, status.get_status_by_id(status.UNDEFINED_ID))

    if result.is_zero():
        result.status = status.get_status_by_id(status.STOP_ID)
    else:
        result.status = status.get_status_by_id(status.RUN_ID)

    return result
