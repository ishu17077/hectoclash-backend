<script lang="ts">
    import { user } from "$lib";
    import Equation from "../components/Equation.svelte";
    let { data } = $props();

    user.set(data.currentUser);
</script>

<h2 class="pointer-events-none select-none whitespace-pre-wrap bg-gradient-to-b from-black to-gray-100/99 bg-clip-text text-center text-[14rem] font-semibold leading-none text-transparent dark:from-white dark:to-slate-900/10">
    HectoClash
</h2>

<div class="flex flex-col justify-center items-center">
    
    <a href="/game" class="w-full sm:w-fit h-fit hover:bg-white hover:text-black text-6xl">
        <h2>Play Now!!</h2>
    </a>    

    <span class="text-xl w-[50%] text-center" style="color: #ffffff80; font-size: 19px;">Six digits from 1 to 9 are given. The target is always 100. All six digits must be used in the order given, along with operators +,-,*,/,(,) and pow(base,exp). Single digits can be combined into larger numbers.</span>
    <br>
    <br>
    <br>
    <br>

    <h1>Let's test your skills</h1>
    <Equation 
            data={{
                sequence: ["1", "2", "3", "4", "5", "6"],
                operators: ["+", "-", "*", "/", "(", ")"],
                hints: "",
                showHint: false
            }} 
        />
</div>