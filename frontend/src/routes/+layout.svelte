<script lang="ts">
	import '../app.css';
    import Header from '../components/Header.svelte';
    import DotPattern from '../components/DotPattern.svelte';

	let { children } = $props();
</script>

<svelte:head>
	<title>HectoClash</title>
</svelte:head>

<main>
	<Header/>

	{@render children()}

	<DotPattern/>
</main>