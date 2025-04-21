"use client";
import Image from "next/image";
import { useRouter } from "next/navigation";
import { FormEvent, useState } from "react";
import Img from "@assets/image.jpg";
import { LoginRequest } from "./requests";
import { authStore } from "@/context/auth_store";

import Loading from "../../_components/svg_loadings";
import AlertPopup from "@/app/_components/alert";

const LoginComponent = () => {
  const router = useRouter();
  const [showAlert, setShowAlert] = useState(false);

  const setUser = authStore((state) => state.setUser);
  const [loader, setLoader] = useState(false);
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = async (e : FormEvent<HTMLFormElement>
  ) => {
    e.preventDefault();
    setLoader(true)
    try {
      
      const response = await LoginRequest({
        username: username,
        password: password,
      });
      const user = response.data.user
      setUser({
        id : user.id,
        name : user.name,
        isLoggedIn : true,
        username : user.username,
      })
      setTimeout(() => {
        setLoader(false)
      }, 400);
      router.push("/dashboard");
    } catch (error) {
      setShowAlert(true)
      setLoader(false)
    }
  };

  return (
    <div className="min-h-screen bg-[#e4e3de] flex items-center justify-center">
      <div className="bg-[#dad7cd] shadow-lg rounded-3xl flex w-full max-w-6xl overflow-hidden">
        {/* Left Section */}
        <div className="w-full md:w-1/2 p-10 flex flex-col justify-center">
          <div className="mb-6">
            <h2 className="text-sm text-gray-400 font-medium">
              START FOR FREE
            </h2>
            <h1 className="text-4xl font-bold mt-2">
              Sign In<span className="text-blue-500">.</span>
            </h1>
          </div>

          <form className="space-y-4" onSubmit={(e : FormEvent<HTMLFormElement>) => handleSubmit(e)}>
            <input
              type="text"
              placeholder="Username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              className="w-full px-4 py-3 border border-gray-300 rounded-xl focus:outline-none focus:ring-2 focus:ring-blue-400"
            />
            <input
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              placeholder="Password"
              className="w-full px-4 py-3 border border-gray-300 rounded-xl focus:outline-none focus:ring-2 focus:ring-blue-400"
            />
            <div className="flex gap-4 mt-4">
              <button
                type="submit"
                className="flex-1 px-4 py-3 rounded-xl bg-[#3a5a40] text-white hover:bg-blue-600 shadow-md"
              >
                {loader ?
                  <Loading/>
                  : <></>}
                Submit
              </button>
            </div>
          </form>
        </div>

        {/* Right Section - Image */}
        <div className="hidden md:block w-1/2 relative">
          <Image
            src={Img}
            alt="Scenic mountain"
            fill
            className="object-cover"
          />
        </div>
      </div>
      {showAlert && (
        <AlertPopup
          message="Credential Invalid!"
          type="error"
          onClose={() => setShowAlert(false)}
        />
      )}
    </div>
  );
};

export default LoginComponent;
