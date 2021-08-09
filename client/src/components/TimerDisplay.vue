<script setup>
import { defineProps } from "vue";

defineProps({
	position: String,
	time: String,
});
</script>

<template>
	<p>{{ position }} {{ time }}</p>
</template>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Atkinson+Hyperlegible&display=swap");

p {
	font-family: "Atkinson Hyperlegible", sans-serif;
	font-size: 4rem;
}
</style>
