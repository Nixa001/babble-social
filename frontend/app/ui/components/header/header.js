import Image from "next/image";

export function Header(params) {
  return (
    <div className="w-full fixed z-10 h-16 p-4 flex justify-between items-center border border-white">
      <a className="">
        <Image
          src="/assets/icons/comment.png"
          alt="logo"
          width={40}
          height={40}
        />
      </a>
      <div className="flex bg-white color-black p-2">
        <i className="material-icons">search</i>
        <input type="search" placeholder="Search ..." name="search" />
      </div>
      <div className="">
        <i className="material-icons chat">chat</i>
        <i className="material-icons notification">notifications</i>
      </div>
    </div>
  );
}
