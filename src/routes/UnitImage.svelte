<script lang="ts">
	import { PUBLIC_API_ENDPOINT_BASE } from '$env/static/public'

	export let image: string
</script>

{#if image.endsWith('.mp4')}
	<!-- svelte-ignore a11y_media_has_caption -->
	<video src={`${PUBLIC_API_ENDPOINT_BASE}/storage/${image}`} controls playsinline></video>
{:else}
	<img src={`${PUBLIC_API_ENDPOINT_BASE}/storage/${image}`} alt={image} />
{/if}

<style lang="postcss">
	img,
	video {
		@apply block mx-auto max-w-[min(100%,1000px)] mb-3;
	}
</style>
