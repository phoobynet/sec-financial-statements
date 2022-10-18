import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App'
import './index.css'
import { HashRouter } from 'react-router-dom'
import { database } from './database'
import { getData } from './api/serverHttp'
import { Company } from './types/Company'

async function loadCompanies (): Promise<void> {
  const isEmpty = await database.companies.count().then(c => c === 0)

  if (isEmpty) {
    const companies = await getData<Company[]>('/companies')

    await database.companies.bulkPut(companies)
  }
}

async function loadQuery (): Promise<void> {
  const isEmpty = await database.companiesPreviousSearch.count().then(c => c === 0)

  if (isEmpty) {
    await database.companiesPreviousSearch.put({
      id: 'query',
      query: '',
    })
  }
}

function start () {
  ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
    <React.StrictMode>
      <HashRouter>
        <App />
      </HashRouter>
    </React.StrictMode>,
  )
}

loadCompanies().then(loadQuery).then(start).catch(console.error)


