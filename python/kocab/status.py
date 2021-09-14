from typing import Match


class Status:
    def __init__(self, id: int, name: str):
        self.id = id
        self.name = name


def get_status_by_id(id: int) -> Status:
    if id == 1:
        return Status(1, "default")
    elif id == 2:
        return Status(2, "run")
    elif id == 8:
        return Status(8, "stop")
    else:
        return Status(0, "undefined")


def parse_raw_status(rawID: str) -> Status:
    return get_status_by_id(int(rawID))
