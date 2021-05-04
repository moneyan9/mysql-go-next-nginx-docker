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
import Link from 'next/link'
import { useState } from 'react'
import type { Group } from 'src/models/group'
import useSWR from 'swr'

import formStyles from '../../styles/form.module.scss'

const Index = () => {
  const [errorMessage, setErrorMessage] = useState('')

  const { data, mutate } = useSWR('http://localhost/api/groups', (url: string) => {
    return axios(url).then((res) => {
      return res.data as Group[]
    })
  })

  const deleteGroup = async (id: number) => {
    try {
      await axios.delete(`http://localhost/api/groups/${id}`)
      mutate()
    } catch (error) {
      setErrorMessage(error.message)
    }
  }
  return (
    <>
      <h1>Group List</h1>

      <Link href="/groups/new">
        <Button variant="outlined">
          <PersonAdd />
          Create Group
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
            {data?.map((group) => {
              return (
                <TableRow key={group.id}>
                  <TableCell component="th" scope="row">
                    {group.id}
                  </TableCell>
                  <TableCell>{group.name}</TableCell>
                  <TableCell>
                    <Link href="/groups/[id]" as={`/groups/${group.id}`}>
                      <IconButton>
                        <Edit />
                      </IconButton>
                    </Link>
                  </TableCell>
                  <TableCell>
                    <IconButton
                      onClick={() => {
                        return deleteGroup(group.id)
                      }}
                    >
                      <Delete />
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
