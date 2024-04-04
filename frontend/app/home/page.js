"use client";
import React from "react";
import { QueryClientProvider } from "react-query";
import { queryClient } from "./groups/page";
import HomePage from "../ui/home/page";
import { useSession } from "../api/api";

const Page = () => {
  const { session, errSess } = useSession(),
    sessionId = session?.session["user_id"];
  // console.log("sess is here => ", session, "with err >", errSess);
   console.log("i got session => ", sessionId);
  return (
    <div className="">
      <QueryClientProvider client={queryClient}>
        <HomePage id={sessionId} />
      </QueryClientProvider>
    </div>
  );
};

export default Page;
