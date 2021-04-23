import type { NextPage } from 'next'
import type { AppProps } from 'next/dist/next-server/lib/router/router'

const MyApp: NextPage<AppProps> = ({ Component, pageProps }: AppProps) => {
  return <Component {...pageProps} />
}

// eslint-disable-next-line import/no-default-export
export default MyApp
