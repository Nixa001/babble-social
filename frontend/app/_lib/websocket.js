import React, { createContext, useState, useEffect } from 'react';

export const WebSocketContext = createContext(null);


export const WebSocketProvider = ({ children }) => {
  const [socket, setSocket] = useState(null);
  const [allMessages, setAllMessages] = useState([])
  

  const sendMessageToServer = (message) => {
    if (socket) {
      socket.send(JSON.stringify(message));
    } else {
      console.error("La connection Websocket n'est pas etablie");
    }
  }
  useEffect(() => {
    console.log(allMessages);
 }, [])

  const readMessage = (message) => {
    const data = JSON.parse(message);
    console.log(data.Data);
    console.log(data.Type);
    if (data.Type === "id-receiver-event" && data.Data !== null){
      setAllMessages(prevMessages => [...prevMessages, data.Data]);
    }
  }

  useEffect(() => {
    // Vérifie si une connexion WebSocket existe déjà
    if (!socket) {
      const ws = new WebSocket('ws://localhost:8080/ws');

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

    return () => {
      // Ferme la connexion WebSocket seulement si elle existe
      if (socket) {
        socket.close();
      }
    };
  }, [socket]);

  return (
    <WebSocketContext.Provider value={{ socket, readMessage, sendMessageToServer, allMessages }}>
      {children}
    </WebSocketContext.Provider>
  );
};