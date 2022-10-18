import Fuse from 'fuse.js'
import { Company } from '../types/Company'
import { database } from '../database'
import { useState } from 'react'
import { getData } from '../api/serverHttp'

let fuse: Fuse<Company> | null = null

export const useCompanySearch = () => {
  const [results, setResults] = useState<Company[]>([])
  const [searching, setSearching] = useState(false)
  const [hasResults, setHasResults] = useState(false)

  const search = async (query: string) => {
    setSearching(true)
    if (!fuse) {
      const isEmpty = await database.companies.count().then(c => c === 0)

      if (isEmpty) {
        const companies = await getData<Company[]>('/companies')

        await database.companies.bulkPut(companies)
      }

      fuse = new Fuse<Company>(await database.companies.toArray(), {
        keys: ['name', 'symbol', 'cik'],
      })
    }

    const _results = fuse.search(query, {
      limit: 20,
    }).map(({ item }) => item)

    await database.companiesSearchResults.clear()
    await database.companiesSearchResults.bulkPut(_results)
    await database.companiesPreviousSearch.clear()
    await database.companiesPreviousSearch.put(query)

    setResults(_results)
    setHasResults(_results.length > 0)
    setSearching(false)
  }

  return {
    search,
    searching,
    hasResults,
    results,
  }
}