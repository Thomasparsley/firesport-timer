import tkinter as tk

COLOR_WHITE: str = 'white'


class Window:
    window = tk.Tk()

    def __init__(self, title: str, width: int, height: int):
        self.title = title
        self.width = width
        self.height = height

        self.window.title(self.title)
        self.window.geometry(str(self.width) + "x" + str(self.height))
        self.window.configure(background=COLOR_WHITE)

    def main_loop(self, custom_func=None):
        while True:
            if custom_func is not None:
                custom_func()

            self.window.update_idletasks()
            self.window.update()
