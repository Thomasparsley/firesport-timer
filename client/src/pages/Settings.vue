<script>
export default {
  props: {
    settings: Object,
  },
};
</script>

<template>
  <div class="wrapper" v-if="settings.isConnected">
    <button v-if="!settings.isRunning" @click="startCloseTimer">
      Začít snímat časomíru
    </button>
    <button v-else @click="startCloseTimer">Ukončit snímání časomíry</button>

    <button v-if="settings.isRunning" @click="resetTimer">
      Resetovat čas na časomíře
    </button>

    <label>
      <p>Název portu</p>
      <input type="text" v-model="settings.portName" placeholder="COM4" />
    </label>
  </div>
</template>

<style lang="scss" scoped>
div {
  margin-top: 4rem;
}
</style>