<script lang="ts">
	import { ArrowDown } from 'lucide-svelte'

	let currentIndex = 0

	function handleClick() {
		currentIndex++;
		const windowH = window.innerHeight
	
		window.scrollTo({
			top: windowH * currentIndex + 20,
			behavior: 'smooth'
		})
	}
</script>

<button on:click|preventDefault={handleClick} class="fixed bottom-3 right-3">
	<ArrowDown />
</button>
