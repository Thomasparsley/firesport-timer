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

    def __eq__(self, other) -> bool:
        if isinstance(other, Status):
            return self.id == other.id
        else:
            return False


def get_status_by_id(id: int) -> Status:
    if id == DEFAULT_ID:
        return Status(DEFAULT_ID, DEFAULT_NAME)
    elif id == RUN_ID:
        return Status(RUN_ID, RUN_NAME)
    elif id == STOP_ID:
        return Status(STOP_ID, STOP_NAME)
    else:
        return Status(UNDEFINED_ID, UNDEFINED_NAME)


def parse_raw_status(raw_id: str) -> Status:
    return get_status_by_id(int(raw_id))
