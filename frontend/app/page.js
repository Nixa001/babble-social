import Image from "next/image";
import Login from "./ui/login/login";
import { Landing } from "./ui/components/landingpage/landing-page";
import { ToastContainer } from "react-toastify";
export default function Home() {
  return (
    <div>
      <Landing />
    </div>
  );
}
