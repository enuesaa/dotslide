<script lang="ts">
	export let description: string
	$: paragraphs = description.split(/\n{2,}/)
</script>

<section>
	{#each paragraphs as paragraph}
		<p>{paragraph}</p>
	{/each}
</section>

<style lang="postcss">
	section {
		@apply mt-1 mb-2 px-2 font-zenmaru text-4xl w-full;
	}
	p {
		letter-spacing: 0.1px;
		line-height: 1.3;
		overflow-wrap: anywhere;
		word-break: normal;
		line-break: strict;
		white-space: pre-wrap;
		font-feature-settings: 'palt';
		@apply font-medium font-zenmaru text-blackgray/80;
		margin-bottom: 0.6em;
	}
	p:last-child {
		margin-bottom: 0;
	}
</style>