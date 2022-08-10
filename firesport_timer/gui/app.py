import dearpygui.dearpygui as dpg  # type: ignore

from .settings import Settings


class App:
    __slots__ = ["viewport_width", "viewport_height", "settings"]

    def __init__(self, viewport_width: int, viewport_height: int):
        self.viewport_width = viewport_width
        self.viewport_height = viewport_height

        self.settings = Settings()

    def start(self):
        dpg.create_context()  # type: ignore

        self.__create_menu()

        dpg.create_viewport(  # type: ignore
            title="Firesport timer",
            width=self.viewport_width,
            height=self.viewport_height,
        )
        dpg.setup_dearpygui()  # type: ignore
        dpg.show_viewport()  # type: ignore
        dpg.start_dearpygui()

    def close(self):
        dpg.destroy_context()  # type: ignore

    def open_timer(self):
        self.settings.open_timer()

    def __create_menu(self):
        with dpg.viewport_menu_bar():  # type: ignore
            dpg.add_menu_item(label="Časomíra", callback=self.open_timer) # type: ignore
            dpg.add_menu_item(label="Nastavení", callback=self.settings.open)  # type: ignore
