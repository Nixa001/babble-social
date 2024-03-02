import { Inter } from "next/font/google";
import "./globals.css";

const inter = Inter({ subsets: ["latin"] });

export const metadata = {
  title: "Social-Network",
  description: "Welcome to PrismaChat, where every word is a work of art",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      {/* <link
        rel="stylesheet"
        href="https://fonts.googleapis.com/icon?family=Material+Icons"
      /> */}
      <body className={inter.className}>{children}</body>
    </html>
  );
}
