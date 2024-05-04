<script lang="ts">
    export let num: number;
    export let start: string = "";
    export let end: string = "";
</script>

<div>
    <span>{num}</span>
    <input
        class="chopFormInput"
        bind:value={start}
        id="chopFormInput{num}_start"
        name="{num}_start"
    />
    <span class="line">~</span>
    <input
        class="chopFormInput"
        bind:value={end}
        id="chopFormInput{num}_end"
        name="{num}_end"
    />
</div>

<style>
    .chopFormInput {
        border: none;
        border-radius: 3px;
        outline: none;
        margin: 5px 10px;
        padding: auto;
        width: 80px;
        height: 30px;
    }

    .line {
        color: black;
        font-size: 20px;
    }
</style>
