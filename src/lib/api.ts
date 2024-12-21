import { type Unit } from './types'
import { PUBLIC_API_ENDPOINT_BASE } from '$env/static/public'

export const fetchSlides = async (): Promise<Unit[]> => {
  type Response = {
    slides: Unit[],
  }
  const res = await fetch(`${PUBLIC_API_ENDPOINT_BASE}/api/dotslide`, {
    headers: {
      Accept: 'application/json',
    },
  })

  const resbody = await res.json() as Response
  const slides = resbody.slides

  return slides
}
