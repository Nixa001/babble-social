import { Header } from "../ui/components/header/header";
import Navbar from "../ui/components/navbar/navbar";
import Sidebar from "../ui/components/sidebarRight/sidebar";
import { displayFollowers } from "./utils";

export default function Layout({ children }) {
  return (
    <div className="h-full">
      <div className="fixed">
        <Header />
      </div>
      <div className="md:flex md:flex-row flex flex-col-reverse justify-between h-screen md:justify-between md:h-full">

        <div className="md:mt-20">
          <Navbar />
        </div>

        <div className="mt-20">
          {children}
        </div>

        <div className="md:mt-20 hidden md:block">
          <Sidebar  />
        </div>

      </div>
    </div>
  );
}
