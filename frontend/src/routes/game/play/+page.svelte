<script lang="ts">
    import { score, tscore, user } from "$lib";
    import Confetti from "../../../components/Confetti.svelte";
    import Equation from "../../../components/Equation.svelte";
    import Timer from "../../../components/Timer.svelte";
    import { onMount } from "svelte";
    import { goto } from "$app/navigation"; // Import the goto function

    let { data } = $props();

    user.set(data.currentUser);

    let currentIndex = $state(0);
    let timer: NodeJS.Timeout;
    let timeLeft = $state(60); // Initialize with the total duration (60 seconds)
    let showDialog = $state(false); // State to control the visibility of the dialog

    function nextProblem() {
        if (currentIndex < data.problems.length - 1) {
            currentIndex++;
            resetTimer();
        } else {
            // Show the dialog when the last problem is completed
            showDialog = true;
        }
    }

    function resetTimer() {
        clearTimeout(timer);
        timeLeft = 60; // Reset time left to 60 seconds
        timer = setInterval(() => {
            if (timeLeft > 0) {
                timeLeft--;
            } else {
                clearInterval(timer);
                if (currentIndex === data.problems.length - 1) {
                    // Show the dialog when the timer runs out on the last problem
                    showDialog = true;
                } else {
                    nextProblem();
                }
            }
        }, 1000);
    }

    function calculateScore() {
        // Add score based on the time left
        score.update((currentScore) => currentScore + timeLeft * 10); // Example: 10 points per second left
    }

    function redirectToGame() {
        goto("/game"); // Redirect to the game page
    }

    onMount(() => {
        resetTimer();
    });

    import { onDestroy } from "svelte";
    onDestroy(() => {
        clearTimeout(timer);
    });

    $effect(() => {
        if ($score == 100) {
            setTimeout(() => {
                calculateScore();
                $tscore += $score; // Update score based on time left
                nextProblem();
            }, 2000);
        }
    });
</script>

{#if $score == 100}
    <Confetti />
{/if}

<main class="flex flex-col justify-center items-center gap-10">
    {#key currentIndex}
        <h1>Problem {currentIndex + 1}</h1>

        <Timer duration={60} />

        <Equation
            data={{
                sequence:
                    data.problems[currentIndex].problems[0].problem_components,
                operators: data.problems[currentIndex].problems[0].operators,
                hints: data.problems[currentIndex].problems[0].hint,
                showHint: true,
            }}
        />
    {/key}

    <!-- Dialog -->
    {#if showDialog}
        <div class="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center">
            <div class="text-3xl p-6 rounded-lg shadow-lg text-center space-y-4">
                <Confetti/>

                <h2 style="font-size: 6rem">Game Over!</h2>
                <p class="text-gray-300">You score was <strong>{$tscore}</strong></p>
                <button
                    class="hover:bg-white hover:text-black"
                    onclick={redirectToGame}
                >
                    <h1>OK</h1>
                </button>
            </div>
        </div>
    {/if}
</main>