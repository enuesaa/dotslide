<script lang="ts">
	import UnitBody from './UnitBody.svelte'
	import type { Unit } from '$lib/types'

	export let unit: Unit
</script>

<div class="h-screen">
	<UnitBody {unit} />

	{#if unit.left !== undefined || unit.right !== undefined}
		<section class="flex gap-10">
			<div class="w-1/2">
				{#if unit.left !== undefined}
					<UnitBody unit={unit.left} />
				{/if}
			</div>
			<div class="w-1/2">
				{#if unit.right !== undefined}
					<UnitBody unit={unit.right} />
				{/if}
			</div>
		</section>
	{/if}
	{#if unit.meta !== undefined}
		<UnitBody unit={unit.meta} />
	{/if}
	{#if unit.metaLeft !== undefined || unit.metaRight !== undefined}
		<section class="flex gap-10">
			<div class="w-1/2">
				{#if unit.metaLeft !== undefined}
					<UnitBody unit={unit.metaLeft} />
				{/if}
			</div>
			<div class="w-1/2">
				{#if unit.metaRight !== undefined}
					<UnitBody unit={unit.metaRight} />
				{/if}
			</div>
		</section>
	{/if}
	{#if unit.footer !== undefined}
		<UnitBody unit={unit.footer} />
	{/if}
</div>

