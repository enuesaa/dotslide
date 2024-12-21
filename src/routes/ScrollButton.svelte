<script lang="ts">
	import { onMount } from 'svelte'

	let currentIndex = 0
	let screenHeight = 0

	onMount(() => {
		screenHeight = window.innerHeight
	})

	function handleClick() {
		currentIndex++;
		window.scrollTo({
			top: screenHeight * currentIndex,
			behavior: 'smooth'
		})
	}
</script>

<button on:click|preventDefault={handleClick} class="fixed bottom-3 right-3">
	scroll
</button>
