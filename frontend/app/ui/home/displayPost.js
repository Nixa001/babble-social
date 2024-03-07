import React, { useState, useEffect } from 'react';
import Image from 'next/image';
import { IoIosHeartDislike } from "react-icons/io";
import { useRouter } from 'next/router';

const AddPost = () => {
  const [postData, setPostData] = useState("post");
  const [userData, setUserData] = useState("data");
  // const router = useRouter();

  useEffect(() => {
  }, []);

  const handleLikeClick = () => {
  };

  const handleDislikeClick = () => {
  };

  const handleCommentClick = () => {
    // router.push(`/post/${post.Id}/comments`);
  };

  const handleProfileClick = () => {
    // router.push(`/user/${userData.user.Id}`);
  };

  return (
    <div className="post_div mb-5">
      <div className="post_div_top flex flex-col gap-1 w-fit justify-center ">
        <div className="header_pos w-fit">
          <div className="info_post flex items-center gap-2">
            <Image
              src="/assets/profilibg.jpg"
              alt="Profile picture"
              onClick={handleProfileClick}
              className="profile_pic rounded-full"

              width={50}
              height={50}
            />
            <div className='flex flex-col'>

              <div className='flex gap-2'>
                <h3 className="user_name_post break-words max-w-[600px] w-[80%] font-bold">
                  Nicolas Cor Faye
                </h3>
                <span className="username_post italic text-primary">
                  @nixa
                </span>
              </div>
              <p className="title_post  break-words w-[100%] md:max-w-[300px] max-w-[300px] lg:max-w-[600px]">
                Ceci est mon titre </p>
            </div>
          </div>
        </div>
        <Image
          src={`/assets/imagepost.jpg`}
          alt="Post image"
          className="img_post w-fit"
          width={700}
          height={200}
        />

      </div>
      {/* <IoIosHeartDislike /> */}
      <div className="footer_post w-fit flex gap-4">
        <button className="like_post flex gap-1 items-center" onClick={handleLikeClick}>
          <Image
            src='/assets/icons/likew.png'
            alt="Like icon"
            width={40}
            height={40}
          />
          <span>190</span>
        </button>
        <button className="dislike_post flex gap-1 items-center" onClick={handleDislikeClick}>
          <Image
            src='/assets/icons/dislikew.png'

            alt="Dislike icon"
            width={40}
            height={40}
          />
          <span>20</span>
        </button>
        <button className="comment_post flex gap-1 items-center" onClick={handleCommentClick}>
          <Image
            src="/assets/icons/comment.png"
            alt="Comment icon"
            width={40}
            height={40}
          />
          <span>3</span>
        </button>
      </div>
    </div >
  );
};

export default AddPost;
