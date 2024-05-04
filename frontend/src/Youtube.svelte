<script lang="ts">
    import { onMount } from "svelte";

    export let player: YT.Player;
    export let initialVideoId = "";

    const ytPlayerId = "youtube-player";

    const load = () => {
        player = new YT.Player(ytPlayerId, {
            height: "100%",
            width: "100%",
            videoId: initialVideoId,
        });
    };

    onMount(() => {
        if (window.YT) {
            load();
        } else {
            window.onYouTubeIframeAPIReady = load;
        }
    });
</script>

<svelte:head>
    <script src="https://www.youtube.com/iframe_api"></script>
</svelte:head>

<div id={ytPlayerId} />
