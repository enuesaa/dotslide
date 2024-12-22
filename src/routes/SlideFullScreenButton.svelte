<script lang="ts">
	import { onMount } from 'svelte'
	import { Maximize } from 'lucide-svelte'

	let isFullScreen = false;

	function handleClick() {
		document.documentElement.requestFullscreen()
	}

	const onFullScreenChange = () => {
		isFullScreen = document.fullscreenElement !== null
	}

	onMount(() => {
		document.addEventListener('fullscreenchange', onFullScreenChange)

		return () => {
			document.removeEventListener('fullscreenchange', onFullScreenChange)
		}
	})
</script>

{#if isFullScreen === false}
	<button on:click={handleClick} class="fixed top-3 right-3">
		<Maximize />
	</button>
{/if}
