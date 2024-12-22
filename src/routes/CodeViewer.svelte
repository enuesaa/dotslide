<script lang="ts">
	import { getTreeViewCtl, getViewing } from '$lib/tree'
	import { type UnitFile } from '$lib/types'
	import { onMount } from 'svelte'
	import Code from './Code.svelte'
	import CodeTree from './CodeTree.svelte'

	export let files: UnitFile[]
	let firstOpen = 'README.md'
	$: {
		if (files.length > 0) {
			firstOpen = files[0].filename
		}
	}

	const { tree } = getTreeViewCtl()
	const viewing = getViewing()
	onMount(() => {
		for (const file of files) {
			if (file.filename === firstOpen) {
				viewing.set(file)
				break
			}
		}
	})

	function calcExt(filename: string): string {
		return filename.includes('.') ? filename.split('.').pop()! : 'txt'
	}
</script>

<section>
	<ul class="flex-none pr-2 pt-2 pb-2 border-r-editorsep border-r" {...$tree}>
		<CodeTree {files} />
	</ul>
	<div class="flex-auto">
		{#key $viewing}
			{#if $viewing !== undefined}
				<Code language={calcExt($viewing.filename)} code={$viewing.code} />
			{/if}
		{/key}
	</div>
</section>

<style lang="postcss">
	section {
		@apply bg-editorbg text-editortext text-base;
		@apply flex overflow-y-scroll min-h-12;
	}
</style>
