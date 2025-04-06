<script lang="ts">
    import Sortable from "sortablejs";
    import { evaluate } from "mathjs";
    import { onMount } from "svelte";
    import { score } from "$lib";

    let { data } = $props();

    let originalSequence = data.sequence.map(String);
    let sequence = [...originalSequence];
    let result = $state<number | null>(null);

    let sequenceContainer: HTMLElement;
    let operatorsContainer: HTMLElement;

    function values() {
        return Array.from(sequenceContainer.querySelectorAll(".number")).map(
            (div) => {
                return div.innerHTML == "×" ? "*" : div.innerHTML;
            },
        );
    }

    let hintVisible = $state(false);

    function showHint() {
        hintVisible = !hintVisible;
    }

    onMount(() => {
        Sortable.create(sequenceContainer, {
            group: {
                name: "equation",
                pull: true,
                put: true,
            },
            filter: ".non-draggable",
            delay: 100,
            animation: 150,
            onChange: () => {
                try {
                    result = evaluate(values().join(""));
                } catch {
                    result = null;
                }
            },
        });

        Sortable.create(operatorsContainer, {
            group: {
                name: "equation",
                pull: "clone",
                put: true,
            },
            sort: true,
            animation: 150,
            setData: (dataTransfer, dragEl) => {
                dataTransfer.setData("Text", dragEl.textContent!);
            },
            onAdd: (evt) => {
                // evt.item.style.display = "none";
                evt.item.parentNode?.removeChild(evt.item);
                try {
                    result = evaluate(values().join(""));
                } catch {
                    result = null;
                }
            },
        });
    });

    $effect(() => {
        score.set(result || -1);
    });
</script>

<main class="flex flex-col gap-5 p-10 w-full justify-center items-center">
    <h1 class="text-6xl">
        {result == -1 ? "Invalid" : `Sum is ${result || 0}`}
    </h1>

    <dir
        bind:this={sequenceContainer}
        class="flex flex-row gap-2 justify-center"
    >
        {#each sequence as number}
            <div
                class="number sm:w-16 sm:h-16 flex items-center justify-center text-2xl rounded-xl
            {data.operators.includes(number) ? 'cursor-no-drop' : 'cursor-grab'}
            {data.operators.includes(number)
                    ? 'sm:bg-[#4D9C02]'
                    : 'sm:bg-[#4c6449] non-draggable'}"
            >
                {number}
            </div>
        {/each}
    </dir>

    {#if data.showHint}
        <button
            class="flex flex-row gap-4 w-full justify-center items-center text-xl text-white cursor-pointer"
        >
            {#if hintVisible}
                <div class="text-2xl">
                    {data.hints}
                </div>
            {:else}
                <svg
                    width="30"
                    height="30"
                    viewBox="0 0 50 50"
                    fill="none"
                    xmlns="http://www.w3.org/2000/svg"
                >
                    <path
                        d="M28.4532 2.69919C26.4336 0.72263 23.6407 0.765598 21.6426 2.76365L2.88676 21.498C0.888715 23.4746 0.867231 26.3105 2.86528 28.3086L21.793 47.1933C23.8125 49.1914 26.6055 49.1269 28.6036 47.1504L47.3594 28.3945C49.3575 26.3965 49.3789 23.5605 47.3809 21.584L28.4532 2.69919ZM28.0235 3.68747L46.3926 22.0136C48.1329 23.7324 48.1543 26.1601 46.4356 27.9004L28.1094 46.248C26.3477 48.0097 23.9844 47.9883 22.2227 46.2265L3.85356 27.8789C2.11332 26.1386 2.09184 23.7109 3.81059 21.9707L22.1368 3.66599C23.877 1.92575 26.2832 1.94724 28.0235 3.68747ZM24.6504 29.748C24.9512 29.748 25.1661 29.5117 25.1661 29.1465V28.3515C25.1661 26.955 25.7461 26.1172 27.5079 24.914C29.7852 23.3457 30.5586 22.0781 30.5586 20.1875C30.5586 17.6738 28.3028 15.7617 25.0586 15.7617C21.793 15.7617 19.709 17.7597 19.4727 20.0371C19.4512 20.2949 19.4297 20.5312 19.4297 20.8105C19.4297 21.1543 19.6661 21.3691 19.9883 21.3691C20.3321 21.3691 20.5469 21.1543 20.5684 20.8105C20.7188 18.6191 22.3516 16.7929 24.9942 16.7929C27.7227 16.7929 29.4415 18.2539 29.4415 20.2519C29.4415 21.7343 28.7325 22.7441 26.584 24.2265C24.7364 25.5156 24.0918 26.6113 24.0918 28.2656V29.1679C24.0918 29.5117 24.3282 29.748 24.6504 29.748ZM24.6075 35.3125C25.3809 35.3125 25.961 34.6679 25.961 33.959C25.961 33.2285 25.3809 32.6054 24.6075 32.6054C23.8555 32.6054 23.254 33.2285 23.254 33.959C23.254 34.6679 23.8555 35.3125 24.6075 35.3125Z"
                        fill="black"
                        stroke="white"
                        stroke-width="2"
                    />
                </svg>
                <!-- svelte-ignore node_invalid_placement_ssr -->
                <button onclick={showHint} class="text-2xl">Show me!</button>
            {/if}
        </button>
    {/if}

    <div
        bind:this={operatorsContainer}
        class="flex flex-row gap-2 justify-center"
    >
        {#each data.operators as operator}
            <div
                class="number w-8 h-8 sm:w-16 sm:h-16 bg-[#4D9C02] flex items-center justify-center text-2xl rounded-xl cursor-grab"
            >
                {operator}
            </div>
        {/each}
    </div>
</main>
