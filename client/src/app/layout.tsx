import type { Metadata } from "next";
import "./globals.css";
import { Auth0Provider } from '@auth0/nextjs-auth0';
import NavBar from "./components/NavBar";

export const metadata: Metadata = {
  title: "Auth0 Demo",
  description: "Demo app for working with Auth0 and Next.js",
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <Auth0Provider>
      <html lang="en">
        <body>
          <NavBar />
          {children}
        </body>
      </html>
    </Auth0Provider>
  );
}
