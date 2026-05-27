<script lang="ts">
	import { PUBLIC_API_ENDPOINT_BASE } from '$env/static/public'
	import { parseAttrs } from '$lib/attrs'

	export let image: string
	const { content: filename, attrs } = parseAttrs(image)
</script>

<div class={attrs?.size === 'large' ? 'mb-3 relative large' : 'mb-3 relative'}>
	{#if filename.endsWith('.mp4')}
		<!-- svelte-ignore a11y_media_has_caption -->
		<video src={`${PUBLIC_API_ENDPOINT_BASE}/storage/${filename}`} controls playsinline></video>
	{:else}
		<img src={`${PUBLIC_API_ENDPOINT_BASE}/storage/${filename}`} alt={filename}/>
	{/if}
</div>

<style lang="postcss">
	img,
	video {
		@apply block mx-auto mb-3 w-full;
		max-width: min(100%, 1000px);
	}

	.large {
		width: 100vw;
		margin-left: calc(50% - 50vw);
	}
	.large img,
	.large video {
		max-width: min(100vw, 1300px);
	}
</style>
