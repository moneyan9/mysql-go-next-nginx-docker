import Head from 'next/head'
import { useState } from 'react'
import { SidebarContainer } from 'src/components/sidebar-container'
import type { SidbarMode } from 'src/contexts/sidebarContext'
import { SidebarContext } from 'src/contexts/sidebarContext'

import { Header } from './header'

type LayoutProps = {
  loading: boolean
  children: React.ReactNode
}

export const Layout = ({ children }: LayoutProps) => {
  const [sidebarMode, setSidebarMode] = useState<SidbarMode>('Fixed')

  const contentMouseMove = (e: React.MouseEvent<HTMLElement, MouseEvent>) => {
    if (sidebarMode !== 'Fixed') {
      const showMenuRange = window.innerWidth / 10
      if (e.clientX < showMenuRange && sidebarMode === 'Hide') {
        setSidebarMode('Absolute')
        return
      } else if (e.clientX > showMenuRange && sidebarMode === 'Absolute') {
        setSidebarMode('Hide')
        return
      }
    }
  }

  return (
    <>
      <Head>
        <title>Next.js Sample</title>
      </Head>

      <SidebarContext.Provider value={{ sidebarMode, setSidebarMode }}>
        <div className="layout-container">
          {sidebarMode !== 'Hide' && (
            <div className={sidebarMode === 'Fixed' ? 'sidebar-fixed' : 'sidebar-absolute'}>
              <SidebarContainer />
            </div>
          )}
          <div
            className="content"
            onMouseMove={(e: React.MouseEvent<HTMLElement, MouseEvent>) => {
              contentMouseMove(e)
            }}
          >
            <div className="header">
              <Header />
            </div>
            <div className="article">{children}</div>
          </div>
        </div>
      </SidebarContext.Provider>

      <style jsx>{`
        .layout-container {
          display: flex;
          flex-direction: row;
        }
        .sidebar-fixed {
          min-width: 240px;
          height: 100vh;
          background-color: #f7f6f3;
        }
        .sidebar-absolute {
          width: 0px;
        }
        .content {
          width: 100%;
        }
        .header {
        }
        .article {
          max-width: 80%;
          margin: 1.5rem auto;
        }
      `}</style>
      <style jsx global>{`
        body {
          margin: 0;
          background-color: #fffdfd;
          color: #666;
          font-family: -apple-system, 'Segoe UI';
        }
      `}</style>
    </>
  )
}

// eslint-disable-next-line import/no-default-export
export default Layout
