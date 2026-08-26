<script lang="ts">
	import type { Unit } from '$lib/types'
	import { parseAttrs } from '$lib/attrs'

	export let unit: Unit

	$: ({ content: cap, attrs } = parseAttrs(unit.cap ?? ''))
	$: symbol = attrs?.symbol ?? '#'
</script>

<div class="font-semibold mt-5 mb-3 font-zenmaru text-6xl text-left pl-2 block w-full whitespace-pre-wrap">
	{symbol} {cap}
</div>