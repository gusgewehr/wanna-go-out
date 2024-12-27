<script setup>
import HelloWorld from './components/HelloWorld.vue'
import UnclickableButton from './components/UnclickableButton.vue'
</script>

<template>
  <header>
    <img alt="Vue logo" class="logo" src="./assets/logo.svg" width="125" height="125" />

    <div class="wrapper">
      <HelloWorld msg="Deseja se cadastrar?" />
    </div>
  </header>

  <main>
    
    <div class="wrapper">
      <UnclickableButton />
    </div>


  </main>


</template>

<style scoped>
header {
  line-height: 1.5;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

@media (min-width: 1024px) {
  header,main{
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
    height: 100vh;

  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper, main .wrapper  {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }
}
</style>
