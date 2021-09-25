UNDEFINED_ID = 0
UNDEFINED_NAME = "undefined"
DEFAULT_ID = 1
DEFAULT_NAME = "default"
RUN_ID = 2
RUN_NAME = "run"
STOP_ID = 8
STOP_NAME = "stop"


class Status:
    def __init__(self, id: int, name: str):
        self.id = id
        self.name = name

    def __str__(self):
        return self.name


def get_status_by_id(id: int) -> Status:
    if id == 1:
        return Status(DEFAULT_ID, DEFAULT_NAME)
    elif id == 2:
        return Status(RUN_ID, RUN_NAME)
    elif id == 8:
        return Status(STOP_ID, STOP_NAME)
    else:
        return Status(UNDEFINED_ID, UNDEFINED_NAME)


def parse_raw_status(raw_id: str) -> Status:
    return get_status_by_id(int(raw_id))
