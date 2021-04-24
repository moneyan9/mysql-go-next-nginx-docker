import { Button, TextField } from '@material-ui/core'
import axios from 'axios'
import type { User } from 'models/user'
import Router from 'next/router'
import { useState } from 'react'
import { FormProvider, useForm } from 'react-hook-form'

import formStyles from '../../styles/form.module.scss'

const New = () => {
  const methods = useForm()

  const {
    handleSubmit,
    register,
    formState: { errors },
  } = methods

  const [errorMessage, setErrorMessage] = useState('')

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

      <FormProvider {...methods}>
        <form onSubmit={handleSubmit(addUser)} className={formStyles.form}>
          <div>
            <TextField type="text" label="Name" {...register('name', { required: 'Name is required' })} />
            {errors.name && (
              <span role="alert" className={formStyles.error}>
                {errors.name.message}
              </span>
            )}
          </div>

          <div className={formStyles.submit}>
            <Button type="submit" variant="outlined">
              Create
            </Button>
          </div>
        </form>
      </FormProvider>

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
