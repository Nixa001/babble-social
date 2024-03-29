"use client";
import { useEffect, useState } from "react";
const NEXT_PUBLIC_API_URL = `http://localhost:8080`;

export async function getSession() {
  return fetch(`${NEXT_PUBLIC_API_URL}/auth/session`, {
    method: "GET",
    credentials: "include",
  })
    .then((response) => response.json())
    .catch((error) => {
      console.error(error);
      throw new Error("An error occurred while fetching session data.");
    });
}

export async function loginUser(email, password) {
  return fetch(`${NEXT_PUBLIC_API_URL}/auth/signin`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ email, password }),
  })
    .then((response) => {
      if (response.status === 401) {
        return { error: "Invalid email or password." };
      }
      if (response.ok) {
        return response.json().then((json) => {
          console.log(json);
          return { error: null, data: json };
        });
      }
      throw new Error("An error occurred.");
    })
    .catch((error) => {
      console.error(error);
      return { error: "An error occurred. Please try again." };
    });
}
export async function logoutUser() {
  let token = localStorage.getItem("token") || "none";

  return fetch(`${NEXT_PUBLIC_API_URL}/auth/signout`, {
    method: "DELETE",
    headers: {
      Authorization: JSON.stringify({ token }),
      accept: "application/json",
    },
  })
    .then((response) => response.json())
    .catch((error) => {
      console.error(error);
      return { error: "Not authorized!!!" };
    });
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

  return fetch(`${NEXT_PUBLIC_API_URL}/auth/signup`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  })
    .then((response) => {
      if (response.status === 401) {
        return { error: "Invalid credentials." };
      }
      if (response.ok) {
        return response.json().then((json) => ({ error: null, data: json }));
      }
      throw new Error("An error occurred.");
    })
    .catch((error) => {
      console.error(error);
      return { error: "An error occurred. Please try again." };
    });
}

export async function getProfile() {
  return fetch(`${NEXT_PUBLIC_API_URL}/auth/profile`, {
    method: "GET",
    headers: {
      Authorization: localStorage.getItem("token"),
    },
  })
    .then((response) => response.json())
    .catch((error) => {
      console.error(error);
      throw new Error("An error occurred while fetching profile data.");
    });
}

export function useSession() {
  const [session, setSession] = useState(null);
  const [error, setError] = useState(null);

  useEffect(() => {
    const token = localStorage.getItem("token") || null;

    async function fetchSessionData() {
      try {
        const response = await fetch(`${NEXT_PUBLIC_API_URL}/auth/session`, {
          method: "GET",
          headers: {
            Authorization: token,
          },
        });

        if (!response.ok) {
          throw new Error(
            `Failed to fetch session data. Status: ${response.status}`
          );
        }

        const sessionData = await response.json();
        setSession(sessionData);
      } catch (error) {
        console.error(error);
        setError(error);
      }
    }

    fetchSessionData();
  }, []);

  return { session, error };
}
