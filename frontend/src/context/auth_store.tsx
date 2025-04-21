import { create } from "zustand";
import { persist } from "zustand/middleware";
import UserState from "./interfaces/user";

interface AuthState {
  user: UserState | undefined; 
  isAuth: boolean;
  errors: any;
	hydrated: boolean;
}

type Actions = {
  logout(): void;
  setUser(user: UserState): void;
	setHydrated(): void;
};

export const authStore = create<Actions & AuthState>()(
  persist(
    (set) => ({
      user: undefined,
      isAuth: false,
			hydrated: false,
      errors: null,
      logout: () => set({ user: undefined }),
      setUser: (user: UserState) => set({ user }),
			setHydrated: () => set({ hydrated: true }),
    }),
    {
      name: "auth_store",
			onRehydrateStorage: () => (state) => {
        state?.setHydrated()
      },
    }
  )
);
