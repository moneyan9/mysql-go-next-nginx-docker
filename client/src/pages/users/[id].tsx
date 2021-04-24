import { Button, TextField } from '@material-ui/core'
import axios from 'axios'
import type { User } from 'models/user'
import Router from 'next/router'
import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'
import { FormProvider, useForm } from 'react-hook-form'
import useSWR from 'swr'

import formStyles from '../../styles/form.module.scss'

const Edit = () => {
  const router = useRouter()
  const { id } = router.query
  const { data } = useSWR(`http://localhost/api/users/${id}`, (url: string) => {
    return axios(url).then((res) => {
      return res.data as User
    })
  })

  const methods = useForm()
  const {
    handleSubmit,
    register,
    reset,
    formState: { errors },
  } = methods
  const [errorMessage, setErrorMessage] = useState('')

  useEffect(() => {
    reset(data)
  }, [reset, data])

  const updateUser = async (user: User) => {
    try {
      await axios.put(`http://localhost/api/users/${id}`, user)
      Router.push('/users')
    } catch (error) {
      setErrorMessage(error.message)
    }
  }

  return (
    <>
      <h1>Edit User</h1>

      <FormProvider {...methods}>
        <form onSubmit={handleSubmit(updateUser)} className={formStyles.form}>
          <div>
            <TextField type="text" label="ID" {...register('id', { valueAsNumber: true })} disabled />
          </div>
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
              Update
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
export default Edit
