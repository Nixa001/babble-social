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
  return fetch(`${NEXT_PUBLIC_API_URL}/auth/signout`, {
    method: "DELETE",
    headers: {
      Authorization: token,
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
export function useSession() {
  const [session, setSession] = useState(null);
  const [error, setError] = useState(null);

  useEffect(() => {
    const token = localStorage.getItem("token") || null;
    console.log("Token in useSession", token);
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

export async function getProfileById(id, sessionId, token) {
  try {
    const response = await fetch(
      `${NEXT_PUBLIC_API_URL}/profile/user?id=${id}&sessionId=${sessionId}`,
      {
        method: "GET",
        headers: {
          Authorization: token,
        },
      }
    );
    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Erreur ", error);
    return Promise.reject(error);
  }
}

export async function followUser(id, sessionId, token) {
  try {
    const response = await fetch(`${NEXT_PUBLIC_API_URL}/follow?id=${id}`, {
      method: "POST",
      headers: {
        Authorization: token,
      },
      body: JSON.stringify({
        followed_id: String(id),
        follower_id: String(sessionId),
      }),
    });
    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Erreur ", error);
    return Promise.reject(error);
  }
}

export async function unfollowUser(id, sessionId, token) {
  try {
    const response = await fetch(`${NEXT_PUBLIC_API_URL}/unfollow?id=${id}`, {
      method: "POST",
      headers: {
        Authorization: token,
      },
      body: JSON.stringify({ id, sessionId }),
    });
    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Erreur ", error);
    return Promise.reject(error);
  }
}

export async function profileTypeUser(sessionId, user_type, token) {
  try {
    const response = await fetch(`${NEXT_PUBLIC_API_URL}/profile/type`, {
      method: "POST",
      headers: {
        Authorization: token,
      },
      body: JSON.stringify({ sessionId: String(sessionId), user_type }),
    });
    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Erreur ", error);
    return Promise.reject(error);
  }
}
