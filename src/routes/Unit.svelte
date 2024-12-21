<script lang="ts">
	import UnitNav from './UnitNav.svelte'
	import UnitBody from './UnitBody.svelte'
	import type { Unit } from '$lib/types'

	export let unit: Unit
</script>

<div class="h-screen">
	{#if unit.title != undefined}
		<UnitNav {unit} />
	{/if}
	<UnitBody {unit} />

	{#if unit.left != undefined || unit.right != undefined}
		<section class="flex gap-10">
			<div class="w-1/2">
				{#if unit.left != undefined}
					<svelte:self unit={unit.left} />
				{/if}
			</div>
			<div class="w-1/2">
				{#if unit.right != undefined}
					<svelte:self unit={unit.right} />
				{/if}
			</div>
		</section>
	{/if}
</div>

