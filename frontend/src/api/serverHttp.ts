import axios, { AxiosRequestConfig } from 'axios'

const serverHttp = axios.create({
  baseURL: 'http://localhost:3000/api',
})

export const getData = async <T> (url: string, config: AxiosRequestConfig = {}): Promise<T> => {
  return serverHttp.get<T>(url, config).then((response) => response.data)
}