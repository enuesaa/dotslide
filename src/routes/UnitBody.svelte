<script lang="ts">
	import type { Unit } from '$lib/types'
	import { createTreeViewCtl, createViewing, type TreeData } from '$lib/tree'
	import CodeViewer from './CodeViewer.svelte'
	import UnitDescription from './UnitDescription.svelte'
	import UnitLinks from './UnitLinks.svelte'
	import UnitImage from './UnitImage.svelte'
	import UnitTerminal from './UnitTerminal.svelte'
	import UnitSep from './UnitSep.svelte'
	import UnitTitle from './UnitTitle.svelte'
	import UnitCap from './UnitCap.svelte'

	createTreeViewCtl()
	createViewing()

	export let unit: Unit
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

	<!-- {#if unit.links !== undefined}
		<UnitLinks links={unit.links} />
	{/if} -->

	{#if unit?.image !== undefined}
		<UnitImage image={unit.image} />
	{/if}

	<!-- <div class="rounded-lg overflow-hidden mt-1">
		{#if unit.open !== undefined}
			<UnitSep text="</>" treeData={files} enableDownloader />
			<CodeViewer treeData={files} firstOpen={unit.open} />
		{/if}

		{#if unit.console !== undefined}
			<UnitSep text="ターミナル" />
			<UnitTerminal content={unit.console} />
		{/if}
	</div> -->
</section>

<style lang="postcss">
	section {
		@apply w-full relative mt-1 overflow-hidden;
	}
</style>
