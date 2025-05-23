export type Unit = {
	cover?: string
  cap?: string
	title?: string
	description?: string
	links?: UnitLink[]
	files?: UnitFile[]
	terminal?: string
	image?: string
	console?: string
	left?: Unit
	right?: Unit
	meta?: Unit
	metaLeft?: Unit
  metaRight?: Unit
	footer?: Unit
}

export type UnitLink = {
	title: string
	url: string
}

export type UnitFile = {
	filename: string
	code: string
}
