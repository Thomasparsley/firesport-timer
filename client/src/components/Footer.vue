<script setup>
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
			<a href="https://github.com/Thomasparsley/firesport-timer" target="_blank">GitHub</a>
		</div>
	</footer>
</template>

<style lang="scss" scoped>
footer {
    position: fixed;

    bottom: 0;
    left: 0;

    width: calc(100% - (2*0.8rem));
    padding: 0.4rem 0.8rem;

    display: flex;
    justify-content: space-between;
    gap: 1rem;


    font-size: 1.2rem;
    color: white;
    background-color: #1e2122;

    a {
        color: white;
    }
}
</style>
