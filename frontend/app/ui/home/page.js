"use client";
import React, { useState, useEffect } from "react";
import { postForm } from "./postForm";
import AddPost from "./displayPost";

const HomePage = () => {
  const props = {
    data: "This is a test",
    title: "Test Title",
  };
  const [data, setData] = useState(props.data);
  const [ws, setWS] = useState(null);
  useEffect(() => {
    const newWS = new WebSocket("ws://localhost:8080/socket");
    newWS.onerror = (err) => console.error(err);
    newWS.onopen = () => {
      setWS(newWS);
      // send data to server when connected
      newWS.send(JSON.stringify({ type: "GET", request: "get10post" }));
    };
    newWS.onmessage = (msg) => {
      setData(JSON.parse(msg.data));
      console.log(`data is here => : ${data}`);
    };
  }, []);
  return (
    <div className=" md:w-[400px] lg:w-[650px] xl:w-[800px] 2xl:w-[1000px] w-screen">
      <div className="flex justify-center mb-4">{postForm()}</div>
      <div className="post_div_top flex flex-col items-center">
        <AddPost
          postData={postData}
          onLikeClick={onLikeClick}
          onDislikeClick={onDislikeClick}
          onCommentClick={onCommentClick}
          onProfileClick={onProfileClick}
        />
        <AddPost
          postData={postData2}
          onLikeClick={onLikeClick}
          onDislikeClick={onDislikeClick}
          onCommentClick={onCommentClick}
          onProfileClick={onProfileClick}
        />
      </div>
    </div>
  );
};

const postData = {
  profilePicture: "/assets/profilibg.jpg",
  userName: "Nicolas Cor Faye",
  userHandle: "@nixa",
  timePosted: "1h",
  hashtags: ["Tech", "Sport"],
  title: "Ceci est mon titre",
  postImage: "/assets/imagepost.jpg",
  likesCount: 190,
  dislikesCount: 20,
  commentsCount: 3,
};
const postData2 = {
  profilePicture: "/assets/profilibg.jpg",
  userName: "Maurice Dassylva",
  userHandle: "@Maurice",
  timePosted: "2h",
  hashtags: ["Tech", "Sport"],
  title: "Ceci est mon titre",
  postImage: "/assets/imagepost2.jpg",
  likesCount: 19,
  dislikesCount: 20,
  commentsCount: 3,
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
