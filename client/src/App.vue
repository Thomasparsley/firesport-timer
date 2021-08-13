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

console.log({
	appAddress,
	appPort,
	appVersion,
});

const liveTimer = reactive({
	left: "00:00.000",
	right: "00:00.000",
});

const timer = reactive({
	isConnected: true,
	isRunning: false,
	portName: "",
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
		timer.isConnected = true;
	};

	socket.onerror = (event) => {
		console.log("Socket error: ", event);
	};

	socket.onclose = (event) => {
		timer.isConnected = false;
	};
}

function startCloseTimer() {
	if (timer.isRunning) {
        fetch("/api/close", {
			method: "POST",
			credentials: "same-origin",
			headers: {
				"Content-Type": "application/json",
			},
		});

		timer.isRunning = false;
	} else {
		if (!timer.isConnected || timer.portName == "") {
			return;
		}

		fetch("/api/start", {
			method: "POST",
			credentials: "same-origin",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify({
				port: timer.portName,
			}),
		});

		timer.isRunning = true;
	}
}

function resetTimer() {
	if (timer.isConnected && timer.isRunning) {
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
	<div class="timerGrid">
		<TimerDisplay position="L" :time="liveTimer.left" />
		<TimerDisplay position="R" :time="liveTimer.right" />
	</div>

	<div class="buttonGrid" v-if="timer.isConnected">
		<button v-if="!timer.isRunning" @click="startCloseTimer">
			Začít snímat časomíru
		</button>
		<button v-else @click="startCloseTimer">
			Ukončit snímání časomíry
		</button>

		<button v-if="timer.isRunning" @click="resetTimer">Resetovat čas na časomíře</button>

		<label>
			<p>Název portu</p>
			<input type="text" v-model="timer.portName" placeholder="COM4" />
		</label>
	</div>

	<Footer :app="appVersion" />
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

.timerGrid,
.buttonGrid {
	display: flex;
	flex-direction: row;
	justify-content: center;
	align-items: center;
	flex-wrap: wrap;
}

.timerGrid {
	gap: 3.5rem;
}

.buttonGrid {
	gap: 1rem;
}

button {
	font-size: 1.25rem;
}

label p {
	margin: 0;
}
</style>
