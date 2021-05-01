import { Button, TextField } from '@material-ui/core'
import axios from 'axios'
import Router from 'next/router'
import { useState } from 'react'
import { useForm } from 'react-hook-form'
import type { User } from 'src/models/user'
import formStyles from 'src/styles/form.module.scss'

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
