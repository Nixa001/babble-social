"use client";
import { getSession } from "@/app/api/api.js";
import Profile from "@/app/ui/components/profile.cp/profile.cp.js";
import { QueryClient, QueryClientProvider } from "react-query";
const queryClient = new QueryClient();
const page = () => {
  const { session, errSess } = getSession();
  console.log("session => ", session);
  console.log("errSess => ", errSess);
  const sessionId = session?.session["user_id"];
  console.log("i got session in group => ", sessionId);
  return (
    <QueryClientProvider client={queryClient}>
      <Profile sessionId={sessionId} />
    </QueryClientProvider>
  );
};

export default page;
