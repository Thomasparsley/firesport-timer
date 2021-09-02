<script>
import Header from "./components/Header.vue";
import Footer from "./components/Footer.vue";

import HomePage from "./pages/Home.vue";
import Settings from "./pages/Settings.vue";

const bodyDataset = document.body.dataset;

export default {
  name: "App",
  components: {
    Header,
    Footer,

    HomePage,
    Settings,
  },

  data() {
    return {
      appVersion: bodyDataset.appVersion || "X.X.X",
      appAddress: bodyDataset.appAddress || "127.0.0.1",
      appPort: bodyDataset.appPort || "3000",

      menu: {
        homePage: true,
        settings: false,
      },

      liveTimer: {
        countdown: "0:00.000",
        lineOne: "0:00.000",
        lineTwo: "0:00.000",
        lineThree: "0:00.000",
        lineFour: "0:00.000",
      },

      settings: {
        isConnected: true,
        isRunning: false,
        portName: "",
        lines: {
          oneOn: true,
          twoOn: true,
          threeOn: true,
          fourOn: true,
        },
      },
    };
  },

  mounted() {
    if (this.appAddress && this.appPort) {
      const socket = new WebSocket(
        `ws://${this.appAddress}:${this.appPort}/ws`
      );

      // read from socket and update the timer display
      socket.onmessage = (event) => {
        const data = JSON.parse(event.data);

        this.liveTimer.countdown = data.countdown;
        this.liveTimer.lineOne = data.lineOne;
        this.liveTimer.lineTwo = data.lineTwo;
        this.liveTimer.lineThree = data.lineThree;
        this.liveTimer.lineFour = data.lineFour;
      };

      socket.onopen = () => {
        this.settings.isConnected = true;
      };

      socket.onerror = (event) => {
        console.log("Socket error: ", event);
      };

      socket.onclose = (event) => {
        this.settings.isConnected = false;
      };
    }
  },
};
</script>

<template>
  <Header :menu="menu" />

  <HomePage v-if="menu.homePage" :settings="settings" :liveTimer="liveTimer" />
  <Settings v-if="menu.settings" :settings="settings" />

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
  width: calc(100% - 6rem);
  margin: 0 3rem;

  @media only screen and (min-width: 800px) {
    width: 100%;
    margin: auto;
  }
}
</style>
