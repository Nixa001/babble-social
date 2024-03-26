import { NEXT_PUBLIC_API_URL } from "../api.js";

export async function getProfile() {
    const response = await fetch(`${NEXT_PUBLIC_API_URL}/profile`, {
        method: 'GET',
        credentials: 'include',
    });
    return response.json();
}
