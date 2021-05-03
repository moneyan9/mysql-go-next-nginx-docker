import { Button, TextField } from '@material-ui/core'
import axios from 'axios'
import Router from 'next/router'
import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'
import { useForm } from 'react-hook-form'
import type { Group } from 'src/models/group'
import useSWR from 'swr'

import formStyles from '../../styles/form.module.scss'

const Edit = () => {
  const router = useRouter()
  const { id } = router.query
  const { data } = useSWR<Group>(`http://localhost/api/groups/${id}`)

  const {
    register,
    reset,
    formState: { errors },
    handleSubmit,
  } = useForm()
  const [errorMessage, setErrorMessage] = useState('')

  const { ref: idRef, ...idProps } = register('id')
  const { ref: nameRef, ...nameProps } = register('name', {
    required: 'Name is required',
  })

  const updateGroup = async (group: Group) => {
    try {
      group.id = Number.parseInt(id as string)
      await axios.put(`http://localhost/api/groups/${id}`, group)
      Router.push('/groups')
    } catch (error) {
      setErrorMessage(error.message)
    }
  }

  useEffect(() => {
    if (data) {
      reset(data)
    }
  }, [reset, data])

  useEffect
  return (
    <>
      <h1>Edit Group</h1>

      <form onSubmit={handleSubmit(updateGroup)} className={formStyles.form}>
        <div>
          <TextField type="text" label="ID" inputRef={idRef} {...idProps} disabled />
        </div>
        <div>
          <TextField
            type="text"
            label="Name"
            inputRef={nameRef}
            {...nameProps}
            error={!!errors.name}
            helperText={errors?.name?.message}
          />
        </div>

        <div className={formStyles.submit}>
          <Button type="submit" variant="outlined">
            Update
          </Button>
        </div>
      </form>

      {errorMessage && (
        <p role="alert" className={formStyles.errorMessage}>
          {errorMessage}
        </p>
      )}
    </>
  )
}

// eslint-disable-next-line import/no-default-export
export default Edit
