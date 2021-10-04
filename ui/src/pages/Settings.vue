<script>
export default {
  props: {
    settings: Object,
  },

  methods: {
    startCloseTimer() {
      if (this.settings.isRunning) {
        fetch("/api/close", {
          method: "POST",
          credentials: "same-origin",
          headers: {
            "Content-Type": "application/json",
          },
        });

        this.settings.isRunning = false;
      } else {
        if (!this.settings.isConnected || this.settings.portName == "") {
          return;
        }

        fetch("/api/start", {
          method: "POST",
          credentials: "same-origin",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            port: this.settings.portName,
          }),
        });

        this.settings.isRunning = true;
      }
    },
  },
};
</script>

<template>
  <div class="wrapper" v-if="settings.isConnected">
    <button v-if="!settings.isRunning" @click="startCloseTimer">
      Začít snímat časomíru
    </button>
    <button v-else @click="startCloseTimer">Ukončit snímání časomíry</button>

    <label>
      <p>Název portu</p>
      <input type="text" v-model="settings.portName" placeholder="COM4" />
    </label>

    <label>
        <p>Počet drah</p>
        <input type="number" v-model="settings.lines" min="1" max="4" />
    </label>
  </div>
</template>

<style lang="scss" scoped>
div {
  margin-top: 4rem;
}
</style>