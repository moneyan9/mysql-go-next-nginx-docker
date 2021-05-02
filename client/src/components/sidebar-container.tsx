import { useContext } from 'react'
import { Sidebar } from 'src/components/sidebar'
import { SidebarContext } from 'src/contexts/sidebarContext'

export const SidebarContainer = () => {
  const { sidebarMode } = useContext(SidebarContext)

  return (
    <>
      <div className={sidebarMode === 'Fixed' ? 'fixed-container' : 'absolute-container'}>
        <Sidebar />
      </div>
      <style jsx>{`
        .fixed-container {
          position: static;
          padding: 0 0.8em;
          transition-duration: 0.1s;
          transform: translateY(0vh);
        }
        .absolute-container {
          position: relative;
          padding: 0 0.8em;
          width: 240px;
          height: 60vh;
          background-color: white;
          box-shadow: rgba(15, 15, 15, 0.05) 0px 0px 0px 1px, rgba(15, 15, 15, 0.1) 0px 3px 6px,
            rgba(15, 15, 15, 0.2) 0px 9px 24px;
          transition-duration: 0.1s;
          transform: translateY(5vh);
          z-index: 9999;
          opacity: 0.98;
        }
      `}</style>
    </>
  )
}
