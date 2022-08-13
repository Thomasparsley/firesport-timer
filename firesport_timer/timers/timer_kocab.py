from typing import Final
""" import random """

import serial  # type: ignore

from .timer import Timer
from .line import Line

MINIMAL_LEN_OF_DATA: Final = 27
RESET_COMMAND: Final = "#RST"
READ_COMMAND: Final = "#APP:cw:data?"

"""
    2:0:1:0:1:0:1:0:1:0:0:0:0:1
	Neutral state

	2:0:4:6270:2:6270:2:0:1:0:0:0:0:2
	Left and right target running

	2:0:8:32010:8:29470:8:0:1:0:0:0:0:8
	Final
	The left target is the resulting time

	2:0:4:5390:8:18070:2:0:1:0:0:0:0:2
	Only the right target runs, the left target is knocked down

	2:300000:1:0:1:0:1:0:1:0:0:0:0:1
	Countdown
"""


class KocabTimer(Timer):
    serial = None

    def __init__(
        self,
        countdown: Line = Line(),
        line_one: Line = Line(),
        line_two: Line = Line(),
        line_three: Line = Line(),
        line_four: Line = Line(),
    ):
        self.countdown = countdown
        self.line_one = line_one
        self.line_two = line_two
        self.line_three = line_three
        self.line_four = line_four

    @classmethod
    def with_default_state(cls):
        return cls(
            countdown=Line(),
            line_one=Line(),
            line_two=Line(),
            line_three=Line(),
            line_four=Line(),
        )

    @classmethod
    def parse_raw_data(cls, raw_data: str):
        timer = cls()
        timer.apply_raw_data(raw_data)
        return timer

    def apply_raw_bytes_data(self, raw_data: bytes):
        self.apply_raw_data(str(raw_data))

    def apply_raw_data(self, raw_data: str):
        if len(raw_data) < MINIMAL_LEN_OF_DATA:
            return

        if raw_data[0] == ":":
            return

        if str(raw_data[0]).isalpha():
            return

        raw_data_split = str(raw_data).split(":")

        countdown = Line()
        line_one = Line.from_raw_data(raw_data_split[3], raw_data_split[4])
        line_two = Line.from_raw_data(raw_data_split[5], raw_data_split[6])
        line_three = Line.from_raw_data(raw_data_split[7], raw_data_split[8])
        line_four = Line.from_raw_data(raw_data_split[9], raw_data_split[10])

        if (
            not line_one.is_zero()
            or not line_two.is_zero()
            or not line_three.is_zero()
            or not line_four.is_zero()
        ):
            countdown = Line()
        else:
            countdown = Line.from_raw_countdown_data(raw_data_split[1])
            line_one = Line()
            line_two = Line()
            line_three = Line()
            line_four = Line()

        self.countdown = countdown
        self.line_one = line_one
        self.line_two = line_two
        self.line_three = line_three
        self.line_four = line_four

    def add_serial(self, port: str):
        self.serial = serial.Serial(
            port, 115200, timeout=0.1, bytesize=serial.EIGHTBITS
        )

    def send_reset(self):
        if self.serial and self.serial.is_open:
            command = RESET_COMMAND + "\n"
            self.serial.write(bytes(command, "utf-8"))  # type: ignore

    def send_read_data(self):
        if self.serial and self.serial.is_open:
            command = READ_COMMAND + "\n"
            self.serial.write(bytes(command, "utf-8"))  # type: ignore
            raw = self.serial.readline()
            self.apply_raw_bytes_data(raw)

    def update_timer(self):
        if self.serial and self.serial.is_open:
            self.send_read_data()
        else:
            """ choise = random.choice([
                "2:0:1:0:1:0:1:0:1:0:0:0:0:1",
                "2:0:4:6270:2:6270:2:0:1:0:0:0:0:2",
                "2:0:8:32010:8:29470:8:0:1:0:0:0:0:8",
                "2:0:4:5390:8:18070:2:0:1:0:0:0:0:2",
                "2:300000:1:0:1:0:1:0:1:0:0:0:0:1",
            ])
            self.apply_raw_data(choise) """
