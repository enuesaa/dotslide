<script lang="ts">
	// see https://github.com/metonym/svelte-highlight/issues/246
	import highlightjs from 'highlight.js/lib/common'
	import 'highlight.js/styles/github-dark-dimmed.min.css'

	export let language: string
	export let code: string

	const hl = highlightjs.getLanguage(language)
		? highlightjs.highlight(code, { language })
		: highlightjs.highlightAuto(code)
</script>

<pre><code class="hljs" style="padding: 7px 12px 12px 12px;">{@html hl.value}</code></pre>

<style lang="postcss">
	pre {
		@apply h-full;
	}
	code {
		@apply h-full;
	}
</style>
