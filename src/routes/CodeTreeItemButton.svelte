<script lang="ts">
	import { melt } from '@melt-ui/svelte'
	import { getTreeViewCtl, getViewing, type TreeFile } from '$lib/tree'

	export let data: TreeFile
	const { item } = getTreeViewCtl()
	const viewing = getViewing()

	function hanldeClick() {
		if (data.isDir) {
			return
		}
		viewing.set(data)
	}
</script>

<button
	use:melt={$item({
		id: data.filename,
		hasChildren: data.isDir,
	})}
	on:click|preventDefault={hanldeClick}
	disabled={data.isDir === true}
	class={$viewing?.filename === data.filename ? 'bg-editorsep/50 border-editortext/50 border-[0.5px]' : ''}
>
	{data.title}
</button>

<style lang="postcss">
	button {
		@apply inline-block py-[1px] mt-[1px] px-1 text-left;
		@apply rounded-sm select-none;
	}
</style>
