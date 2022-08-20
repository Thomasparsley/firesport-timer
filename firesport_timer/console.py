import random
from time import sleep
from timers.kocab_dual150 import Dual150Timer

timer = Dual150Timer()


while True:
    choise = random.choice([
        "2:0:1:0:1:0:1:0:1:0:0:0:0:1",
        "2:0:4:6270:2:6270:2:0:1:0:0:0:0:2",
        "2:0:8:32010:8:29470:8:0:1:0:0:0:0:8",
        "2:0:4:5390:8:18070:2:0:1:0:0:0:0:2",
        "2:300000:1:0:1:0:1:0:1:0:0:0:0:1",
    ])
    choise = bytes("2:0:8:32010:8:29470:8:0:1:0:0:0:0:8", "utf-8")

    timer.apply_raw_bytes_data(choise)
    print(timer)
    sleep(0.3)
