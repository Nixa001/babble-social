// 'use client';
import Image from "next/image";
import Login from "./ui/login/login";
import { Landing } from "./ui/components/landingpage/landing-page";
// import { WebSocketProvider } from "./_lib/websocket";

export default function Home() {
  return (
    // <WebSocketProvider> 
    <div>
       <Landing />
    </div>
    //  </WebSocketProvider>
  );
}
