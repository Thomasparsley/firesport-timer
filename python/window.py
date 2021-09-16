import tkinter as tk

BACKGROUND_WHITE = "white"
BACKGROUND_BLACK = "black"

THEME_DARK = "dark"
THEME_LIGHT = "light"


class Window:
    window = tk.Tk()
    theme: str = THEME_LIGHT
    theme_changed: bool = False

    def __init__(self, title: str, width: int, height: int):
        self.title = title
        self.width = width
        self.height = height

        self.window.title(self.title)
        self.window.geometry(str(self.width) + "x" + str(self.height))
        self.window.configure(background=BACKGROUND_WHITE)

    def change_theme(self):
        if self.theme == THEME_LIGHT:
            self.theme = THEME_DARK
        else:
            self.theme = THEME_LIGHT

        self.theme_changed = True

    def main_loop(self, widget_list: list):
        while True:
            """ if self.theme_changed:
                self.theme_changed = False
                self.window.configure(background=self.theme)

                for windget in widget_list:
                    windget.configure(background=self.theme) """

            self.window.update_idletasks()
            self.window.update()
