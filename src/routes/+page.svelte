<script lang="ts">
	import Unit from './Unit.svelte'

	import { fetchSlides } from '$lib/api'
	import { createQuery } from '@tanstack/svelte-query'

	const slides = createQuery({
		queryKey: ['slides'],
		queryFn: fetchSlides,
	})
</script>

{#if $slides.isSuccess}
	{#each $slides.data as unit}
		<Unit {unit} />
	{/each}
{/if}
