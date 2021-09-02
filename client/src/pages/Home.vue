<script>
import TimerDisplay from "../components/TimerDisplay.vue";

export default {
    components: {
        TimerDisplay
    },

    props: {
        settings: Object,
        liveTimer: Object,
    },

    methods: {
        resetTimer() {
            if (this.settings.isConnected && this.settings.isRunning) {
                fetch("/api/reset", {
                method: "POST",
                credentials: "same-origin",
                headers: {
                    "Content-Type": "application/json",
                },
                });
            }
        }
    }
}
</script>

<template>
  <div class="wrapper">
    <div class="timer-grid">
      <TimerDisplay title="Odpočet" :time="liveTimer.countdown" />

      <TimerDisplay
        v-if="settings.lines.oneOn"
        title="Dráha jedna (Levý terč)"
        :time="liveTimer.lineOne"
      />
      <TimerDisplay
        v-if="settings.lines.twoOn"
        title="Dráha dva (Pravý terč)"
        :time="liveTimer.lineTwo"
      />
      <TimerDisplay
        v-if="settings.lines.lineThree"
        title="Dráha tři"
        time="00:00.000"
      />
      <TimerDisplay
        v-if="settings.lines.lineFour"
        title="Dráha čtyři"
        time="00:00.000"
      />
    </div>

    <button v-if="settings.isRunning" @click="resetTimer">
      Resetovat čas na časomíře
    </button>
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

button {
  cursor: pointer;

  margin-top: 2rem;
  padding: 1.5rem 3rem;
  font-size: 2rem;

  background-color: white;
  box-shadow: rgba(0, 0, 0, 0.15) 1.95px 1.95px 2.6px;
  border: none;
  border-radius: 1rem;

  &:hover {
    background-color: rgb(124, 124, 124);
    color: white;
  }
}
</style>
