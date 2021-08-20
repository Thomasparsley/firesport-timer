<script setup>
import TimerDisplay from "./components/TimerDisplay.vue";
import Footer from "./components/Footer.vue";

import { reactive } from "vue";

// This starter template is using Vue 3 experimental <script setup> SFCs
// Check out https://github.com/vuejs/rfcs/blob/master/active-rfcs/0040-script-setup.md

const bodyDataset = document.body.dataset;

const appVersion = bodyDataset.appVersion || "X.X.X";
const appAddress = bodyDataset.appAddress || "127.0.0.1";
const appPort = bodyDataset.appPort || "3000";

const liveTimer = reactive({
  left: "00:00.000",
  right: "00:00.000",
});

const settings = reactive({
  isConnected: true,
  isRunning: false,
  portName: "",
  lines: {
    oneOn: true,
    twoOn: true,
    threeOn: false,
    fourOn: false,
  },
});

if (appAddress && appPort) {
  const socket = new WebSocket(`ws://${appAddress}:${appPort}/ws`);

  // read from socket and update the timer display
  socket.onmessage = (event) => {
    const data = JSON.parse(event.data);

    if (
      data.lineOne &&
      data.lineOne == "00:00.000" &&
      data.lineTwo &&
      data.lineTwo == "00:00.000"
    ) {
      liveTimer.left = data.countdown;
      liveTimer.right = data.countdown;
    } else {
      liveTimer.left = data.lineOne;
      liveTimer.right = data.lineTwo;
    }
  };

  socket.onopen = () => {
    settings.isConnected = true;
  };

  socket.onerror = (event) => {
    console.log("Socket error: ", event);
  };

  socket.onclose = (event) => {
    settings.isConnected = false;
  };
}

function startCloseTimer() {
  if (settings.isRunning) {
    fetch("/api/close", {
      method: "POST",
      credentials: "same-origin",
      headers: {
        "Content-Type": "application/json",
      },
    });

    settings.isRunning = false;
  } else {
    if (!settings.isConnected || settings.portName == "") {
      return;
    }

    fetch("/api/start", {
      method: "POST",
      credentials: "same-origin",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        port: settings.portName,
      }),
    });

    settings.isRunning = true;
  }
}

function resetTimer() {
  if (settings.isConnected && settings.isRunning) {
    fetch("/api/reset", {
      method: "POST",
      credentials: "same-origin",
      headers: {
        "Content-Type": "application/json",
      },
    });
  }
}
</script>

<template>
  <Footer :app="appVersion" />
</template>

<style lang="scss">
body {
  margin: 0;
  background-color: #f0f0f0;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;

  position: relative;
}

.wrapper {
  max-width: 70vmax;
  width: 100%;
  margin: auto;
}
</style>
