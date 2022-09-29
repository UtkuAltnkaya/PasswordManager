import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost:8080/',
  withCredentials: true,
});

export const RequestToServer = async ({ ...option }: any) => {
  if (option?.method === 'POST' || option?.method === 'DELETE')
    api.defaults.headers.common['X-Csrf-Token'] = document.cookie.split('csrf_=')[1];

  const onSuccess = (response: any) => response.data;
  const onError = (error: any) => Promise.reject(error.response.data);

  try {
    const response = await api(option);
    return onSuccess(response);
  } catch (error) {
    return onError(error);
  }
};

export default api;
