<script lang="ts">
	import UnitBody from './UnitBody.svelte'
	import type { Unit } from '$lib/types'

	export let unit: Unit
</script>

<div class="h-screen">
	<UnitBody {unit} />

	{#if unit.center !== undefined}
		<section class="flex gap-10">
			<div class="w-1/3">
				{#if unit.left !== undefined}
					<UnitBody unit={unit.left} />
				{/if}
			</div>
			<div class="w-1/3">
				{#if unit.center !== undefined}
					<UnitBody unit={unit.center} />
				{/if}
			</div>
			<div class="w-1/3">
				{#if unit.right !== undefined}
					<UnitBody unit={unit.right} />
				{/if}
			</div>
		</section>
	{:else if unit.left !== undefined || unit.right !== undefined}
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

	{#if unit.metaCenter !== undefined}
		<section class="flex gap-10">
			<div class="w-1/3">
				{#if unit.metaLeft !== undefined}
					<UnitBody unit={unit.metaLeft} />
				{/if}
			</div>
			<div class="w-1/3">
				{#if unit.metaCenter !== undefined}
					<UnitBody unit={unit.metaCenter} />
				{/if}
			</div>
			<div class="w-1/3">
				{#if unit.metaRight !== undefined}
					<UnitBody unit={unit.metaRight} />
				{/if}
			</div>
		</section>
	{:else if unit.metaLeft !== undefined || unit.metaRight !== undefined}
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

