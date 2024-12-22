export type Unit = {
  cap?: string
	title?: string
	description?: string
	links?: UnitLink[]
	files?: UnitFile[]
	terminal?: string
	image?: string
	console?: string
	left?: Unit
  center?: Unit
	right?: Unit
	metaLeft?: Unit
  metaCenter?: Unit
  metaRight?: Unit
}

export type UnitLink = {
	title: string
	url: string
}

export type UnitFile = {
	filename: string
	code: string
}
