<script lang="ts">
	import PageTitle from './PageTitle.svelte'
	import Description from './Description.svelte'
	import PagePublishedBar from './PagePublishedBar.svelte'
	import Unit from './Unit.svelte'

	import { fetchSlides } from '$lib/api'
	import { createQuery } from '@tanstack/svelte-query'

	const slides = createQuery({
		queryKey: ['slides'],
		queryFn: fetchSlides,
	})
</script>

{#if $slides.isSuccess}
	{#each $slides.data as slide}
		<p>タイトル: {slide.cap}</p>
	{/each}
{/if}

<!-- <svelte:head>
	<title>{data.project.title}</title>
	<meta name="description" content={`${data.project.title} | lab.enuesaa.dev`} />
</svelte:head>

<PageTitle title={data.project.title} />
<PagePublishedBar published={data.project.published} />
<Description content={data.project.description} />

{#each data.project.units as unit}
	<Unit project={data.project} {unit} files={data.unitfiles[unit.title]} />
{/each} -->
