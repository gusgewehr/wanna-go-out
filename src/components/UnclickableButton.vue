<script setup>
import { ref, inject } from 'vue'
import axios from 'axios'

var count = ref(0)
var affirmativeMessage = ref("Sim!")
var posY = ref(100)
var posX = ref(100)
var positionUncBtn = ref('static')
var calculatedWidth = ref(100)
var calculatedHeight = ref(70)

const isAccepted = inject("isAccepted")
const apiUrl = inject("apiUrl")



function randomPos(event) {
 
  positionUncBtn.value = 'fixed'

  var width = window.innerWidth
  var height = window.innerHeight

  if (width > height) {
    posY.value = (Math.random() * height)
    posX.value = (Math.random() * (width/2) )
  }
  else {
    posY.value = ((Math.random() * (height/2)) + height/2)
    posX.value = (Math.random() * width)
  }

  console.log(posX.value, width, (posX.value >= width-100))

  if(posX.value >= width-100) {
    posX.value = posX.value-100
  }

  if(posY.value >= height-70) {
    posY.value = posY.value-70
  }
  

  if(count.value > 1){
    calculatedWidth.value += 10
    calculatedHeight.value += 10

    affirmativeMessage.value = affirmativeMessage.value.concat('!')
  }
  
  
  count.value++
}


function affirmativeClick()  {
  isAccepted.value = true

  axios.post(apiUrl.value+'/message', {
    "text": affirmativeMessage.value
  }) .then(function (response) {
    console.log(response);
  })
  .catch(function (error) {
    console.log(error);
  });


}

</script>

<template>
  <div class="wrapper">
    <button @click="affirmativeClick()" class="affirmative button" :class="{ changeColor: count > 1}" :style="{width: calculatedWidth+'px', height:calculatedHeight+'px'} ">
    {{ affirmativeMessage }}
    </button>

    <button class="unclickable button" @mouseover="randomPos()" @mousedown="randomPos()"  v-bind:style="{ position:positionUncBtn, top: posY+'px', right: posX+'px'}">
      Não
    </button>
  </div>

</template>

<style scoped>
.wrapper{
  height: 100%;
  width: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: 35px
}


.button{
  background-color: #4a4a4a; /* Cor de fundo do botão */
  color: var(--base-color); /* Cor do texto */
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
  background-color: var(--secondary-bg-color)
}

.changeColor{
  animation: changeColor 3s infinite; /* Adiciona animação */
}

@keyframes changeColor {
            0% {
                background-color: #E94F37; /* Cor inicial */
            }
            33% {
                background-color: #932F6D; /* Cor intermediária 1 */
            }
            66% {
                background-color: #907F9F; /* Cor intermediária 2 */
            }
            100% {
                background-color: #E94F37; /* Volta à cor inicial */
            }
        }


</style>