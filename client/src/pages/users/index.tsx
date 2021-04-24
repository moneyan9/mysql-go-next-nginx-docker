import {
  Button,
  IconButton,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from '@material-ui/core'
import { Delete, Edit, PersonAdd } from '@material-ui/icons'
import axios from 'axios'
import type { User } from 'models/user'
import Link from 'next/link'
import { useState } from 'react'
import useSWR from 'swr'

import formStyles from '../../styles/form.module.scss'

const Index = () => {
  const [errorMessage, setErrorMessage] = useState('')

  const { data, mutate } = useSWR('http://localhost/api/users', (url: string) => {
    return axios(url).then((res) => {
      return res.data as User[]
    })
  })

  const deleteUser = async (id: number) => {
    try {
      await axios.delete(`http://localhost/api/users/${id}`)
      mutate()
    } catch (error) {
      setErrorMessage(error.message)
    }
  }
  return (
    <>
      <h1>User List</h1>

      <Link href="/users/new">
        <Button variant="outlined">
          <PersonAdd />
          Create User
        </Button>
      </Link>

      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>id</TableCell>
              <TableCell>name</TableCell>
              <TableCell></TableCell>
              <TableCell></TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {data?.map((user) => {
              return (
                <TableRow key={user.id}>
                  <TableCell component="th" scope="row">
                    {user.id}
                  </TableCell>
                  <TableCell>{user.name}</TableCell>
                  <TableCell>
                    <Link href="/users/[id]" as={`/users/${user.id}`}>
                      <IconButton>
                        <Edit />
                      </IconButton>
                    </Link>
                  </TableCell>
                  <TableCell>
                    <IconButton>
                      <Delete
                        onClick={() => {
                          return deleteUser(user.id)
                        }}
                      />
                    </IconButton>
                  </TableCell>
                </TableRow>
              )
            })}
          </TableBody>
        </Table>
      </TableContainer>

      {errorMessage && (
        <p role="alert" className={formStyles.errorMessage}>
          {errorMessage}
        </p>
      )}
    </>
  )
}

// eslint-disable-next-line import/no-default-export
export default Index
