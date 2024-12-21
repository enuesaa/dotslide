import type { TreeData } from './tree'

export type Config = {
	title: string
	description: string
	published: string
	units: Unit[]
}

export type Unit = {
  cap?: string
	title?: string
	description?: string
	image?: string
	console?: string
	left?: Unit
  center?: Unit
	right?: Unit
	metaLeft?: Unit
  metaCenter?: Unit
  metaRight?: Unit
}

export type Project = Config & {
	name: string
}

export type UnitFiles = Record<string, TreeData[]> // per unit
