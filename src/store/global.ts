import { User } from "@/types/user";
import { create } from "zustand";

interface State {
  isDarkMode: boolean;
  user: User;
}

interface Action {
  setIsDarkMode: (isDarkMode: boolean) => void;
  getIsDarkMode: (isDarkMode: boolean) => boolean;
  setUser: (user: User) => void;
  getUser: () => User;
}

const useGlobalStore = create<State & Action>((set, get) => ({
  user: {} as User,
  setUser: (user: User) => set({ user }),
  getUser: () => get().user,
  isDarkMode: localStorage.getItem("isDarkMode") === "true",
  setIsDarkMode: (isDarkMode: boolean) => set({ isDarkMode }),
  getIsDarkMode: () => get().isDarkMode,
}));

export default useGlobalStore;
