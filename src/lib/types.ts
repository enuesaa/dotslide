export type Unit = {
  cap?: string
	title?: string
	description?: string
	links?: {
		title: string
		url: string
	}[]
	image?: string
	console?: string
	left?: Unit
  center?: Unit
	right?: Unit
	metaLeft?: Unit
  metaCenter?: Unit
  metaRight?: Unit
}
