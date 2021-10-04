<script>
export default {
  props: {
    menu: Object,
  },

  methods: {
    menuDispecher(menuID) {
      switch (menuID) {
        case 0:
          this.menu.homePage = true;
          this.menu.settings = false;
          break;
        case 1:
          this.menu.settings = true;
          this.menu.homePage = false;
          break;
      }
    },
  },
};
</script>

<template>
  <header>
    <nav>
      <ul>
        <li @click.prevent="menuDispecher(0)">Zobrazení času</li>
        <li @click.prevent="menuDispecher(1)">Nastavení</li>
      </ul>
    </nav>

    <div></div>
  </header>
</template>

<style lang="scss" scoped>
header {
  width: calc(100% - (2 * 0.8rem));
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

ul {
  padding: 0;
  margin: 0;
  list-style: none;

  display: flex;
  gap: 1.5rem;

  li {
    padding: 0.4rem 0.8rem;
    background-color: rgba(white, 0.2);

    cursor: pointer;

    &:hover {
      color: black;
      background-color: white;
    }
  }
}
</style>