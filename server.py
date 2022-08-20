import pathlib
from typing import Any

from fastapi import FastAPI, WebSocket, Request, WebSocketDisconnect
from websockets.exceptions import ConnectionClosed

from vel.template import Jinja2Templating
from vel.livereload import LiveReload


STREAM_DELAY = 0.03  # second
RETRY_TIMEOUT = 15000  # milisecond

app = FastAPI()
templating = Jinja2Templating(directory="templates")

live_reload = LiveReload(pathlib.Path() / "templates")
live_reload.start()
live_reload.add_endpoint(app, "/ws/livereload")
templating.env.globals["livereload"] = live_reload.jinja_filter("ws://localhost:8000/ws/livereload")  # type: ignore

left_time: str = "0:00.000"
right_time: str = "0:00.000"


class WSHub:
    def __init__(self):
        self.connections: list[WebSocket] = []

    async def connect(self, connection: WebSocket):
        await connection.accept()
        self.connections.append(connection)

    def disconnect(self, connection: WebSocket):
        self.connections.remove(connection)

    async def broadcast_json(self, json_message: dict[str, Any]):
        for connection in self.connections:
            await connection.send_json(json_message)

    async def send_personal_json(
        self, websocket: WebSocket, json_message: dict[str, Any]
    ):
        await websocket.send_json(json_message)


web_socket_hub = WSHub()


@app.websocket("/ws")
async def reciver(connection: WebSocket):
    global left_time
    global right_time

    await web_socket_hub.connect(connection)

    try:
        while True:
            data = await connection.receive_json()
            left_time = data["left_time"]
            right_time = data["right_time"]
    except WebSocketDisconnect:
        pass
    except ConnectionClosed:
        pass
    finally:
        web_socket_hub.disconnect(connection)


@app.get("/live_timer")
async def message_stream(request: Request):
    global left_time
    global right_time

    return templating.RenderResponse(request, "partial.html", {
        "left_time": left_time,
        "right_time": right_time,
    })


@app.get("/")
async def index(request: Request):
    global left_time
    global right_time

    return templating.RenderResponse(request, "index.html", {
        "left_time": left_time,
        "right_time": right_time,
    })
