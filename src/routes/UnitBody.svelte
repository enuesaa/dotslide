<script lang="ts">
	import type { Unit } from '$lib/types'
	import UnitDescription from './UnitDescription.svelte'
	import UnitLinks from './UnitLinks.svelte'
	import UnitImage from './UnitImage.svelte'
	import UnitTitle from './UnitTitle.svelte'
	import UnitCap from './UnitCap.svelte'
	import CodeViewer from './CodeViewer.svelte'
	import { createTreeViewCtl, createViewing } from '$lib/tree'
	import UnitSep from './UnitSep.svelte'
	import UnitFiles from './UnitFiles.svelte'

	export let unit: Unit

	createTreeViewCtl()
	createViewing()
</script>

<section>
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

	<!-- 
	{#if unit.console !== undefined}
		<UnitSep text="</>" />
		<UnitTerminal content={unit.console} />
	{/if}
	-->
</section>

<style lang="postcss">
	section {
		@apply w-full relative mt-1 overflow-hidden;
	}
</style>
