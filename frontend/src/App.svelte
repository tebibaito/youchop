<script lang="ts">
  import { onMount } from "svelte";
  import { Download, GetChatCount } from "../wailsjs/go/main/App";
  import Youtube from "./Youtube.svelte";
  import ChopForm from "./ChopForm.svelte";
  import type { ytchop } from "../wailsjs/go/models";
  import ChatChart from "./ChatChart.svelte";

  let player: YT.Player;

  let inputString: string;

  let chopsCnt = 1;

  let currentTime = 0;

  let videoId = "";

  let chartLavel = [];
  let chartCount = [];
  let chartUpdate;

  function loadVideo(): void {
    currentTime = 0;
    if (inputString.startsWith("http")) {
      const paramString = inputString.split("?")[1];
      const searchParams = new URLSearchParams(paramString);
      videoId = searchParams.get("v");
    } else {
      videoId = inputString;
    }
    player.loadVideoById(videoId, 0);
    player.stopVideo();
    setInterval(() => {
      currentTime = player.getCurrentTime();
    }, 500);

    const url = getUrl();
    GetChatCount(url).then((countData) => {
      countData.forEach((e) => {
        chartLavel.push(e.time);
        chartCount.push(e.count);
      });
      console.log(countData);
      chartUpdate();
    });
  }

  function addChopForm(): void {
    chopsCnt += 1;
    new ChopForm({
      target: document.getElementById("chopForms"),
      props: {
        num: chopsCnt,
      },
    });
    loadChopFormsOptions();
  }

  function loadChopFormsOptions(): void {
    const chopForms = document.getElementsByClassName("chopFormInput");
    const select = document.getElementById("formSelect");
    select.innerHTML = "";
    for (const cf of chopForms) {
      const option = document.createElement("option");
      option.value = cf.id;
      option.text = cf.getAttribute("name");
      option.className = "chopFormOption";
      select.append(option);
    }
  }

  function timeToCodeString(time: number): string {
    let code: string;
    let hour = ("00" + Math.trunc(time / 3600)).slice(-2);
    let minute = ("00" + Math.trunc((time / 60) % 60)).slice(-2);
    let second = ("00" + Math.trunc(time % 60)).slice(-2);
    code = `${hour}:${minute}:${second}`;
    return code;
  }

  function setTime(): void {
    const select = document.getElementById("formSelect") as HTMLSelectElement;
    const idx = select.selectedIndex;
    const id = select.options[idx].value;
    const input = document.getElementById(id) as HTMLInputElement;
    const timecode = timeToCodeString(currentTime);
    input.value = timecode;
  }

  function download(): void {
    document.getElementById("result").innerText = "Downloading...";
    const url = getUrl();
    let chopDataMap = new Map<number, ytchop.ChopData>();
    const chopFormList = Array.from(
      document.getElementsByClassName("chopFormInput"),
    );
    chopFormList.forEach((e) => {
      const inputElm = e as HTMLInputElement;
      const name = inputElm.getAttribute("name");
      const index = Number(name.split("_")[0]);
      const time = inputElm.value;
      const paramName = name.split("_")[1];
      if (chopDataMap.has(index)) {
        chopDataMap.get(index)[paramName] = time;
      } else {
        const c: ytchop.ChopData = {
          start: "",
          end: "",
        };
        c[paramName] = time;
        chopDataMap.set(index, c);
      }
    });
    const chopDataList = Array.from(chopDataMap.values());
    console.log(url);
    console.log(chopDataList);
    Download(url, chopDataList).then((val) => {
      document.getElementById("result").innerText = val;
    });
  }

  function getUrl(): string {
    if (inputString.startsWith("http")) {
      return inputString;
    } else {
      return "https://www.youtube.com/watch?v=" + videoId;
    }
  }

  function setVideoTime(time) {
    player.seekTo(time, true);
  }

  function skipTime(sec: number): void {
    player.seekTo(currentTime + sec, true);
  }

  onMount(() => {
    loadChopFormsOptions();
  });
</script>

<main>
  <div class="input-box" id="input">
    <input
      autocomplete="off"
      bind:value={inputString}
      class="input"
      id="inputString"
      type="text"
    />
    <button class="btn" on:click={loadVideo}>Load</button>
  </div>
  <div class="container">
    <div class="row1">
      <div class="player">
        <Youtube bind:player></Youtube>
      </div>
      <div class="controller">
        <p class="timecode">{timeToCodeString(currentTime)}</p>
        <button
          on:click={() => {
            skipTime(-30);
          }}
          id="skipBtn1"
          class="skipBtn">-30sec</button
        >
        <button
          on:click={() => {
            skipTime(30);
          }}
          id="skipBtn2"
          class="skipBtn">+30sec</button
        >
        <select id="formSelect"></select>
        <button on:click={setTime} id="setBtn">Set</button>
      </div>
      <div class="chart">
        <ChatChart
          labels={chartLavel}
          counts={chartCount}
          bind:update={chartUpdate}
          {setVideoTime}
        />
      </div>
    </div>
    <div class="row2">
      <div id="chopForms">
        <ChopForm num={1}></ChopForm>
      </div>
      <button on:click={addChopForm} class="addBtn">Add</button>
      <br />
      <button on:click={download} class="downloadBtn">Download</button>
      <p id="result"></p>
    </div>
  </div>
</main>

<style>
  main {
    color: #333333;
  }

  .container {
    display: flex;
  }

  .row1 {
    width: 90%;
    height: 100%;
    margin: 10px;
  }

  .row2 {
    background-color: #bdbdbd;
    width: 40%;
    margin: 10px;
  }

  .player {
    height: 24em;
  }

  .controller {
    display: flex;
    background-color: #bdbdbd;
    height: 60px;
  }

  .chart {
    background-color: #bdbdbd;
    height: 14em;
    width: 100%;
  }

  .input-box .btn {
    width: 80px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 10px 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }

  .input-box .input {
    width: 80%;
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .skipBtn {
    border-radius: 10px;
    border: none;
    background-color: #dddddd;
    height: 50%;
    width: 60px;
    margin-top: 2%;
    margin-bottom: 2%;
    margin-left: 1%;
    margin-right: 1%;
  }

  #formSelect {
    background-color: rgba(255, 255, 255, 1);
    border-radius: 5px;
    width: 100px;
    height: 30px;
    margin-top: 2%;
    margin-bottom: 2%;
    margin-left: 2%;
    margin-right: 2%;
  }

  .timecode {
    font-size: 24px;
    margin-top: 2%;
    margin-bottom: 2%;
    margin-left: 2%;
    margin-right: 2%;
  }

  #setBtn {
    font-size: 16px;
    border-radius: 10px;
    border: none;
    background-color: #facf5a;
    height: 50%;
    width: 60px;
    margin-top: 2%;
    margin-bottom: 2%;
    margin-left: 2%;
    margin-right: 2%;
  }

  .downloadBtn {
    font-size: 16px;
    border-radius: 10px;
    border: none;
    background-color: #facf5a;
    height: 25px;
    margin-top: 2%;
    margin-bottom: 2%;
    margin-left: 2%;
    margin-right: 2%;
  }

  .addBtn {
    border-radius: 5px;
    border: none;
    background-color: #dddddd;
    height: 25px;
    margin-top: 2%;
    margin-bottom: 2%;
    margin-left: 2%;
    margin-right: 2%;
  }
</style>
