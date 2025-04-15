<script lang="ts">
	import type { Unit } from '$lib/types'
	import { SquareChevronDown } from 'lucide-svelte'

	export let unit: Unit

</script>

<div class="font-semibold text-3xl mt-5 mb-3 pl-2 font-zenmaru flex gap-[3px] text-left">
	<span class="text-3xl leading-[0.8] w-7">
		<SquareChevronDown class="inline align-middle" />
	</span>
	<span class="break-all inline-block">{unit.title}</span>
</div>

