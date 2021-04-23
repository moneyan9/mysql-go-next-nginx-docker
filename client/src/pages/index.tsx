import { Button } from '@material-ui/core'
import type { NextPage } from 'next'
import Link from 'next/link'

const Home: NextPage = () => {
  return (
    <>
      <Link href="/users">
        <Button variant="outlined">MOVE TO USERS PAGE</Button>
      </Link>
    </>
  )
}

// eslint-disable-next-line import/no-default-export
export default Home
