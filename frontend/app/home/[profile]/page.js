"use client";
import { useSession } from "@/app/api/api.js";
import Profile from "@/app/ui/components/profile.cp/profile.cp.js";
import { QueryClient, QueryClientProvider } from "react-query";
const queryClient = new QueryClient();
const page = () => {
  const { session, errSess } = useSession();
  console.log("session => ", session);
  console.log("errSess => ", errSess);
  const sessionId = session?.session["user_id"];
  const sessionToken = session?.session["token"];
  console.log("i got session in group => ", sessionId);
  return (
    <QueryClientProvider client={queryClient}>
      <Profile sessionId={sessionId} sessionToken={sessionToken} />
    </QueryClientProvider>
  );
};

export default page;
