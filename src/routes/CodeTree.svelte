<script lang="ts">
	import { getTreeViewCtl, type TreeFile } from '$lib/tree'
	import { type UnitFile } from '$lib/types'
	import { melt } from '@melt-ui/svelte'
	import CodeTreeItemButton from './CodeTreeItemButton.svelte'

	export let files: UnitFile[]
	export let prefix: string = ''
	export let isChild: boolean = false
	const { group } = getTreeViewCtl()

	let filtered: TreeFile[] = []
	$: {
		filtered = []
		for (const file of files) {
			if (file.filename.startsWith(prefix)) {
				const rest = file.filename.slice(prefix.length)
				if (rest.includes('/')) {
					const dirname = rest.split('/')[0]
					if (filtered.filter(f => f.title === dirname).length > 0) {
						continue
					}
					filtered.push({
						filename: `${prefix}${dirname}/`,
						code: '',
						title: dirname,
						isDir: true,
					})
				} else {
					filtered.push({
						...file,
						title: rest,
						isDir: false,
					})
				}
			}
		}
	}
</script>

{#each filtered as data}
	<li class="block pl-2">
		{#if isChild}
			<span class="inline-block opacity-50">└</span>
		{/if}
		<CodeTreeItemButton {data} />

		{#if data.isDir}
			<ul use:melt={$group({ id: data.filename })}>
				<svelte:self {files} prefix={data.filename} isChild />
			</ul>
		{/if}
	</li>
{/each}
