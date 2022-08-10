from abc import ABC, abstractmethod


class Timer(ABC):
    name: str
    is_open: bool

    @abstractmethod
    def make_gui(self):
        pass
