<script lang="ts">
	import './app.css'
	import { QueryClient, QueryClientProvider } from '@tanstack/svelte-query'

	const queryClient = new QueryClient()
</script>

<QueryClientProvider client={queryClient}>
	<main class="min-h-screen w-[95%] max-w-[1500px] mx-auto pb-8">
		<slot />
	</main>
</QueryClientProvider>
