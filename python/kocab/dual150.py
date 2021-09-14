from . import line

RESET_COMMAND = "#RST"
READ_COMMAND = "#APP:cw:data?"

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


class Dual150:
    def __init__(self, countdown: line.Line, line_one: line.Line, line_two: line.Line, line_three: line.Line, line_four: line.Line):
        self.countdown = countdown
        self.line_one = line_one
        self.line_two = line_two
        self.line_three = line_three
        self.line_four = line_four

    def __str__(self) -> str:
        return """Dual150 (
        countdown={};
        line_one={};
        line_two={};
        line_three={};
        line_four={}
        )""".format(
            self.countdown, self.line_one, self.line_two, self.line_three, self.line_four
        )


def new() -> Dual150:
    return Dual150(countdown=line.Line().set_default(),
                   line_one=line.Line().set_default(),
                   line_two=line.Line().set_default(),
                   line_three=line.Line().set_default(),
                   line_four=line.Line().set_default())


def parse_raw_data(raw_data: str) -> Dual150:
    """
    Parses the raw data from the dual150 into a Dual150 object
    :param raw_data: The raw data from the dual150
    :return: A Dual150 object with the data
    """
    # if len of raw_data is less then 27, then throw an error
    if len(raw_data) < 27:
        raise Exception("The raw_data is too short")

    # if the first character is a ':', then throw an error.
    # raw_data cant start with a ':'
    if raw_data[0] == ':':
        raise Exception("The raw_data starts with a ':'")

    # if the first character is not a a-Z, then throw an error
    if str(raw_data[0]).isalpha():
        raise Exception("The raw_data does not start with a a-Z")

    raw_data_split = str(raw_data).split(":")

    countdown = line.new()
    line_one = line.parse(raw_data_split[3], raw_data_split[4])
    line_two = line.parse(raw_data_split[5], raw_data_split[6])
    line_three = line.parse(raw_data_split[7], raw_data_split[8])
    line_four = line.parse(raw_data_split[9], raw_data_split[10])

    if not line_one.is_zero() or not line_two.is_zero() or not line_three.is_zero() or not line_four.is_zero():
        countdown.set_default()
    else:
        countdown = line.parse_countdown(raw_data_split[1])

        line_one.set_default()
        line_two.set_default()
        line_three.set_default()
        line_four.set_default()

    return Dual150(countdown=countdown,
                   line_one=line_one,
                   line_two=line_two,
                   line_three=line_three,
                   line_four=line_four)
