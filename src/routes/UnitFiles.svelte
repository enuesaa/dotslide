<script lang="ts">
	import type { UnitFile } from '$lib/types'
	import CodeViewer from './CodeViewer.svelte'
	import UnitSep from './UnitSep.svelte'

	export let files: UnitFile[]
</script>

<div class="overflow-hidden rounded-lg">
	<UnitSep text="</>" />
	<CodeViewer {files} />	
</div>
