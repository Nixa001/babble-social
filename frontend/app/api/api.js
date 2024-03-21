
/* In a browser environment,`location.host` returns the host name and port number of the current URL.*/
export const NEXT_PUBLIC_API_URL =`localhost:8080`; ;



export async function getSession() {
  const response = await fetch(`${NEXT_PUBLIC_API_URL}/auth/session`, {
    method: 'GET',
    credentials: 'include',
  });
  return response.json();
}