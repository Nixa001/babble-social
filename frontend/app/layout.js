import "./globals.css";

export const metadata = {
  title: "Social-Network",
  description: "Welcome to SNK, where every word is a work of art",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      {/* <link
        rel="stylesheet"
        href="https://fonts.googleapis.com/icon?family=Material+Icons"
      /> */}
      <body className="">{children}</body>
    </html>
  );
}
