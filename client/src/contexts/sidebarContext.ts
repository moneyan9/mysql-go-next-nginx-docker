import { createContext } from 'react'

export type SidbarMode = 'Hide' | 'Absolute' | 'Fixed'
export interface SidebarContextType {
  sidebarMode: SidbarMode
  setSidebarMode: (value: SidbarMode) => void
}

const initialSidebarContext: SidebarContextType = {
  sidebarMode: 'Fixed',
  setSidebarMode: () => {
    //do nothing
  },
}

export const SidebarContext = createContext<SidebarContextType>(initialSidebarContext)
