// Page utils.js
// import React, { createContext, useContext, useEffect, useState } from "react";

// // Créer un contexte
// const ApiContext = createContext();

// // Créer un composant fournisseur pour encapsuler vos fonctions de lecture/écriture
// export const ApiProvider = ({ children }) => {
//   const [messages, setMessages] = useState([]); // Pour stocker les messages
//   const wsUrl = "ws://localhost:8080/ws";
//   const [socket, setSocket] = useState(null);

//   useEffect(() => {
//     const newSocket = new WebSocket(wsUrl);

//     newSocket.onopen = () => {
//       console.log("Socket is open");
//     };

//     newSocket.onclose = () => {
//       console.log("Socket is closed");
//     };

//     setSocket(newSocket);

//     // Nettoyage de la connexion WebSocket lors du démontage du composant
//     return () => {
//       if (newSocket) {
//         newSocket.close();
//       }
//     };
//   }, []);

//   // Fonction pour envoyer un message via l'API
//   const sendMessage = (message) => {
//     if (socket && socket.readyState === WebSocket.OPEN) {
//       console.log("Envoie du message");
//       console.log("Hello ", String(message));
//       socket.send(JSON.stringify(message));
//     } else {
//       console.error("WebSocket is not open");
//     }
//   };

//   // Fonction pour récupérer les messages via l'API
//   const readMessages = () => {
//     if (socket && socket.readyState === WebSocket.OPEN) {
//       // console.log("WebSocket is ready");

//       // Écoute des messages entrants
//       socket.onmessage = (event) => {
//         const message = JSON.parse(event.data);
//         // console.log("Message received:", message);

//         // Exécuter des actions en fonction du message reçu
//         // Par exemple, mettre à jour l'état des messages dans le contexte
//         setMessages((prevMessages) => [...prevMessages, message]);
//       };
//     } else {
//       console.log("WebSocket is not ready");
//     }
//   };

// console.log("Messages updated out useEffect:", messages);
// useEffect(() => {
//   // Imprime les messages chaque fois qu'ils changent
//   console.log("Messages updated:", messages);
// }, [messages]); // Surveillance des changements de l'état messages

//   return (
//     <ApiContext.Provider value={{ sendMessage, readMessages, messages }}>
//       {children}
//     </ApiContext.Provider>
//   );
// };

// // Utilisez un hook personnalisé pour accéder au contexte dans vos composants
// export const useApi = () => useContext(ApiContext);

// Exemple d'utilisation de fetch pour envoyer une requête au serveur
export async function getSessionUser() {
  const token = localStorage.getItem("token");
  const response = await fetch(
    "http://localhost:8080/auth/usersessions?token=" + token,
    {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    }
  );

  if (!response.ok) {
    throw new Error("Failed to fetch session");
  }

  const user = await response.json();
  //  console.log("user in session", user); // Afficher les informations de l'utilisateur
  return user; // Retourner les données de l'utilisateur
}
