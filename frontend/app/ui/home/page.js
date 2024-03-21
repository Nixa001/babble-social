"use client";
import React, { useState, useEffect, useContext } from "react";
import { postForm } from "./postForm";
import DisplayPost from "./displayPost";
import { websocketProvider } from "@/app/home/page";

const HomePage = ({ data }) => {
  console.log("test =>  ", data);
  //const [posts, setPosts] = useState(data || []);
  // const ws = useContext(websocketProvider);
  // ws.onmessage = (e) => {
  //   console.log("in child homepage => ", e.data);
  // };

  return (
    <div className=" md:w-[400px] lg:w-[650px] xl:w-[800px] 2xl:w-[1000px] w-screen">
      <div className="flex justify-center mb-4">{postForm()}</div>
      <div className="post_div_top flex flex-col items-center">
        <DisplayPost
          postData={postData}
          onLikeClick={onLikeClick}
          onDislikeClick={onDislikeClick}
          onCommentClick={onCommentClick}
          onProfileClick={onProfileClick}
        />
        <DisplayPost
          postData={postData}
          onLikeClick={onLikeClick}
          onDislikeClick={onDislikeClick}
          onCommentClick={onCommentClick}
          onProfileClick={onProfileClick}
        />
      </div>
    </div>
  );
};

export default HomePage;

const onLikeClick = () => {
  alert("like");
};

const onDislikeClick = () => {
  alert("dislike");
};

const onCommentClick = () => {
  alert("Comment disp");
};

const onProfileClick = () => {
  alert("profile disp");
};


let postData = {
  "id": 3,
  "content": "This is the content of the third post.",
  "media": "imagepost.jpg",
  "date": "2024-03-05",
  "userId": 1,
  "fullname": "Madike Yade",
  "username": "dickss",
  "user": {
    "id": 1,
    "first_name": "Madike",
    "last_name": "Yade",
    "user_name": "dickss",
    "gender": "Male",
    "user_type": "private",
    "birth_date": "2000-01-01",
    "avatar": "profilibg.jpg",
    "about_me": "about me...",
    "password": "1234",
    "email": "dickss@gmail.com"
  },
  "groupId": 0,
  "privacy": "almost"
}