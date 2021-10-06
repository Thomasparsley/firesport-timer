import unittest

import ms


class TestMs(unittest.TestCase):
    def test_new(self):
        tests = [
            (ms.new(), ms.Ms(0)),
            (ms.new(1), ms.Ms(1))
        ]

        for input, expected in tests:
            self.assertEqual(input, expected)

    def test_lt(self):
        self.assertTrue(ms.new() < ms.new(1))
        self.assertFalse(ms.new() > ms.new(1))

    def test_add(self):
        self.assertEqual(ms.new() + ms.new(1), ms.new(1))
        self.assertEqual(ms.new() + 1, ms.new(1))

    def test_sub(self):
        self.assertEqual(ms.new(5) - ms.new(1), ms.new(4))
        self.assertEqual(ms.new(5) - 1, ms.new(4))

    def test_mul(self):
        self.assertEqual(ms.new() * ms.new(2), ms.new())
        self.assertEqual(ms.new() * 2, ms.new())
        self.assertEqual(ms.new() * 0, ms.new())
        self.assertEqual(ms.new(2) * ms.new(2), ms.new(4))
        self.assertEqual(ms.new(2) * 2, ms.new(4))

    def test_format(self):
        tests = [
            (ms.new(0), "0.000"),
            (ms.new(6), "0.006"),
            (ms.new(60), "0.060"),
            (ms.new(600), "0.600"),
            (ms.new(6000), "6.000"),
            (ms.new(14000), "14.000"),
            (ms.new(48248), "48.248"),
            (ms.new(60000), "1:00.000"),
            (ms.new(120000), "2:00.000"),
            (ms.new(144567), "2:24.567"),
        ]

        for input, expected in tests:
            self.assertEqual(input.format(), expected)

    def test_str_to_ms(self):
        tests = [
            ("0", ms.new(0)),
            ("0.000", ms.new(0)),
            ("6", ms.new(6)),
            ("249", ms.new(249)),
            ("6.000", ms.new(6000)),
            ("14.000", ms.new(14000)),
            ("48.248", ms.new(48248)),
            ("1:00.000", ms.new(60000)),
            ("2:00.000", ms.new(120000)),
            ("2:24.567", ms.new(144567)),
        ]

        for input, expected in tests:
            self.assertEqual(ms.str_to_ms(input), expected)

    def test_is_zero(self):
        tests = [
            (ms.new(0), True),
            (ms.new(1), False)
        ]

        for input, expected in tests:
            if isinstance(input, ms.Ms):
                self.assertEqual(input.is_zero(), expected)
