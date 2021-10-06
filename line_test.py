import unittest

import status
import line
import ms


class TestLine(unittest.TestCase):
    def test_new(self):
        self.assertEqual(line.new(), line.Line(
            ms.new(), status.get_status_by_id(status.DEFAULT_ID)))
