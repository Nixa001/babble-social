"use client"
import React from "react";
import Button from "../components/button/button";
import { FaImage } from "react-icons/fa6";

const HomePage = () => {
  return (
    <div className=" md:w-[400px] lg:w-[650px] xl:w-[800px] 2xl:w-[1000px] w-screen">
      <div className="">
        {postForm()}
        </div>
      <div className="post_div_top"></div>
    </div>
  );
};

export default HomePage;

const handlePost = () => {
  alert("ok");
};

const postForm = () => {
  return (
    <form
      className="flex flex-col gap-1  ml-2 mr-2"
      action=""
      method="POST"
      data-form="post"
      encType="multipart/form-data"
    >
      <textarea
        className="resize-none border focus:outline-none focus:border text-bg rounded-md p-1"
        name="title_post"
        placeholder="Let's post something ..."
        required=""
        defaultValue={""}
      />
      <div className="flex items-center justify-end">
        <div className="flex gap-1 mr-2">
          <div>
            <input
              type="checkbox"
              id="check"
              defaultValue="technologie"
              name="techno"
            />
            <label htmlFor="check">Tech</label>
          </div>
          <div>
            <input
              type="checkbox"
              id="checks"
              defaultValue="sport"
              name="sport"
            />
            <label htmlFor="checks">Sport</label>
          </div>
          <div>
            <input
              type="checkbox"
              id="checkS"
              defaultValue="sante"
              name="sante"
            />
            <label htmlFor="checkS">Sante</label>
          </div>
          <div>
            <input
              type="checkbox"
              id="checkM"
              defaultValue="musique"
              name="music"
            />
            <label htmlFor="checkM">Musique</label>
          </div>
          <div>
            <input
              type="checkbox"
              id="checkN"
              defaultValue="news"
              name="news"
            />
            <label htmlFor="checkN">News</label>
          </div>
          <div>
            <input
              type="checkbox"
              id="checko"
              defaultValue="other"
              name="other"
              defaultChecked=""
            />
            <label htmlFor="checko">Other</label>
          </div>
        </div>

      <div>
        <select>
          <option value="1">1</option>
          <option value="2">2</option>
          <option value="3">3</option>
          <option value="4">4</option>
          <option value="5">5</option>
        </select>
      </div>

        <label htmlFor="image_post" className=" cursor-pointer mr-2">
          <FaImage className="w-8 h-8 " />
        </label>
        <input type="file" name="image_post" id="image_post" hidden />
        <input
          className="bg-second text-lg font-bold pl-3 pr-3 rounded-sm cursor-pointer hover:bg-primary"
          type="submit"
          Value="Post"
        />
        {/* < Button text="Log In" onClick={handlePost()} /> */}
      </div>
    </form>
  );
};
