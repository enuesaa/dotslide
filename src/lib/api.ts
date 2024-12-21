import { type Unit } from './types'

export const fetchSlides = async (): Promise<Unit[]> => {
  const res = await fetch(`http://localhost:3000/api/slide`, {
    headers: {
      Accept: 'application/json',
    },
  })

  const resbody: {slides: Unit[]} = await res.json()
  const slides = resbody.slides

  console.log(slides)

  return slides
}
