import type { ReactNode } from 'react'

interface Props {
  children: ReactNode
}

const Layout = ({ children }: Props) => {
  return (
    <>
      <div className="container">
        {children}
        <style jsx global>{`
          body {
            margin: 0;
            background-color: #fffdfd;
            color: #666;
            font-family: -apple-system, 'Segoe UI';
          }
        `}</style>
        <style jsx>{`
          .container {
            max-width: 60vw;
            margin: 4vh auto;
          }
        `}</style>
      </div>
    </>
  )
}

// eslint-disable-next-line import/no-default-export
export default Layout
