<script lang="ts">
	import UnitNav from './UnitNav.svelte'
	import UnitBody from './UnitBody.svelte'
	import type { Project, Unit } from '$lib/types'
	import type { TreeData } from '$lib/tree'

	export let project: Project
	export let unit: Unit
	export let files: TreeData[]
</script>

<div class="h-screen">
	{#if unit.title != undefined}
		<UnitNav {unit} />
	{/if}
	<UnitBody {project} {unit} {files} />

	{#if unit.left != undefined || unit.right != undefined}
		<section class="flex gap-10">
			<div class="w-1/2">
				{#if unit.left != undefined}
					<svelte:self {project} unit={unit.left} {files} />
				{/if}
			</div>
			<div class="w-1/2">
				{#if unit.right != undefined}
					<svelte:self {project} unit={unit.right} {files} />
				{/if}
			</div>
		</section>
	{/if}
</div>

