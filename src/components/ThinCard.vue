<script setup>
defineProps({
  icon: String,
  text: String
})

import { ref, inject } from 'vue'
import axios from 'axios'

const apiUrl = inject("apiUrl")
const pickedHangout = inject("pickedHangout")

function chooseHangout(pick){
    pickedHangout.value = pick.value

    axios.post(apiUrl.value+'/message', {
    "text": pick
  }) .then(function (response) {
    console.log(response);
  })
  .catch(function (error) {
    console.log(error);
  });

}
</script>

<template>
    <div class="card" @click="chooseHangout(text)">
        <span class="icon">
            <font-awesome-icon :icon="icon" />
        </span>
        <div class="text">
            {{text}}
        </div>
    </div>
</template>

<style scoped>
.icon{
    width: 30%;
    height: 100%;
    border-right: solid var(--color-border) 1px;
    font-size: 2em;
    display: flex;
    justify-content: left;
    align-items: center;
}

.text{
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;

}

.card{
    height: 75px;
    width: 100%;
    border-radius: 10px;
    filter: drop-shadow(0 5px 10px 0 var(--base-color));
    background-color: var(--base-color);
    padding: 20px;
    z-index: 0;
    overflow: hidden;
    transition: 0.6s ease-in;
    position: relative;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
    color: black;
    cursor: pointer;
}


.card::before {
  content: "";
  position: absolute;
  z-index: -1;
  top: -15px;
  right: -15px;
  background: var(--secondary-bg-color);
  height:220px;
  width: 25px;
  border-radius: 32px;
  transform: scale(1);
  transform-origin: 50% 50%;
  transition: transform 0.25s ease-out;
}

.card:hover::before, .card:active::before{
    transition-delay:0.2s ;

  transform: scale(80);
}

.card:hover, .card:active {
    color: var(--base-color);;

    .icon{
        transition-delay:0.2s ;
        border-right: solid var(--base-color) 1px;
    }

}

</style>