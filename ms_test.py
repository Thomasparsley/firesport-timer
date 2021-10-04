import unittest

from ms import *


class TestMs(unittest.TestCase):
    def test_lt(self):
        self.assertTrue(Ms(0) < Ms(1))
        self.assertFalse(Ms(0) > Ms(1))

    def test_add(self):
        self.assertEqual(Ms(0) + Ms(1), Ms(1))
        self.assertEqual(Ms(0) + 1, Ms(1))

    def test_sub(self):
        self.assertEqual(Ms(5) - Ms(1), Ms(4))
        self.assertEqual(Ms(5) - 1, Ms(4))

    def test_mul(self):
        self.assertEqual(Ms(0) * Ms(2), Ms(0))
        self.assertEqual(Ms(0) * 2, Ms(0))
        self.assertEqual(Ms(0) * 0, Ms(0))
        self.assertEqual(Ms(2) * Ms(2), Ms(4))
        self.assertEqual(Ms(2) * 2, Ms(4))

    def test_format(self):
        test = [
            (Ms(0), "0.000"),
            (Ms(6), "0.006"),
            (Ms(60), "0.060"),
            (Ms(600), "0.600"),
            (Ms(6000), "6.000"),
            (Ms(14000), "14.000"),
            (Ms(48248), "48.248"),
            (Ms(60000), "1:00.000"),
            (Ms(120000), "2:00.000"),
            (Ms(144567), "2:24.567"),
        ]

        for input, expected in test:
            self.assertEqual(input.format(), expected)

    def test_str_to_ms(self):
        test = [
            ("0", Ms(0)),
            ("0.000", Ms(0)),
            ("6", Ms(6)),
            ("249", Ms(249)),
            ("6.000", Ms(6000)),
            ("14.000", Ms(14000)),
            ("48.248", Ms(48248)),
            ("1:00.000", Ms(60000)),
            ("2:00.000", Ms(120000)),
            ("2:24.567", Ms(144567)),
        ]

        for input, expected in test:
            self.assertEqual(str_to_ms(input), expected)
