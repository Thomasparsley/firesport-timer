import dearpygui.dearpygui as dpg  # type: ignore

from .timer_kocab import KocabTimer


class Dual150Timer(KocabTimer):
    name: str = "Dual-150 (TRV Kocab)"

    def make_gui(self):
        with dpg.window(label=self.name, on_close=self.on_close):  # type: ignore
            if self.is_open:
                return
            else:
                self.is_open = True

            with dpg.group(horizontal=True): # type: ignore
                with dpg.group(): # type: ignore
                    dpg.add_text("Levy terc:")  # type: ignore
                    dpg.add_text(str(self.line_one))  # type: ignore

                with dpg.group(): # type: ignore
                    dpg.add_text("Pravy terc:")  # type: ignore
                    dpg.add_text(str(self.line_two))  # type: ignore

    def on_close(self):
        self.is_open = False
