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
  const [post, setpost] = useState({
    ...postData,
  });
  const [isVisibleComment, setIsVisibleComment] = useState(false)

  const handleLikeClick = () => {
    onLikeClick();
    setpost((prevState) => ({
      ...prevState,
      likesCount: prevState.likesCount + 1,
    }));
  };
  // console.log(post);
  const handleDislikeClick = () => {
    onDislikeClick();
    setpost((prevState) => ({
      ...prevState,
      dislikesCount: prevState.dislikesCount + 1,
    }));
  };

  const handleCommentClick = () => {
    onCommentClick();
  };

  const handleProfileClick = () => {
    onProfileClick();
  };

  return (
    <div className=" post_div mb-5 border border-gray-700 p-2 rounded-md">
      <div className="post_div_top flex flex-col gap-1 w-fit justify-center ">
        <div className="header_pos w-fit">
          <div className="info_post flex items-start gap-2">
            <Image
              src={`/assets/${post.user.avatar}`}
              alt="Profile picture"
              onClick={handleProfileClick}
              className="profile_pic rounded-full cursor-pointer hover:opacity-60"
              width={50}
              height={50}
            />
            <div className='flex flex-col'>
              <div className='flex gap-2 items-center w-[100%]'>
                <h3 className="user_name_post break-words max-w-[600px] w-[80%] font-bold">
                  {post.user.first_name + " "}
                  {post.user.last_name}
                </h3>
                <span className="username_post italic text-primary">
                  {post.user.user_name ?
                    '@' + post.user.user_name : ('')
                  }
                </span>
                <div className='flex'>

                  {post.privacy === "private" ? (
                    <SiGnuprivacyguard className="text-2xl" />
                  ) : post.privacy === "almost" ? (
                    <MdPrivacyTip className="text-2xl" />
                  ) : (
                    <BiWorld className="text-2xl" />
                  )}

                </div>
              </div>
              <div className='flex gap-4 text-sm'>
                <span>{post.date}</span>
                <div className='flex gap-1 italic text-primary'>
                  {/* {post.name.map((tag) => (
                    <span key={tag} className="mr-1">#{tag}</span>
                  ))} */}
                </div>
              </div>
              <p className="title_post  break-words w-[100%] md:max-w-[300px] max-w-[300px] lg:max-w-[600px]">
                {post.content}
              </p>
            </div>
          </div>
        </div>
        {post.media ? (
          <Image
            src={`/assets/${post.media}`}
            alt="Post image"
            className="img_post max-w-full hover:shadow-xl overflow-hidden cursor-pointer transition duration-300 ease-linear scale-95 hover:scale-100"
            width={700}
            height={200}
          />) : ("")}

      </div>
      <div className="footer_post w-fit flex gap-4">
        <button className="like_post flex gap-1 items-center" onClick={handleLikeClick}>
          <Image
            src='/assets/icons/likew.png'
            alt="Like icon"
            width={30}
            height={30}
          />
          <span>{post.likes}</span>
        </button>
        <button className="dislike_post flex gap-1 items-center" onClick={handleDislikeClick}>
          <Image
            src='/assets/icons/dislikew.png'
            alt="Dislike icon"
            width={30}
            height={30}
          />
          <span>{post.dislikes}</span>
        </button>
        <button className="comment_post flex gap-1 items-center" onClick={() => {
          setIsVisibleComment(true)
        }}>
          <Image
            src="/assets/icons/comment.png"
            alt="Comment icon"
            width={30}
            height={30}
          />
          <span>{post.comments ? post.comments.length : ("0")}</span>
        </button>
      </div>
      <DisplayComments isVisible={isVisibleComment} onClose={() => setIsVisibleComment(false)} />

    </div >
  );
};

export default DisplayPost;
