"use client";
import React from "react";
import HomePage from "../ui/home/page";

const newWS = new WebSocket("ws://localhost:8080/socket");
export const websocketProvider = React.createContext(newWS);
newWS.onopen = () => {
  console.log("open");
  // send data to server when connected
  newWS.send(JSON.stringify({ type: "GET", request: "get10post" }));
};
const Page = () => {
  newWS.onerror = (err) => console.error(err);
  newWS.onmessage = (msg) => {
    console.log("here is problem");
    // setData(JSON.parse(msg.data));
    console.log(`data is here => : ${msg.data}`);
  };

  return (
    <div className="">
      {/* <QueryClientProvider client={queryClient}> */}
        <websocketProvider>
          <HomePage data="this is a test" />
        </websocketProvider>
      {/* </QueryClientProvider> */}

    </div>
  );
};

export default Page;
