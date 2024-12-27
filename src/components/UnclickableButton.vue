<script setup>
import { ref } from 'vue'

var count = ref(0)
var affirmativeMessage = ref("Sim!")
var posY = ref(100)
var posX = ref(100)
var positionUncBtn = ref('relative')
var calculatedWidth = ref(100)
var calculatedHeight = ref(70)



function randomPos(event) {
 
  positionUncBtn.value = 'fixed'
  posY.value = Math.random() * window.innerHeight
  posX.value = Math.random() * window.innerWidth

  if(count.value > 1){
    calculatedWidth.value += 10
    calculatedHeight.value += 10

    affirmativeMessage.value = affirmativeMessage.value.concat('!')
  }
  
  
  count.value++
}

</script>

<template>
  <button class="affirmative button" :class="{ changeColor: count > 1}" :style="{width: calculatedWidth+'px', height:calculatedHeight+'px'} ">
    {{ affirmativeMessage }}
  </button>

  <button class="unclickable button" @mouseover="randomPos()" @mousedown="randomPos()"  v-bind:style="{ position:positionUncBtn, top: posY+'px', right: posX+'px'}">
    Não
  </button>
</template>

<style scoped>


.button{
  background-color: #4a4a4a; /* Cor de fundo do botão */
  color: #ffffff; /* Cor do texto */
  border: none;
  border-radius: 4px; /* Bordas arredondadas */
  padding: 10px 20px;
  font-size: 16px;
  font-weight: bold;
  cursor: pointer;
  box-shadow: 0px 2px 4px rgba(0, 0, 0, 0.5); /* Sombra */
  text-align: center;
  height: 70px;
  width: 100px;
}

.affirmative{
  background-color: hsla(160, 100%, 37%, 1); /* Cor de fundo do botão */
}

.changeColor{
  animation: changeColor 3s infinite; /* Adiciona animação */
}

@keyframes changeColor {
            0% {
                background-color: #00F0A0; /* Cor inicial */
            }
            33% {
                background-color: #BD7E00; /* Cor intermediária 1 */
            }
            66% {
                background-color: #7E00BD; /* Cor intermediária 2 */
            }
            100% {
                background-color: #00BD7E; /* Volta à cor inicial */
            }
        }


</style>