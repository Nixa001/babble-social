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
    <div className="post_div">
      <div className="post_div_top">
        <div className="header_post">
          <div className="info_post">
            <Image
              src="/assets/profilibg.jpg"
              alt="Profile picture"
              onClick={handleProfileClick}
              className="profile_pic"

              width={50}
              height={50}
            />
            <h3 className="user_name_post">
              Nicolas Cor Faye
            </h3>
            <span className="username_post">
              nixa
            </span>
          </div>
          <h4 className="title_post">Ceci est mon titre</h4>
          <h4 className="title_post">Ceci est ma description</h4>
        </div>
        <Image
          src={`/assets/imagepost.jpg`}
          alt="Post image"
          className="img_post"
          width={300}
          height={200}
        />

      </div>
        {/* <IoIosHeartDislike /> */}
      <div className="footer_post">
        <button className="like_post" onClick={handleLikeClick}>
          <Image
            src='/assets/icons/likew.png'

            alt="Like icon"
            width={20}
            height={20}
          />
          <span>190</span>
        </button>
        <button className="dislike_post" onClick={handleDislikeClick}>
          <Image
            src='/assets/icons/dislikew.png'

            alt="Dislike icon"
            width={20}
            height={20}
          />
          <span>20</span>
        </button>
        <button className="comment_post" onClick={handleCommentClick}>
          <Image
            src="/assets/icons/comment.png"
            alt="Comment icon"
            width={20}
            height={20}
          />
          <span>3</span>
        </button>
      </div>
    </div>
  );
};

export default AddPost;
