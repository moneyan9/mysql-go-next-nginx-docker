import ArrowRightIcon from '@material-ui/icons/ArrowRight'
import MenuIcon from '@material-ui/icons/Menu'
import { useContext } from 'react'
import { SidebarContext } from 'src/contexts/sidebarContext'

export const Header = () => {
  const { sidebarMode, setSidebarMode } = useContext(SidebarContext)
  const getMenuIcon = () => {
    if (sidebarMode === 'Hide') {
      return <MenuIcon />
    } else if (sidebarMode === 'Absolute') {
      return <ArrowRightIcon />
    }
  }

  return (
    <>
      <div className="container">
        <div
          className="open-sidebar"
          onClick={() => {
            return setSidebarMode('Fixed')
          }}
        >
          {getMenuIcon()}
        </div>
        <div className="flex"></div>
      </div>
      <style jsx>{`
        .container {
          position: sticky;
          top: 0;
          height: 40px;
          display: flex;
          flex-direction: row;
          align-items: center;
        }
        .open-sidebar {
          flex: 0 1 50px;
          text-align: center;
          cursor: pointer;
        }
        .flex {
          flex: 1 1 auto;
          text-align: center;
        }
      `}</style>
    </>
  )
}
