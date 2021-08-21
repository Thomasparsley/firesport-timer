<script setup>
import { defineProps } from "vue";

defineProps({
  title: String,
  time: String,
});

function MinutesSeconds(time) {
  // Remove last three chars from time
  return time.slice(0, -3);
}

function miliseconds(time) {
  // Get last three chars from time
  return time.substr(time.length - 3);
}
</script>

<template>
  <div class="div">
    <span class="title">
      {{ title }}
    </span>

    <div class="time">
      <span>{{ MinutesSeconds(time) }}</span>
      <span class="small">{{ miliseconds(time) }}</span>
    </div>
  </div>
</template>

<style lang="scss" scoped>
@import url("https://fonts.googleapis.com/css2?family=Atkinson+Hyperlegible&display=swap");

.div {
  background-color: white;
  box-shadow: rgba(0, 0, 0, 0.15) 1.95px 1.95px 2.6px;

  border-radius: 1rem;

  padding-top: 0.75rem;
  padding-left: 0.75rem;
  padding-right: 0.75rem;

  @media only screen and (min-width: 1024px) {
    padding-top: 1.25rem;
    padding-left: 1.25rem;
    padding-right: 1.25rem;
  }

  .title {
    font-size: 1.5vmax;
    text-transform: uppercase;
  }

  .time {
    margin-top: -0.5rem;

    font-family: "Atkinson Hyperlegible", sans-serif;
    font-size: 6.5vmax;

    .small {
      font-size: 4.5vmax;
    }
  }
}
</style>
