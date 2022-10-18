import { useParams } from 'react-router-dom'

export default function Company () {
  const params = useParams<{ symbol: string }>()
  return (
    <div>
      Company {params.symbol}
    </div>
  )
}