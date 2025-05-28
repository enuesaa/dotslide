<script lang="ts">
	export let description: string
</script>

<section>
	<pre>{description}</pre>
</section>

<style lang="postcss">
	section {
		@apply mt-1 mb-2 px-2 font-zenmaru text-3xl w-full;
	}
	section pre {
		letter-spacing: 0.1px;
		line-height: 1.5;
		overflow-wrap: anywhere;
		word-break: normal;
		line-break: strict;
		white-space: pre-wrap;
		font-feature-settings: 'palt';
		@apply font-medium font-zenmaru text-blackgray/80;
	}
</style>
