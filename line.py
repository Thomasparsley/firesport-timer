import ms
import status


class Line:
    def __init__(self, time: ms.Ms, status: status.Status):
        self.time = time
        self.status = status

    def __str__(self):
        return str(self.time) + " " + str(self.status)

    def set_default(self):
        self.time = ms.Ms(0)
        self.status = status.get_status_by_id(status.DEFAULT_ID)

    def is_zero(self) -> bool:
        return self.time.value == 0

    def format_time(self) -> str:
        """
        Formats a time in milliseconds into a 1:02.03 format
        """
        ms = int(str(self.time.microsecond)[:2])
        return "{}:{:02d}.{:02d}".format(self.time.minute, self.time.second, ms)


def new() -> Line:
    return Line(ms.Ms(0), status.get_status_by_id(status.DEFAULT_ID))


def parse(raw_time: str, raw_id: str) -> Line:
    milisecons = (int(raw_time) % 1000) * 1000
    seconds = (int(raw_time) // 1000) % 60
    minutes = (int(raw_time) // 1000 // 60) % 60
    hours = (int(raw_time) // 1000 // 60 // 60) % 24

    line_time = datetime(1, 1, 1, hours, minutes, seconds, milisecons)
    line_status = status.parse_raw_status(raw_id)

    return Line(line_time, line_status)


def parse_countdown(raw_time: str) -> Line:
    milisecons = (int(raw_time) % 1000) * 1000
    seconds = (int(raw_time) // 1000) % 60
    minutes = (int(raw_time) // 1000 // 60) % 60
    hours = (int(raw_time) // 1000 // 60 // 60) % 24

    line_time = datetime(1, 1, 1, hours, minutes, seconds, milisecons)
    result = Line(line_time, status.get_status_by_id(status.UNDEFINED_ID))

    if result.is_zero():
        result.status = status.get_status_by_id(status.STOP_ID)
    else:
        result.status = status.get_status_by_id(status.RUN_ID)

    return result
