'use client'
import React from "react";
import Login from "../ui/login/login";
import { WebSocketProvider } from "../_lib/websocket";

const Page = () => {

  return (
    <WebSocketProvider>
      <>
        <Login />{" "}
      </>
    </WebSocketProvider>
  );
};

export default Page;
