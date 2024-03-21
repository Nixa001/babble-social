import { NEXT_PUBLIC_API_URL } from "../api.js";

export async function loginUser(data) {
    const response = await fetch(`${NEXT_PUBLIC_API_URL}/auth/login`, {
        method: 'POST',
        credentials: 'include',
        headers: {
        'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    });
    return response.json();
}

export async function logoutUser() {
    const response = await fetch(`${NEXT_PUBLIC_API_URL}/auth/logout`, {
        method: 'GET',
        credentials: 'include',
    });
    return response.json();
}