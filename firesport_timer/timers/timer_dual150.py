import dearpygui.dearpygui as dpg  # type: ignore

from .timer import Timer


class Dual150Timer(Timer):
    name: str = "Dual-150 (TRV Kocab)"

    def make_gui(self):
        with dpg.window(  # type: ignore
            label=self.name,
        ):
            pass
