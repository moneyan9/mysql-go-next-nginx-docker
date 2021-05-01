import type { NextPage } from 'next'
import type { AppProps } from 'next/app'
import Layout from 'src/components/layout'

const MyApp: NextPage<AppProps> = ({ Component, pageProps }: AppProps) => {
  return (
    <Layout>
      <Component {...pageProps} />
    </Layout>
  )
}

// eslint-disable-next-line import/no-default-export
export default MyApp
