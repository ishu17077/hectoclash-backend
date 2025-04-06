<script lang="ts">
    import { onDestroy } from "svelte";

    let { duration } = $props();

    let time = $state(duration);

    function formatDuration(seconds: number) {
        const mins = Math.floor(seconds / 60);
        const secs = Math.floor(seconds % 60);
        return `${mins}:${secs.toString().padStart(2, "0")}`;
    }

    const intervalId = setInterval(() => {
        if (time > 0) {
            time = Math.max(0, time - 1 / 60);
        } else {
            clearInterval(intervalId);
        }
    }, 1000 / 60);

    onDestroy(() => {
        clearInterval(intervalId);
    });

    function getProgressColor() {
        const percentage = time / duration;
        if (percentage > 0.5) {
            return "#4D9C02";
        } else if (percentage > 0.25) {
            return "#FFA500";
        } else {
            return "#FF0000";
        }
    }
</script>

<div class="flex flex-col items-center">
    <h1>
        {#if time != 0}
            {formatDuration(time)}
        {:else}
            Time's Up!
        {/if}
    </h1>
    
    <div class="flex w-96 justify-center">
        <div
            class="h-2 rounded-full transition-colors"
            style="width: {(time / duration) *
                100}%; background-color: {getProgressColor()};"
        ></div>
    </div>
</div>
