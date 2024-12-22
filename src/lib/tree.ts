import { createTreeView } from '@melt-ui/svelte'
import { writable, type Writable } from 'svelte/store'
import { setContext, getContext } from 'svelte'
import type { UnitFile } from './types'

export type TreeFile = UnitFile & {
	title: string
	isDir: boolean
}

export const createTreeViewCtl = () => {
	setContext(
		'treeView',
		createTreeView({
			forceVisible: true,
		})
	)
}
export const getTreeViewCtl = () => {
	const treeView = getContext<ReturnType<typeof createTreeView>>('treeView')
	return treeView.elements
}
export const createViewing = () => {
	setContext('viewing', writable<UnitFile | undefined>(undefined))
}
export const getViewing = () => {
	return getContext<Writable<UnitFile | undefined>>('viewing')
}
