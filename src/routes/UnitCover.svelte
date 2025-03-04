<script lang="ts">
	import type { Unit } from '$lib/types'

	export let unit: Unit
</script>

<div class="font-semibold mt-32 mb-3 font-zenmaru text-6xl text-left pl-2 block w-full whitespace-pre-wrap">
	{unit.cover}
</div>

<style lang="postcss">
	div {
		line-height: 1.3;
	}
</style>