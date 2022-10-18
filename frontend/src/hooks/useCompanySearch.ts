import Fuse from 'fuse.js'
import { Company } from '../types/Company'
import { database, getQuery, storeQuery } from '../database'
import { useEffect, useState } from 'react'

let fuse: Fuse<Company> | null = null

export const useCompanySearch = () => {
  const [results, setResults] = useState<Company[]>([])
  const [searching, setSearching] = useState(false)
  const [hasResults, setHasResults] = useState(false)
  const [cachedQuery, setCachedQuery] = useState('')

  useEffect(() => {
    void (async () => {
      const q = await getQuery()

      setCachedQuery(q)

      if (q) {
        await search(q)
      }
    })()
  }, [])

  const search = async (q: string) => {
    setSearching(true)
    if (!fuse) {
      fuse = new Fuse<Company>(await database.companies.toArray(), {
        keys: ['name', 'symbol', 'cik'],
      })
    }

    const _results = fuse.search(q, {
      limit: 20,
    }).map(({ item }) => item)

    await storeQuery(q)
    setCachedQuery(q)
    setResults(_results)
    setHasResults(_results.length > 0)
    setSearching(false)
  }

  return {
    search,
    searching,
    hasResults,
    results,
    cachedQuery,
  }
}