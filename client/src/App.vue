<script setup>
import Update from "./components/Update.vue";

import TimerDisplay from "./components/TimerDisplay.vue";
import Footer from "./components/Footer.vue";

import { reactive } from "vue";

// This starter template is using Vue 3 experimental <script setup> SFCs
// Check out https://github.com/vuejs/rfcs/blob/master/active-rfcs/0040-script-setup.md

const bodyDataset = document.body.dataset;

const clieantVersion = "1.0.0";
const appVersion = bodyDataset.appVersion || "X.X.X";
const appAddress = bodyDataset.appAddress || "127.0.0.1";
const appPort = bodyDataset.appPort || "3000";

const latestAppVersion = "1.0.0";
const isUpdateAvailable = appVersion !== latestAppVersion;

console.log({
	appAddress,
	appPort,
	appVersion,
	clieantVersion,
	latestAppVersion,
	isUpdateAvailable,
});

const liveTimer = reactive({
	left: "00:00:00",
	right: "00:00:00",
});

if (appAddress && appPort) {
	const socket = new WebSocket(`ws://${appAddress}:${appPort}/api/timer`);
}
</script>

<template>
	<Update v-if="isUpdateAvailable" />

	<div>
		<TimerDisplay position="L" :time="liveTimer.left" />
		<TimerDisplay position="R" :time="liveTimer.right" />
	</div>

	<Footer :client="clieantVersion" :app="appVersion" />
</template>

<style>
body {
	margin: 0;
}

#app {
	font-family: Avenir, Helvetica, Arial, sans-serif;
	-webkit-font-smoothing: antialiased;
	-moz-osx-font-smoothing: grayscale;
	color: #2c3e50;

	position: relative;
}
</style>
