<script setup>



import HelloWorld from './components/HelloWorld.vue'
import UnclickableButton from './components/UnclickableButton.vue'
import CardList from './components/cardList.vue'


import { ref, provide } from 'vue'

var width = window.innerWidth
var height = window.innerHeight
var videoWidth = ref(1232)


if (width <= height){
  videoWidth = width -64
}


const isAccetped = ref(false)
const pickedHangout = ref("")
const apiUrl  =ref("http://localhost:8080/api/v1")

provide("isAccepted", isAccetped)
provide("pickedHangout", pickedHangout)
provide("apiUrl", apiUrl)


</script>

<template>
  <header class="top-half" v-if="pickedHangout == ''">
    <img alt="Vue logo" class="logo" src="./assets/logo.webp" width="250" height="250" v-if="!isAccetped" />
    <img alt="Vue logo" class="logo" src="./assets/logo2.webp" width="250" height="250" v-if="isAccetped" />

    <div class="wrapper">
      <HelloWorld msg="Deseja se cadastrar?" msg2="Cadastre-se gratuitamente e ganhe benefícios" v-if="!isAccetped" />

      <HelloWorld msg="Qual tipo de conta?" msg2="Para sabermos o que lhe recomendar" v-if="isAccetped" />
    </div>
    
  </header>

  <main v-if="pickedHangout == ''">



    <CardList v-if="pickedHangout == '' && isAccetped" />    
    
    <UnclickableButton v-if="!isAccetped"/>
    


  </main>

  <iframe v-if="pickedHangout != ''" autoplay :width="videoWidth" src="https://www.youtube.com/embed/Q0G-SThjW0c?autoplay=1" title="happy happy happy" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>


</template>



<style scoped>
header {
  line-height: 1.5;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

@media (max-width: 1024px) {
  .top-half{
    margin-bottom: 50px;
  }
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
    flex-wrap: wrap;
    justify-content: center;
    align-items: center;
    gap: 20px
  }
}
</style>
