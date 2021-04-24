import { Button, TextField } from '@material-ui/core'
import axios from 'axios'
import type { User } from 'models/user'
import Router from 'next/router'
import { useState } from 'react'
import { useForm } from 'react-hook-form'

import formStyles from '../../styles/form.module.scss'

const New = () => {
  const {
    handleSubmit,
    register,
    formState: { errors },
  } = useForm()
  const [errorMessage, setErrorMessage] = useState('')

  const { ref, ...inputProps } = register('name', {
    required: 'Name is required',
  })

  const addUser = async (user: User) => {
    try {
      await axios.post('http://localhost/api/users', user)
      Router.push('/users')
    } catch (error) {
      setErrorMessage(error.message)
    }
  }

  return (
    <>
      <h1>Create User</h1>

      <form onSubmit={handleSubmit(addUser)} className={formStyles.form}>
        <div>
          <TextField
            type="text"
            label="Name"
            error={!!errors.name}
            helperText={errors?.name?.message}
            inputRef={ref}
            {...inputProps}
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
