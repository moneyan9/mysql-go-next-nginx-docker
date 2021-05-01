import { Button, TextField } from '@material-ui/core'
import axios from 'axios'
import Router from 'next/router'
import { useState } from 'react'
import { useForm } from 'react-hook-form'
import type { Group } from 'src/models/group'

import formStyles from '../../styles/form.module.scss'

const New = () => {
  const {
    register,
    formState: { errors },
    handleSubmit,
  } = useForm()
  const [errorMessage, setErrorMessage] = useState('')

  const { ref: nameRef, ...nameProps } = register('name', {
    required: 'Name is required',
  })

  const addGroup = async (group: Group) => {
    try {
      await axios.post('http://localhost/api/groups', group)
      Router.push('/groups')
    } catch (error) {
      setErrorMessage(error.message)
    }
  }

  return (
    <>
      <h1>Create Group</h1>

      <form onSubmit={handleSubmit(addGroup)} className={formStyles.form}>
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
            Create
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
export default New
