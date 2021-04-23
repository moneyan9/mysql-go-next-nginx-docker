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
import useSWR from 'swr'

const Index = () => {
  const { data } = useSWR('http://localhost/api/users', (url: string) => {
    return axios(url).then((res) => {
      return res.data
    })
  })

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
            {data?.map((row: any) => {
              return (
                <TableRow key={data.id}>
                  <TableCell component="th" scope="row">
                    {row.id}
                  </TableCell>
                  <TableCell>{row.name}</TableCell>
                  <TableCell>
                    <IconButton>
                      <Edit />
                    </IconButton>
                  </TableCell>
                  <TableCell>
                    <IconButton>
                      <Delete />
                    </IconButton>
                  </TableCell>
                </TableRow>
              )
            })}
          </TableBody>
        </Table>
      </TableContainer>
    </>
  )
}

// eslint-disable-next-line import/no-default-export
export default Index
