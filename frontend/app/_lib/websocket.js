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
  //  useEffect(()=>{
  //   const token  = localStorage.getItem("token");
  //   console.log("token", token);
  //  },[token])

  const readMessage = (message) => {
    const data = JSON.parse(message);
    console.log(data.Data);
    console.log(data);
    if (data.Type === "id-receiver-event" && data.Data !== null) {
      setAllMessages(prevMessages => [...prevMessages, data.Data]);
    }

  }

  useEffect(() => {
    // Fonction pour vérifier le token et établir la connexion WebSocket
    const checkTokenAndConnect = () => {
       const token = localStorage.getItem("token");
       if (token && !socket) {
         console.log("token", token);
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
    <WebSocketContext.Provider value={{ socket, readMessage, sendMessageToServer, allMessages }}>
      {children}
    </WebSocketContext.Provider>
  );
};