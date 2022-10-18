import { useEffect, useState } from 'react'
import { useAsync, useDebounce } from 'react-use'
import { useCompanySearch } from '../../hooks/useCompanySearch'
import { getQuery } from '../../database'

export default function Dashboard () {
  const [query, setQuery] = useState('')
  const [debouncedQuery, setDebouncedQuery] = useState('')

  const { search, results, searching, hasResults } = useCompanySearch()

  useEffect(() => {
    void (async () => {
      const q = await getQuery()
      setQuery(q)
    })()
  }, [])

  useDebounce(() => {
    setDebouncedQuery(query)
  }, 1_000, [query])

  useAsync(async () => {
    if (debouncedQuery) {
      await search(debouncedQuery)
    }
  }, [debouncedQuery])

  return (
    <div className={'mx-4 mt-4'}>
      <div id={'search'}>
        <input
          className={'input input-bordered w-full'}
          placeholder={'Search for a company...'}
          value={query}
          onChange={e => setQuery(e.currentTarget.value)}
        ></input>
      </div>
      <div>
        {searching && (
          <div className={'w-full font-bold text-3xl flex justify-center mt-10 opacity-30'}>Searching...</div>)}
        {!searching && !hasResults && (
          <div className={'w-full font-bold text-3xl flex justify-center mt-10 opacity-30'}>Nothing found</div>)}
        {!searching && hasResults && (
          <table className={'table table-compact hover mt-5 w-full cursor-pointer'}>
            <thead>
            <tr>
              <th className={'text-left'}>CIK</th>
              <th className={'text-left'}>Symbol</th>
              <th className={'text-left'}>Name</th>
              <th className={'text-left'}>Exchange</th>
              <th className={'text-left'}>Office</th>
              <th className={'text-left'}>Industry</th>
            </tr>
            </thead>
            <tbody>
            {results.map(r => (
              <tr
                key={r.cik}
                className={'hover'}
              >
                <td>{r.cik}</td>
                <td className={'font-bold'}>{r.symbol}</td>
                <td>{r.name}</td>
                <td>{r.exchange}</td>
                <td>{r.office}</td>
                <td>{r.industry}</td>
              </tr>
            ))}
            </tbody>
          </table>
        )}
      </div>
    </div>
  )
}