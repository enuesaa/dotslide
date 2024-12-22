<script lang="ts">
	import Code from './Code.svelte'
	import UnitSep from './UnitSep.svelte'

	export let content: string
</script>

<div class="overflow-hidden rounded-lg">
	<UnitSep text="</>" />
	<Code language="console" code={content} />
</div>
