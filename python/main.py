""" import tkinter as tk


def a():
    print("a")


def main():
    root = tk.Tk()
    root.title("Firesport timer")

    my_label = tk.Label(root, text="Hello World")
    my_label.pack()

    my_button = tk.Button(root, text="Click me", command=a)
    my_button.pack()

    root.mainloop()

    super_loop()


def super_loop():
    pass


if __name__ == '__main__':
    main() """

import os
import serial

from kocab import dual150

ser = serial.Serial('COM4', 115200, timeout=0.1)
for _ in range(250):
    write_command = dual150.READ_COMMAND + "\n"
    ser.write(bytes(write_command, 'utf-8'))
    raw = ser.readline()

    print(raw)
    d = dual150.parse_raw_data(raw)
    print(d)
