import Dexie from 'dexie'
import { Company } from '../types/Company'

class Database extends Dexie {
  companies!: Dexie.Table<Company>
  companiesSearchResults!: Dexie.Table<Company>
  companiesPreviousSearch!: Dexie.Table<{ query: string }>

  constructor () {
    super('SECFinancialStatements')
    this.version(2).stores({
      companies: '++,cik,name,exchange,office,industry,symbol',
      companiesSearchResults: '++,cik,name,exchange,office,industry,symbol',
      companiesPreviousSearch: '++,query',
    })
  }
}

export const database = new Database()