<script lang="ts">
	const paths = [
		{
			herf: "/game/play?type=quick",
			value: "Quick Match",
			description: "A fast-paced single-player match.",
			people: 1,
			time: 60,
		},
		{
			herf: "/game/play?type=duel",
			value: "Duel Match",
			description:
				"A competitive two-player match where players face off head-to-head.",
			people: 2,
			time: 60,
		},
	];
</script>

<main>
	<div class="flex flex-col gap-10 justify-center items-center m-10">
		{#each paths as path}
			<div
				class="group/item flex flex-col gap-10 justify-center items-center m-10"
			>
				<div class="w-full sm:w-fit h-fit hover:bg-white hover:text-black">
					<a href={path.herf} class="text-7xl">
						<h2>{path.value}</h2>
					</a>
				</div>

				<span
					class="group/edit group-hover/item:visible text-2xl"
					style="color: #fffff50;">{path.description}</span
				>
			</div>
		{/each}
	</div>
</main>
