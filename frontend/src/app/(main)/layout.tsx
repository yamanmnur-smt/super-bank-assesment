'use client';
import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import "@/app/globals.css";
import { authStore } from "@/context/auth_store";
import { useEffect } from "react";
import { redirect, useRouter } from "next/navigation";
import Sidebar from "./_components/sidebar";
import Navbar from "./_components/navbar";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  // const router = useRouter();
  
  // const isAuth = authStore((state) => state.isAuth);
  // const hydrated = authStore((state) => state.hydrated);
  // useEffect(() => {
   
  //   if (hydrated && !isAuth) {
  //     router.replace("/login");
  //   }

  // }, [hydrated, isAuth]);

  return (
    <section className="flex w-full h-full bg-[#e4e3de] relative">
      <Sidebar />
      <div className="flex-col flex w-full">
      <Navbar />
        <div className="overflow-y-scroll rounded-xl mb-4 w-full h-full no-scrollbar">
          <div className="flex flex-1 flex-col bg-white/30 w-full bg-blur p-3">
            {children}
          </div>
        </div>
      </div>
    </section>
  );
}
