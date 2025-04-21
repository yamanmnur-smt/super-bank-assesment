import { useState } from "react";
import { CircleUserRound, User, Lock, Settings, LogOut } from "lucide-react";
import { authStore } from "@/context/auth_store";
import { LogoutRequest } from "./request";
import { useRouter } from "next/navigation";

export default function Navbar() {
	const router = useRouter();

	const [isDropdownOpen, setIsDropdownOpen] = useState(false);

	const toggleDropdown = () => {
		setIsDropdownOpen(!isDropdownOpen);
	};

	const logout = authStore((state) => state.logout);
	
	const submitLogout = async () => {
		
		logout();
		await LogoutRequest()
		router.push("/login")
	};

	return (
		<div className="sticky top-0 bg-[#e4e3de] z-20 h-20 w-full">
			<div className="flex items-center justify-between h-full px-4">
				<div className="text-lg font-semibold text-gray-500">
					<p>Dashboard</p>
					<span className="text-sm">CMS / Dashboard</span>
				</div>
				<div className="relative flex items-center">
					<p className="text-sm text-gray-500 mr-3">Yaman M Nur</p>
					<button
						onClick={toggleDropdown}
						className="relative w-10 h-10 flex items-center justify-center bg-[#588157] text-white font-bold rounded-full"
					>
						<span>CK</span>
					</button>

					{isDropdownOpen && (
						<div className="absolute top-14 right-0 w-56 rounded-xl transition-all transition-discrete shadow-lg bg-[#3a5a40] p-4 backdrop-blur-sm text-white z-50">
							<div className="space-y-3">
								<button
									onClick={() => alert("User Profile")}
									className="flex items-center w-full px-3 py-2 rounded-lg hover:bg-white/20 transition"
								>
									<User className="mr-2 h-4 w-4" /> User Profile
								</button>
								
								<button
									onClick={submitLogout}
									className="cursor-pointer flex items-center w-full px-3 py-2 rounded-lg hover:bg-white/20 transition"
								>
									<LogOut className="mr-2 h-4 w-4" /> Sign Out
								</button>
							</div>
						</div>
					)}
				</div>
			</div>
		</div>
	);
}
