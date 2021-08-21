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
        left: "00:00.000",
        right: "00:00.000",
      },

      settings: {
        isConnected: true,
        isRunning: false,
        portName: "",
        lines: {
          oneOn: true,
          twoOn: true,
          threeOn: false,
          fourOn: false,
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

        if (
          data.lineOne &&
          data.lineOne == "00:00.000" &&
          data.lineTwo &&
          data.lineTwo == "00:00.000"
        ) {
          this.liveTimer.left = data.countdown;
          this.liveTimer.right = data.countdown;
        } else {
          this.liveTimer.left = data.lineOne;
          this.liveTimer.right = data.lineTwo;
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
