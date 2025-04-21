
export interface FetchInterceptorOptions extends RequestInit {}

export async function fetchWithInterceptor<T = any>(
  url: string,
  options: FetchInterceptorOptions = {}
): Promise<T> {
  let token: string | null = null;

  if (typeof window !== 'undefined') {
    const auth_store_string = localStorage.getItem('auth_store');
    const auth_store = auth_store_string ? JSON.parse(auth_store_string) : null;

    if(auth_store !== null) {
      token = auth_store?.state?.token
    }
  }

  const modifiedOptions: RequestInit = {
    ...options,
    headers: {
      ...options.headers,
      Authorization: token ? `Bearer ${token}` : '',
      'Content-Type': 'application/json',
    },
  };

  try {
    const response = await fetch(url, modifiedOptions);

    const data: T = await response.json();
    return data;
  } catch (error: any) {
    console.error('Fetch Error:', error.message);
    throw error;
  }
}
