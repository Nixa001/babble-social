"use client";
import { useSession } from "@/app/api/api";
import Group from "@/app/ui/home/groups/group.js/group";
import { useRouter } from "next/navigation";
import React from "react";
import { QueryClient } from "react-query";
import { QueryClientProvider } from "react-query";
const queryClient = new QueryClient();

const page = () => {
  const { session, errSess } = useSession();
  const sessionId = session?.session["user_id"];
  // console.log("i got session in group => ", sessionId);
  return (
    <>
      <QueryClientProvider client={queryClient}>
        <Group sessionID={sessionId} />
      </QueryClientProvider>
    </>
  );
};

export default page;
