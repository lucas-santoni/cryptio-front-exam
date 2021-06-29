import Head from 'next/head'
import Link from 'next/link'
import { useState } from 'react'
import axios from 'axios'

import styles from './index.module.css'

const REMOTE_BACKEND_PING_URL = 'http://localhost:8080/ping'
const GOOD_PONG_RESPONSE = "pong\n"

export default function Home() {
  const [connected, setConnected] = useState<boolean>(false)

  useState(() => {
    axios.get<string>(REMOTE_BACKEND_PING_URL)
      .then(res => setConnected(res.data === GOOD_PONG_RESPONSE))
      .catch(_ => setConnected(false))
  })

  return (
    <div className={styles.container}>
      <Head>
        <title>Cryptio Front Test</title>
        <meta name="description" content="Do your best!" />
      </Head>

      <nav className={styles.nav}>
        <Link href="/">Cryptio</Link>
        <Link href="/">Menu 1</Link>
        <Link href="/">Menu 2</Link>
      </nav>

      <main className={styles.main}>
        <h1>Dashboard</h1>
        <p>
          Hello! I guess the widgets you design will go somewhere here...
        </p>
        <p>
          Don't forget to read <Link href="https://github.com/lucas-santoni/cryptio-technical-interview">the assignment</Link> carefully!
        </p>
      </main>

      <footer className={styles.footer}>
        <p>
          { connected && 'You are connected!' }
          { !connected && 'You are not connected...' }
        </p>
      </footer>
    </div>
  )
}
