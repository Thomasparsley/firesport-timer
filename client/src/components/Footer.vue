<script setup>
import { defineProps } from "vue";

defineProps({
	app: String,
});
</script>

<template>
	<footer>
		<div>
			<span>Verze aplikace: V{{ app }}</span>
		</div>

		<div>
			<a href="https://thomasparsley.cz/" target="_blank">Autor</a>
			|
			<a target="_blank">GitHub</a>
		</div>
	</footer>
</template>

<style scoped>
footer {
	position: fixed;

	left: 0;
	bottom: 0;

	width: calc(100% - (0.8rem * 2));

	display: flex;
	justify-content: space-between;

	padding: 0.4rem 0.8rem;

	color: white;
	background-color: black;
}

a {
	color: aqua;
}
</style>
