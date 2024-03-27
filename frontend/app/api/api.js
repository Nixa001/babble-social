"use server"

import { redirect } from "next/navigation.js";

const NEXT_PUBLIC_API_URL = `localhost:8080`;;

export async function getSession() {
  const response = await fetch(`${NEXT_PUBLIC_API_URL}/auth/session`, {
    method: 'GET',
    credentials: 'include',
  });
  return response.json();
}
export async function loginUser(state, formData) {
    console.log("loginUser");
  console.log(formData);
console.log(state);
    let email = formData.get("email");
    let password = formData.get("password");
    console.log(email);
  console.log(password);
  try {
    const response = await fetch(`localhost:8080/auth/signin`, {
        method: 'POST',
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
    });
    if (response.status === 401) {
      return { error: 'Invalid email or password.' };
    }
    if (response.ok) {
      redirect('/home');
    }
  }
  catch (error) {
    console.error(error);
    return { error: 'An error occurred. Please try again.' };
  }
}
export async function logoutUser() {
    const response = await fetch(`localhost:8080/auth/signout`, {
      method: 'GET',
      Autorisation: localStorage.getItem('token'),
        credentials: 'include',
    });
    return response.json();
}

export async function registerUser(data) {
    const response = await fetch(`localhost:8080/auth/signup`, {
        method: 'POST',
        credentials: 'include',
        headers: {
        'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    });
    return response.json();
}