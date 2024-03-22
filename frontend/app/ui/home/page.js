"use client";
import React, { useState, useEffect, useContext } from "react";
import { postForm } from "./postForm";
import AddPost from "./displayPost";
import { useQuery } from "react-query";

const HomePage = () => {
  const [posts, setPosts] = useState([]);
  const [postList, setPostList] = useState([]);
  const fetchPosts = async () => {
    try {
      const response = await fetch("http://localhost:8080/post");
      const data = await response.json();
      return { posts: data };
    } catch (error) {
      console.error("Error while querying posts ", error);
      return Promise.reject(error);
    }
  };

  useQuery("posts", fetchPosts, {
    enabled: true,
    refetchInterval: 1000,
    staleTime: 500,
    onSuccess: (newData) => {
      setPosts(newData.posts);
      if (newData.posts.length >0)
      setPostList(
        newData.posts.map((x) => {
          x = (
            <AddPost
              postData={x}
              onLikeClick={onLikeClick}
              onDislikeClick={onDislikeClick}
              onCommentClick={onCommentClick}
              onProfileClick={onProfileClick}
            />
          );
        })
      );
      console.log("debug => ", posts);
      console.log("debugList => ", postList);
    },
    onError: (error) => {
      console.error("Query error in posts:", error);
    },
  });
  return (
    <div className=" md:w-[400px] lg:w-[650px] xl:w-[800px] 2xl:w-[1000px] w-screen">
      <div className="flex justify-center mb-4">{postForm()}</div>
      <div className="post_div_top flex flex-col items-center">{postList}</div>
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
