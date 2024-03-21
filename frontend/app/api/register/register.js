import { NEXT_PUBLIC_API_URL } from "../api.js";

export async function registerUser(data) {
    const response = await fetch(`${NEXT_PUBLIC_API_URL}/auth/register`, {
        method: 'POST',
        credentials: 'include',
        headers: {
        'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    });
    return response.json();
}