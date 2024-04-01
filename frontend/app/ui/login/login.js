"use client";
import Image from "next/image";
import Link from "next/link";
import Button from "../components/button/button";

export default function Login() {
  const handleLogin = () => {
    alert("Login");
  };

  return (
    <div className="w-screen h-screen flex">
      <div className="flex flex-col items-center w-full sm:w-6/12">
        <div className="header_login flex items-center justify-around w-full">
          <Image src="/assets/logo.png" alt="logo" width={100} height={100} />
          <div>
            Don&apos;t have an account ?
            <Link
              href="/register"
              className="text-primary hover:text-second cursor-pointer"
            >
              Sign Up.{" "}
            </Link>
          </div>
        </div>
        <div className="mt-40 sm:mt-56 flex flex-col gap-3 w-8/12 max-w-96">
          <h1 className="text-center font-bold text-4xl">Welcome Back</h1>
          <div className="text-center login_other">
            <h4 className="font-bold text-xl mb-2">Login in to account</h4>
            <div className="flex items-center justify-center gap-4">
              <Image
                src="/assets/login/google.svg"
                alt="google"
                width={40}
                height={40}
              />
              <Image
                src="/assets/login/githubb.svg"
                alt="google"
                width={40}
                height={40}
              />
            </div>
          </div>
          <p className="error_login_msg" />

          <form
            className="form_login flex flex-col gap-3"
            method="POST"
            data-form="login"
          >
            <input
              type="text"
              id="email"
              name="email"
              placeholder="Email or Username"
              className="h-10 rounded pl-2 border border-border_color text-bg"
            />

            <input
              type="password"
              id="password"
              name="password"
              placeholder="Password"
              className="h-10 rounded pl-2 border border-border_color text-bg"
            />
            <Button text="Log In" onClick={handleLogin} />
          </form>
        </div>
      </div>
      <div className="bg-[url('/assets/login/bg.jpg')] bg-cover bg-center w-6/12 h-screen hidden sm:block"></div>{" "}
    </div>
  );
}
