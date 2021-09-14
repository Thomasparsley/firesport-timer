from datetime import datetime

import status as status


class Line:
    def __init__(self, time: datetime, status: status.Status):
        self.time = time
        self.status = status

    def set_default(self):
        self.time = datetime(1, 1, 1, 0, 0, 0, 0)
        self.status = status.get_status_by_id(0)

    def is_zero_time(self) -> bool:
        return self.time.hour == 0 and self.time.minute == 0 and self.time.second == 0 and self.time.microsecond == 0


def parse(rawTime: str, rawId: str) -> Line:
    line_time = datetime(1, 1, 1, 0, 0, 0, int(rawTime))
    line_status = status.parse_raw_status(rawId)

    return Line(line_time, line_status)
