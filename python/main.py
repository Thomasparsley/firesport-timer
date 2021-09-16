import tkinter as tk

import window


def main():
    root = window.Window(
        title="Firesport timer",
        width=1280,
        height=720,
    )

    my_label = tk.Label(root.window, text="Hello World")
    my_label.pack()

    my_button = tk.Button(
        master=root.window, text="Change mode", command=root.change_theme)
    my_button.pack()

    root.main_loop(widget_list=[
        my_label,
        my_button
    ])
    print("Goodbye")


if __name__ == '__main__':
    main()

""" import os
import serial

from kocab import dual150

ser = serial.Serial('COM4', 115200, timeout=0.1)
for _ in range(250):
    write_command = dual150.READ_COMMAND + "\n"
    ser.write(bytes(write_command, 'utf-8'))
    raw = ser.readline()

    print(raw)
    d = dual150.parse_raw_data(raw)
    print(d) """
