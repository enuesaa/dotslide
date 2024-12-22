<script lang="ts">
	import { getTreeViewCtl } from '$lib/tree'
	import { type UnitFile } from '$lib/types'
	import { melt } from '@melt-ui/svelte'
	import CodeTreeItemButton from './CodeTreeItemButton.svelte'

	export let files: UnitFile[]
	// export let isChild: boolean = false
	const { group } = getTreeViewCtl()

</script>

{#each files as data}
	<li class="block pl-2">
		<!-- {#if isChild}
			<span class="inline-block opacity-50">└</span>
		{/if} -->
		<CodeTreeItemButton {data} />

		<!-- {#if data.children.length > 0}
			<ul use:melt={$group({ id: data.id })}>
				<svelte:self treeData={data.children} isChild />
			</ul>
		{/if} -->
	</li>
{/each}
