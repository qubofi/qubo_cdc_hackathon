import React from 'react';
import axios from 'axios';

export enum RequestStatus {
    IDLE = 'idle',
    LOADING = 'loading',
    ERROR = 'error',
    SUCCESS = 'success',
}

export interface RequestProps {
    url: string;
    method: 'get' | 'post' | 'put' | 'delete';
    initialData?: any;
    initialState?: RequestStatus;
}

function useRequest(props: RequestProps) {
  const { 
    url, 
    method, 
    initialData = [], 
    initialState = RequestStatus.IDLE
} = props;
  const [data, setData] = React.useState<any>(initialData);
  const [error, setError] = React.useState<string>('');
  const [status, setStatus] = React.useState<RequestStatus>(initialState);

  const internalFetch = React.useCallback(async () => {
    setStatus(RequestStatus.LOADING);
    try {
        const response = await axios({
            url,
            method,
        });
        setData(response.data);
        setStatus(RequestStatus.SUCCESS);
    } catch (error: any) {
        setError(error);
        setStatus(RequestStatus.ERROR);
    }
  }, [url, method]);

  React.useEffect(() => {
    internalFetch();
  }, [url, method]);

  return { data, error, status };
}

export default useRequest;