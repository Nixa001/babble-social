"use client"
import React from "react";
import { postForm } from "./postForm";
import AddPost from "./displayPost";

const HomePage = () => {
  return (
    <div className=" md:w-[400px] lg:w-[650px] xl:w-[800px] 2xl:w-[1000px] w-screen">
      <div className="">
        {postForm()}
      </div>
      <div className="post_div_top">
        {AddPost()}
      </div>
    </div>
  );
};

export default HomePage;

const handlePost = () => {
  alert("ok");
};



const followers = [
  { name: 'Vindour', src: "/assets/profilibg.jpg", alt: "profil" },
  { name: 'ibg', src: "/assets/profilibg.jpg", alt: "profil", },
  { name: 'dicks', src: "/assets/profilibg.jpg", alt: "profil", },
  { name: 'Vindcour', src: "/assets/profilibg.jpg", alt: "profil" },
  { name: 'ibgs', src: "/assets/profilibg.jpg", alt: "profil", },
  { name: 'dickss', src: "/assets/profilibg.jpg", alt: "profil", },
];
