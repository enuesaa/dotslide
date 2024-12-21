<script lang="ts">
	import './app.css'
	import { QueryClient, QueryClientProvider } from '@tanstack/svelte-query'

	const queryClient = new QueryClient()
</script>

<QueryClientProvider client={queryClient}>
	<main class="min-h-screen container mx-auto px-1 py-8">
		<slot />
	</main>
</QueryClientProvider>
