<script lang="ts">
	import { goto } from '$app/navigation'
	import { ArrowLeft, ArrowRight } from 'lucide-svelte'

	export let current = 0
	export let end = 0

	function handleKeyDown(e: KeyboardEvent) {
		if (e.key === 'ArrowLeft') {
			if (current > 0) {
				goto(`?slide=${current - 1}`)
			}
        } else if (e.key === 'ArrowRight') {
			if (current < end) {
				goto(`?slide=${current + 1}`)
			}
        } 
	}
</script>

<svelte:window on:keydown={handleKeyDown} />

<div class="flex fixed bottom-3 right-3">
	{#if current > 0}
		<a href={`?slide=${current - 1}`} class="block w-1/2">
			<ArrowLeft />
		</a>
	{:else}
		<div class="w-1/2 opacity-10">
			<ArrowLeft />
		</div>
	{/if}
	{#if current < end}
		<a href={`?slide=${current + 1}`} class="block w-1/2">
			<ArrowRight />
		</a>
	{:else}
		<div class="w-1/2 opacity-10">
			<ArrowRight />
		</div>
	{/if}
</div>

