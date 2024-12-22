<script lang="ts">
	import Unit from './Unit.svelte'
	import SlideNextButton from './SlideNextButton.svelte'
	import SlideFullScreenButton from './SlideFullScreenButton.svelte'
	import SlideEnd from './SlideEnd.svelte'
	import { page } from '$app/stores'

	import { fetchSlides } from '$lib/api'
	import { createQuery } from '@tanstack/svelte-query'

	let slideNumber = 0
	$: slideNumber = Number($page.url.searchParams.get('slide'))

	const slides = createQuery({
		queryKey: ['slides'],
		queryFn: fetchSlides,
	})
</script>

{#if $slides.isSuccess}
	{#each $slides.data as unit, i}
		{#if i === slideNumber}
			<Unit {unit} />
		{/if}
	{/each}

	{#if slideNumber >= $slides.data.length}
		<SlideEnd />
	{/if}

	<SlideNextButton current={slideNumber} end={$slides.data.length} />
	<SlideFullScreenButton />
{/if}
