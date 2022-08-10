import dearpygui.dearpygui as dpg  # type: ignore

from timers.timer import Timer
from timers.timers import Timers


class Settings:
    __slots__ = ["is_open", "timer"]

    def __init__(self) -> None:
        self.is_open = False
        self.timer: Timer | None = None

    def open(self) -> None:
        if not self.is_open:
            self.is_open = True
            self.__draw_gui()

    def open_timer(self) -> None:
        if self.timer and not self.timer.is_open:
            self.timer.make_gui()

    def __draw_gui(self) -> None:
        with dpg.window(  # type: ignore
            label="Nastavení",
            on_close=self.__close,
            width=500,
            height=300,
        ):
            dpg.add_slider_int(  # type: ignore
                default_value=2,
                min_value=1,
                max_value=8,
                label="Počet snímaných drah",
            )
            timers = [timer.value.name for timer in Timers]
            dpg.add_listbox(items=timers, label="Typ časomíry")  # type: ignore

            dpg.add_separator(indent=32) # type: ignore

            dpg.add_button( # type: ignore
                label="Nastavit"
            )

    def __close(self):
        self.is_open = False
