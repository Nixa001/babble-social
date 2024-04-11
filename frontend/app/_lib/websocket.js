import React, { createContext, useState, useEffect } from 'react';

export const WebSocketContext = createContext(null);

export const WebSocketProvider = ({ children }) => {
  //===========
  const [messages, setMessages] = useState([]); // Pour stocker les messages
//===============================
  const [socket, setSocket] = useState(null);
  const [allMessages, setAllMessages] = useState([])

  //==========================================

  const sendMessage = (message) => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      // console.log("Envoie du message");
      // console.log("Hello ", String(message));
      socket.send(JSON.stringify(message));
    } else {
      console.error("WebSocket is not open");
    }
  };

  // Fonction pour récupérer les messages via l'API
  const readMessages = () => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      // console.log("WebSocket is ready");

      // Écoute des messages entrants
      socket.onmessage = (event) => {
        const message = JSON.parse(event.data);
        // console.log("Message received:", message);

        // Exécuter des actions en fonction du message reçu
        // Par exemple, mettre à jour l'état des messages dans le contexte
        setMessages((prevMessages) => [...prevMessages, message]);
      };
    } else {
      console.log("WebSocket is not ready");
    }
  };

  //======================================================


  const sendMessageToServer = (message) => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      socket.send(JSON.stringify(message));
    } else {
      console.error("La connection Websocket n'est pas etablie");
    }
  }
  const resetAllMessages = () => {
    
    setAllMessages([]);
};
  useEffect(() => {
    // console.log(allMessages);
  }, [allMessages])

  const readMessage = (message) => {
    const data = JSON.parse(message);
    // console.log(data.Data);
    if (data.Type === "id-receiver-event") {
      // console.log("===", data.Data);
      setAllMessages(data.Data);
    }
    if (data.Type === "idGroup-receiver-event") {
      setAllMessages(data.Data);
    }
    if (data.Type === "message-user-event") {
      setAllMessages(prevMessages => [...prevMessages, data.Data]);
      // console.log("message-user = ", data.Data);
    }
    if (data.Type === "message-group-event") {
      setAllMessages(prevMessages => [...prevMessages, data.Data]);
    }
  }

  useEffect(() => {
    // Fonction pour vérifier le token et établir la connexion WebSocket
    const checkTokenAndConnect = () => {
      const token = localStorage.getItem("token");
      if (token && !socket) {
        // console.log("token", token);
        const ws = new WebSocket(`ws://localhost:8080/ws?token=${token}`);

        ws.onopen = (event) => {
          console.log("Connexion websocket ouverte", event);
        };

        ws.onmessage = (event) => {
          readMessage(event.data);
          // Vous pouvez gérer les messages ici ou les propager à travers le contexte
        };

        ws.onerror = (error) => {
          console.log("Erreur websocket", error);
        };

        ws.onclose = (event) => {
          console.log("Connexion websocket ferme", event);
        };

        setSocket(ws);
      }
    };

    // Vérifie toutes les secondes si le token est disponible et établit la connexion si nécessaire
    const intervalId = setInterval(checkTokenAndConnect, 1000);

    return () => clearInterval(intervalId);
  }, [socket]);


  return (
    <WebSocketContext.Provider value={{sendMessage, readMessages, messages, socket, readMessage, sendMessageToServer, allMessages, resetAllMessages }}>
      {children}
    </WebSocketContext.Provider>
  );
};