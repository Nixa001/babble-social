"use client";
import React from "react";
import { QueryClientProvider } from "react-query";
import { queryClient } from "./groups/page";
import HomePage from "../ui/home/page";
import { useSession } from "../api/api";
import { useRouter } from "next/navigation";

const Page = () => {
  const router = useRouter(),
    { session, errSess } = useSession();
  if (errSess) alert(errSess);
  const token = localStorage.getItem("token") || null;
  if (token == null) router.push("/");
  const sessionId = session?.session["user_id"];
  return (
    sessionId != null && (
      <div className="">
        <QueryClientProvider client={queryClient}>
          <HomePage id={sessionId} />
        </QueryClientProvider>
      </div>
    )
  );
};

export default Page;
