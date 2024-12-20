<script lang="ts">
	import ProjectLink from './ProjectLink.svelte'
	import type { PageData } from './$types'

	export let data: PageData
</script>

{#each data.projects as project}
	<ProjectLink href={`/prototype/${project.name}`} title={project.title} />
{/each}
