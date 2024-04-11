"use client";
import { useEffect, useState } from "react";
const NEXT_PUBLIC_API_URL = `http://localhost:8080`;

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
          localStorage.setItem("token", json.token);
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
export async function logoutUser(token) {
  console.log(token);

  return fetch(`${NEXT_PUBLIC_API_URL}/auth/signout`, {
    method: "DELETE",
    headers: {
      Authorization: JSON.stringify({ token }),
    },
  })
    .then((response) => response.json())
    .catch((error) => {
      console.log(error);
      return { error: "Not authorized!!!" };
    });
}

export async function registerUser(formData) {
  return fetch(`${NEXT_PUBLIC_API_URL}/auth/signup`, {
    method: "POST",
    body: formData,
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
  return fetch(`${NEXT_PUBLIC_API_URL}/profile`, {
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
