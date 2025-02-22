import { create } from "zustand";

interface State {
  isDarkMode: boolean;
}

interface Action {
  setIsDarkMode: (isDarkMode: boolean) => void;
  getIsDarkMode: (isDarkMode: boolean) => boolean;
}

const useGlobalStore = create<State & Action>((set, get) => ({
  isDarkMode: false,
  setIsDarkMode: (isDarkMode: boolean) => set({ isDarkMode }),
  getIsDarkMode: () => get().isDarkMode,
}));

export default useGlobalStore;
