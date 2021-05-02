import ArrowLeftIcon from '@material-ui/icons/ArrowLeft'
import Link from 'next/link'
import { useContext } from 'react'
import { SidebarContext } from 'src/contexts/sidebarContext'

export const Sidebar = () => {
  const { sidebarMode, setSidebarMode } = useContext(SidebarContext)

  return (
    <>
      <div className="sidebar">
        <div className="sidebar-header">
          <div className="user"></div>
          <div
            className="close-sidebar"
            onClick={() => {
              return setSidebarMode('Absolute')
            }}
          >
            {sidebarMode === 'Fixed' && <ArrowLeftIcon />}
          </div>
        </div>
        <div className="sidebar-nav">
          <div>
            <Link href="/">
              <a>Home</a>
            </Link>
          </div>
          <div>
            <Link href="/users">
              <a>Users Page</a>
            </Link>
          </div>
          <div>
            <Link href="/groups">
              <a>Groups Page</a>
            </Link>
          </div>
        </div>
      </div>
      <style jsx>{`
        .sidebar {
          width: 100%;
        }
        .sidebar-header {
          width: 100%;
          height: 40px;
          display: flex;
          flex-direction: row;
          align-items: center;
        }
        .sidebar-nav {
          width: 100%;
          height: 100%;
          padding: 0 0.4em;
        }
        .user {
          flex: 1 1 auto;
          font-style: oblique;
        }
        .close-sidebar {
          flex: 0 1 24px;
          display: none;
        }
        .sidebar:hover .close-sidebar {
          display: inline-block;
        }
        a {
          color: #666;
          text-decoration: none;
        }
      `}</style>
    </>
  )
}
