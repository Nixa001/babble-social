import Image from "next/image";
import Login from "./ui/components/login/page";

export default function Home() {
  
  
  return (
    <div>
      <Login handleOnclick={handleLogin} />{" "}
    </div>
  );
}

  const handleLogin = ()=>{
    alert("Login")
  }