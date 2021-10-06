import serial
import json
import time
import asyncio
import websockets
import os

import kocab.dual150 as dual150

ser = serial.Serial('COM4', 115200, timeout=0.1, bytesize=serial.EIGHTBITS, )

if not ser.is_open:
    raise Exception('Serial port not open')

for _ in range(50):
    write_command = dual150.READ_COMMAND + "\n"
    ser.write(bytes(write_command, 'utf-8'))
    raw = ser.readline()
    print(raw)


async def main():
    i = 0
    while True:
        if i == 0:
            d = dual150.parse_raw_data("2:0:1:0:1:0:1:0:1:0:0:0:0:1")
        elif i == 1:
            d = dual150.parse_raw_data("2:0:4:6270:2:6270:2:0:1:0:0:0:0:2")
        elif i == 2:
            d = dual150.parse_raw_data("2:0:8:32010:8:29470:8:0:1:0:0:0:0:8")
        elif i == 3:
            d = dual150.parse_raw_data("2:0:4:5390:8:18070:2:0:1:0:0:0:0:2")
        elif i == 4:
            d = dual150.parse_raw_data(
                "2:0:4:5390:8:18070:8:698148:2:14235:8:0:0:2")
        else:
            d = dual150.parse_raw_data("2:300000:1:0:1:0:1:0:1:0:0:0:0:1")
            i = -1

        i += 1

        # connect to web socket
        """ try:
            async with websockets.connect("ws://127.0.0.1:4999/ws/timer") as ws:
                timer = {
                    "left": d.line_one.format_time(),
                    "right": d.line_two.format_time()
                }

                await ws.send(json.dumps(timer))
        except Exception as e:
            print("Error: ", e) """

        print(d)

        # sleep for 1s
        time.sleep(1)

        clearConsole()


if __name__ == "__main__":
    asyncio.run(main())
