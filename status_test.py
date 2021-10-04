import unittest

from status import *


class TestStatsu(unittest.TestCase):
    def test_get_status_by_id(self):
        tests = [
            (UNDEFINED_ID, Status(UNDEFINED_ID, UNDEFINED_NAME)),
            (DEFAULT_ID, Status(DEFAULT_ID, DEFAULT_NAME)),
            (RUN_ID, Status(RUN_ID, RUN_NAME)),
            (STOP_ID, Status(STOP_ID, STOP_NAME)),
        ]

        for input, expected in tests:
            self.assertEqual(get_status_by_id(input), expected)

    def test_parse_raw_status(self):
        tests = [
            (str(UNDEFINED_ID), Status(UNDEFINED_ID, UNDEFINED_NAME)),
            (str(DEFAULT_ID), Status(DEFAULT_ID, DEFAULT_NAME)),
            (str(RUN_ID), Status(RUN_ID, RUN_NAME)),
            (str(STOP_ID), Status(STOP_ID, STOP_NAME)),
        ]

        for input, expected in tests:
            self.assertEqual(parse_raw_status(input), expected)
