<script lang="ts">
	import type { Unit } from '$lib/types'
	import UnitDescription from './UnitDescription.svelte'
	import UnitLinks from './UnitLinks.svelte'
	import UnitImage from './UnitImage.svelte'
	import UnitTitle from './UnitTitle.svelte'
	import UnitCover from './UnitCover.svelte'
	import UnitCap from './UnitCap.svelte'
	import { createTreeViewCtl, createViewing } from '$lib/tree'
	import UnitFiles from './UnitFiles.svelte'
	import UnitTerminal from './UnitTerminal.svelte'

	export let unit: Unit

	createTreeViewCtl()
	createViewing()
</script>

<section>
	{#if unit?.cover !== undefined}
		<UnitCover {unit} />
	{/if}
	{#if unit?.cap !== undefined}
		<UnitCap {unit} />
	{/if}
	{#if unit?.title !== undefined}
		<UnitTitle {unit} />
	{/if}
	{#if unit?.description !== undefined}
		<UnitDescription description={unit.description} />
	{/if}

	{#if unit.links !== undefined}
		<UnitLinks links={unit.links} />
	{/if}

	{#if unit?.image !== undefined}
		<UnitImage image={unit.image} />
	{/if}

	{#if unit?.files !== undefined && unit.files.length > 0}
		<UnitFiles files={unit.files} />
	{/if}

	{#if unit?.terminal !== undefined}
		<UnitTerminal content={unit.terminal} />
	{/if}
</section>

<style lang="postcss">
	section {
		@apply w-full relative mt-1 overflow-hidden;
	}
</style>
