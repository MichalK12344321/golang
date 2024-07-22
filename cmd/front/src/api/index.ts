
import axios, { AxiosError, AxiosInstance } from 'axios';
import * as rax from 'retry-axios';

const instance : AxiosInstance = axios.create({
  raxConfig:{
    onRetryAttempt: (error: AxiosError) => console.error(error)
  },
  transformResponse: (data) => {
    return data;
  },
});
rax.attach(instance);

export const axiosInstance = instance;
export * from './client/cdg/client.generated'
export * from './client/scheduler/client.generated'

