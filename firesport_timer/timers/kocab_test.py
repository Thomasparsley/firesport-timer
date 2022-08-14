from .kocab_dual150 import Dual150Timer


def test_apply_raw_bytes_data():
    input = bytes("2:300000:1:0:1:0:1:0:1:0:0:0:0:1", "utf-8")
    timer = Dual150Timer()
    timer.apply_raw_bytes_data(input)
    assert str(timer.countdown) == "5:00.000"

    input = bytes("2:0:8:32010:8:29470:8:0:1:0:0:0:0:8", "utf-8")
    timer = Dual150Timer()
    timer.apply_raw_bytes_data(input)
    assert str(timer.line_one) == "0:32.010"
    assert str(timer.line_two) == "0:29.470"


def test_apply_raw_data():
    input = "2:300000:1:0:1:0:1:0:1:0:0:0:0:1"
    timer = Dual150Timer()
    timer.apply_raw_data(input)
    assert str(timer.countdown) == "5:00.000"

    input = "2:0:8:32010:8:29470:8:0:1:0:0:0:0:8"
    timer = Dual150Timer()
    timer.apply_raw_data(input)
    assert str(timer.line_one) == "0:32.010"
    assert str(timer.line_two) == "0:29.470"