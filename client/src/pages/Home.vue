<script setup>
import TimerDisplay from "../components/TimerDisplay.vue";

import { defineProps } from "vue";

defineProps({
  settings: Object,
  liveTimer: Object,
});

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
  <div class="wrapper">
    <div class="timer-grid">
      <TimerDisplay title="Odpočet" time="00:00.000" />

      <TimerDisplay
        v-if="settings.lines.oneOn"
        title="Dráha jedna (Levý terč)"
        :time="liveTimer.left"
      />
      <TimerDisplay
        v-if="settings.lines.twoOn"
        title="Dráha dva (Pravý terč)"
        :time="liveTimer.right"
      />
      <TimerDisplay
        v-if="settings.lines.threeOn"
        title="Dráha tři"
        time="00:00.000"
      />
      <TimerDisplay
        v-if="settings.lines.fourOn"
        title="Dráha čtyři"
        time="00:00.000"
      />
    </div>
  </div>
</template>

<style lang="scss" scoped>
.timer-grid {
  margin-top: 4rem;

  display: grid;
  grid-template-columns: 1fr;
  @media only screen and (min-width: 800px) {
    grid-template-columns: repeat(2, 1fr);
  }
  gap: 1.25rem;

  & > div {
    &:first-child {
      @media only screen and (min-width: 800px) {
        grid-column: 1 / 3;
      }
    }
  }
}
</style>
