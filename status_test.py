import unittest

from status import *


class TestStatsu(unittest.TestCase):
    def test_get_status_by_id(self):
        test = [
            (UNDEFINED_ID, Status(UNDEFINED_ID, UNDEFINED_NAME)),
            (DEFAULT_ID, Status(DEFAULT_ID, DEFAULT_NAME)),
            (RUN_ID, Status(RUN_ID, RUN_NAME)),
            (STOP_ID, Status(STOP_ID, STOP_NAME)),
        ]

        for input, expected in test:
            self.assertEqual(get_status_by_id(input), expected)
