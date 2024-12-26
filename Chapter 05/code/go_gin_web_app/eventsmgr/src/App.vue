<template>
    <div id="app">
        <h1>Event List</h1>
        <div id="widget-container"></div>
        <form @submit.prevent="sendItem()">
            <input
                type="text"
                size="50"
                v-model="geventitem"
                placeholder="Enter new item"
            />
            <input type="submit" value="Submit" />
        </form>
        <ul>
            <li v-for="item in geventlist" v-bind:key="item">{{ item }}</li>
        </ul>
        <div>{{ message }}</div>
    </div>
</template>

<script>
import axios from "axios";
const appData = {
    geventlist: ["Evening Party : 6 p.m."],
    token: "",
    message: "",
};
export default {
    name: "App",
    data() {
        return appData;
    },
    methods: {
        getList: getList,
        sendItem: sendItem,
    },
    mounted: function () {
        getList();
    },
};

function getList() {
    axios.get("/api/events").then((res) => {
        appData.geventlist = res.data.list;
    });
}

async function sendItem() {
    const params = new URLSearchParams();
    params.append("item", this.geventitem);
    await axios
        .post("/api/events", params)
        .then(function () {
            getList();
        })
        .catch(function (error) {
            appData.geventlist = [error.message];
        });
}
</script>

<style>
#app {
    font-family: Avenir, Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    margin-top: 60px;
}
</style>
