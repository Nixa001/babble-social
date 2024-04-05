"use client";
import { useSession } from "@/app/api/api.js";
import Profile from "@/app/ui/components/profile.cp/profile.cp.js";
import { QueryClient, QueryClientProvider } from "react-query";
const queryClient = new QueryClient();
const page = () => {
  const { session, errSess } = useSession();
  const sessionId = session?.session["user_id"];
  console.log("i got session in group => ", sessionId);
  return (
    <QueryClientProvider client={queryClient}>
      <Profile sessionId={sessionId} />
    </QueryClientProvider>
  );
};

export default page;
