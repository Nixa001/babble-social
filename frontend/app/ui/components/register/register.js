"use client";
import React from "react";
import Image from "next/image";
import Link from "next/link";

function Register() {
  const handleRegister = () => {
    alert("Register");
  };

  return (
    <div className="w-screen h-screen flex">
      <div className="flex flex-col items-center w-full sm:w-6/12">
        <div className="header_login flex items-center justify-between w-full">
          <Image
            src="/assets/icons/comment.png"
            alt="logo"
            width={40}
            height={40}
          />{" "}
          <div>
            You have already an account?{" "}
            <Link
              href="/login"
              className="text-primary hover:text-second cursor-pointer"
            >
              Sign In .{" "}
            </Link>
          </div>
        </div>
        <div className="mt-24 flex flex-col gap-3 w-8/12 max-w-96">
          <h1 className="text-center font-bold text-4xl">Sign Up</h1>
          <div className="flex items-center justify-center gap-4">
            <img src="/assets/login/google.svg" alt="google" />
            <img src="/assets/login/githubb.svg" alt="github" />
          </div>
          <p className="error_login_msg" />

          <form
            className="w-full flex flex-col gap-3"
            // method="POST"
            data-form="login"
          >
            <input
              type="text"
              id="firstname"
              name="firstname"
              placeholder="Firstname"
              className="h-10 rounded pl-2 text-bg"
            />
            <input
              type="text"
              id="lastname"
              name="lastname"
              placeholder="Lastname"
              className="h-10 rounded pl-2 text-bg"
            />

            <input
              type="email"
              id="email"
              name="email"
              placeholder="E-mail"
              className="h-10 rounded pl-2 text-bg"
            />
            <input
              type="text"
              id="username"
              name="username"
              placeholder="Username"
              className="h-10 rounded pl-2 text-bg"
            />
            <input
              type="date"
              id="dateofbirth"
              name="dateofbirth"
              placeholder="Date of birth"
              className="h-10 rounded pl-2 text-bg"
            />
            <input
              type="password"
              id="password"
              name="password"
              placeholder="Password"
              className="h-10 rounded pl-2 text-bg"
            />
            <input
              type="password"
              id="password"
              name="password"
              placeholder="Confirm your password"
              className="h-10 rounded pl-2 text-bg"
            />
            <textarea
              type="text"
              id="aboutme"
              name="aboutme"
              placeholder="About me"
              className="h-20 pt-6 rounded pl-2 text-bg resize-none"
            />
            <input
              type="file"
              id="avatar"
              name="avatar"
              placeholder="Avatar"
              className="h-10 rounded pl-2 text-bg"
            />
            <Link href="/home">
            <button
              className="hover:bg-second bg-primary cursor-pointer border-none w-full h-10 rounded font-bold text-center"
              // onClick={() => {
              //   handleRegister();
              // }}
            >
              Log In
            </button>
            </Link>
          </form>
        </div>
      </div>
      <div className="bg-[url('/assets/login/bg.jpg')] bg-cover bg-center w-6/12 h-screen hidden sm:block"></div>{" "}
    </div>
  );
}

export default Register;
