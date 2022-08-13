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
                    self.gui_left_text = dpg.add_text(str(self.line_one))  # type: ignore

                with dpg.group(): # type: ignore
                    dpg.add_text("Pravy terc:")  # type: ignore
                    self.gui_right_text = dpg.add_text(str(self.line_two))  # type: ignore

            dpg.add_separator() # type: ignore
            
            with dpg.group(): # type: ignore
                dpg.add_text("Odpocet:")  # type: ignore
                self.gui_countdown_text = dpg.add_text(str(self.countdown))  # type: ignore

            dpg.add_separator() # type: ignore
            dpg.add_button( # type: ignore
                label="Reset",
                callback=self.send_reset,
            )

    def on_close(self):
        self.is_open = False

    def update_timer(self):
        super().update_timer()
        
        dpg.set_value( # type: ignore
            self.gui_left_text,
            str(self.line_one),
        )
        dpg.set_value( # type: ignore
            self.gui_right_text,
            str(self.line_two),
        )
        dpg.set_value( # type: ignore
            self.gui_countdown_text,
            str(self.countdown),
        )