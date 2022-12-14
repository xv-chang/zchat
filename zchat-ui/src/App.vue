

<template>
  <div class="menu">
    <div class="logo">ZCHAT</div>
    <div class="avatar" @click="changeAvatar()">
      <img :src="avatar" draggable="false" />
    </div>
    <div class="name" @click="changeNickName()">{{ nickName }}</div>
  </div>
  <div class="chat">
    <div
      class="chat-item"
      v-for="item in list"
      :key="item.id"
      :class="item.nickName == nickName ? 'self' : 'other'"
    >
      <div class="avatar">
        <img :src="item.avatar" draggable="false" />
      </div>
      <div class="msg-card">
        <div class="nick" v-if="item.nickName != nickName">
          {{ item.nickName }}
        </div>
        <div class="msg">
          <p @click="handleCopy(item)">{{ item.text }}</p>
          <time>{{ item.sendTime }}</time>
        </div>
      </div>
    </div>
  </div>
  <input
    class="textarea"
    v-model="text"
    @keydown.enter="sendText()"
    type="search"
    required
    placeholder="输入内容，回车发送"
  />
</template>


<script setup>
import randomName from "./utils/random-name";
import { onMounted, ref } from "vue";
const apiURL = import.meta.env.VITE_API_URL;
const wsURL = apiURL.replace("https:", "wss:").replace("http:", "ws:");

const getNow = () => {
  var d = new Date();
  return `${d.getFullYear()}/${
    d.getMonth() + 1
  }/${d.getDate()} ${d.getHours()}:${d.getMinutes()}`;
};
const getRandomPic = () => {
  let randomNum = Math.floor(Math.random() * 7);
  let imgUrl = `/avatar/${randomNum}.jpg`;
  return imgUrl;
};

let nickNameValue =
  localStorage.getItem("nickname") || randomName.getNickName();
let avatarValue = localStorage.getItem("avatar") || getRandomPic();
let clickCount = 0;

localStorage.setItem("nickname", nickNameValue);
localStorage.setItem("avatar", avatarValue);

const list = ref([]);
const text = ref("");

const nickName = ref(nickNameValue);
const avatar = ref(avatarValue);
const ws = new WebSocket(`${wsURL}/ws`);

ws.onmessage = (evt) => {
  let recv = JSON.parse(evt.data);
  list.value.push(recv);
  scrollToLast();
};
ws.onclose = (evt) => {
  console.log("onclose", evt);
};
ws.onopen = (evt) => {
  console.log("onopen", evt);
};

let scrollTimeout = null;
const scrollToLast = () => {
  if (scrollTimeout) {
    clearTimeout(scrollTimeout);
  }
  scrollTimeout = setTimeout(() => {
    document.querySelector(".chat").scrollTop =
      document.querySelector(".chat").scrollHeight;
    scrollTimeout = null;
  }, 500);
};

const fetchMessages = () => {
  fetch(`${apiURL}/messages`)
    .then((res) => res.json())
    .then((res) => {
      list.value = res.sort((a, b) => a.id - b.id);
      scrollToLast();
    });
};

const changeNickName = () => {
  let newVal = randomName.getNickName();
  localStorage.setItem("nickname", newVal);
  nickName.value = newVal;
};

const changeAvatar = () => {
  let newVal = getRandomPic();
  localStorage.setItem("avatar", newVal);
  avatar.value = newVal;
};

const sendText = () => {
  if (text.value) {
    let msg = {
      avatar: avatar.value,
      nickName: nickName.value,
      text: text.value,
      sendTime: getNow(),
    };
    ws.send(JSON.stringify(msg));
    text.value = "";
    scrollToLast();
  }
};

let clickTimer = null;

const handleCopy = (item) => {
  clickCount++;

  if (clickCount == 1) {
    setTimeout(() => {
      clickCount = 0;
    }, 300);
  }
  if (clickCount == 2) {
    const input = document.createElement("input");
    input.setAttribute("readonly", "readonly");
    input.value = item.text;
    document.body.appendChild(input);
    input.setSelectionRange(0, 9999); // 防止 ios 下没有全选内容而无法复制
    input.select();
    document.execCommand("copy");
    document.body.removeChild(input);
    alert("已复制内存到剪切板");
  }
};

onMounted(() => {
  fetchMessages();
});
</script>


