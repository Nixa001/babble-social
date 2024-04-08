"use client";
import { useState } from "react";
import { useQuery } from "react-query";
import { PostForm } from "./postForm";
import DisplayPost from "./displayPost";

const HomePage = ({ id }) => {
  const [posts, setPosts] = useState([]),
    [fetchState, setFetchState] = useState(true),
    userID = new FormData();
  userID.append("userID", id);
  userID.append("type", "loadPosts");
  const options = {
    method: "POST",
    body: userID,
  };
  const fetchPosts = async () => {
    try {
      const response = await fetch("http://localhost:8080/post", options);
      const data = await response.json();
      return { posts: data.data, errType: data.status };
    } catch (error) {
      console.error("Error while querying posts ", error);
      setFetchState(false);
      return Promise.reject(error);
    }
  };

  useQuery("posts", fetchPosts, {
    enabled: fetchState,
    refetchInterval: 1000,
    staleTime: 500,
    onSuccess: (newData) => {
      //  if (newData.errType == 400) setFetchState(false);
      setPosts(newData.posts);
      console.log("debug => ", newData);
    },
    onError: (error) => {
      console.error("Query error in posts:", error);
      setFetchState(false);
    },
  });
  return (
    <div className=" md:w-[400px] lg:w-[650px] xl:w-[800px] 2xl:w-[1000px] w-screen">
      <div className="flex justify-center mb-4">{PostForm(id)}</div>
      <div className="post_div_top flex flex-col items-center">
        {posts?.map((e) => (
          <DisplayPost
            key={e.ID}
            postData={e}
            idUser={id}
            onLikeClick={onLikeClick}
            onDislikeClick={onDislikeClick}
            onCommentClick={onCommentClick}
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

let postData = {
  id: 3,
  content: "This is the content of the third post.",
  media: "imagepost.jpg",
  date: "2024-03-05",
  userId: 1,
  fullname: "Madike Yade",
  username: "dickss",
  user: {
    id: 1,
    first_name: "Madike",
    last_name: "Yade",
    user_name: "dickss",
    gender: "Male",
    user_type: "private",
    birth_date: "2000-01-01",
    avatar: "profilibg.jpg",
    about_me: "about me...",
    password: "1234",
    email: "dickss@gmail.com",
  },
  groupId: 0,
  privacy: "almost",
};
