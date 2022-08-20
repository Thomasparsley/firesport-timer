""" try:
            async with websockets.connect("ws://127.0.0.1:4999/ws/timer") as ws:
                timer = {
                    "left": d.line_one.format_time(),
                    "right": d.line_two.format_time()
                }
                await ws.send(json.dumps(timer))
        except Exception as e:
            print("Error: ", e) """