'use client'
import { authStore } from "@/context/auth_store";
import { LayoutDashboard, Users, EllipsisVertical, ChevronLeft } from "lucide-react";
import { AppRouterInstance } from "next/dist/shared/lib/app-router-context.shared-runtime";
import { usePathname, useRouter  } from "next/navigation";

import React from "react";

interface SidebarMenuProps {
  icon: React.ComponentType<React.SVGAttributes<SVGSVGElement>>
  label: string; 
  active?: boolean;
  link: string;
  iconProps?: React.SVGProps<SVGSVGElement>;
}
export default function Sidebar() {
    const activePath = usePathname()
    const router = useRouter()
    const userState = authStore((state) => state.user);

    const routes: SidebarMenuProps[] = [
      {
        icon : LayoutDashboard,
        label : 'Dashboard',
        active : false,
        iconProps : {className: "w-5 h-5"},
        link : '/dashboard'
      },
      {
        icon : Users,
        label : 'Customers',
        active : false,
        iconProps : {className: "w-5 h-5"},
        link : '/customers'
      }
    ]


    return (
      <div className="relative w-64  text-gray-400 flex flex-col p-4 shadow-xl mx-4 my-4 rounded-2xl bg-[#3a5a40]">
         <button className="cursor-pointer absolute -right-3 top-1/12 transform -translate-y-1/2 bg-white/10 hover:bg-white/20 text-white w-6 h-17 rounded-md flex items-center justify-center shadow-md backdrop-blur-sm">
          <ChevronLeft className="w-4 h-4" />
        </button>
        {/* Logo */}
        <div className="mb-5 text-[#a3b18a] font-bold text-2xl">
          <span className="text-3xl">Superbank </span>
        </div>

        <div className="flex items-center space-x-3 border-b border-white pb-8 mb-5">
          <button
						className="relative w-18 h-17 flex items-center justify-center bg-gradient-to-br from-[#a3b18a] via-[#a3b18a] to-[#588157] text-white font-bold rounded-full"
					>
						<span>YM</span>
					</button>
          <div className="flex flex-col ">
            <span className="text-sm text-white ">Hello 👋</span>
            <span className="text-lg text-white font-bold">{userState?.name}</span>
          </div>
        </div>
        {/* Navigation Links */}
        <div className="mb-4 flex">
          <EllipsisVertical className="text-white text-sm" size={20} />
          <span className="text-sm text-gray-100 mr-2 font-semibold">menu </span>
        </div>
        <nav className="flex-1 space-y-1">
          {routes.map((item, index) => (
            <SidebarLink key={index} router={router} label={item.label} link={item.link} icon={item.icon} iconProps={item.iconProps} active={activePath === item.link} />
          ))}
        </nav>
       
      </div>
    );
  }

  
  function SidebarLink({ label, active = false, icon: IconComponent, iconProps = {}, link, router}: SidebarMenuProps & {
    router : AppRouterInstance
  }) {
 
    return (
      <a
        onClick={() => {
          router.push(link)
        }}
        className={
          `cursor-pointer flex items-center space-x-3 rounded-md px-2 py-2 mb-3 text-sm font-medium transition-colors duration-150 gradient ${
          active ? "bg-white/30 text-white" : "hover:bg-white/30 hover:text-gray-100 text-gray-100"
        }`}
      >

        <IconComponent {...iconProps} />
        <span>{label}</span>
      </a>
    );
  }
  
  function TeamBadge({ initial, name }: { initial: string; name: string }) {
    return (
      <div className="flex items-center space-x-3 mb-2">
        <div className="flex items-center justify-center w-6 h-6 rounded-full bg-indigo-500 text-sm font-bold">
          {initial}
        </div>
        <span className="text-sm text-indigo-100">{name}</span>
      </div>
    );
  }
  