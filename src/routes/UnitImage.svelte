<script lang="ts">
	import { PUBLIC_API_ENDPOINT_BASE } from '$env/static/public'

	export let image: string
</script>

<img src={`${PUBLIC_API_ENDPOINT_BASE}/storage/${image}`} alt={image} />

<style lang="postcss">
	img {
		@apply block mx-auto max-w-[min(100%,1000px)] mb-3;
	}
</style>
