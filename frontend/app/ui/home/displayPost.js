"use client";
import React, { useState } from "react";
import Image from "next/image";
import { MdPrivacyTip } from "react-icons/md";
import { BiWorld } from "react-icons/bi";
import { SiGnuprivacyguard } from "react-icons/si";
import { DisplayComments } from "../components/modals/displayComment";

const DisplayPost = ({
  postData,
  onLikeClick,
  onDislikeClick,
  onCommentClick,
  onProfileClick,
}) => {
  const [postDataState, setPostDataState] = useState(postData);
  const [isVisibleComment, setIsVisibleComment] = useState(false);

  // const handleLikeClick = () => {
  //   onLikeClick();
  //   setPostDataState((prevState) => ({
  //     ...prevState,
  //     likesCount: prevState.likesCount + 1,
  //   }));
  // };

  // const handleDislikeClick = () => {
  //   onDislikeClick();
  //   setPostDataState((prevState) => ({
  //     ...prevState,
  //     dislikesCount: prevState.dislikesCount + 1,
  //   }));
  // };

  // const handleCommentClick = () => {
  //   onCommentClick();
  // };

  const handleProfileClick = () => {
    onProfileClick();
  };
  //console.log("in display 🔔", postData);
  return (
    <div className="w-fit post_div mb-5 border border-gray-700 p-2 rounded-md">
      <div className="post_div_top flex flex-col gap-1 w-fit justify-center ">
        <div className="header_pos w-fit">
          <div className="info_post flex items-start gap-2">
            <img
              src={postDataState.Avatar || "/assets/profilibg.jpg"}
              alt="Profile picture"
              onClick={handleProfileClick}
              className="profile_pic rounded-full cursor-pointer hover:opacity-60"
              width={50}
              height={50}
            />
            <div className="flex flex-col">
              <div className="flex gap-2 items-center w-[100%]">
                <h3 className="user_name_post break-words max-w-[600px] w-[80%] font-bold">
                  {postDataState.FullName}
                </h3>
                <span className="username_post italic text-primary">
                  {postDataState.UserName}
                </span>
                <div className="flex">
                  <MdPrivacyTip className="text-2xl" />
                  <SiGnuprivacyguard className="text-2xl" />
                  <BiWorld className="text-2xl" />
                </div>
              </div>
              <div className="flex gap-4 text-sm">
                <span>{postDataState.Date}</span>
                <div className="flex gap-1 italic text-primary">
                  {postDataState.Categories.split(",").map((tag) => (
                    <span key={tag} className="mr-1">
                      #{tag}
                    </span>
                  ))}
                </div>
              </div>
              <p className="title_post  break-words w-[100%] md:max-w-[300px] max-w-[300px] lg:max-w-[600px]">
                {postDataState.Content}
              </p>
            </div>
          </div>
        </div>
        {postDataState.Media != "NULL" && (
          <img
            src={postDataState.Media}
            alt="Post image"
            className="img_post max-w-full hover:shadow-xl overflow-hidden cursor-pointer transition duration-300 ease-linear scale-95 hover:scale-100"
            width={700}
            height={200}
          />
        )}
      </div>
      <div className="footer_post w-fit flex gap-4">
        <button
          className="comment_post flex gap-1 items-center"
          onClick={() => {
            setIsVisibleComment(true);
          }}>
          <Image
            src="/assets/icons/comment.png"
            alt="Comment icon"
            width={30}
            height={30}
          />
          <span>{postDataState.Comments}</span>
        </button>
      </div>
        <DisplayComments
          isVisible={isVisibleComment}
          postId={postDataState.ID}
          onClose={() => setIsVisibleComment(false)}
        />
    </div>
  );
};

export default DisplayPost;
