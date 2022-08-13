from .ms import Ms


def test_eq():
    assert Ms() == Ms()
    assert Ms(0) == Ms(0)
    assert Ms(1) == Ms(1)
    assert Ms(25) == Ms(25)

    assert Ms() != Ms(25)
    assert Ms(0) != Ms(25)


def test_lt():
    assert Ms() < Ms(1)


def test_gt():
    assert Ms(5) > Ms(1)


def test_add():
    assert Ms() + Ms(1) == Ms(1)
    assert Ms() + 1 == Ms(1)


def test_sub():
    assert Ms(2) - Ms(1) == Ms(1)
    assert Ms(2) - 1 == Ms(1)


def test_mul():
    assert Ms() * Ms(2) == Ms()
    assert Ms() * 2 == Ms()
    assert Ms() * 0 == Ms()
    assert Ms(2) * Ms(2) == Ms(4)
    assert Ms(2) * 2 == Ms(4)


def test_str():
    tests = [
        (Ms(), "0:00.000"),
        (Ms(0), "0:00.000"),
        (Ms(6), "0:00.006"),
        (Ms(60), "0:00.060"),
        (Ms(600), "0:00.600"),
        (Ms(6000), "0:06.000"),
        (Ms(14000), "0:14.000"),
        (Ms(48248), "0:48.248"),
        (Ms(60000), "1:00.000"),
        (Ms(120000), "2:00.000"),
        (Ms(144567), "2:24.567"),
    ]

    for input, expected in tests:
        assert str(input) == expected


def test_is_zero():
    tests = [(Ms(), True), (Ms(0), True), (Ms(1), False)]

    for input, expected in tests:
        assert input.is_zero() == expected


def test_from_str():
    assert Ms.from_str("144567") == Ms(144567)
