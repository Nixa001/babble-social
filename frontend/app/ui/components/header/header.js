import Image from "next/image";
import Link from "next/link";

export function Header() {
  return (
    <div className="w-full fixed z-10 h-16 p-4 flex justify-between items-center border border-border_color">
      <Link href="/home" className="">
        <Image
          src="/assets/icons/comment.png"
          alt="logo"
          width={40}
          height={40}
        />
      </Link>
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
