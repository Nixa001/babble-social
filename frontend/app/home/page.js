"use client";
import { useState } from "react";
import HomePage from "../ui/home/page";
{
  useState;
}

const Page = () => {
  const props = {
    data: "This is a test",
    title: "Test Title",
  };
  const [data, setData] = useState(props.data);
  const [ws, setWS] = useState(null);

  const newWS = new WebSocket("ws://localhost:8080/socket");
  newWS.onerror = (err) => console.error(err);
  newWS.onopen = () => {
    setWS(newWS);
    // send data to server when connected
    newWS.send(JSON.stringify({ type: "GET", request: "get10post" }));
  };
  newWS.onmessage = (msg) => {
    setData(JSON.parse(msg.data));
    console.log(`data is here => : ${data}`);
  };

  return (
    <div className="">
      <HomePage data="this is a test" />
    </div>
  );
};

export default Page;
