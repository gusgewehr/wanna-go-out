<script setup>


import Card from './ThinCard.vue'


var hangoutList = [
// {
//     icon: "fa-solid fa-film",
//     text: "Netflix & chill",
// },
{
    icon: "fa-solid fa-wine-glass-empty",
    text: "Fine Dine",
},
{
    icon: "fa-solid fa-record-vinyl",
    text: "Pub & Dine"
},
{
    icon: "fa-solid fa-bag-shopping",
    text: "Shopping & Movie",
}
]

</script>

<template>
    
    <div class="list">
        <p v-for="item in hangoutList">
            <Card :icon="item.icon" :text="item.text"/>
        </p>
    </div>

</template>

<style>
.list{
    display: flex;
    flex-direction: column;
    width: 100%;
    height: 100%;
    gap: 35px;
    justify-content: center;
}

</style>