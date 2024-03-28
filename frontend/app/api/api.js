"use server";

import { redirect } from "next/navigation.js";

const NEXT_PUBLIC_API_URL = `http://localhost:8080`;

export async function getSession() {
  const response = await fetch(`${NEXT_PUBLIC_API_URL}/auth/session`, {
    method: "GET",
    credentials: "include",
  });
  return response.json();
}
export async function loginUser(state, formData) {
  let email = formData.get("email");
  let password = formData.get("password");
  try {
    const response = await fetch(`${NEXT_PUBLIC_API_URL}/auth/signin`, {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ email, password }),
    });
    if (response.status === 401) {
      return { error: "Invalid email or password." };
    }
    if (response.ok) {
      return { error: "ok" };
    }
  } catch (error) {
    console.error(error);
    return { error: "An error occurred. Please try again." };
  }
}
export async function logoutUser() {
  try {
    const response = await fetch(`${NEXT_PUBLIC_API_URL}/auth/signout`, {
      method: "GET",
      Autorisation: localStorage.getItem("token"),
      credentials: "include",
    });
    if (response.ok) {
      redirect("/home");
    }
    return response.json();
  } catch (error) {
    console.error(error);
    return { error: "Not authorize !!!" };
  }
}

export async function registerUser(state, formData) {
  let data = {
    first_name: formData.get("firstname"),
    last_name: formData.get("lastname"),
    birth_date: formData.get("dateofbirth"),
    avatar: formData.get("avatar"),
    user_name: formData.get("username"),
    email: formData.get("email"),
    password: formData.get("password"),
    about_me: formData.get("aboutme"),
  };
  console.log(data);
  try {
    const response = await fetch(`${NEXT_PUBLIC_API_URL}/auth/signup`, {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    });
    if (response.status === 401) {
      return { error: "Invalid credentials." };
    }
    if (response.ok) {
      console.log("registerUser ok");
      return { error: "ok" };
    }
  } catch (error) {
    console.log(error);
    return { error: "An error occurred. Please try again." };
  }
}

export async function getProfile() {
  const response = await fetch(`${NEXT_PUBLIC_API_URL}/auth/profile`, {
    method: "GET",
    credentials: "include",
  });
  return response.json();
}

export async function getUserByToken(token) {
  const response = await fetch(`${NEXT_PUBLIC_API_URL}/auth/user`, {
    method: "GET",
    headers: {
      Authorisation: localStorage.getItem("token"),
    },
  });
  return response.json();
}
