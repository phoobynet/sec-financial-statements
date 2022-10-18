import Dexie from 'dexie'
import { Company } from '../types/Company'

class Database extends Dexie {
  companies!: Dexie.Table<Company>
  companiesPreviousSearch!: Dexie.Table<{ id: string, query: string }>

  constructor () {
    super('SECFinancialStatements')
    this.version(2).stores({
      companies: '++,cik,name,exchange,office,industry,symbol',
      companiesPreviousSearch: '++id,query',
    })
  }
}

export const database = new Database()

export const getQuery = async (): Promise<string> => {
  const query = await database.companiesPreviousSearch.where('id').equals('query').first()

  return query?.query ?? ''
}

export const storeQuery = async (query: string): Promise<void> => {
  await database.companiesPreviousSearch.update('query', { query })
}