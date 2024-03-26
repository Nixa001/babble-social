"use client";
import React, { useState, useEffect, useContext } from "react";
import { postForm } from "./postForm";
import AddPost from "./displayPost";
import { useQuery } from "react-query";

const HomePage = () => {
  const [posts, setPosts] = useState([]);
  const fetchPosts = async () => {
    try {
      const response = await fetch("http://localhost:8080/post");
      const data = await response.json();
      return { posts: data.data };
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
      // console.log("debug => ", posts);
    },
    onError: (error) => {
      console.error("Query error in posts:", error);
    },
  });
  return (
    <div className=" md:w-[400px] lg:w-[650px] xl:w-[800px] 2xl:w-[1000px] w-screen">
      <div className="flex justify-center mb-4">{postForm()}</div>
      <div className="post_div_top flex flex-col items-center">
        {posts.map((e) => (
          <AddPost
            key={e.ID}
            postData={e}
            onLikeClick={onLikeClick}
            onDislikeClick={onDislikeClick}
            onCommentClick={onCommentClick}
            onProfileClick={onProfileClick}
          />
        ))}
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